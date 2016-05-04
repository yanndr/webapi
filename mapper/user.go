package mapper

import (
	"fmt"
	"time"

	"github.com/yanndr/webapi/database"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, password string) (int, error) {

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	var lastInsertId int
	err := database.DBCon.QueryRow("INSERT INTO \"User\"(\"Username\",\"Password\",\"Created\") VALUES($1,$2,$3) returning \"Id\";", username, hashedPassword, time.Now()).Scan(&lastInsertId)

	if err != nil {
		return 0, err
	}

	fmt.Println("last inserted id =", lastInsertId)

	return lastInsertId, nil
}
