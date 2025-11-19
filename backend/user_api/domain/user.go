package domain

type User struct {
	ID          int    `gorm:"column:id_usuarios;primaryKey;autoIncrement"`
	Nombre      string `gorm:"type:varchar(45);not null"`
	Email       string `gorm:"type:varchar(45);unique;not null"`
	Password    string `gorm:"column:password"`
	TipoUsuario string `gorm:"column:tipo_usuarios"`
}

type Users []User