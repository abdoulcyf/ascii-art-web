package handlers

import (
	"errors"
	"net/http"
	"text/template"
)

// assertArtWebHandler handles requests to generate ASCII art.
// func AssertArtWebHandler(w http.ResponseWriter, r *http.Request) error {
// 	err := r.ParseForm()
// 	if err != nil {
// 		http.Error(w, "Bad Request", http.StatusBadRequest)
// 		return nil
// 	}
// 	text := r.Form.Get("text")
// 	banner := r.Form.Get("banner")

// 	switch banner {
// 	case shadowBanner, standardBanner, thinkerBanner:
// 		asciiArt, err := GenerateAsciiArt(text, banner)
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Error generating ASCII art: %v", err), http.StatusInternalServerError)
// 			return nil
// 		}

// 		result := Banner{
// 			Banner: asciiArt,
// 			HomePagePath: "/",
// 		}
// 		return WriteHTMLResult(w, http.StatusOK, result.Banner, result.HomePagePath)
// 	default:
// 		http.Error(w, "Banner not found", http.StatusNotFound)
// 		return nil
// 	}
// }

// GenerateAsciiArtHandler handles requests to generate and display ASCII art.
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) error {
	//-----------------------------------------------------------
	errParseForm := r.ParseForm()
	if errParseForm != nil {
		errMsg := "----<-GenerateAsciiArtHandler------<--ParseForm----" + errParseForm.Error()
		logMsg := "Error in parsing form"
		logger.Error(logMsg + errMsg)
		return errors.New(errMsg)
	}
	logger.Info("Form parsed successfully")
	//-----------------------------------
	text := r.Form.Get("text")
	banner := r.Form.Get("banner")
	//---------------------------------
	switch banner {
	case shadowBanner, standardBanner, thinkerBanner:
		asciiArt, errGenerateAscii := GenerateAsciiArt(text, banner)
		if errGenerateAscii != nil {
			errMsg := "----<-GenerateAsciiArtHandler------<--GenerateAsciiArt----" + errGenerateAscii.Error()
			logMsg := "Error in generating ASCII banner"
			logger.Error(logMsg + errMsg)
			return errors.New(errMsg)
		}
		logger.Info("ASCII Banner generated successfully")
		// Parse HTML template
		tmpl := template.Must(template.ParseFiles(asciiTemplateAdrr))
		//prepare data for template
		data := Banner{
			Banner: asciiArt,
			Url:    homePagePath,
		}
		// Execute ASCII art template
		errAscciTemp := tmpl.Execute(w, data)
		if errAscciTemp != nil {
			errMsg := "----<-GenerateAsciiArtHandler------<--ParseFiles----" + errAscciTemp.Error()
			logMsg := "Error executing ASCII art template"
			logger.Error(logMsg + errMsg)
		} else {
			logger.Info("ASCII Art Template parsed successfully")
		}
		return nil
	default:
		ErrorPage(w, BannerNotFound)
		return nil
	}
}
