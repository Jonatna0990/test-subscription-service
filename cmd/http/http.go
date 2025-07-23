package http

import (
	"github.com/Jonatna0990/test-subscription-service/internal/app"
	"github.com/spf13/cobra"
)

var configFilePath string

func RunHTTP() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "http",
		Short: "Run http server",
	}
	cmd.PersistentFlags().StringVar(&configFilePath, "config", "", "config file path")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {

		if err := initApp(); err != nil {
			return err
		}
		a, err := app.GetGlobalApp()
		if err != nil {
			return err
		}

		if err := a.StartHTTPServer(); err != nil {
			return err
		}

		return nil
	}

	return cmd
}

func initApp() error {

	a, err := app.NewApp(configFilePath)
	if err != nil {
		return err
	}
	app.SetGlobalApp(a)
	return nil
}
