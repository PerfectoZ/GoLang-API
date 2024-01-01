package main

import (
	"database/sql"
	"github.com/PerfectoZ/GoLang-API/api"
	db "github.com/PerfectoZ/GoLang-API/db/sqlc"
	"github.com/PerfectoZ/GoLang-API/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to Database", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("Cannot start server:", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("Cannot start server:", err)
	}
}
