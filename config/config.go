package config

import (
	"encoding/json"
	"os"
)

type Conf struct {
	Host     string   `json:"host"`
	Database Database `json:"database"`
}

type Database struct {
	Type     string `json:"type"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	DBName   string `json:"db_name"`
}

var AppConfig Conf

func SetupEnv(fileLocation string) {
	raw, err := os.ReadFile(fileLocation)
	if err != nil {
		panic("Error occured while reading config:" + err.Error())
	}
	json.Unmarshal(raw, &AppConfig)
}
