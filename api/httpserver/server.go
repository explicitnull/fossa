package httpserver

import (
	"context"
	"errors"
	"fossa/service/ticket"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const port = ":8080"

type Config struct {
	Address               string `env:"ADDRESS" yaml:"address"`
	Port                  string `env:"PORT" yaml:"port"`
	AuthenticationEnabled bool   `env:"AUTHENTICATION_ENABLED" yaml:"authentication_enabled"`
	// JwtPublicKey          string `env:"JWT_PUBLIC_KEY" yaml:"jwt_public_key"`
	// JwtPublicKeyParsed    crypto.PublicKey
}

const (
	gracefulTimeout = time.Second * 15
)

type Server struct {
	config Config

	stdserver *http.Server
	engine    *gin.Engine

	ticketService *ticket.Service
}

func New(config Config, ticketService *ticket.Service) *Server {
	// Disabling gin logs
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	stdserver := &http.Server{
		Addr:    port,
		Handler: engine,

		//TODO: set read/write timeouts
	}

	return &Server{
		config:        config,
		stdserver:     stdserver,
		engine:        engine,
		ticketService: ticketService,
	}
}
func (s *Server) setupRoutes() {
	// add metrics and liveness check
	apiv1 := s.engine.Group("/api/v1")

	ticketsGroup := apiv1.Group("/tickets")
	ticketsGroup.GET("/", s.GetTickets)
}

func (s *Server) Run() {
	log.Println("Running server on ", port)

	s.setupRoutes()

	err := s.stdserver.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalln("Can't listen and serve: ", err)
	}
}

func (s *Server) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), gracefulTimeout)
	defer cancel()

	err := s.stdserver.Shutdown(ctx)
	if err != nil {
		log.Println("Can't shutdown http server: ", err)
	}
}
