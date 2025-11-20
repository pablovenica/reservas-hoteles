package cache

import (
    "encoding/json"
    "fmt"
    "time"

    "github.com/bradfitz/gomemcache/memcache"
    "github.com/karlseguin/ccache/v2"
)

var (
    localCache *ccache.Cache
    memcached  *memcache.Client
)

// ----------------------------------------------------------------------
// INIT
// ----------------------------------------------------------------------

func Init() {
    localCache = ccache.New(ccache.Configure().
        MaxSize(5000).
        ItemsToPrune(100))
    memcached = memcache.New("memcached:11211")
}

// ----------------------------------------------------------------------
// KEYS
// ----------------------------------------------------------------------
var versionKeys = make(map[string]int)

func getPrefixVersion(prefix string) int {
    if v, ok := versionKeys[prefix]; ok {
        return v
    }
    versionKeys[prefix] = 1
    return 1
}

func MakeSearchKey(query string, page, size int) string {
    prefix := "search"
    version := getPrefixVersion(prefix)
    return fmt.Sprintf("%s:v%d:%s:p%d:s%d", prefix, version, query, page, size)
}

// ----------------------------------------------------------------------
// GET – CCache → Memcached → Solr
// ----------------------------------------------------------------------
func Get[T any](key string) (*T, bool) {
    start := time.Now()

    // 1) Buscar en CCache
    if item := localCache.Get(key); item != nil && !item.Expired() {
        if value, ok := item.Value().(*T); ok {
            fmt.Printf("[CACHE HIT] CCache key=%s tiempo=%s\n", key, time.Since(start))
            return value, true
        }
    }

    // 2) Buscar en Memcached
    data, err := memcached.Get(key)
    if err == nil && data != nil {
        var value T
        if err := json.Unmarshal(data.Value, &value); err == nil {
            localCache.Set(key, &value, 5*time.Minute)
            fmt.Printf("[CACHE HIT] Memcached key=%s tiempo=%s\n", key, time.Since(start))
            return &value, true
        }
    }

    fmt.Printf("[CACHE MISS] key=%s tiempo=%s\n", key, time.Since(start))
    return nil, false
}

// ----------------------------------------------------------------------
// SET
// ----------------------------------------------------------------------
func Set[T any](key string, value T) error {
    start := time.Now()

    // 1) Guardar en CCache
    localCache.Set(key, &value, 5*time.Minute)

    // 2) Guardar en Memcached
    jsonBytes, err := json.Marshal(value)
    if err != nil {
        return err
    }

    err = memcached.Set(&memcache.Item{
        Key:        key,
        Value:      jsonBytes,
        Expiration: int32(10 * 60),
    })

    fmt.Printf("[CACHE SET] key=%s tiempo=%s\n", key, time.Since(start))
    return err
}

// ----------------------------------------------------------------------
// CLEAR
// ----------------------------------------------------------------------
func ClearPrefix(prefix string) {
    versionKeys[prefix]++
    items := localCache.Items()
    for _, item := range items {
        if len(item.Key()) >= len(prefix) && item.Key()[:len(prefix)] == prefix {
            localCache.Delete(item.Key())
        }
    }
}

func InvalidateSearchCache() {
    ClearPrefix("search")
}
