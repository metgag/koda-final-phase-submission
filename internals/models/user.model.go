package models

import (
	"mime/multipart"
	"time"
)

type PostBody struct {
	Content *string               `form:"content"`
	Image   *multipart.FileHeader `form:"image"`
}

type Post struct {
	PostID       int       `json:"post_id"`
	Content      *string   `json:"content"`
	Image        *string   `json:"image"`
	CreatorID    int       `json:"creator_id"`
	CreatorUname string    `json:"creator_uname"`
	CreatedAt    time.Time `json:"created_at"`
}

type CommentPostBody struct {
	Comment string `json:"comment"`
