package dto

type Hotel struct {
    ID       string `json:"id,omitempty"`
    Titulo   string `json:"titulo"`
    Nivel    int    `json:"nivel"`
    Estado   bool   `json:"estado"`
    Duracion string `json:"duracion"`
    Precio   string `json:"precio"`
}

type Hotels []Hotel
