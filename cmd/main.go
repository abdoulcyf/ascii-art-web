package main

import (
	"fmt"
	"html/template"
	cli "main/pkg/cli"
	arr "main/pkg/plot/arr"
	draw "main/pkg/plot/draw"
	"net/http"
)

// Data structure to hold the ASCII art details
type AsciiArtWeb struct {
	Text   string
	Banner string
}

// Function variable to hold ASCII art generator
var generateAsciiArt func(text string) string

func main() {

	directory := "."
	patternFileName := "standard.txt"
	//patternFileName := "thinkertoy.txt"
	chLength := 8
	var firstCh byte = ' '
	var lastCh byte = '~'
	//----------------------------------------
	patternContent, errPattern := cli.ReadFileToStr(directory, patternFileName)

	if errPattern != nil {
		fmt.Println("Error : loading pattern file failed")
		return
	}
	//--------------------------------
	patternMap, errPatternMap := arr.ContentToMap(patternContent, chLength, firstCh, lastCh)

	if errPatternMap != nil {
		fmt.Println(errPatternMap)
		return
	}
	//--------------------------------------------
	cliStr, errClistr := cli.ReadCli()

	if errClistr != nil {
		//fmt.Println(errClistr)
		return
	}
	//----------------------------------------------

	finalStr := draw.StrMaker(cliStr, patternMap, chLength)
	fmt.Println(finalStr)

	//Define the ASCII art generator function
	generateAsciiArt = func(text string) string {
		return draw.StrMaker(text, patternMap, chLength)
	}
	//-----------------------------------------

	//Define HTTP routes
	http.HandleFunc("/", MainPageHandler)
	http.HandleFunc("/ascii-art", AsciiArtHandler)

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}

const (
	shadowBanner     = "shadow"
	standardBanner   = "standard"
	thinkertoyBanner = "thinkertoy"
)

// Render the main page using a template
func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("main").Parse(`
        <!DOCTYPE html>
        <html>
        <head>
            <title>ASCII Art Web</title>
            <style>
                /* Your CSS styles here */
            </style>
        </head>
        <body>
            <h1>ASCII Art Web</h1>
            <form method="post" action="/ascii-art">
                <label for="text">Text:</label><br>
                <input type="text" id="text" name="text"><br>
                <label for="banner">Banner:</label><br>
                <select id="banner" name="banner">
                    <option value="shadow">Shadow</option>
                    <option value="standard">Standard</option>
                    <option value="thinkertoy">Thinkertoy</option>
                </select><br><br>
                <button type="submit">Generate ASCII Art</button>
            </form>
        </body>
        </html>
    `))
	tmpl.Execute(w, nil)
}

// Handle the POST request to generate ASCII art
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	r.ParseForm()
	text := r.Form.Get("text")
	banner := r.Form.Get("banner")

	// Generate ASCII art based on the selected banner
	var asciiArt string
	switch banner {
	// case shadowBanner:
	// 	asciiArt = generateShadowAsciiArt(text)
	case standardBanner:
		asciiArt = generateAsciiArt(text)
	// case thinkertoyBanner:
	// 	asciiArt = generateThinkertoyAsciiArt(text)
	default:
		// Use standard banner by default
		asciiArt = generateAsciiArt(text)
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
                /* Your CSS styles here */
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

// Function to generate ASCII art using the shadow banner
// func generateShadowAsciiArt(text string) string {
// 	// Implement your ASCII art generation logic for the shadow banner here
// }

// // Function to generate ASCII art using the thinkertoy banner
// func generateThinkertoyAsciiArt(text string) string {
// 	// Implement your ASCII art generation logic for the thinkertoy banner here
// }
