package handlers

import (
	"net/http"
	"text/template"
)

// Render the main page using a template
func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("main").Parse(`
        <!DOCTYPE html>
        <html>
        <head>
            <title>ASCII Art Web</title>
            <style>
			body {
				background-color: #f7f7f7;
				font-family: Arial, sans-serif;
				width: 600px;
				margin: auto;
			}
	
			.container {
				max-width: 400px;
				padding: 20px;
				background-color: #fff;
				border-radius: 8px;
				box-shadow: 0px 0px 20px rgba(0, 0, 0, 0.1);
				text-align: center;
			}
	
			h1 {
				color: #333;
				margin-bottom: 20px;
				text-aligned: center;
			}
	
			form {
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
            </style>
        </head>
        <body>
            <h1>ASCII Art Web</h1>
            <form method="post" action="/ascii-art">
                <label for="text">Text:</label><br>
                <input type="text" id="text" name="text" placeholder="Type something" required><br>
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
