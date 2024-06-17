package request

import "mime/multipart"

type ProductRequest struct {
	UUID     string
	Name     string                `form:"name" binding:"required"`
	ImageURL *multipart.FileHeader `form:"image_file" binding:"required"`
	AdminID  uint                  `binding:"required"`
	// Variants
}
