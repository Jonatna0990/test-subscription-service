package app

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

// initConfig загружает конфигурацию приложения из файла, путь к которому
// задаётся через переменную окружения CONFIG_PATH.
func (a *App) initConfig() error {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatalf("CONFIG_PATH is not set")
		return errors.New("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
		return err
	}

	if err := cleanenv.ReadConfig(configPath, &a.config); err != nil {
		log.Fatalf("cannot read config: %s", err)

		return err
	}

	return nil
}
