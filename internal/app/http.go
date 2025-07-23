package app

import (
	"errors"
	"fmt"
	"github.com/Flussen/swagger-fiber-v3"
	handler "github.com/Jonatna0990/test-subscription-service/internal/http/handler/subscriptions"
	"github.com/Jonatna0990/test-subscription-service/internal/repository/subscriptions"
	"github.com/Jonatna0990/test-subscription-service/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"log"
)

type structValidator struct {
	validate *validator.Validate
}

// Validate method implementation
func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

// StartHTTPServer — публичный метод для запуска HTTP-сервера
func (a *App) StartHTTPServer() error {
	fmt.Println("Starting HTTP server")
	go func() {
		err := a.startHTTPServer()
		if err != nil {
			log.Fatal(err)
		}
	}()
	return nil
}

// startHTTPServer — внутренняя проверка и запуск сервера
func (a *App) startHTTPServer() error {
	if a.server == nil {
		err := errors.New("server is empty")
		return err
	}

	repo := subscriptions.NewRepository(a.postgres.Pool, a.logger)
	uc := usecase.New(repo)
	h := handler.New(uc)

	addr := fmt.Sprintf("%s:%d", a.config.HTTPServer.Host, a.config.HTTPServer.Port)
	a.server.Use(cors.New())
	a.server.Use(logger.New())
	h.RegisterRoutes(a.server)
	a.server.Get("/swagger/*", swagger.HandlerDefault)
	return a.server.Listen(addr)
}

// initHttpServer — инициализирует Fiber-сервер с настройками
func (a *App) initHttpServer() {
	a.server = fiber.New(fiber.Config{
		// TODO вынести в параметры конфигурации
		CaseSensitive:   true,
		StrictRouting:   false,
		AppName:         a.config.App.Name,
		StructValidator: &structValidator{validate: validator.New()},
	})

	a.logger.Info("Http server initialized")
}
