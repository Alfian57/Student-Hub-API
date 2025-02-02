package main

import (
	"github/Alfian57/student-hub-api/internal/db"
	"github/Alfian57/student-hub-api/internal/store"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	// Logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugarLogger := logger.Sugar()

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		sugarLogger.Fatalw("failed to load .env file", "error", err.Error())
	}

	// Config
	cfg := config{
		appUrl: os.Getenv("APP_URL"),
		addr:   ":8080",
		env:    os.Getenv("APP_ENV"),
		db: dbConfig{
			addr: os.Getenv("DB_ADDR"),
		},
	}

	// Connect to DB
	db, err := db.New(cfg.db.addr)
	if err != nil {
		sugarLogger.Errorw("failed to connect to db", "error", err.Error())
	}
	defer db.Close()
	sugarLogger.Info("DB Connected")

	// Create new store
	store := store.NewStorage(db)
	app := &application{
		config: cfg,
		store:  store,
		logger: sugarLogger,
	}

	r := app.mount()
	err = app.run(r)
	if err != nil {
		sugarLogger.Errorw("failed to run application", "error", err.Error())
	}
}
