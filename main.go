package main

import (
	"fmt"
	"log"
	"satang-test/app/config"
	"satang-test/app/withGETH"
	"satang-test/pkg/database"
)

const (
	monitorHexAddress = "0x28C6c06298d514Db089934071355E5743bf21d60"
)

func main() {
	conf, err := config.ReadConfig()
	if err != nil {
		log.Panic(err)
	}
	mongodb := database.NewMongoDB(conf.Database.Url(), conf.Database.Name)
	if err := mongodb.ConnectDatabase(); err != nil {
		log.Panic(err)
	}

	fmt.Println("connect DB successfully")

	if err := mongodb.PingDatabase(); err != nil {
		log.Panic(err)
	}
	fmt.Println("ping DB successfully")
	repo := withGETH.NewRepository(mongodb.Database())
	service := withGETH.NewService(conf.InfuraWSUrl, monitorHexAddress, repo)
	service.SubNewHeads()
	defer service.UnSubNewHeads()

	service.Monitoring()
}
