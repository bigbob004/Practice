package web

import "net/http"

func GetHTTPServer() http.Server {
	return http.Server{
		Addr:    ":8080",
		Handler: GetHTTPHandler(),
	}
}
