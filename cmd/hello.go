package cmd

import (
	"github.com/spf13/cobra"
	"go-app-template/config"
	"go-app-template/logger"
)

var (
	objHello struct {
		name string
	}

	helloCmd = &cobra.Command{
		Use:   "hello",
		Short: "A hello to a name.",
		Run:   helloLogic,
	}
)

func helloLogic(cmd *cobra.Command, args []string) {
	logger.Info("Environment " + config.Environment)
	logger.Info("Hello " + objHello.name)
}

func init() {
	rootCmd.AddCommand(helloCmd)
	flags := helloCmd.Flags()
	flags.StringVarP(&objHello.name, "name", "",
		"", "Specify the name")
}