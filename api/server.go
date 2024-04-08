package api

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/ediallocyf/asciiartweb/util"
)

type Server struct {
	listener string
}

type AsciiArtWeb struct {
	Text   string
	Banner string
}

const (
	shadowBanner   = "shadow"
	standardBanner = "standard"
	thinkerBanner  = "thinkertoy"
)

// WriteHTML writes an HTML response with the provided status code and content.
func WriteHTML(w http.ResponseWriter, status int) error {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(status)
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	return tmpl.Execute(w, nil)
}

// WriteHTMLResult writes an HTML response for displaying ASCII art result.
func WriteHTMLResult(w http.ResponseWriter, status int, asciiArt string) error {
	w.Header().Add("Content-type", "text/html")
	w.WriteHeader(status)
	tmpl := template.Must(template.ParseFiles("templates/ascii.html"))
	return tmpl.Execute(w, struct{ Banner string }{asciiArt})
}

func NewAPIServer(listener string) *Server {
	return &Server{
		listener: listener,
	}
}

// MainPageHandler handles requests to the main page
func (s *Server) MainPageHandler(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return nil
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return nil
	}

	return s.MainPageGetHandler(w, r)
}

// MainPageGetHandler handles GET requests to the main page.
func (s *Server) MainPageGetHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	return WriteHTML(w, http.StatusOK)
}

// assertArtWebHandler handles requests to generate ASCII art.
func (s *Server) assertArtWebHandler(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return nil
	}
	text := r.Form.Get("text")
	banner := r.Form.Get("banner")

	switch banner {
	case shadowBanner, standardBanner, thinkerBanner:
		asciiArt, err := GenerateAsciiArt(text, banner)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error generating ASCII art: %v", err), http.StatusInternalServerError)
			return nil
		}
		result := AsciiArtWeb{
			Text:   text,
			Banner: asciiArt,
		}
		return WriteHTMLResult(w, http.StatusOK, result.Banner)
	default:
		http.Error(w, "Banner not found", http.StatusNotFound)
		return nil
	}
}

// GenerateAsciiArt generates ASCII art based on the provided text and banner type.
func GenerateAsciiArt(text, bannerType string) (string, error) {
	var patternFileName string

	switch bannerType {
	case shadowBanner:
		patternFileName = "shadow.txt"
	case standardBanner:
		patternFileName = "standard.txt"
	case thinkerBanner:
		patternFileName = "thinkertoy.txt"
	default:
		return "", fmt.Errorf("unknown banner type: %s", bannerType)
	}

	chLength := 8
	patternContent, errPattern := util.ReadFileToStr(patternFileName)
	if errPattern != nil {
		return "", fmt.Errorf("error loading pattern file: %v", errPattern)
	}
	patternMap, errPatternMap := util.ContentToMap(patternContent, chLength)
	if errPatternMap != nil {
		return "", fmt.Errorf("error creating pattern map: %v", errPatternMap)
	}
	return util.StrMaker(text, patternMap, chLength), nil
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type Error struct {
	Error string
}

// makeHTTPHandlerFunc creates an HTTP handler function from an API function.
func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			err := WriteHTML(w, http.StatusBadRequest)
			if err != nil {
				return
			}
		}
	}
}

// Run starts the HTTP server.
func (s *Server) Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", makeHTTPHandlerFunc(s.MainPageHandler))
	mux.HandleFunc("/ascii-art", makeHTTPHandlerFunc(s.assertArtWebHandler))
	err := http.ListenAndServe(s.listener, mux)
	if err != nil {
		return
	}
}
