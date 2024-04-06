
# Description

Ascii-art-web is a web application that allows users to generate ASCII art using different banners (shadow, standard, thinkertoy) through a graphical user interface. It provides HTTP endpoints for handling requests to display the main page, receive user input, and generate ASCII art.

# Authors

1. Elhadj Abdoul Diallo - student at 01 founders
2. Reza - student at 01 founders

# Usage: How to run

To run the Ascii-art-web server:

* Make sure you have [Golang](https://go.dev/doc/install) installed on your machine.
* Clone the repository: git clone [repository_url]
* Navigate to the project subdirectory: cd ascii-art-web/cmd
* Run the server: go run main.go
* Access the application in your web browser at http://localhost:8080`

# Implementation details: algorithm

The Ascii-art-web application follows these implementation details:

The server is written in Go.  
HTML templates are stored in the templates directory in the project root.  
The server handles GET requests to the root path ('/') by rendering the main page template, which includes a text input, radio buttons to switch between banners, and a button to submit the form via a POST request.  
The server handles POST requests to '/ascii-art' by parsing the form data, generating ASCII art based on the provided text and selected banner, and rendering the result on a separate page to the main page.  
The server returns appropriate HTTP status codes for different scenarios:  
OK (200) for successful requests.  
Not Found (404) for missing templates or banners.
Bad Request (400) for incorrect requests.  
Internal Server Error (500) for unhandled errors.  