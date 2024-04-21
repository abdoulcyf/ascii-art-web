package servers

import (
	"log/slog"
	"net/http"
	"os"
)

type Server struct {
	listener string
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type Error struct {
	Error string
}

var (
	errMsg string
	logMsg string
)

var (
	logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
)

const (
	shadowBanner   = "shadow"
	standardBanner = "standard"
	thinkerBanner  = "thinkertoy"

	homePagePath            = "/"
	StatusNotFound          = "404 not found"
	getRequest              = "GET"
	methodNotAllowed        = "Method not allowed"
	badRequest              = "Bad Request"
	ErrorGeneratingAsciiArt = "Error generating ASCII art:"
	BannerNotFound          = "Banner not found"

	shadowPatternFileName     = "shadow.txt"
	standardPatternfileName   = "standard.txt"
	thinkerToyPatternFileName = "thinkertoy.txt"

	ascciArtPagePath = "/ascii-art"

	portAddressNumber = ":8080"
)
