package main

import (
	"github.com/Jonatna0990/test-subscription-service/cmd/app"
	_ "github.com/Jonatna0990/test-subscription-service/docs"
	"github.com/Jonatna0990/test-subscription-service/internal/app"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var configFilePath string

func initApp() {
	a, err := app.NewApp(configFilePath)
	if err != nil {
		log.Fatal("Fail to create app: ", err)
	}

	app.SetGlobalApp(a)
}

// @title Subscription Service API
// @version 1.0
// @description This is a sample swagger for subscription service API
// @host localhost:3001
// @BasePath /
// @schemes         http
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

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop
}
