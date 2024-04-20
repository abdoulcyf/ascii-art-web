package handlers

import "net/http"

// MainPageGetHandler handles GET requests to the main page.
func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		http.Error(w, methodNotAllowed, http.StatusMethodNotAllowed)
	}

	return writeHomePageContent(w, http.StatusOK)
}
