package cache

import (
    "time"

    "github.com/karlseguin/ccache"
    "hotels_api/dto"
)

var Cache = ccache.New(ccache.Configure().MaxSize(200))

// Guarda lista de hoteles
func SetHotels(key string, hotels dto.Hotels, ttl time.Duration) {
    Cache.Set(key, hotels, ttl)
}

// Obtiene lista de hoteles
func GetHotels(key string) (dto.Hotels, bool) {
    item := Cache.Get(key)
    if item == nil || item.Expired() {
        return nil, false
    }
    return item.Value().(dto.Hotels), true
}

// Borra un item
func Delete(key string) {
    Cache.Delete(key)
}

// Invalida TODAS las claves relacionadas a hoteles
func Invalidate() {
    Cache.Delete("hotels_all") // lista completa

    // Podrías también borrar por prefijo, pero ccache no soporta eso, así que:
    // O borras claves individuales cuando actualizas/eliminás
    // Aquí mantenemos lo básico: lista general.
}
