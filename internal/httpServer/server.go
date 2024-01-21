package httpserver

import (
	"context"
	"english_bot_admin/database"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

func NewServer() *Server {
	app := fiber.New()
	return &Server{app: app}
}

func (s *Server) Run() error {
	var (
		err error
	)

	dbURI := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")

	ctx := context.Background()
	db, err := database.NewConnection(ctx, dbName, dbURI)
	if err != nil {
		log.Fatal(err)
	}

	err = MapHandlers(db, s)
	if err != nil {
		return err
	}

	err = s.app.Listen(":3000")
	if err != nil {
		return err
	}
	return nil
}
