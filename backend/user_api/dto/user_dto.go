package dto

type UserDto struct {
	ID          int    `json:"id_usuarios"`
	Nombre      string `json:"nombre"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	TipoUsuario string `json:"tipo_usuario"`
}
