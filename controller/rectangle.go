package controller

import (
	"net/http"

	"shape-api/form"
	"shape-api/repo"

	"github.com/gin-gonic/gin"
)

//UserController ...
type RectangleController struct{}

var rectangleForm = new(form.CreateRectangleForm)
var rectangleRepo = new(repo.RectangelRepo)

//Create rectangle ...
func (u RectangleController) CreateRectangle(c *gin.Context) {
	var CreateRecForm form.CreateRectangleForm

	if err := c.ShouldBindJSON(&CreateRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	user, err := rectangleRepo.CreateRectangle(CreateRecForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully create rectangle", "user": user})
}

//Get area rectangle ...
func (u RectangleController) GetListRectangle(c *gin.Context) {
	var listRecForm form.GetListRectangleForm

	if err := c.ShouldBindJSON(&listRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	area, err := rectangleRepo.GetListRectangle(listRecForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully get list rectangle", "area": area})
}

//Update rectangle ...
func (u RectangleController) UpdateRectangle(c *gin.Context) {
	var updateRecForm form.UpdateRectangleForm

	if err := c.ShouldBindJSON(&updateRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	user, err := rectangleRepo.UpdateRectangle(updateRecForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully update rectangle", "user": user})
}

//Delete rectangle ...
func (u RectangleController) DeleteRectangle(c *gin.Context) {
	var deleteRecForm form.DeleteRectangleForm

	if err := c.ShouldBindJSON(&deleteRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	user, err := rectangleRepo.DeleteRectangle(deleteRecForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully delete rectangle", "user": user})
}

//Get area rectangle ...
func (u RectangleController) GetArea(c *gin.Context) {
	var areaRecForm form.AreaRectangleForm

	if err := c.ShouldBindJSON(&areaRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	area, err := rectangleRepo.GetArea(areaRecForm.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully get rectangle area", "area": area})
}

//Get area rectangle ...
func (u RectangleController) GetPerimeter(c *gin.Context) {
	var areaRecForm form.AreaRectangleForm

	if err := c.ShouldBindJSON(&areaRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	area, err := rectangleRepo.GetPerimeter(areaRecForm.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully get rectangle perimeter", "area": area})
}
