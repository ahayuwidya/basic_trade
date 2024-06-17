package controllers

import (
	"basic_trade/database"
	"basic_trade/helpers"
	"basic_trade/models/entity"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	appJSON = "application/json"
)

func AdminRegister(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	Admin := entity.Admin{}
	Admins := []entity.Admin{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Admin)
	} else {
		ctx.ShouldBind(&Admin)
	}

	err := db.Where("email = ?", &Admin.Email).First(&Admins).Error
	if err != nil {
		// Generate new UUID and set the generated UUID as Admin's UUID
		newUUID := uuid.New()
		Admin.UUID = newUUID.String()

		err := db.Debug().Create(&Admin).Error
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    Admin,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": "Email is already registered.",
		})
	}

}

func AdminLogin(ctx *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(ctx)
	Admin := entity.Admin{}

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Admin)
	} else {
		ctx.ShouldBind(&Admin)
	}

	password := Admin.Password

	err := db.Debug().Where("email = ?", Admin.Email).Take(&Admin).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid email.",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(Admin.Password), []byte(password))
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "Invalid password.",
		})
		return
	}

	token := helpers.GenerateToken(Admin.ID, Admin.Email)
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
