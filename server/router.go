package server

import (
	"shape-api/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controller.HealthController)

	router.GET("/health", health.Status)
	// router.Use(middleware.AuthMiddleware())

	userGroup := router.Group("/user")
	{
		user := new(controller.UserController)
		userGroup.POST("/register", user.Register)
		userGroup.POST("/login", user.Login)

		// router.Use(middleware.AuthMiddleware())

	}

	shapeGroup := router.Group("/shape")
	// shapeGroup.Use(middleware.AuthMiddleware())
	{
		rectangleGroup := shapeGroup.Group("/rectangle")

		rectangle := new(controller.RectangleController)
		rectangleGroup.POST("/create", rectangle.CreateRectangle)
		rectangleGroup.POST("/getList", rectangle.GetListRectangle)
		rectangleGroup.POST("/update", rectangle.UpdateRectangle)
		rectangleGroup.POST("/delete", rectangle.DeleteRectangle)
		rectangleGroup.POST("/getInfo/area", rectangle.GetArea)
		rectangleGroup.POST("/getInfo/perimeter", rectangle.GetPerimeter)

	}

	{
		squareGroup := shapeGroup.Group("/square")

		square := new(controller.SquareController)
		squareGroup.POST("/create", square.CreateSquare)
		squareGroup.POST("/getList", square.GetListSquare)
		squareGroup.POST("/update", square.UpdateSquare)
		squareGroup.POST("/delete", square.DeleteSquare)
		squareGroup.POST("/getInfo/area", square.GetArea)
		squareGroup.POST("/getInfo/perimeter", square.GetPerimeter)

	}
	return router
}
