package server

import (
	"net/http"

	"github.com/goharbor/harbor/src/server/router"
)

func registerRoutes() {
	// API version
	router.NewRoute().Method(http.MethodGet).Path("/api/version").HandlerFunc(GetAPIVersion)
}