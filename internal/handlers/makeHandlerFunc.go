package handlers

import "net/http"

// makeHTTPHandlerFunc creates an HTTP handler function from an API function.
func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			err := writeHomePageContent(w, http.StatusBadRequest)
			if err != nil {
				return
			}
		}
	}
}
