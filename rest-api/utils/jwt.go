package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const jwtSecretKey = "your-secret-key"

func GenerateJWTToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
		"iat":    time.Now().Unix(),
	})

	return token.SignedString([]byte(jwtSecretKey))
}

func ValidateJWTToken(tokenString string) (int64, error) {

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(jwtSecretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse token.")
	}

	if !parsedToken.Valid {
		return 0, errors.New("Invalid token.")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Could not parse claims.")
	}

	/* email:= claims["email"].(string) */
	//userId := claims["userId"].(float64)

	fmt.Println("FLOAT ?:", claims["userId"].(float64))
	fmt.Println(claims["userId"])
	/* 	if email == "" || userId == 0 {
		return 0, errors.New("Invalid token claims.")
	} */
	return int64(claims["userId"].(float64)), nil
}
