package controller

import (
	"user_api/services"
	"user_api/utils"
	"net/http"
	"user_api/dto"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	user, err := services.Login(body.Email, body.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
		print(body.Email)
		print(body.Password)
	}
token, err := utils.GenerarJWT(user.ID, user.TipoUsuario)
if err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error generando token"})
    return
}
	ctx.JSON(http.StatusOK, gin.H{
		"token":  token,
		"userID": user.ID,
		"rol":    user.TipoUsuario,
	})
}
func CrearUsuario(ctx *gin.Context) {
    var body struct {
        Nombre      string `json:"nombre"`
        Email       string `json:"email"`
        Password    string `json:"password"`
        TipoUsuario string `json:"tipo_usuario"`
    }

    if err := ctx.ShouldBindJSON(&body); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    user, err := services.CrearUsuario(dto.UserDto{
        Nombre:      body.Nombre,
        Email:       body.Email,
        Password:    body.Password,
        TipoUsuario: body.TipoUsuario,
    })
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"user": user})
}

