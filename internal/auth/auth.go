package auth

import (
	"GoDriver/internal/users"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = "asidouhbaspifdhauiosdfnasdpfiondfunfidsjnfdipujnsio"

type Claims struct {
	UserId   int64  `json: "user_id"`
	Username string `json: "username"`
	jwt.RegisteredClaims
}

type Credentials struct {
	Username string `json: "username"`
	Password string `json: "password"`
}

func createToken(user *users.User) (string, error) {

	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		UserId:   user.ID,
		Username: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(jwtSecret)

}

func Auth(rw http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	u, err := users.Autenticate(creds.Username, creds.Password)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusUnauthorized)
		return
	}

	token, err := createToken(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Write([]byte(token))
	//decode do payload
	//validar que o usuario existe
	//gerar o token
	//retornar o token
}
