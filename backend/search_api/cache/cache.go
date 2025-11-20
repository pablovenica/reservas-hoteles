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
    // Cache local (in-memory, súper rápida)
    localCache = ccache.New(ccache.Configure().
        MaxSize(5000).
        ItemsToPrune(100))

    // Cache distribuida (Memcached)
    memcached = memcache.New("memcached:11211") // Cambiar si tu host es diferente
}

// ----------------------------------------------------------------------
// GENERACIÓN DE KEYS
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
// GET – lee desde CCache y Memcached
// ----------------------------------------------------------------------

func Get[T any](key string) (*T, bool) {

    // 1) Buscar en CCache
    if item := localCache.Get(key); item != nil && !item.Expired() {
        if value, ok := item.Value().(*T); ok {
            return value, true
        }
    }

    // 2) Buscar en Memcached
    data, err := memcached.Get(key)
    if err == nil && data != nil {
        var value T
        if err := json.Unmarshal(data.Value, &value); err == nil {

            // Hidratar CCache
            localCache.Set(key, &value, 5*time.Minute)

            return &value, true
        }
    }

    return nil, false
}

// ----------------------------------------------------------------------
// SET – guarda en CCache y Memcached
// ----------------------------------------------------------------------

func Set[T any](key string, value T) error {

    // 1) Guardar en CCache
    localCache.Set(key, &value, 5*time.Minute)

    // 2) Guardar en Memcached
    jsonBytes, err := json.Marshal(value)
    if err != nil {
        return err
    }

    return memcached.Set(&memcache.Item{
        Key:        key,
        Value:      jsonBytes,
        Expiration: int32(10 * 60),
    })
}

func ClearPrefix(prefix string) {

    // 1) Subir la versión → invalida Memcached automáticamente
    versionKeys[prefix]++

    // 2) Borrar CCache completo del prefijo
    items := localCache.Items()
    for _, item := range items {
        if len(item.Key()) >= len(prefix) && item.Key()[:len(prefix)] == prefix {
            localCache.Delete(item.Key())
        }
    }
}

// ----------------------------------------------------------------------
// INVALIDATE – borra cache del prefijo "search"
// ----------------------------------------------------------------------

func InvalidateSearchCache() {
    ClearPrefix("search")
}
