package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/trenchesdeveloper/mcp-server-store/configs"
)

func main() {
	// Load configuration
	cfg := configs.LoadConfig()

	// Configure logging
	logger := logrus.New()
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		level = logrus.DebugLevel
	}
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	// Log to stderr so stdout stays clean for JSON-RPC
	logger.SetOutput(os.Stderr)

	logger.Info("Starting MCP Server...")

}
