package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"search_api/domain"
	"search_api/utils"
	"strings"
)

var solrURL = os.Getenv("SOLR_BASE_URL")

func AddOrUpdateHotel(h domain.Hotel) error {

	doc := fmt.Sprintf(`
	{
	  "add": {
		"doc": {
		  "id": "%s",
		  "nombre": ["%s"],
		  "provincia": ["%s"],
		  "descripcion": ["%s"],
		  "imagen": ["%s"],
		  "direccion": ["%s"],
		  "precio": [%f]
		}
	  }
	}`, h.ID, h.Nombre, h.Provincia, h.Descripcion, h.Imagen, h.Direccion, h.Precio)

	req, _ := http.NewRequest("POST", solrURL+"/update?commit=true",
		strings.NewReader(doc))

	req.Header.Set("Content-Type", "application/json")
	_, err := utils.HttpClient.Do(req)
	return err
}

func DeleteHotel(id string) error {
	body := fmt.Sprintf(`{"delete":{"id":"%s"}}`, id)

	req, _ := http.NewRequest("POST", solrURL+"/update?commit=true",
		strings.NewReader(body))

	req.Header.Set("Content-Type", "application/json")
	_, err := utils.HttpClient.Do(req)
	return err
}

func SearchHotels(q string, page, size int) (int, []domain.Hotel, error) {

	params := url.Values{}
	if q == "" {
		params.Set("q", "*:*")
	} else {
		params.Set("q", fmt.Sprintf("nombre:%s OR provincia:%s", q, q))
	}

	params.Set("start", fmt.Sprintf("%d", (page-1)*size))
	params.Set("rows", fmt.Sprintf("%d", size))
	params.Set("wt", "json")

	resp, err := utils.HttpClient.Get(solrURL + "/select?" + params.Encode())
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Response struct {
			NumFound int `json:"numFound"`
			Docs     []struct {
				ID          string    `json:"id"`
				Nombre      []string  `json:"nombre"`
				Imagen      []string  `json:"imagen"`
				Descripcion []string  `json:"descripcion"`
				Provincia   []string  `json:"provincia"`
				Direccion   []string  `json:"direccion"`
				Precio      []float64 `json:"precio"`
			} `json:"docs"`
		} `json:"response"`
	}

	json.NewDecoder(resp.Body).Decode(&result)

	var hotels []domain.Hotel

	for _, d := range result.Response.Docs {
		hotels = append(hotels, domain.Hotel{
			ID:          d.ID,
			Nombre:      firstOrEmpty(d.Nombre),
			Imagen:      firstOrEmpty(d.Imagen),
			Descripcion: firstOrEmpty(d.Descripcion),
			Provincia:   firstOrEmpty(d.Provincia),
			Direccion:   firstOrEmpty(d.Direccion),
			Precio:      firstOrZero(d.Precio),
		})
	}

	return result.Response.NumFound, hotels, nil
}

func firstOrEmpty(arr []string) string {
	if len(arr) > 0 {
		return arr[0]
	}
	return ""
}

func firstOrZero(arr []float64) float64 {
	if len(arr) > 0 {
		return arr[0]
	}
	return 0
}
