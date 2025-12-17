package main

import (
	"database/sql"
	"log"
	"time"

	"age-calculator/config"
	"age-calculator/internal/logger"
	"age-calculator/internal/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.Load()

	if cfg.DBURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := sql.Open("mysql", cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	// wait for MySQL to be ready
	for i := 1; i <= 10; i++ {
		if err := db.Ping(); err == nil {
			break
		}
		log.Printf("waiting for database... (%d/10)", i)
		time.Sleep(2 * time.Second)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("database connection failed:", err)
	}

	logg := logger.New()
	defer logg.Sync()

	app := fiber.New()

	// global middleware
	app.Use(logger.RequestID())
	app.Use(logger.RequestLogger(logg))

	// routes
	routes.Register(app, db)

	logg.Info("server started on :3000")
	log.Fatal(app.Listen(":3000"))
}
