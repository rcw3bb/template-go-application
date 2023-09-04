// Package cmd must hold all the command implementation.
// Author: Ron Webb
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-app-template/config"
	"go-app-template/config/appinfo"
	"go-app-template/logger"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   appinfo.Application,
		Short: appinfo.Application + " v" + appinfo.Version,
		Run:   rootLogic,
	}
)

func rootLogic(cmd *cobra.Command, args []string) {
	flags := cmd.Flags()
	version, _ := flags.GetBool("version")

	env := config.Environment
	appName := appinfo.Application
	if env != "" {
		appName += " (" + env + ")"
	}

	if version {
		logger.Info(appName + " v" + appinfo.Version)
		return
	}

	// Default logic.
	logger.Info(appName)

	prop1 := config.GetString("prop1")
	logger.Info("Prop1=" + prop1)

	prop2 := config.GetString("prop2")
	logger.Info("Prop2=" + prop2)

	logger.Debug("In debug mode.")
}

func init() {
	flags := rootCmd.Flags()
	flags.BoolP("version", "v", false, "Show the version")

	rootCmd.PersistentFlags().BoolVarP(&logger.IsDebugEnabled, "debug", "", false, "Show debug information")
	rootCmd.PersistentFlags().StringVarP(&config.Environment, "environment", "e",
		"", "Specify the environment")

	cobra.OnInitialize(logger.InitLogger, config.InitConfig)
}

// Execute the entry point for all the commands.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
