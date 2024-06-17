package controllers

import (
	"basic_trade/database"
	"basic_trade/helpers"
	"basic_trade/models/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	jwt5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateVariant(ctx *gin.Context) {
	db := database.GetDB()

	adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	adminID := uint(adminData["id"].(float64))

	Product := entity.Product{}
	Variant := entity.Variant{}

	contentType := helpers.GetContentType(ctx)
	if contentType == appJSON {
		ctx.ShouldBindJSON(&Variant)
	} else {
		ctx.ShouldBind(&Variant)
	}

	Product.AdminID = adminID
	newUUID := uuid.New()
	Variant.UUID = newUUID.String()

	err := db.Debug().Create(&Variant).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Variant,
	})
}

func GetVariant(ctx *gin.Context) {
	db := database.GetDB()
	Variants := []entity.Variant{}

	err := db.Debug().Find(&Variants).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request.",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Variants,
	})
}

func GetVariantbyUUID(ctx *gin.Context) {
	db := database.GetDB()
	Variants := entity.Variant{}
	variantUUID := ctx.Param("variantUUID")

	err := db.Debug().Where("uuid = ?", variantUUID).First(&Variants).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": Variants,
	})
}

func UpdateVariantbyUUID(ctx *gin.Context) {
	db := database.GetDB()
	// adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	contentType := helpers.GetContentType(ctx)

	Variants := []entity.Variant{}
	updatedVariant := entity.Variant{}
	variantUUID := ctx.Param("variantUUID")
	updatedVariant.UUID = variantUUID

	if contentType == appJSON {
		ctx.ShouldBindJSON(&updatedVariant)
	} else {
		ctx.ShouldBind(&updatedVariant)
	}

	err := db.Debug().Where("uuid = ?", variantUUID).First(&Variants).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Model(&Variants).Where("uuid = ?", variantUUID).Updates(&updatedVariant).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": updatedVariant,
	})
}

func DeleteVariantbyUUID(ctx *gin.Context) {
	db := database.GetDB()
	// adminData := ctx.MustGet("adminData").(jwt5.MapClaims)
	contentType := helpers.GetContentType(ctx)

	Variants := []entity.Variant{}
	variantToDelete := entity.Variant{}
	variantUUID := ctx.Param("variantUUID")
	variantToDelete.UUID = variantUUID

	if contentType == appJSON {
		ctx.ShouldBindJSON(&variantToDelete)
	} else {
		ctx.ShouldBind(&variantToDelete)
	}

	err := db.Debug().Where("uuid = ?", variantUUID).First(&Variants).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Model(&Variants).Where("uuid = ?", variantUUID).Delete(&variantToDelete).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted record.",
	})
}
