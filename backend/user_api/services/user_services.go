// usuarios_services.go
package services

import (
	userRepository "user_api/repository"
	"user_api/domain"
	"user_api/dto"
	"errors"
)

func GetUsuarioDtoById(id int) (dto.UserDto, error) {
	userDomain, err := userRepository.GetUserById(id)
	if err != nil {
		return dto.UserDto{}, err
	}
	return mapUsuarioToDto(userDomain), nil
}

func Login(email, password string) (dto.UserDto, error) {
	userDomain, valido := userRepository.VerifyCredentials(email, password)
	if !valido {
		return dto.UserDto{}, errors.New("credenciales inválidas")
	}
	return mapUsuarioToDto(userDomain), nil
}

func mapUsuarioToDto(u domain.User) dto.UserDto {
	return dto.UserDto{
		ID:          u.ID,
		Nombre:      u.Nombre,
		Email:       u.Email,
		TipoUsuario: u.TipoUsuario,
	}
}
func CrearUsuario(u dto.UserDto) (dto.UserDto, error) {
    userDomain := domain.User{
        Nombre:      u.Nombre,
        Email:       u.Email,
        Password:    u.Password,
        TipoUsuario: u.TipoUsuario,
    }

    creado, err := userRepository.CreateUser(userDomain)
    if err != nil {
        return dto.UserDto{}, err
    }

    return mapUsuarioToDto(creado), nil
}
