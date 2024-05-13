package servers

import (
	"fmt"
	"net/http"
)


type MyError struct {
	code int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.code, e.Message)
}


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
