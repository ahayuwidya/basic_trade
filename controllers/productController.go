package controllers

import (
	"basic_trade/database"
	"basic_trade/helpers"
	"basic_trade/models/entity"
	"basic_trade/models/request"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()

	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	contentType := helpers.GetContentType(ctx)
	adminID := uint(adminData["id"].(float64))

	productReq := request.ProductRequest{}
	if contentType == appJSON {
		ctx.ShouldBindJSON(&productReq)
	} else {
		ctx.ShouldBind(&productReq)
	}

	productReq.AdminID = adminID
	newUUID := uuid.New()
	productReq.UUID = newUUID.String()

	if helpers.IsValidImageSize(int(productReq.ImageURL.Size)) {
		if helpers.IsValidImageExtension(productReq.ImageURL.Filename) {
			fileName := helpers.RemoveExtension(productReq.ImageURL.Filename)
			uploadResult, err := helpers.UploadFile(productReq.ImageURL, fileName)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			Product := entity.Product{
				UUID:     productReq.UUID,
				Name:     productReq.Name,
				ImageURL: uploadResult,
				AdminID:  productReq.AdminID,
			}

			err = db.Debug().Create(&Product).Error
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error":   "Bad request",
					"message": err.Error(),
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"data": Product,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": "File extension should be in JPG, JPEG, PNG or SVG.",
			})
		}

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": "File size should be less than 5MB.",
		})
	}
}

func GetProduct(ctx *gin.Context) {
	db := database.GetDB()
	Products := []entity.Product{}

	err := db.Debug().Find(&Products).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request.",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Products,
	})
}

func GetProductbyUUID(ctx *gin.Context) {
	db := database.GetDB()
	Products := []entity.Product{}
	productUUID := ctx.Param("productUUID")

	err := db.Debug().Where("uuid = ?", productUUID).First(&Products).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Products,
	})
}

func UpdateProductbyUUID(ctx *gin.Context) {
	db := database.GetDB()

	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	contentType := helpers.GetContentType(ctx)
	adminID := uint(adminData["id"].(float64))
	productUUID := ctx.Param("productUUID")

	Products := []entity.Product{}
	updatedProductReq := request.ProductRequest{}
	if contentType == appJSON {
		ctx.ShouldBindJSON(&updatedProductReq)
	} else {
		ctx.ShouldBind(&updatedProductReq)
	}

	err := db.Debug().Where("uuid = ?", productUUID).First(&Products).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err,
		})
		return
	}

	updatedProductReq.AdminID = adminID

	if helpers.IsValidImageSize(int(updatedProductReq.ImageURL.Size)) {
		if helpers.IsValidImageExtension(updatedProductReq.ImageURL.Filename) {
			fileName := helpers.RemoveExtension(updatedProductReq.ImageURL.Filename)
			uploadResult, err := helpers.UploadFile(updatedProductReq.ImageURL, fileName)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			updatedProduct := entity.Product{
				UUID:     updatedProductReq.UUID,
				Name:     updatedProductReq.Name,
				ImageURL: uploadResult,
				AdminID:  updatedProductReq.AdminID,
			}

			err = db.Debug().Model(&Products).Where("uuid = ?", productUUID).Updates(&updatedProduct).Error
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error":   "Bad request",
					"message": err.Error(),
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"data": updatedProduct,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": "File extension should be in JPG, JPEG, PNG or SVG.",
			})
		}

	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": "File size should be less than 5MB.",
		})
	}
}

func DeleteProductbyUUID(ctx *gin.Context) {
	db := database.GetDB()

	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	contentType := helpers.GetContentType(ctx)

	Products := []entity.Product{}
	productToDelete := entity.Product{}
	productUUID := ctx.Param("productUUID")

	productToDelete.AdminID = uint(adminData["id"].(float64))
	productToDelete.UUID = productUUID

	if contentType == appJSON {
		ctx.ShouldBindJSON(&productToDelete)
	} else {
		ctx.ShouldBind(&productToDelete)
	}

	err := db.Debug().Model(&Products).Where("uuid = ?", productUUID).Delete(&productToDelete).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Successfully deleted record.",
		})
	}
}
