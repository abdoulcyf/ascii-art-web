package servers

import "net/http"

type server struct {
	listener string
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

type Error struct {
	Error string
}

const (
	homePagePath = "/"
	asciiArtPage = "/ascii-art"

	//portNumber = ":8080"

	baseStaticDir = "/static/"
	staticFilesAddress = "../../static"
)
