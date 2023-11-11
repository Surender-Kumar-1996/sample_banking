package main

import (
	"github.com/Surender-Kumar-1996/sample_banking/app"
	"github.com/Surender-Kumar-1996/sample_banking/config"
	logger "github.com/Surender-Kumar-1996/sample_banking/logger"
)

func main() {
	conf := config.ReadConfig()
	logger.Info("Starting the application....")
	app.Start(conf)

}
