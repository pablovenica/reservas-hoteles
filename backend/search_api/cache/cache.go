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

func MakeSearchKey(query string, page, size int) string {
    return fmt.Sprintf("search:%s:p%d:s%d", query, page, size)
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

// ----------------------------------------------------------------------
// INVALIDATE – borra cache local (simple)
// ----------------------------------------------------------------------

func InvalidateSearchCache() {
    localCache.Clear()
    // Memcached NO se borra completa porque puede contener otras keys.
    // Si querés borrar todo, pedirlo explícitamente.
}
