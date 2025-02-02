package main

import (
	"github/Alfian57/student-hub-api/internal/db"
	"github/Alfian57/student-hub-api/internal/store"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := db.New(os.Getenv("DB_ADDR"))
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	log.Println("DB Connected")

	store := store.NewStorage(conn)
	db.Seed(store)
}
