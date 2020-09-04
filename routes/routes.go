package routes

import (
	"github.com/103cuong/gorm_kit/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	group := router.Group("/cats")
	group.GET("/", controllers.GetCats)
	group.GET("/:id", controllers.GetCatByID)
	group.POST("/", controllers.CreateCat)
	group.PUT("/:id", controllers.UpdateCat)
	group.DELETE("/:id", controllers.DeleteCat)
	return router
}
