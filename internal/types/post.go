package types

type CreatePostRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}

type UpdatePostRequest struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateCommentRequest struct {
	UserID  uint   `json:"user_id"`
	PostID  uint   `json:"post_id"`
	Content string `json:"content"`
}

type LikePostRequest struct {
	UserID uint `json:"user_id"`
	PostID uint `json:"post_id"`
}
