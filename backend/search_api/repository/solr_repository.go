package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"search_api/domain"
)

var solrURL = os.Getenv("SOLR_BASE_URL")

type solrResponse struct {
	Response struct {
		NumFound int `json:"numFound"`
		Docs     []struct {
			ID          string    `json:"id"`
			Nombre      []string  `json:"nombre"`
			Provincia   []string  `json:"provincia"`
			Descripcion []string  `json:"descripcion"`
			Imagen      []string  `json:"imagen"`
			Precio      []float64 `json:"precio"`
			Direccion   []string  `json:"direccion"`
		} `json:"docs"`
	} `json:"response"`
}

func SearchSolr(q string, page, size int) (int, []domain.Hotel, error) {

	var encodedQuery string

	if strings.TrimSpace(q) == "" {
		encodedQuery = url.QueryEscape("*:*")
	} else {
		query := fmt.Sprintf(
			"nombre:*%[1]s* OR provincia:*%[1]s* OR descripcion:*%[1]s* OR direccion:*%[1]s*",
			q,
		)
		encodedQuery = url.QueryEscape(query)
	}

	params := fmt.Sprintf(
		"q=%s&start=%d&rows=%d",
		encodedQuery,
		(page-1)*size,
		size,
	)

	urlFinal := fmt.Sprintf("%s/select?%s&wt=json", solrURL, params)

	resp, err := http.Get(urlFinal)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	var data solrResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, nil, err
	}

	hotels := []domain.Hotel{}

	for _, d := range data.Response.Docs {
		hotels = append(hotels, domain.Hotel{
			ID:          d.ID,
			Nombre:      safeString(d.Nombre),
			Provincia:   safeString(d.Provincia),
			Descripcion: safeString(d.Descripcion),
			Imagen:      safeString(d.Imagen),
			Precio:      safeFloat(d.Precio),
			Direccion:   safeString(d.Direccion),
		})
	}

	return data.Response.NumFound, hotels, nil
}

func safeString(arr []string) string {
	if len(arr) > 0 {
		return arr[0]
	}
	return ""
}

func safeFloat(arr []float64) float64 {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}

func AddOrUpdateHotel(h domain.Hotel) error {

	payload := fmt.Sprintf(`{
		"add": {
			"doc": {
				"id": "%s",
				"nombre": ["%s"],
				"provincia": ["%s"],
				"descripcion": ["%s"],
				"imagen": ["%s"],
				"precio": [%f],
				"direccion": ["%s"]
			}
		}
	}`, h.ID, h.Nombre, h.Provincia, h.Descripcion, h.Imagen, h.Precio, h.Direccion)

	urlFinal := fmt.Sprintf("%s/update?commit=true", solrURL)

	_, err := http.Post(urlFinal, "application/json", strings.NewReader(payload))
	return err
}

func DeleteHotelFromSolr(id string) error {

	payload := fmt.Sprintf(`{
		"delete": { "id": "%s" }
	}`, id)

	urlFinal := fmt.Sprintf("%s/update?commit=true", solrURL)

	_, err := http.Post(urlFinal,
		"application/json",
		strings.NewReader(payload),
	)
	return err
}
