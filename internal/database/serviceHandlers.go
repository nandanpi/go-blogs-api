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
	return user
}

func (d *service) CreatePost(c *gin.Context, req types.CreatePostRequest) {
	resp := make(map[string]string)
	result := d.db.Select("Title", "Content", "UserID").Create(&BlogPost{Title: req.Title, Content: req.Content, UserID: req.UserID})
	if result.Error != nil {
		resp["message"] = "Failed to create post"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp["message"] = "Post created successfully"

	c.JSON(http.StatusOK, resp)
}

func (d *service) GetPost(c *gin.Context, id uint) *BlogPost {
	resp := make(map[string]string)
	post := &BlogPost{}
	result := d.db.Where("id = ?", id).First(&post)
	if result.Error != nil {
		resp["message"] = "Post not found"
		c.JSON(http.StatusNotFound, resp)
		return nil
	}
	return post
}

func (d *service) GetAllPosts(c *gin.Context) []BlogPost {
	resp := make(map[string]string)
	posts := []BlogPost{}
	result := d.db.Find(&posts)
	if result.Error != nil {
		resp["message"] = "Failed to get posts"
		c.JSON(http.StatusInternalServerError, resp)
		return nil
	}
	return posts
}

func (d *service) UpdatePost(c *gin.Context, req types.UpdatePostRequest) {
	resp := make(map[string]string)
	post := &BlogPost{}
	result := d.db.Where("id = ?", req.ID).First(&post)
	if result.Error != nil {
		resp["message"] = "Post not found"
		c.JSON(http.StatusNotFound, resp)
		return
	}
	post.Title = req.Title
	post.Content = req.Content
	result = d.db.Save(&post)
	if result.Error != nil {
		resp["message"] = "Failed to update post"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp["message"] = "Post updated successfully"
}

func (d *service) DeletePost(c *gin.Context, id uint) {
	resp := make(map[string]string)
	post := &BlogPost{}
	result := d.db.Where("id = ?", id).First(&post)
	if result.Error != nil {
		resp["message"] = "Post not found"
		c.JSON(http.StatusNotFound, resp)
		return
	}
	result = d.db.Delete(&post)
	if result.Error != nil {
		resp["message"] = "Failed to delete post"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp["message"] = "Post deleted successfully"
	c.JSON(http.StatusOK, resp)
}

func (d *service) CreateComment(c *gin.Context, req types.CreateCommentRequest) {
	resp := make(map[string]string)
	result := d.db.Select("UserID", "PostID", "Content").Create(&Comments{UserID: req.UserID, PostID: req.PostID, Content: req.Content})
	if result.Error != nil {
		resp["message"] = "Failed to create comment"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp["message"] = "Comment created successfully"

	c.JSON(http.StatusOK, resp)
}

func (d *service) GetCommentByID(c *gin.Context, commentID uint) Comments {
	resp := make(map[string]string)
	comment := Comments{}
	result := d.db.Where("id = ?", commentID).First(&comment)
	if result.Error != nil {
		resp["message"] = "Comment not found"
		c.JSON(http.StatusNotFound, resp)
		return Comments{}
	}
	return comment
}

func (d *service) GetComments(c *gin.Context, postID uint) []Comments {
	resp := make(map[string]string)
	comments := []Comments{}
	result := d.db.Where("post_id = ?", postID).Find(&comments)
	if result.Error != nil {
		resp["message"] = "Failed to get comments"
		c.JSON(http.StatusInternalServerError, resp)
		return nil
	}
	return comments
}

func (d *service) DeleteComment(c *gin.Context, id uint) {
	resp := make(map[string]string)
	comment := &Comments{}
	result := d.db.Where("id = ?", id).First(&comment)
	if result.Error != nil {
		resp["message"] = "Comment not found"
		c.JSON(http.StatusNotFound, resp)
		return
	}
	result = d.db.Delete(&comment)
	if result.Error != nil {
		resp["message"] = "Failed to delete comment"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp["message"] = "Comment deleted successfully"
	c.JSON(http.StatusOK, resp)
}

func (d *service) LikePost(c *gin.Context, req types.LikePostRequest) {
	resp := make(map[string]string)
	result := d.db.Select("UserID", "PostID").Create(&Likes{UserID: req.UserID, PostID: req.PostID})
	if result.Error != nil {
		resp["message"] = "Failed to like post"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp["message"] = "Post liked successfully"

	c.JSON(http.StatusOK, resp)
}

func (d *service) UnlikePost(c *gin.Context, req types.LikePostRequest) {
	resp := make(map[string]string)
	like := &Likes{}
	result := d.db.Where("user_id = ? AND post_id = ?", req.UserID, req.PostID).First(&like)
	if result.Error != nil {
		resp["message"] = "Like not found"
		c.JSON(http.StatusNotFound, resp)
		return
	}
	result = d.db.Delete(&like)
	if result.Error != nil {
		resp["message"] = "Failed to unlike post"
		c.JSON(http.StatusInternalServerError, resp)
		return
	}
	resp["message"] = "Post unliked successfully"
	c.JSON(http.StatusOK, resp)
}
