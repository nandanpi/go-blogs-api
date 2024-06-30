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
	r.GET("/user/:username", auth.ProtectedRoute(s.GetUserHandler))
	r.POST("/createPost", auth.ProtectedRoute(s.CreatePostHandler))
	r.GET("/post/:id", s.GetPostHandler)
	r.GET("/posts", s.GetAllPostsHandler)
	r.POST("/updatePost/:id", auth.ProtectedRoute(s.UpdatePostHandler))
	r.DELETE("/deletePost/:id", auth.ProtectedRoute(s.DeletePostHandler))
	r.POST("/addComment", auth.ProtectedRoute(s.AddCommentHandler))
	r.GET("/getComments", s.GetCommentsHandler)
	r.DELETE("deleteComment/:id", auth.ProtectedRoute(s.DeleteCommentHandler))
	r.POST("/likePost", auth.ProtectedRoute(s.LikePostHandler))
	r.POST("/unlikePost", auth.ProtectedRoute(s.UnlikePostHandler))

	return r
}
