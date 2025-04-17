package httpServer

import "github.com/gofiber/fiber/v2"

type Server struct {
	cfg *config.Config
	fiber *fiber.App
}

func NewServer (cfg *config.Config) *Server {
	return &Server{
		fiber:fiber.New(fiber.Config{DisableStartMesage: true})
		cfg: cfg,
	}
}
