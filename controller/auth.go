package controller

import (
	"encoding/json"
	"net/http"
	"database/sql"
	"fmt"
	"time"

	"github.com/yanndr/webapi/model"
	"github.com/yanndr/webapi/service"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const (
    DB_USER     = "postgres"
    DB_PASSWORD = "Antibes06"
    DB_NAME     = "Ft"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(model.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	responseStatus, token := service.Login(requestUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func Register(w http.ResponseWriter, r *http.Request){
	
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
    db, err := sql.Open("postgres", dbinfo)
    checkErr(err)
    defer db.Close()
	
	
	registerRequest := model.Register{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&registerRequest)
	
	
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), 10)

	
	var lastInsertId int
    err = db.QueryRow("INSERT INTO \"User\"(\"Username\",\"Password\",\"Created\") VALUES($1,$2,$3) returning \"Id\";", registerRequest.Username, hashedPassword, time.Now()).Scan(&lastInsertId)
    checkErr(err)
    fmt.Println("last inserted id =", lastInsertId)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
