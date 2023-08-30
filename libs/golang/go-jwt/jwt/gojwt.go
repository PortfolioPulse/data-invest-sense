package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
     "github.com/go-chi/jwtauth"
)

// GenerateJWT generates a JWT token from the provided JSON schema using the TokenAuth instance.
func GenerateSchemaJWT(tokenAuth *jwtauth.JWTAuth, jsonSchema map[string]interface{}) (string, error) {
	// Check if JsonSchema is not empty
	if len(jsonSchema) == 0 {
		return "", errors.New("JsonSchema is empty")
	}

	// Create a new JWT token using jwtauth.Encode
	_, tokenString, err := tokenAuth.Encode(jwt.MapClaims{
		"iat":         time.Now().Unix(),
		"json_schema": jsonSchema,
	})
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
