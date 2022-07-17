package controller

import (
	"net/http"

	"shape-api/form"
	"shape-api/repo"

	"github.com/gin-gonic/gin"
)

//UserController ...
type TriangleController struct{}

var triangleForm = new(form.CreateTriangleForm)
var triangleRepo = new(repo.TriangleRepo)

//Create triangle ...
func (u TriangleController) CreateTriangle(c *gin.Context) {
	var CreateRecForm form.CreateTriangleForm

	if err := c.ShouldBindJSON(&CreateRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	user, err := triangleRepo.CreateTriangle(CreateRecForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully create triangle", "user": user})
}

//Get area triangle ...
func (u TriangleController) GetListTriangle(c *gin.Context) {
	var listRecForm form.GetListTriangleForm

	if err := c.ShouldBindJSON(&listRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	area, err := triangleRepo.GetListTriangle(listRecForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully get list triangle", "area": area})
}

//Update triangle ...
func (u TriangleController) UpdateTriangle(c *gin.Context) {
	var updateRecForm form.UpdateTriangleForm

	if err := c.ShouldBindJSON(&updateRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	user, err := triangleRepo.UpdateTriangle(updateRecForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully update triangle", "user": user})
}

//Delete triangle ...
func (u TriangleController) DeleteTriangle(c *gin.Context) {
	var deleteRecForm form.DeleteTriangleForm

	if err := c.ShouldBindJSON(&deleteRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	user, err := triangleRepo.DeleteTriangle(deleteRecForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully delete triangle", "user": user})
}

//Get area triangle ...
func (u TriangleController) GetArea(c *gin.Context) {
	var areaRecForm form.AreaTriangleForm

	if err := c.ShouldBindJSON(&areaRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	area, err := triangleRepo.GetArea(areaRecForm.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully get triangle area", "area": area})
}

//Get area triangle ...
func (u TriangleController) GetPerimeter(c *gin.Context) {
	var areaRecForm form.AreaTriangleForm

	if err := c.ShouldBindJSON(&areaRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	area, err := triangleRepo.GetPerimeter(areaRecForm.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully get triangle perimeter", "area": area})
}
