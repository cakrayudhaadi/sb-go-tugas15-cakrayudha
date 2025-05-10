package routers

import (
	"tugas15/services"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	router.POST("/bioskop", services.CreateBioskop)
	router.GET("/bioskop", services.GetAllBioskop)
	router.GET("/bioskop/:id", services.GetBioskop)
	router.PUT("/bioskop/:id", services.UpdateBioskop)
	router.DELETE("/bioskop/:id", services.DeleteBioskop)

	router.Run(":8080")
}
