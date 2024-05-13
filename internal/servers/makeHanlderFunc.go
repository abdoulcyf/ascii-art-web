package servers

import (
	"net/http"
	"github.com/ediallocyf/asciiartweb/internal/handlers"
)

// makeHTTPHandlerFunc creates an HTTP handler function from an API function.
func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			err := handlers.HomeHadler(w, http.StatusBadRequest)
			if err != nil {
				return
			}
		}
	}
}
