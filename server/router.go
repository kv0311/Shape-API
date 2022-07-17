package server

import (
	"shape-api/controller"
	"shape-api/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controller.HealthController)

	router.GET("/health", health.Status)

	userGroup := router.Group("/user")
	{
		user := new(controller.UserController)
		userGroup.POST("/register", user.Register)
		userGroup.POST("/login", user.Login)
	}

	shapeGroup := router.Group("/shape")
	shapeGroup.Use(middleware.AuthMiddleware())
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

	{
		diamondGroup := shapeGroup.Group("/diamond")
		diamond := new(controller.DiamondController)
		diamondGroup.POST("/create", diamond.CreateDiamond)
		diamondGroup.POST("/getList", diamond.GetListDiamond)
		diamondGroup.POST("/update", diamond.UpdateDiamond)
		diamondGroup.POST("/delete", diamond.DeleteDiamond)
		diamondGroup.POST("/getInfo/area", diamond.GetArea)
		diamondGroup.POST("/getInfo/perimeter", diamond.GetPerimeter)
	}

	{
		triangleGroup := shapeGroup.Group("/triangle")
		triangle := new(controller.TriangleController)
		triangleGroup.POST("/create", triangle.CreateTriangle)
		triangleGroup.POST("/getList", triangle.GetListTriangle)
		triangleGroup.POST("/update", triangle.UpdateTriangle)
		triangleGroup.POST("/delete", triangle.DeleteTriangle)
		triangleGroup.POST("/getInfo/area", triangle.GetArea)
		triangleGroup.POST("/getInfo/perimeter", triangle.GetPerimeter)
	}
	return router
}
