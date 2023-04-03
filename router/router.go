package router

import (
	"net/http"

	"git.brilliantmonkey.net/frigate/frigate-clips/config"
	"git.brilliantmonkey.net/frigate/frigate-clips/controller"
	"github.com/gorilla/mux"
)

func getFileHandler(clipDir string) http.Handler {
	return http.StripPrefix("/clips/v1/", http.FileServer(http.Dir(clipDir)))
}

func CreateRouter(config *config.AppConfig) *mux.Router {
	router := mux.NewRouter()

	router.
		PathPrefix("/clips/v1/").
		Path("/health").Methods(http.MethodGet, http.MethodOptions).
		Handler(http.HandlerFunc(controller.Health)).
		Name("health")

	router.
		PathPrefix("/clips/v1/").
		Methods(http.MethodGet, http.MethodOptions).
		Handler(getFileHandler(config.FFMPEG.OutputDir)).
		Name("serve")

	return router
}
