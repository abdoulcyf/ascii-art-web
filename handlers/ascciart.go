package handlers

import (
	"net/http"
	"text/template"
)

// Data structure to hold the ASCII art details
type AsciiArtWeb struct {
	Text   string
	Banner string
}

const (
	shadowBanner     = "shadow"
	standardBanner   = "standard"
	thinkertoyBanner = "thinkertoy"
)

// Handle the POST request to generate ASCII art
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	r.ParseForm()
	text := r.Form.Get("text")
	banner := r.Form.Get("banner")

	// Generate ASCII art based on the selected banner
	var asciiArt string
	switch banner {
	case shadowBanner:
		asciiArt = GenerateShadowAsciiArt(text)
	case standardBanner:
		asciiArt = GenerateAsciiArt(text)
	case thinkertoyBanner:
		asciiArt = GenerateThinkertoyAsciiArt(text)
	default:
		// Use standard banner by default
		asciiArt = GenerateAsciiArt(text)
	}

	// Render the result using a template
	result := AsciiArtWeb{
		Text:   text,
		Banner: asciiArt,
	}

	tmpl := template.Must(template.New("result").Parse(`
        <!DOCTYPE html>
        <html>
        <head>
            <title>ASCII Art Result</title>
            <style>
              body{
				text-align: center;
				background-color:#4CAF50;
				color: white;
			  }
            </style>
        </head>
        <body>
            <h1>ASCII Art Result</h1>
            <pre>{{.Text}}</pre>
            <pre>{{.Banner}}</pre>
        </body>
        </html>
    `))
	tmpl.Execute(w, result)
}
