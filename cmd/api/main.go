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

	cfg := config{
		appUrl: os.Getenv("APP_URL"),
		addr:   ":8080",
		env:    os.Getenv("APP_ENV"),
		db: dbConfig{
			addr: os.Getenv("DB_ADDR"),
		},
	}

	db, err := db.New(cfg.db.addr)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Println("DB Connected")

	store := store.NewStorage(db)
	app := &application{
		config: cfg,
		store:  store,
	}

	r := app.mount()
	err = app.run(r)
	if err != nil {
		log.Fatal(err)
	}
}
