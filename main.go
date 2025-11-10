package main

import (
	"database/sql"
	"log"
	"mySimpleBank/api"
	db "mySimpleBank/db/sqlc"
	"mySimpleBank/util"

	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfing(".")
	if err != nil {
		log.Fatal("Connot read config", err)
	}
	dbDriver, dbSource, serverAddress := config.DBDriver, config.DBSource, config.ServerAddress
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to DB", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start the server", err)
	}
}
