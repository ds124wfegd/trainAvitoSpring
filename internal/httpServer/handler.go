package httpServer

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (server *Server) MapHandlers(app *fiber.App) error {
	db, err := storagePostgres.InitPsqlDB(s.cfg)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err := storagePostgres.CreateTable(db)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return nil
}
