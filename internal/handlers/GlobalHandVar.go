package handlers

import (
	"log/slog"
	"net/http"
	"os"
)

var _ os.DirEntry    // for debugging, delete when done
var _ slog.LogValuer // for debugging, delete when done

type Server struct {
	listener string
}

type AsciiArtWeb struct {
	Text   string
	Banner string
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type Error struct {
	Error string
}

const (
	chLength = 8
	
	shadowBanner   = "shadow"
	standardBanner = "standard"
	thinkerBanner  = "thinkertoy"

	contentType = "text/html"

	portAddressNumber = ":8080"

	templatesFilesAdrress   = "../../templates/"
	homePageFileName        = "index.html"
	hopePageTemplateAdrress = templatesFilesAdrress + homePageFileName

	assciiArtFileName       = "ascii.html"
	ascciArtTemplateAddress = templatesFilesAdrress + assciiArtFileName

	homePagePath            = "/"
	ascciArtPagePath        = "/ascii-art"
	StatusNotFound          = "404 not found"
	getRequest              = "GET"
	methodNotAllowed        = "Method not allowed"
	badRequest              = "Bad Request"
	ErrorGeneratingAsciiArt = "Error generating ASCII art:"
	BannerNotFound          = "Banner not found"

	assetsFilesAddress = "../../assets/"

	shadowPatternFileName    = "shadow.txt"
	shadowPatternFileAddress = assetsFilesAddress + shadowPatternFileName

	standardPatternfileName    = "standard.txt"
	standardPatternfileAdrress = assetsFilesAddress + standardPatternfileName

	thinkerToyPatternFileName    = "thinkertoy.txt"
	thinkerToyPatternFileAddress = assetsFilesAddress + thinkerToyPatternFileName
)

// var (

// 	logger = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
// 		Level: slog.LevelDebug,
// 	})
// )
