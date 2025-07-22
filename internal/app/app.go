package app

import (
	"errors"
	"fmt"
	"github.com/Jonatna0990/test-subscription-service/internal/config"
	"log/slog"
)

var a *App

type App struct {
	config *config.App
	logger *slog.Logger
}

func NewApp(configPath string) (*App, error) {
	fmt.Printf("TESTESTETSET")

	app := &App{}

	app.initConfig()
	return app, nil
}

func SetGlobalApp(app *App) {
	a = app
}

func GetGlobalApp() (*App, error) {
	if a == nil {
		return nil, errors.New("global app is not initialized")
	}
	return a, nil
}
