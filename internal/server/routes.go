package server

import (
	"go-blog-api/internal/auth"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	config := cors.Config{
		AllowAllOrigins: true,

		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	r.Use(cors.New(config))

	r.GET("/", auth.ProtectedRoute(s.RootRouteHandler))
	r.POST("/signup", s.HandleSignUp)
	r.POST("/login", s.HandleLogin)
	r.GET("/user/:username", auth.ProtectedRoute(s.GetUserHandler))
	r.POST("/createPost", auth.ProtectedRoute(s.CreatePostHandler))
	r.GET("/post/:id", s.GetPostHandler)
	r.GET("/posts", s.GetAllPostsHandler)
	r.POST("/updatePost/:id", auth.ProtectedRoute(s.UpdatePostHandler))
	r.DELETE("/deletePost/:id", auth.ProtectedRoute(s.DeletePostHandler))
	r.POST("/addComment", auth.ProtectedRoute(s.AddCommentHandler))
	r.GET("/getComments/:id", s.GetCommentsHandler)
	r.DELETE("deleteComment/:id", auth.ProtectedRoute(s.DeleteCommentHandler))
	r.POST("/likePost", auth.ProtectedRoute(s.LikePostHandler))
	r.POST("/unlikePost", auth.ProtectedRoute(s.UnlikePostHandler))

	return r
}
