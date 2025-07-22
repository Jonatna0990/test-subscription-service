package config

// Config — основная структура конфигурации приложения
type Config struct {
	App        App        `yaml:"app"`
	HTTPServer HTTPServer `yaml:"http_server"`
	Postgres   Postgres   `yaml:"postgres"`
}

// App — конфигурация приложения
type App struct {
	Env            string `yaml:"env"`
	MigrationsPath string `yaml:"migrationsPath"`
}

// HTTPServer — конфигурация HTTP-сервера
type HTTPServer struct {
	Address string `yaml:"address"`
}

// Postgres — конфигурация подключения к PostgreSQL
type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"db_name"`
	SSLMode  string `yaml:"ssl_mode"`
}
