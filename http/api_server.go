package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"git.brilliantmonkey.net/frigate/frigate-clips/config"
	"git.brilliantmonkey.net/frigate/frigate-clips/router"
	"github.com/gorilla/mux"
)

type APIServer struct {
	httpServer *http.Server
	router     *mux.Router
}

func NewAPIServer(config *config.AppConfig) *APIServer {
	router := router.CreateRouter(config)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("0.0.0.0:%s", port),
	}

	return &APIServer{
		server,
		router,
	}
}

func (server *APIServer) Start() error {
	log.Println("Starting API server...")
	err := server.httpServer.ListenAndServe()
	if err == http.ErrServerClosed {
		return nil
	}
	return err
}

func (server *APIServer) Stop() error {
	log.Println("Shutting down API server...")
	return server.httpServer.Shutdown(context.Background())
}
