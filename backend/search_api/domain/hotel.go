package domain

type Hotel struct {
	ID          string
	Nombre      string
	Imagen      string
	Provincia   string
	Descripcion string
	Direccion   string
	Precio      float64
}

type Hotels []Hotel
