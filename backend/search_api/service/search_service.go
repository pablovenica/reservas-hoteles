package service

import (
    "search_api/cache"
    "search_api/domain"
    "search_api/repository"
)

type SearchResult struct {
    Total int             `json:"total"`
    Items []domain.Hotel  `json:"items"`
}

func Search(q string, page, size int) (int, []domain.Hotel, error) {
    key := cache.MakeSearchKey(q, page, size)

    // ----- 1) Cache (CCache + Memcached) -----
    if cached, ok := cache.Get[SearchResult](key); ok {
        return cached.Total, cached.Items, nil
    }

    // ----- 2) Solr -----
    total, items, err := repository.SearchHotels(q, page, size)
    if err != nil {
        return 0, nil, err
    }

    // ----- 3) Guardar en cache -----
    cache.Set(key, SearchResult{
        Total: total,
        Items: items,
    })

    return total, items, nil
}
