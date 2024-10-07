package main

import (
	"log"

	"github.com/Edu4rdoNeves/EasyStrock/internal/api/server"
	"github.com/Edu4rdoNeves/EasyStrock/internal/infrastructure/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../internal/api/config/.env")
	if err != nil {
		log.Fatal("Fail to load .env file")
	}

	database.Start()

	server.Run()
}
