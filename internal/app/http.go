package app

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3"
)

func (a *App) StartHTTPServer() error {

	fmt.Println("Starting HTTP server")
	return nil
}

func (a *App) startHTTPServer() error {
	if a.server == nil {
		err := errors.New("server is empty")
		return err
	}
}

func (a *App) initHttpServer() {
	a.server = fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       a.config.App.Name,
	})
}
