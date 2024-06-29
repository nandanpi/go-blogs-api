package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.RootRouteHandler)
	r.POST("/signup", s.HandleSignUp)
	r.POST("/login", s.HandleLogin)

	return r
}
