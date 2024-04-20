package handlers

import (
	"fmt"
	"net/http"
)

// assertArtWebHandler handles requests to generate ASCII art.
func (s *Server) assertArtWebHandler(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, badRequest, http.StatusBadRequest)
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
		return WriteHTMLAscii(w, http.StatusOK, result.Banner)
	default:
		http.Error(w, BannerNotFound, http.StatusNotFound)
		return nil
	}
}
