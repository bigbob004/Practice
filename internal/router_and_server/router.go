package web

import (
	"net/http"
	"webScraper/internal/routes"
)

// GetHTTPHandler gets the Handler responds to an HTTP request.
func GetHTTPHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/parse", routes.Parse)

	return mux
}
