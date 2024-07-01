package server

import (
	"encoding/json"
	"go-blog-api/internal/auth"
	"go-blog-api/internal/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) RootRouteHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Go Blogs"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) GetUserHandler(c *gin.Context) {
	username := c.Param("username")
	user := s.db.GetUser(c, username)
	if user == nil {
		return
	}
	c.JSON(http.StatusOK, user)
}

func (s *Server) HandleSignUp(c *gin.Context) {
	accReq := &types.AuthRequest{}
	resp := make(map[string]string)
	err := json.NewDecoder(c.Request.Body).Decode(accReq)

	if err != nil {
		resp["message"] = "Invalid request"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	existingUser := s.db.GetUser(c, accReq.Username)
	if existingUser != nil {
		resp["message"] = "User already exists"
		c.JSON(http.StatusConflict, resp)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(accReq.Password), bcrypt.DefaultCost)

	accReq.Password = string(hashedPassword)

	if err != nil {
		resp["message"] = "Failed to hash password"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	s.db.CreateUser(c, *accReq)
	resp["message"] = "Account created"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) HandleLogin(c *gin.Context) {
	loginReq := &types.AuthRequest{}
	json.NewDecoder(c.Request.Body).Decode(loginReq)

	dbUser := s.db.GetUser(c, loginReq.Username)

	resp := make(map[string]string)

	PassErr := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginReq.Password))
	if PassErr != nil {
		resp["message"] = "Wrong Password"
		c.JSON(http.StatusUnauthorized, resp)
		return
	}

	token, err := auth.GenerateJWT(dbUser.ID)
	if err != nil {
		resp["message"] = "Error generating JWT Token"
		c.JSON(http.StatusUnauthorized, resp)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"token": token, "user": map[string]interface{}{"username": dbUser.Username, "id": dbUser.ID}})
}

func (s *Server) CreatePostHandler(c *gin.Context) {
	req := &types.CreatePostRequest{}
	resp := make(map[string]string)
	err := json.NewDecoder(c.Request.Body).Decode(req)

	if err != nil {
		resp["message"] = "Invalid request"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	s.db.CreatePost(c, *req)
	resp["message"] = "Post created"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) GetPostHandler(c *gin.Context) {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	post := s.db.GetPost(c, uint(intID))
	if post == nil {
		c.JSON(http.StatusNotFound, map[string]string{"message": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (s *Server) GetAllPostsHandler(c *gin.Context) {
	posts := s.db.GetAllPosts(c)
	c.JSON(http.StatusOK, posts)
}

func (s *Server) UpdatePostHandler(c *gin.Context) {
	req := &types.UpdatePostRequest{}
	resp := make(map[string]string)
	err := json.NewDecoder(c.Request.Body).Decode(req)

	if err != nil {
		resp["message"] = "Invalid request"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	id := c.Param("id")
	intID, _ := strconv.Atoi(id)
	post := s.db.GetPost(c, uint(intID))
	if post == nil {
		c.JSON(http.StatusNotFound, map[string]string{"message": "Post not found"})
		return
	}

	if post.UserID != req.UserID {
		c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
		return
	}

	s.db.UpdatePost(c, *req)
	resp["message"] = "Post updated"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) DeletePostHandler(c *gin.Context) {
	resp := make(map[string]string)
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	post := s.db.GetPost(c, uint(intID))
	if post == nil {
		c.JSON(http.StatusNotFound, map[string]string{"message": "Post not found"})
		return
	}

	s.db.DeletePost(c, uint(intID))
	resp["message"] = "Post deleted"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) AddCommentHandler(c *gin.Context) {
	req := &types.CreateCommentRequest{}
	resp := make(map[string]string)
	err := json.NewDecoder(c.Request.Body).Decode(req)

	if err != nil {
		resp["message"] = "Invalid request"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	s.db.CreateComment(c, *req)
	resp["message"] = "Comment created"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) GetCommentsHandler(c *gin.Context) {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)
	comments := s.db.GetComments(c, uint(intID))
	c.JSON(http.StatusOK, comments)
}

func (s *Server) DeleteCommentHandler(c *gin.Context) {
	resp := make(map[string]string)
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	s.db.DeleteComment(c, uint(intID))
	resp["message"] = "Comment deleted"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) LikePostHandler(c *gin.Context) {
	resp := make(map[string]string)
	req := &types.LikePostRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(req)
	if err != nil {
		resp["message"] = "Invalid request"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	post := s.db.GetPost(c, req.PostID)
	if post == nil {
		c.JSON(http.StatusNotFound, map[string]string{"message": "Post not found"})
		return
	}

	s.db.LikePost(c, *req)
	resp["message"] = "Post liked"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) UnlikePostHandler(c *gin.Context) {
	resp := make(map[string]string)
	req := &types.LikePostRequest{}
	err := json.NewDecoder(c.Request.Body).Decode(req)
	if err != nil {
		resp["message"] = "Invalid request"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	post := s.db.GetPost(c, req.PostID)
	if post == nil {
		c.JSON(http.StatusNotFound, map[string]string{"message": "Post not found"})
		return
	}

	s.db.UnlikePost(c, *req)
	resp["message"] = "Post unliked"

	c.JSON(http.StatusOK, resp)
}
