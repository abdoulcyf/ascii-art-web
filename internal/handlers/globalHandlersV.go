package handlers

import (
	"log/slog"
	"os"
)

// type Server struct {
// 	listener string
// }

type Banner struct {
	//Text   string
	Banner string
	Url    string
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
