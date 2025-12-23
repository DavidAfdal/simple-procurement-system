package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/DavidAfdal/purchasing-systeam/internal/middelwares"
	"github.com/DavidAfdal/purchasing-systeam/pkg/response"
	"github.com/DavidAfdal/purchasing-systeam/pkg/route"
	"github.com/DavidAfdal/purchasing-systeam/pkg/token"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	Server *http.Server
}

func NewServer(
	publicRoutes, privateRoutes []*route.Route,
	secretKey string,
	tokenUse token.TokenUseCase,
) *Server {

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middelwares.CorsMiddleware())

	r.GET("/", func(c *gin.Context) {
		response.SuccessResponse(c, http.StatusOK, "Hello, World!", nil)
	})

	v1 := r.Group("/api")

	for _, v := range publicRoutes {
		v1.Handle(v.Method, v.Path, v.Handler)
	}

	for _, v := range privateRoutes {
		v1.Handle(
			v.Method,
			v.Path,
			middelwares.JWTProtection(secretKey),
			middelwares.UserContextMiddleware(),
			v.Handler,
		)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return &Server{
		Engine: r,
		Server: srv,
	}
}

func (s *Server) Run() {
	runServer(s)
	gracefulShutdown(s.Server)
}

func runServer(s *Server) {
	go func() {
		log.Println("Server running on :8080")
		if err := s.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen:", err)
		}
	}()
}

func gracefulShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
