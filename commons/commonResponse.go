package commons

import (
	"tugas15/structs"

	"github.com/gin-gonic/gin"
)

func ResponseSuccessWithData(ctx *gin.Context, statusCode int, message string, data interface{}) {
	ctx.JSON(
		statusCode,
		structs.ApiResponse{
			Success: true,
			Message: message,
			Data:    data,
		},
	)
}

func ResponseSuccessWithoutData(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(
		statusCode,
		structs.ApiResponse{
			Success: true,
			Message: message,
		},
	)
}

func ResponseError(ctx *gin.Context, statusCode int, message string) {
	ctx.JSON(
		statusCode,
		structs.ApiResponse{
			Success: true,
			Message: message,
		},
	)
}
