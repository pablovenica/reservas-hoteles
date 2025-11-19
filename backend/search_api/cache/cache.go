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
// INIT – inicializa CCache (local) y Memcached (remota)
// ----------------------------------------------------------------------

func Init() {
	// Cache local en memoria (rápida, dentro del proceso)
	localCache = ccache.New(ccache.Configure().
		MaxSize(5000).
		ItemsToPrune(100))

	// Cache distribuida (Memcached en Docker)
	// IMPORTANTE: el host "memcached" debe existir en docker-compose
	memcached = memcache.New("memcached:11211")
}

// ----------------------------------------------------------------------
// GENERACIÓN DE KEYS
// ----------------------------------------------------------------------

func MakeSearchKey(query string, page, size int) string {
	return fmt.Sprintf("search:%s:p%d:s%d", query, page, size)
}

// ----------------------------------------------------------------------
// GET – lee desde CCache y, si no existe, desde Memcached
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
			// Rehidratar CCache para siguientes lecturas
			localCache.Set(key, &value, 5*time.Minute)
			return &value, true
		}
	}

	return nil, false
}

// ----------------------------------------------------------------------
// SET – guarda en CCache y en Memcached
// ----------------------------------------------------------------------

func Set[T any](key string, value T) error {
	// 1) Guardar en CCache
	localCache.Set(key, &value, 5*time.Minute)

	// 2) Guardar en Memcached (como JSON)
	jsonBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return memcached.Set(&memcache.Item{
		Key:        key,
		Value:      jsonBytes,
		Expiration: int32(10 * 60), // 10 minutos en Memcached
	})
}

// ----------------------------------------------------------------------
// INVALIDATE – borra cache local (simple)
// ----------------------------------------------------------------------

func InvalidateSearchCache() {
	localCache.Clear()
	// Memcached NO se limpia porque podría tener otras keys.
}
