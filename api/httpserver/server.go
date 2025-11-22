package httpserver

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const port = ":8080"

const (
	gracefulTimeout = time.Second * 15
)

type Server struct {
	stdserver *http.Server
	engine    *gin.Engine
}

func New() *Server {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	stdserver := &http.Server{
		Addr:    port,
		Handler: engine,

		//TODO: set read/write timeouts
	}

	return &Server{
		stdserver: stdserver,
		engine:    engine,
	}
}

func (s *Server) Run() {
	// add metrics and liveness check
	apiv1 := s.engine.Group("/api/v1")

	ticketsGroup := apiv1.Group("/tickets")
	ticketsGroup.GET("/test", s.GetTickets)

	log.Println("Running server on", port)

	err := s.stdserver.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln("can't listen and serve:", err)
	}
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), gracefulTimeout)
	defer cancel()

	err := s.stdserver.Shutdown(ctx)
	if err != nil {
		log.Println("can't shutdown http server: ", err)
	}
}
