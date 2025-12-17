package routes

import (
	"database/sql"

	"age-calculator/internal/handler"
	"age-calculator/internal/repository"
	"age-calculator/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App, db *sql.DB) {
	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	app.Post("/users", h.Create)
	app.Get("/users/:id", h.GetByID)
	app.Get("/users", h.GetAll)
	app.Put("/users/:id", h.Update)
	app.Delete("/users/:id", h.Delete)
}
