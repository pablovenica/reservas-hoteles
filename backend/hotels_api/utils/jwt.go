package utils

import (
    "os"

    "github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
    UserID int    `json:"user_id"`
    Rol    string `json:"rol"`
    jwt.RegisteredClaims
}

func ParseJWT(tokenString string) (*JWTClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if err != nil {
        return nil, err
    }

    claims, ok := token.Claims.(*JWTClaims)
    if !ok || !token.Valid {
        return nil, err
    }

    return claims, nil
}
