package handlers

import (
	"log/slog"
	"os"
)

// type Server struct {
// 	listener string
// }

type AsciiArt struct {
	//Text   string
	Banner string
	Url    string
	StyleAdrr string
}

const (
	domain       = "localhost:"
	portNumber   = "8080"
	homePathUrl  = domain + portNumber
	homePagePath = "/"
	actionPath = "/ascii-art"

	shadowBanner   = "shadow"
	standardBanner = "standard"
	thinkerBanner  = "thinkertoy"
	fileExt        = ".txt"

	assetsFileDir = "../../assets/"

	shadowPatterFileAdrr     = assetsFileDir + shadowBanner + fileExt
	standardPatterFileAdrr   = assetsFileDir + standardBanner + fileExt
	thinkertoyPatterFileAdrr = assetsFileDir + thinkerBanner + fileExt

	templateFilesDir = "../../templates/"

	homeTemplateName = "index.html"
	homeTemplateAdrr = templateFilesDir + homeTemplateName

	asciiTemplateName = "ascii.html"
	asciiTemplateAdrr = templateFilesDir + asciiTemplateName

	errorTemplateName = "error.html"
	errorTemplateAdrr = templateFilesDir + errorTemplateName

	BannerNotFound = "Banner Not Found"
	NotFound = "404 - Page Not Found"

	staticFilesDir = "../../static/styles/"

	ascciStyleName = "ascii.css"
	asciiStyleAdrr = staticFilesDir + ascciStyleName
)

var (
	logger = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
)
var (
	errMsg string
	logMsg string
)
