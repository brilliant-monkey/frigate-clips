package router

import (
	"fmt"
	"net/http"

	"git.brilliantmonkey.net/frigate/frigate-clips/config"
	"git.brilliantmonkey.net/frigate/frigate-clips/controller"
	"git.brilliantmonkey.net/frigate/frigate-clips/types"
	"github.com/gorilla/mux"
)

func getRoutes(config *config.AppConfig) []types.Route {
	return []types.Route{
		{Method: http.MethodGet, Path: "/health", HandlerFunc: http.HandlerFunc(controller.Health), Name: "health"},
		{Method: http.MethodGet, Path: "/", HandlerFunc: http.FileServer(http.Dir(config.FFMPEG.OutputDir)), Name: "files"},
	}
}

func CreateRouter(config *config.AppConfig) *mux.Router {
	router := mux.NewRouter()

	for _, route := range getRoutes(config) {
		router.
			Methods(route.Method, http.MethodOptions).
			Path(fmt.Sprint("/clips/v1", route.Path)).
			Handler(http.StripPrefix("/clips/v1/", route.HandlerFunc)).
			Name(route.Name)
	}
	return router
}
