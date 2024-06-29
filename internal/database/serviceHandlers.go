package database

import (
	"go-blog-api/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d *service) CreateUser(c *gin.Context, user types.AuthRequest) {
	resp := make(map[string]string)
	result := d.db.Select("Username", "Password").Create(&Users{Username: user.Username, Password: user.Password})
	if result.Error != nil {
		resp["message"] = "Failed to create user"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp["message"] = "User created successfully"

	c.JSON(http.StatusOK, resp)
}

func (d *service) GetUser(c *gin.Context, username string) *Users {
	user := &Users{}
	resp := make(map[string]string)
	result := d.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		resp["message"] = "User not found"
		c.JSON(http.StatusNotFound, resp)
		return nil
	}
	resp["message"] = "User found"
	c.JSON(http.StatusOK, resp)
	return user
}
