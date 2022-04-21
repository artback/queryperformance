package main

import (
	"context"
	"flag"
	"github.com/artback/queryperformance/pkg/api"
	"github.com/artback/queryperformance/pkg/repository/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	_ "github.com/lib/pq"
	"log"
	"time"
)

var (
	connectionString *string
)

func init() {
	connectionString = flag.String("connection-string", "user=postgres password=password dbname=name sslmode=disable", "connect to postgressql")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	app := fiber.New()
	app.Use(recover.New())
	app.Use(logger.New())
	flag.Parse()
	db, err := postgres.NewConnection(ctx, *connectionString)
	if err != nil {
		log.Fatal(err)
	}
	api.SetupRoutes(db, app)
	if err := app.Listen(":7070"); err != nil {
		log.Fatal(err)
	}
}
