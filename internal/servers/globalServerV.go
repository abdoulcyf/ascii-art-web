package servers

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type MyError struct {
	code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.Message)
}

type server struct {
	listener string
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type Error struct {
	Error string
}

const (
	homePagePath = "/"
	asciiArtPage = "/ascii-art"

	//portNumber = ":8080"

	baseStaticDir      = "/static/"
	staticFilesAddress = "static"
)

var (
	opts   = &slog.HandlerOptions{Level: slog.LevelDebug}
	logger = slog.New(slog.NewTextHandler(os.Stderr, opts))

	errMsg string
	logMsg string
)

// LoggerConfig represents the configuration options for the logger
type LoggerConfig struct {
	Level slog.Level // Log level
}

// NewLogger creates a new logger with the provided configuration
func NewLogger(config LoggerConfig) *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: config.Level,
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, opts))
	slog.SetDefault(logger)
	return logger
}
