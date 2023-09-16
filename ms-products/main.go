package main

import (
	"database/sql"
	"log"

	"github.com/djudju12/order-system/ms-products/configs"
	"github.com/djudju12/order-system/ms-products/controller"
	"github.com/djudju12/order-system/ms-products/service"
	_ "github.com/lib/pq"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot read configurations:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot open db connection:", err)
	}

	service := service.NewProcutService(conn)
	server := controller.NewServer(service)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}