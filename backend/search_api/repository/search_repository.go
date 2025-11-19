package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"search_api/domain"
)

type SearchRepository struct {
	httpClient  *http.Client
	solrBaseURL string
}

func NewSearchRepository(baseURL string) *SearchRepository {
	return &SearchRepository{
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		solrBaseURL: baseURL,
	}
}

type solrDoc struct {
	ID          string    `json:"id"`
	Nombre      []string  `json:"nombre"`
	Imagen      []string  `json:"imagen"`
	Provincia   []string  `json:"provincia"`
	Descripcion []string  `json:"descripcion"`
	Direccion   []string  `json:"direccion"`
	Precio      []float64 `json:"precio"`

	Capacidad []int    `json:"capacidad,omitempty"`
	Servicios []string `json:"servicios,omitempty"`
}

type solrResponse struct {
	Response struct {
		NumFound int       `json:"numFound"`
		Docs     []solrDoc `json:"docs"`
	} `json:"response"`
}

func firstString(values []string) string {
	if len(values) > 0 {
		return values[0]
	}
	return ""
}

func firstFloat(values []float64) float64 {
	if len(values) > 0 {
		return values[0]
	}
	return 0
}

func (r *SearchRepository) SearchHotels(q string, page, pageSize int) ([]domain.Hotel, int, error) {
	if page < 1 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	start := (page - 1) * pageSize

	params := url.Values{}
	if q == "" {
		params.Set("q", "*:*")
	} else {
		params.Set("q", fmt.Sprintf("nombre:%s OR provincia:%s", q, q))
	}

	params.Set("start", strconv.Itoa(start))
	params.Set("rows", strconv.Itoa(pageSize))
	params.Set("wt", "json")

	fullURL := fmt.Sprintf("%s/select?%s", r.solrBaseURL, params.Encode())

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, 0, err
	}

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, 0, fmt.Errorf("Solr devolvió status %d", resp.StatusCode)
	}

	var solrResp solrResponse
	if err := json.NewDecoder(resp.Body).Decode(&solrResp); err != nil {
		return nil, 0, err
	}

	hotels := make([]domain.Hotel, 0, len(solrResp.Response.Docs))
	for _, d := range solrResp.Response.Docs {
		hotels = append(hotels, domain.Hotel{
			ID:          d.ID,
			Nombre:      firstString(d.Nombre),
			Imagen:      firstString(d.Imagen),
			Provincia:   firstString(d.Provincia),
			Descripcion: firstString(d.Descripcion),
			Direccion:   firstString(d.Direccion),
			Precio:      firstFloat(d.Precio),
		})
	}

	return hotels, solrResp.Response.NumFound, nil
}
