package utils

import (
    "os"
    "time"

    "github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
    UserID int    `json:"user_id"`
    Rol    string `json:"rol"`
    jwt.RegisteredClaims
}

// GenerarJWT crea un token con expiración (20.000 segundos)
func GenerarJWT(userID int, rol string) (string, error) {

    expiracion := time.Now().Add(20000 * time.Second)

    claims := &Claims{
        UserID: userID,
        Rol:    rol,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expiracion),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

// ValidarJWT valida el token
func ValidarJWT(tokenStr string) (*Claims, error) {

    token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*Claims)
    if !ok || !token.Valid {
        return nil, err
    }

    return claims, nil
}
