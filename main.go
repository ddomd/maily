package main

import (
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/ddomd/maily/api"
	"github.com/ddomd/maily/grpcapi"
	"github.com/ddomd/maily/internal/mdb"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load ENV file: %s\n", err.Error())
	}

	db, err := sql.Open("sqlite3", os.Getenv("DBPATH"))
	if err != nil {
		log.Fatalf("Failed to open database: %s\n", err.Error())
	}
	defer db.Close()

	mdb := mdb.NewMdb(db)

	jsonPort := os.Getenv("JSON_PORT")
	grpcPort := os.Getenv("GRPC_PORT")

	restServer := api.NewServer(jsonPort, mdb)
	grpcServer := grpcapi.NewServer(grpcPort, mdb)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		log.Printf("Starting GRPC server on port %s...\n", grpcServer.Port)
		grpcServer.Serve()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		log.Printf("Starting REST server on port %s...\n", restServer.Port)
		restServer.Serve()
		wg.Done()
	}()

	wg.Wait()
}