package routes

import (
    "backend/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/register", controllers.Register)
    r.POST("/products", controllers.CreateProduct)
    r.POST("/orders", controllers.CreateOrder)
    r.GET("/orderss", controllers.GetOrders)
    return r
}
