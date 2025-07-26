package logger

import (
	"os"
	"time"

	"github.com/flash-go/flash/logger"
	"github.com/rs/zerolog"
)

// Create logger service
func NewConsole() logger.Logger {
	// Define console logger settings
	settings := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	// Create console logger
	consoleLogger := logger.NewConsole(settings)

	// Return logger service
	return logger.New(consoleLogger)
}
