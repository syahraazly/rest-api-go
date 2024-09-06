package router

import (
	"rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	orderRoutes := r.Group("/orders")
	{
		orderRoutes.POST("/", controllers.CreateOrder)
		orderRoutes.GET("/", controllers.GetAllOrders)
		orderRoutes.GET("/:id", controllers.GetOrderById)
		orderRoutes.PUT("/:id", controllers.UpdateOrder, )
		orderRoutes.DELETE("/:id", controllers.DeleteOrder)
	}

	return r
}
