package models

import "mime/multipart"

type PostBody struct {
	Content *string               `form:"content"`
	Image   *multipart.FileHeader `form:"image"`
}
