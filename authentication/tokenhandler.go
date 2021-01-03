package authentication

import (
	"fmt"
	"gqlGenTutorial/settings"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//GenerateToken : genera un token jwt y asigna un usuario a su claim
func GenerateToken(username string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	fmt.Println(token)

	//Creamos un diccionario para almacenar nuestro claim
	claims := token.Claims.(jwt.MapClaims)

	//Establecemos el token al claim
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	tokenString, err := token.SignedString([]byte(settings.SecretKey))

	if err != nil {
		log.Fatal("Error al generar el key")
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(settings.SecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	}
	return "", err
}
