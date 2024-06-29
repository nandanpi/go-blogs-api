package server

import (
	"go-blog-api/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", auth.ProtectedRoute(s.RootRouteHandler))
	r.POST("/signup", s.HandleSignUp)
	r.POST("/login", s.HandleLogin)

	return r
}
