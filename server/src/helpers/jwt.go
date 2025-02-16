package helpers

import (
    "time"

    "github.com/dgrijalva/jwt-go"
)


func GenerateJWT(userID uint, secretKey string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "exp":  time.Now().Add(time.Hour * 72).Unix(),
        "iat":  time.Now().Unix(),
        "sub":  userID,
    })

    tokenString, err := token.SignedString([]byte(secretKey))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}