package middleware

import (
    "net/http"
    "strings"

    "booking_api/utils"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware(allowedRoles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        // validar JWT
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token requerido"})
            c.Abort()
            return
        }

        tokenString = strings.TrimPrefix(tokenString, "Bearer ")
        claims, err := utils.ParseJWT(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
            c.Abort()
            return
        }

        // Guardamos info del usuario en el contexto
        c.Set("userID", claims.UserID)
        c.Set("rol", claims.Rol)

        // validar rol
        if len(allowedRoles) > 0 {
            allowed := false
            for _, role := range allowedRoles {
                if claims.Rol == role {
                    allowed = true
                    break
                }
            }
            if !allowed {
                c.JSON(http.StatusForbidden, gin.H{"error": "No tienes permiso"})
                c.Abort()
                return
            }
        }

        c.Next()
    }
}
