package config

import (
	"log"

	"gopkg.in/ini.v1"
)

const (
	SERVER   = "server"
	DATABASE = "database"
)

type BankingConfig struct {
	Server struct {
		SerAddress string
		SerPort    string
	}

	Database struct {
		DbUser     string
		DbPassword string
		DbAddr     string
		DbPort     string
		DbName     string
	}
}

func ReadConfig() *BankingConfig {
	iniFile, err := ini.Load("banking.ini")
	if err != nil {
		log.Fatal("Failed to load ini file.", err)
	}
	config := &BankingConfig{}

	//Server Section
	config.Server.SerAddress = iniFile.Section(SERVER).Key("server_address").String()
	config.Server.SerPort = iniFile.Section(SERVER).Key("server_port").String()

	//Database Serction
	config.Database.DbAddr = iniFile.Section(DATABASE).Key("db_addr").String()
	config.Database.DbPort = iniFile.Section(DATABASE).Key("db_port").String()
	config.Database.DbUser = iniFile.Section(DATABASE).Key("db_user").String()
	config.Database.DbName = iniFile.Section(DATABASE).Key("db_name").String()
	config.Database.DbPassword = iniFile.Section(DATABASE).Key("db_password").String()

	return config
}
