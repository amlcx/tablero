package wiring

import (
	"os"

	"github.com/charmbracelet/log"
)

type Dependencies struct {
	Logger *log.Logger
}

func InitDependencies() *Dependencies {
	logger := log.NewWithOptions(os.Stdout, log.Options{
		Level:           log.DebugLevel,
		ReportTimestamp: true,
		ReportCaller:    true,
	})

	return &Dependencies{
		Logger: logger,
	}
}
