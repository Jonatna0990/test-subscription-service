package main

import (
	"github.com/Jonatna0990/test-subscription-service/cmd/app"
	"github.com/Jonatna0990/test-subscription-service/internal/app"
	"github.com/spf13/cobra"
	"log"
)

var configFilePath string

func initApp() {
	a, err := app.NewApp(configFilePath)
	if err != nil {
		log.Fatal("Fail to create app: ", err)
	}

	app.SetGlobalApp(a)
}

// @title Auth Service API
// @version 1.0
// @description This is a sample swagger for Auth service API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {
	rootCmd := &cobra.Command{
		Use:   "main service",
		Short: "Main entry-point command for the application",
	}

	rootCmd.PersistentFlags().StringVar(&configFilePath, "config", "", "config file path")

	cobra.OnInitialize(initApp)

	rootCmd.AddCommand(
		cmd.RunHTTP(),
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("failed to execute root cmd: %v", err)

		return
	}
}
