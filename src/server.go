package server

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ServerConfig struct {
	Port int
}

type Server struct {
	config *ServerConfig
	router *httprouter.Router
}

func NewServer(config *ServerConfig, router *httprouter.Router) *Server {
	sv := Server{
		config: config,
		router: router,
	}

	return &sv
}

func (sv *Server) Run() {
	addr := fmt.Sprintf(":%d", sv.config.Port)

	http.ListenAndServe(addr, sv.router)
}
