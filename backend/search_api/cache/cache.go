package cache

import (
	"encoding/json"
	"log"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/karlseguin/ccache/v2"
)

var (
	LocalCache *ccache.Cache
	Client     *memcache.Client
	version    = "v1"
)

func InitCache() {
	LocalCache = ccache.New(ccache.Configure().MaxSize(10000))
	Client = memcache.New("memcached:11211")
	log.Println("Cache inicializada")
}

func MakeKey(prefix, q string, page, size int) string {
	return version + ":" + prefix + ":" + q + ":p" + string(rune(page)) + ":s" + string(rune(size))
}

func Get[T any](key string) (T, bool) {
	var zero T

	item := LocalCache.Get(key)
	if item != nil && !item.Expired() {
		return item.Value().(T), true
	}

	entry, err := Client.Get(key)
	if err == nil && entry != nil {
		var val T
		if json.Unmarshal(entry.Value, &val) == nil {
			LocalCache.Set(key, val, time.Minute*5)
			return val, true
		}
	}

	return zero, false
}

func Set[T any](key string, value T) {
	LocalCache.Set(key, value, time.Minute*5)

	data, _ := json.Marshal(value)
	Client.Set(&memcache.Item{Key: key, Value: data, Expiration: int32(300)})
}

func ClearPrefix(prefix string) {
	version = version + "_new"
}
