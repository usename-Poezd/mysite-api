package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/usename-Poezd/mysite-api/internal/config"
	"github.com/usename-Poezd/mysite-api/internal/delivery/http"
	"github.com/usename-Poezd/mysite-api/internal/repository"
	"github.com/usename-Poezd/mysite-api/internal/services"
	"log"
)

// @title My site API
// @version 1.0
// @description REST API for my site

// @host localhost:8080
// @BasePath /api/

// @securityDefinitions.apikey UsersAuth
// @in header
// @name Authorization

// Run initializes whole application.
func Run() {
	app := fiber.New()

	conf, err := config.GetConfig(".")
	if err != nil {
		log.Fatalln(err)
	}

	// this Pings the database trying to connect
	// use sqlx.Open() for sql.Open() semantics
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", conf.PostgresHost, conf.PostgresUser, conf.PostgresPassword, conf.PostgresDB))
	if err != nil {
		log.Fatalln(err)
	}

	repositories := repository.NewRepositories(db)
	serv := services.NewServices(repositories)

	handler := http.NewHandler(app, serv)
	handler.Init()

	if err := app.Listen(":8080"); err != nil {
		log.Fatalln(err)
	}
}
