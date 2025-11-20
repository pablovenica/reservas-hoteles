package domain

type Hotel struct {
	ID          string  `json:"id"`
	Nombre      string  `json:"nombre"`
	Provincia   string  `json:"provincia"`
	Descripcion string  `json:"descripcion"`
	Imagen      string  `json:"imagen"`
	Precio      float64 `json:"precio"`
	Direccion   string  `json:"direccion"`
}
