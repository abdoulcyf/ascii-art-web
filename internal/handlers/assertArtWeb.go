package handlers

import (
	"errors"
	"net/http"
)

// assertArtWebHandler handles requests to generate ASCII art.
func AssertArtWebHandler(w http.ResponseWriter, r *http.Request) error {
	errParseForm := r.ParseForm()
	if errParseForm != nil {
		errMsg = "----<-assertArtWebHandler------<--ParseForm----" + errParseForm.Error()
		logMsg = "Error in parsing form"
		logger.Error(logMsg + errMsg)
		return errors.New(errMsg)
	} else {
		logger.Info("Form parsed succesfully")
	}
	text := r.Form.Get("text")
	banner := r.Form.Get("banner")

	switch banner {
	case shadowBanner, standardBanner, thinkerBanner:
		asciiArt, errGenerateAscii := GenerateAsciiArt(text, banner)
		if errGenerateAscii != nil {
			errMsg = "----<-assertArtWebHandler------<--GenerateAsciiArt----" + errGenerateAscii.Error()
			logMsg = "Error in generating ascii banner"
			logger.Error(logMsg + errMsg)
			return errors.New(errMsg)
		} else {
			logger.Info("Ascci Banner generated succussfully")
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
