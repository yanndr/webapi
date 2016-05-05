package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yanndr/webapi/mapper"
	"github.com/yanndr/webapi/model"
	"golang.org/x/crypto/bcrypt"
	//jwt "github.com/dgrijalva/jwt-go"
)

type TokenAuthentication struct {
	Token string `json:"token" form:"token"`
}

func Login(requestUser *model.User) (int, []byte) {
	authBackend := NewJWTAuthenticationBackend()

	if authenticate(requestUser) {
		token, err := authBackend.GenerateToken(requestUser.UUID)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			response, _ := json.Marshal(TokenAuthentication{token})
			return http.StatusOK, response
		}
	}

	return http.StatusUnauthorized, []byte("")
}

func authenticate(requestedUser *model.User) bool {

	user, err := mapper.GetUser(requestedUser.Username)

	if err != nil {
		log.Printf("Error during authentication: %v \n", err.Error())
		return false
	}

	if user == nil {
		log.Printf("No result from the db auth rejected\n")
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestedUser.Password)) == nil
}
