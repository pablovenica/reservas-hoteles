package service

import (
	"encoding/json"
	"fmt"

	"search_api/cache"
	"search_api/domain"
	"search_api/repository"
	"search_api/utils"
)

type SearchResult struct {
	Total int            `json:"total"`
	Items []domain.Hotel `json:"items"`
}

func Search(q string, page, size int) (int, []domain.Hotel, error) {
	key := cache.MakeSearchKey(q, page, size)

	// 1) Intentar leer desde cache (CCache o Memcached)
	if cached, ok := cache.Get[SearchResult](key); ok {
		return cached.Total, cached.Items, nil
	}

	// 2) Si no está en cache, consultar en Solr
	total, items, err := repository.SearchHotels(q, page, size)
	if err != nil {
		return 0, nil, err
	}

	// 3) Guardar en cache para próximas llamadas
	_ = cache.Set(key, SearchResult{
		Total: total,
		Items: items,
	})

	return total, items, nil
}

// NUEVO: reindexa todos los hoteles que existan en hotels_api hacia Solr
func ReindexAllHotels() error {
	// 1) Pedir todos los hoteles a hotels_api
	//    (hostname del servicio en Docker: hotels_api:8082)
	url := "http://hotels_api:8082/hotels"

	resp, err := utils.HttpClient.Get(url)
	if err != nil {
		return fmt.Errorf("error pidiendo hoteles a hotels_api: %w", err)
	}
	defer resp.Body.Close()

	// 2) Decodificar respuesta
	//    Sabemos que /hotels responde:
	//    { "hoteles": [ {id, nombre, ...}, ... ] }
	var data struct {
		Hoteles []domain.Hotel `json:"hoteles"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return fmt.Errorf("error decodificando respuesta de hotels_api: %w", err)
	}

	// 3) Mandar cada hotel a Solr con AddOrUpdateHotel
	for _, h := range data.Hoteles {
		if err := repository.AddOrUpdateHotel(h); err != nil {
			return fmt.Errorf("error indexando hotel %s en Solr: %w", h.ID, err)
		}
	}

	return nil
}
