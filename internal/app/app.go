package app

import (
	"errors"
	"github.com/Jonatna0990/test-subscription-service/internal/config"
	"github.com/gofiber/fiber/v3"
	"log/slog"
)

var a *App

type App struct {
	config *config.Config
	logger *slog.Logger
	server *fiber.App
}

// NewApp создаёт и инициализирует новый экземпляр приложения.
func NewApp(configPath string) (*App, error) {
	app := &App{}
	app.initConfig(configPath)
	app.initLogger()
	app.initHttpServer()
	app.logger.Info("App initialized")
	return app, nil
}

// SetGlobalApp устанавливает глобальный экземпляр приложения(Singleton).
func SetGlobalApp(app *App) {
	a = app
}

// GetGlobalApp возвращает глобальный экземпляр приложения.
func GetGlobalApp() (*App, error) {
	if a == nil {
		return nil, errors.New("global app is not initialized")
	}
	return a, nil
}
