package main

import (
	// "log/slog"
	// "os"

	"github.com/ediallocyf/asciiartweb/internal/servers"
)

// var (
// 	opts   = &slog.HandlerOptions{Level: slog.LevelDebug}
// 	logger = slog.New(slog.NewTextHandler(os.Stderr, opts))

// 	errMsg string
// 	logMsg string
// )

// const (
// 	portAddressNumber = ":8080"
// )

func main() {
	servers.LoadingServerS()
}
