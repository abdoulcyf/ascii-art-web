package handlers

import "net/http"

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
	shadowBanner   = "shadow"
	standardBanner = "standard"
	thinkerBanner  = "thinkertoy"

	contentType = "text/html"

	templatesFilesAdrress   = "../../templates/"
	homePageFileName        = "index.html"
	hopePageTemplateAdrress = templatesFilesAdrress + homePageFileName

	assciiArtFileName       = "ascii.html"
	ascciArtTemplateAddress = templatesFilesAdrress + assciiArtFileName

	homePagePath            = "/"
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
