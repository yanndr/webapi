package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ServerConfiguration struct {
	Port string `json:"port"`
}

//LoadServerConfiguration Retrun the configuration of the Server
func LoadServerConfiguration() (*ServerConfiguration, error) {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := ServerConfiguration{}

	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	return &configuration, nil
}
