// Package main is the entry point of the application.
// Author: Ron Webb
package main

import (
	"go-app-template/cmd"
	"go-app-template/logger"
)

func main() {
	// Must be close to properly close the logger.
	defer logger.Close()

	// Bootstrap the application.
	cmd.Execute()
}
