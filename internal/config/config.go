package config

import "github.com/Jonatna0990/test-subscription-service/pkg/postgres"

// Config — основная структура конфигурации приложения
type Config struct {
	App        App             `yaml:"app"`
	HTTPServer HTTPServer      `yaml:"http_server"`
	Postgres   postgres.Config `yaml:"postgres"`
}

// App — конфигурация приложения
type App struct {
	Env  string `yaml:"env"`
	Name string `yaml:"name"`
}

// HTTPServer — конфигурация HTTP-сервера
type HTTPServer struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
