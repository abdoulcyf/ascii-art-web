package api

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/ediallocyf/asciiartweb/pkg/plot/draw"
	"github.com/ediallocyf/asciiartweb/util"
)

type APIServer struct {
	listenAdrr string
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

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteHTML(w, http.StatusBadRequest, err.Error())
		}
	}
}

func WriteHTML(w http.ResponseWriter, status int, htmlContent string) error {
	w.Header().Add("Content-type", "text/html")
	w.WriteHeader(status)
	tmpl := template.Must(template.New("main").Parse(htmlContent))
	return tmpl.Execute(w, nil)
}

func WriteHTMLResult(w http.ResponseWriter, status int, asciiArt string) error {
	htmlContent := `<!DOCTYPE html>
	<html>
	<head>
		<title>ASCII Art Result</title>
		<style>
			body {
				text-align: center;
				background-color: #4CAF50;
				color: white;
				font-family: monospace;
				margin-top: 30px;
			}
	
			@media (max-width: 600px) {
				body {
					margin-top: 10px;
				}
			}
	
			pre {
				white-space: pre-wrap;
				word-wrap: break-word;
			}
	
			@media (max-width: 400px) {
				/* Apply styles for screens with a maximum width of 400px */
				pre {
					font-size: 14px;
				}
			}
		</style>
	</head>
	<body>
		<h1>ASCII Art Result</h1>
		<pre id="banner">{{.Banner}}</pre>
		<script>
		const banner = document.getElementById("banner");
		console.log("BANNER:", banner);
		</script>
	</body>
	</html>
	`

	w.Header().Add("Content-type", "text/html")
	w.WriteHeader(status)
	tmpl := template.Must(template.New("result").Parse(htmlContent))
	return tmpl.Execute(w, struct{ Banner string }{asciiArt})
}

func NewAPIServer(listenAdrr string) *APIServer {
	return &APIServer{
		listenAdrr: listenAdrr,
	}
}

func (s *APIServer) Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", makeHTTPHandlerFunc(s.MainPageHandler))
	mux.HandleFunc("/ascii-art", makeHTTPHandlerFunc(s.AsscitArtWebHandler))
	http.ListenAndServe(s.listenAdrr, mux)
}

func (s *APIServer) MainPageHandler(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}

	if r.Method == "GET" {
		return s.MainPageGetHandler(w, r)
	} else {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

}

func (s *APIServer) MainPageGetHandler(w http.ResponseWriter, r *http.Request) error {
	htmlContent := `<!DOCTYPE html>
	<html>
	<head>
		<title>ASCII Art Web</title>
		<style>
		body {
            background-color: #f7f7f7;
            font-family: Arial, sans-serif;
            margin: auto;
            padding: 20px;
            max-width: 600px; /* Limit width to 600px */
        }

        .container {
            padding: 20px;
            background-color: #fff;
            border-radius: 8px;
            box-shadow: 0px 0px 20px rgba(0, 0, 0, 0.1);
            text-align: center;
        }

        h1 {
            color: #333;
            margin-bottom: 20px;
            text-align: center;
        }

        form {
            display: flex;
            flex-direction: column;
            align-items: center;
        }

        label {
            color: #555;
            margin-bottom: 5px;
        }

        input[type="text"],
        select {
            width: 100%;
            padding: 10px;
            margin-bottom: 15px;
            border: 1px solid #ccc;
            border-radius: 4px;
            box-sizing: border-box;
        }

        button {
            background-color: #4CAF50;
            color: white;
            padding: 15px 20px;
            border: none;
            border-radius: 4px;
            cursor: pointer;
            transition: background-color 0.3s ease;
        }

        button:hover {
            background-color: #45a049;
        }

        @media (max-width: 600px) {
            /* Apply styles for screens with a maximum width of 600px */
            body {
                padding: 10px;
            }

            .container {
                padding: 10px;
            }

            input[type="text"],
            select,
            button {
                width: auto; /* Adjust width to fit content */
            }
        }
		</style>
	</head>
	<body>
		<h1>ASCII Art Web</h1>
		<form method="POST" action="/ascii-art">
			<label for="text">Text:</label><br>
			<input type="text" id="text" name="text" placeholder="Type something" required><br>
			<label for="banner">Banner:</label><br>
			<select id="banner" name="banner">
				<option value="shadow">Shadow</option>
				<option value="standard">Standard</option>
				<option value="thinkertoy">Thinkertoy</option>
				<!-- <option value="SKit">Skit</option> -->
			</select><br><br>
			<button type="submit">Generate ASCII Art</button>
		</form>
	</body>
	</html>`

	return WriteHTML(w, http.StatusOK, htmlContent)
}

func (s *APIServer) AsscitArtWebHandler(w http.ResponseWriter, r *http.Request) error {
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

	directory := "."
	chLength := 8
	// var firstCh byte = ' '
	// var lastCh byte = '~'
	patternContent, errPattern := util.ReadFileToStr(directory, patternFileName)
	if errPattern != nil {
		return "", fmt.Errorf("error loading pattern file: %v", errPattern)
	}
	patternMap, errPatternMap := util.ContentToMap(patternContent, chLength)
	if errPatternMap != nil {
		return "", fmt.Errorf("error creating pattern map: %v", errPatternMap)
	}
	return draw.StrMaker(text, patternMap, chLength), nil
}
