package handlers

import "net/http"

// MainPageHandler handles requests to the main page
func (s *Server) MainPageHandler(w http.ResponseWriter, r *http.Request) error {
	if r.URL.Path != homePagePath {
		http.Error(w, StatusNotFound, http.StatusNotFound)
		return nil
	}

	if r.Method != getRequest {
		http.Error(w, methodNotAllowed, http.StatusMethodNotAllowed)
		return nil
	}

	return s.homeHandler(w, r)
}
