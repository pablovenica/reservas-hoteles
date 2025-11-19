package dto

type Hotel struct {
	ID          string  `json:"id"`
	Nombre      string  `json:"nombre"`
	Imagen      string  `json:"imagen"`
	Provincia   string  `json:"provincia"`
	Descripcion string  `json:"descripcion"`
	Direccion   string  `json:"direccion"`
	Precio      float64 `json:"precio"`
}

// Respuesta de búsqueda paginada
type HotelSearchResponse struct {
	Total    int     `json:"total"`
	Page     int     `json:"page"`
	PageSize int     `json:"page_size"`
	Results  []Hotel `json:"results"`
}
