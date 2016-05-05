package mapper

import (
	"time"

	"github.com/yanndr/webapi/database"
	"github.com/yanndr/webapi/model"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, password string) (int, error) {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	var lastInsertId int
	err := database.DBCon.QueryRow("INSERT INTO \"User\"(\"Username\",\"Password\",\"Created\") VALUES($1,$2,$3) returning \"Id\";", username, hashedPassword, time.Now()).Scan(&lastInsertId)

	if err != nil {
		return 0, err
	}

	return lastInsertId, nil
}

func GetUser(username string) (*model.User, error) {

	var user = model.User{}
	err := database.DBCon.QueryRow("Select \"Id\", \"Username\",\"Password\" From \"User\" Where \"Username\"=$1;", username).Scan(&user.UUID, &user.Username, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
