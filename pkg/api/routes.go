package api

import (
	"database/sql"
	"github.com/artback/queryperformance/pkg/api/handler/perfhandler"
	"github.com/artback/queryperformance/pkg/repository/postgres"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(db *sql.DB, app *fiber.App) {
	perf := perfhandler.PerfHandler{Repository: postgres.NewDbPerfRepository(db)}
	app.Get("/performance", perf.GetQueryPerformance)
}
