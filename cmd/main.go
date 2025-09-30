package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/metgag/final-assignment/internals/configs"
	"github.com/metgag/final-assignment/internals/routers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("error load .env file: %s", err.Error())
		return
	}

	dbpool, err := configs.InitDB()
	if err != nil {
		log.Printf("unable to create connection pool: %s", err.Error())
		return
	}
	defer dbpool.Close()

	if err := configs.PingDB(dbpool); err != nil {
		log.Printf("pg unable to connect: %s", err.Error())
		return
	}
	log.Printf(
		"connected to db \"%s\" as user \"%s\"",
		os.Getenv("PG_DB"),
		os.Getenv("PG_USER"),
	)

	router := routers.InitRouter(dbpool)
	router.Run(":8090")
}
