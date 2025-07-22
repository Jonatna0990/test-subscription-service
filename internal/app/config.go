package app

import (
	"github.com/Jonatna0990/test-subscription-service/internal/config"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

// initConfig загружает конфигурацию приложения из файла, путь к которому
// задаётся через переменную окружения configPath.
func (a *App) initConfig(configPath string) {
	if configPath == "" {
		log.Fatal("config is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	a.config = &config.Config{}

	if err := cleanenv.ReadConfig(configPath, a.config); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}
}
