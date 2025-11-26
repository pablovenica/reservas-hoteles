package service

import (
	"errors"
	"testing"

	userRepository "user_api/repository"
	"user_api/domain"
	"user_api/dto"
)

// helper para crear un usuario de prueba directamente en el repo
func crearUsuarioDePrueba(t *testing.T, nombre, email, password, tipo string) domain.User {
	t.Helper()

	u := domain.User{
		Nombre:      nombre,
		Email:       email,
		Password:    password,
		TipoUsuario: tipo,
	}

	creado, err := userRepository.CreateUser(u)
	if err != nil {
		t.Fatalf("no se pudo crear usuario de prueba: %v", err)
	}
	return creado
}

// Tests GetUsuarioDtoById

func TestGetUsuarioDtoById_OK(t *testing.T) {

	u := crearUsuarioDePrueba(t,
		"Juan Perez",
		"juan@example.com",
		"secreto",
		"cliente",
	)

	dtoUsuario, err := GetUsuarioDtoById(u.ID)

	if err != nil {
		t.Fatalf("se esperaba usuario sin error, pero se obtuvo error: %v", err)
	}

	if dtoUsuario.ID != u.ID {
		t.Errorf("ID esperado %d, obtenido %d", u.ID, dtoUsuario.ID)
	}
	if dtoUsuario.Nombre != u.Nombre {
		t.Errorf("Nombre esperado %q, obtenido %q", u.Nombre, dtoUsuario.Nombre)
	}
	if dtoUsuario.Email != u.Email {
		t.Errorf("Email esperado %q, obtenido %q", u.Email, dtoUsuario.Email)
	}
	if dtoUsuario.TipoUsuario != u.TipoUsuario {
		t.Errorf("TipoUsuario esperado %q, obtenido %q", u.TipoUsuario, dtoUsuario.TipoUsuario)
	}
}

func TestGetUsuarioDtoById_NotFound(t *testing.T) {
	// Arrange: usamos un ID que razonablemente no exista
	idInexistente := 999999


	_, err := GetUsuarioDtoById(idInexistente)


	if err == nil {
		t.Fatalf("se esperaba error al buscar ID inexistente (%d), pero no hubo error", idInexistente)
	}
}

// Tests de Login


func TestLogin_ValidCredentials(t *testing.T) {
	// Arrange
	email := "login_ok@example.com"
	password := "clave123"
	tipo := "admin"

	u := crearUsuarioDePrueba(t, "Login OK", email, password, tipo)


	dtoUsuario, err := Login(email, password)


	if err != nil {
		t.Fatalf("se esperaba login exitoso, pero se obtuvo error: %v", err)
	}

	if dtoUsuario.ID != u.ID {
		t.Errorf("ID esperado %d, obtenido %d", u.ID, dtoUsuario.ID)
	}
	if dtoUsuario.Email != email {
		t.Errorf("Email esperado %q, obtenido %q", email, dtoUsuario.Email)
	}
	if dtoUsuario.TipoUsuario != tipo {
		t.Errorf("TipoUsuario esperado %q, obtenido %q", tipo, dtoUsuario.TipoUsuario)
	}
}

func TestLogin_InvalidCredentials(t *testing.T) {
	// Caso 1: email inexistente
	_, err := Login("no_existe@example.com", "cualquier")
	if err == nil {
		t.Fatalf("se esperaba error por credenciales inválidas (usuario inexistente), pero no hubo error")
	}

	// Caso 2: password incorrecto (suponiendo que el repo valida password)
	email := "login_fail@example.com"
	passwordCorrecto := "ok123"
	crearUsuarioDePrueba(t, "Login Fail", email, passwordCorrecto, "cliente")

	_, err = Login(email, "password_incorrecto")
	if err == nil {
		t.Fatalf("se esperaba error por credenciales inválidas (password incorrecto), pero no hubo error")
	}
}

// Tests de CrearUsuario

func TestCrearUsuario_OK(t *testing.T) {
	// Arrange
	input := dto.UserDto{
		Nombre:      "Nuevo Usuario",
		Email:       "nuevo@example.com",
		Password:    "pass_nuevo",
		TipoUsuario: "cliente",
	}


	creado, err := CrearUsuario(input)


	if err != nil {
		t.Fatalf("se esperaba creación de usuario sin error, pero se obtuvo: %v", err)
	}

	if creado.ID == 0 {
		t.Errorf("se esperaba que el ID del usuario creado sea distinto de 0")
	}
	if creado.Nombre != input.Nombre {
		t.Errorf("Nombre esperado %q, obtenido %q", input.Nombre, creado.Nombre)
	}
	if creado.Email != input.Email {
		t.Errorf("Email esperado %q, obtenido %q", input.Email, creado.Email)
	}
	if creado.TipoUsuario != input.TipoUsuario {
		t.Errorf("TipoUsuario esperado %q, obtenido %q", input.TipoUsuario, creado.TipoUsuario)
	}

	// Importante: por diseño mapUsuarioToDto NO expone el password
	if creado.Password != "" {
		t.Errorf("no se esperaba que Password se exponga en el DTO (debería estar vacío)")
	}
}


// Tests mapUsuarioToDto 

func TestMapUsuarioToDto(t *testing.T) {
	// Arrange
	u := domain.User{
		ID:          10,
		Nombre:      "Map Test",
		Email:       "map@test.com",
		Password:    "no_deberia_salir",
		TipoUsuario: "cliente",
	}

	// Act
	dtoUsuario := mapUsuarioToDto(u)

	// Assert
	if dtoUsuario.ID != u.ID {
		t.Errorf("ID esperado %d, obtenido %d", u.ID, dtoUsuario.ID)
	}
	if dtoUsuario.Nombre != u.Nombre {
		t.Errorf("Nombre esperado %q, obtenido %q", u.Nombre, dtoUsuario.Nombre)
	}
	if dtoUsuario.Email != u.Email {
		t.Errorf("Email esperado %q, obtenido %q", u.Email, dtoUsuario.Email)
	}
	if dtoUsuario.TipoUsuario != u.TipoUsuario {
		t.Errorf("TipoUsuario esperado %q, obtenido %q", u.TipoUsuario, dtoUsuario.TipoUsuario)
	}

	// Password NO debe mapearse al DTO
	if dtoUsuario.Password != "" {
		t.Errorf("no se esperaba que el Password se incluya en el DTO")
	}
}

