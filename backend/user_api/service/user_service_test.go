package service

import (
	"os"
	"testing"

	"user_api/bd"
	"user_api/domain"
	"user_api/dto"
	userRepository "user_api/repository"
)
func TestMain(m *testing.M) {
    os.Setenv("DB_USER", "root")
    os.Setenv("DB_PASSWORD", "kaneki15..")
    os.Setenv("DB_NAME", "users_db")

    // MUY IMPORTANTE:
    os.Setenv("DB_HOST", "127.0.0.1") // host del Docker visto desde tu PC
    os.Setenv("DB_PORT", "3307")      // el puerto publicado por docker-compose

    bd.Init()
    bd.StartDbEngine()

    code := m.Run()
    os.Exit(code)
}


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

// -----------------------------------------------------------------------------
// Tests de GetUsuarioDtoById
// -----------------------------------------------------------------------------

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
}

func TestGetUsuarioDtoById_NotFound(t *testing.T) {
	_, err := GetUsuarioDtoById(999999)
	if err == nil {
		t.Fatalf("se esperaba error por usuario inexistente")
	}
}

// -----------------------------------------------------------------------------
// Tests de Login
// -----------------------------------------------------------------------------

func TestLogin_ValidCredentials(t *testing.T) {
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
}

func TestLogin_InvalidCredentials(t *testing.T) {
	_, err := Login("no_existe@example.com", "cualquiera")
	if err == nil {
		t.Fatalf("se esperaba error por credenciales inválidas")
	}
}

// -----------------------------------------------------------------------------
// Test de CrearUsuario
// -----------------------------------------------------------------------------

func TestCrearUsuario_OK(t *testing.T) {
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
		t.Errorf("el ID del usuario creado no debería ser 0")
	}
}

// -----------------------------------------------------------------------------
// Test unitario puro de mapUsuarioToDto (no toca repo ni DB)
// -----------------------------------------------------------------------------

func TestMapUsuarioToDto(t *testing.T) {
	u := domain.User{
		ID:          10,
		Nombre:      "Map Test",
		Email:       "map@test.com",
		Password:    "no_deberia_salir",
		TipoUsuario: "cliente",
	}

	dtoUsuario := mapUsuarioToDto(u)

	if dtoUsuario.Password != "" {
		t.Errorf("password no debe exponerse en el DTO")
	}
}
