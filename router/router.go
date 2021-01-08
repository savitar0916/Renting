package router

import (
	Controllers "Renting/controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter -
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/house", Controllers.GetAllHouses)
	router.GET("/house/:houseid", Controllers.GetHouse)
	router.PUT("/house/:houseid", Controllers.UpdateHouse)
	router.DELETE("/house/:houseid", Controllers.DeleteHouse)
	router.POST("/Member", Controllers.AddMember)
	router.POST("//Member/Login", Controllers.Login)

	return router
}
