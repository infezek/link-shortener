package security

import (
	"errors"
	"fmt"
	"net/http"
	"shortener/src/config"
	"strings"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ValidateToken(r *http.Request) error {
	tokenString := ExtractToken(r)
	token, erro := jwt.Parse(tokenString, CheckTokenKey)

	if erro != nil {
		return erro
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("token inválido")
}

func CheckTokenKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado %v", token.Header["alg"])
	}
	return []byte(config.ProjectSettings().SecretKey), nil
}

func ExtractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

type User struct {
	Sub     string `json:"sub"`
	Payload string `json:"payload"`
	jwt.StandardClaims
}

func DecodeToken(r *http.Request) User {
	user := User{}

	tokenString := ExtractToken(r)

	_, err := jwt.ParseWithClaims(tokenString, &user, CheckTokenKey)
	if err != nil {
		fmt.Println("err", err)
		return User{}
	}

	return user
}
