package utils

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/metgag/final-assignment/internals/models"
)

func LogCtxError(ctx *gin.Context, errHead, errClient string, errDev error, statusCode int) {
	log.Printf("%s\n\t%s", errHead, errDev.Error())
	ctx.JSON(statusCode, models.ErrorResponse{
		Success: false,
		Status:  statusCode,
		Error:   errClient,
	})
}

func MwareLogCtxError(ctx *gin.Context, errHead, errClient string, errDev error, statusCode int) {
	log.Printf("%s\n\t%s", errHead, errDev.Error())
	ctx.AbortWithStatusJSON(statusCode, models.ErrorResponse{
		Success: false,
		Status:  statusCode,
		Error:   errClient,
	})
}
