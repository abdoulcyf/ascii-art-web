package handlers

import "net/http"

// MainPageGetHandler handles GET requests to the main page.
func MainPageGetHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	return HomeHadler(w, http.StatusOK)
}
