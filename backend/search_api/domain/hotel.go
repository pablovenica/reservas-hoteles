package domain

type Hotel struct {
	ID          string  `json:"id"`
	Nombre      string  `json:"nombre"`
	Imagen      string  `json:"imagen"`
	Descripcion string  `json:"descripcion"`
	Provincia   string  `json:"provincia"`
	Direccion   string  `json:"direccion"`
	Precio      float64 `json:"precio"`
}
