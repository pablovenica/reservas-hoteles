package service

import (
    "encoding/json"
    "fmt"
    "net/http"
    "search_api/domain"
    "search_api/repository"
    "search_api/cache"
    "search_api/config"
)

type SearchResult struct {
    Total int            `json:"total"`
    Items []domain.Hotel `json:"items"`
}

func Search(q string, page, size int) (int, []domain.Hotel, error) {
    key := cache.MakeSearchKey(q, page, size)

    if cached, ok := cache.Get[SearchResult](key); ok {
        return cached.Total, cached.Items, nil
    }

    total, items, err := repository.SearchHotels(q, page, size)
    if err != nil {
        return 0, nil, err
    }

    cache.Set(key, SearchResult{
        Total: total,
        Items: items,
    })

    return total, items, nil
}

func getHotelFromMainAPI(id string) (*domain.Hotel, error) {
    url := fmt.Sprintf("%s/hotels/%s", config.MainAPIURL, id)

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("error fetching hotel %s: status %d", id, resp.StatusCode)
    }

    var hotel domain.Hotel
    if err := json.NewDecoder(resp.Body).Decode(&hotel); err != nil {
        return nil, err
    }

    return &hotel, nil
}

func IndexHotel(id string) error {
    h, err := getHotelFromMainAPI(id)
    if err != nil {
        return err
    }
    return repository.AddOrUpdateHotel(*h)
}

func UpdateHotel(id string) error {
    return IndexHotel(id)
}

func DeleteHotel(id string) error {
    return repository.DeleteHotel(id)
}

func InvalidateSearchCache() {
    cache.ClearPrefix("search") // ← CORRECCIÓN
}
