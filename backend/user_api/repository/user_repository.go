package repository

import (
	"user_api/domain"
	"crypto/sha256"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetUserById(id int) (domain.User, error) {
	var user domain.User
	err := DB.First(&user, id).Error
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func VerifyCredentials(email, password string) (domain.User, bool) {
	var user domain.User
	hash := sha256.Sum256([]byte(password))
	passwordHash := fmt.Sprintf("%x", hash)

	result := DB.Where("email = ? AND password = ?", email, passwordHash).First(&user)
	print("este es el resultado:")
	print(passwordHash)
	if result.Error != nil {
		log.Warn("Credenciales inválidas o usuario no encontrado.")
		return domain.User{}, false
	}

	log.Debug("Usuario autenticado: ", user.Email)
	return user, true
}
func CreateUser(u domain.User) (domain.User, error) {
    hash := sha256.Sum256([]byte(u.Password))
    u.Password = fmt.Sprintf("%x", hash)

    result := DB.Create(&u)
    if result.Error != nil {
        return domain.User{}, result.Error
    }
    return u, nil
}
