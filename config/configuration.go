package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ServerConfiguration struct {
	Port               string `json:"port"`
	PrivateKeyPath     string `json:"privateKeyPath"`
	PublicKeyPath      string `json:"publicKeyPath"`
	JWTExpirationDelta int
}

var configuration ServerConfiguration

//LoadServerConfiguration Retrun the configuration of the Server
func Init() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)

	configuration = ServerConfiguration{}

	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func Get() ServerConfiguration {
	if &configuration == nil {
		Init()
	}
	return configuration
}
