package auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret []byte = []byte("signkey")

func JWTSign(uuid, fullName string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid":      uuid,
		"full_name": fullName,
	})
	return token.SignedString(hmacSampleSecret)
}

func JWTParse(jwtValue string) (map[string]interface{}, error) {
	token, err := jwt.Parse(jwtValue, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["uuid"], claims["full_name"])
		res := map[string]interface{}(claims)
		return res, nil
	}
	return nil, errors.New("auth: token not valid")
}
