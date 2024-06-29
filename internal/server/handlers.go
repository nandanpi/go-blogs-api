package server

import (
	"encoding/json"
	"go-blog-api/internal/auth"
	"go-blog-api/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) RootRouteHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Go Blogs"

	c.JSON(http.StatusOK, resp)
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

	c.JSON(http.StatusOK, map[string]string{"token": token})
}
