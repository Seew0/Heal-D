package server

import (
	"log"

	"github.com/Seew0/Heal-D/internal/router"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

// NewServer initializes a new server with the router
func NewServer(r *router.Router) *Server {
	ginRouter := gin.Default()

	// Setup all routes
	r.SetupRoutes(ginRouter)

	return &Server{router: ginRouter}
}

// Start runs the server
func (s *Server) Start(port string) {
	log.Printf("Server running on port %s", port)
	if err := s.router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
