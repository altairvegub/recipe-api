package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	UserId string `json:"user id"`
	jwt.StandardClaims
}

func CreateJwtToken(userID string) (string, error) {
	os.Setenv("RECIPE_SECRET_KEY", "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160")

	// Create the Claims
	claims := CustomClaims{
		"test 123",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    "test",
			IssuedAt:	 time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("RECIPE_SECRET_KEY")))
	if err != nil {
		return "", fmt.Errorf("failed to create token: %w", err)
	}

	fmt.Println("jwt created: " + signedToken, err)

	return signedToken, err
}
