package controller

import (
	"net/http"

	"shape-api/form"
	"shape-api/repo"

	"github.com/gin-gonic/gin"
)

//UserController ...
type SquareController struct{}

var squareForm = new(form.CreateSquareForm)

var squareRepo = new(repo.SquareRepo)

//Create square ...
func (u SquareController) CreateSquare(c *gin.Context) {
	var CreateRecForm form.CreateSquareForm

	if err := c.ShouldBindJSON(&CreateRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	user, err := squareRepo.CreateSquare(CreateRecForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully create square", "user": user})
}

//Get area square ...
func (u SquareController) GetListSquare(c *gin.Context) {
	var listRecForm form.GetListSquareForm

	if err := c.ShouldBindJSON(&listRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	area, err := squareRepo.GetListSquare(listRecForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully get list square", "area": area})
}

//Update square ...
func (u SquareController) UpdateSquare(c *gin.Context) {
	var updateRecForm form.UpdateSquareForm

	if err := c.ShouldBindJSON(&updateRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	user, err := squareRepo.UpdateSquare(updateRecForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully update square", "user": user})
}

//Delete square ...
func (u SquareController) DeleteSquare(c *gin.Context) {
	var deleteRecForm form.DeleteSquareForm

	if err := c.ShouldBindJSON(&deleteRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	user, err := squareRepo.DeleteSquare(deleteRecForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully delete square", "user": user})
}

//Get area square ...
func (u SquareController) GetArea(c *gin.Context) {
	var areaRecForm form.AreaSquareForm

	if err := c.ShouldBindJSON(&areaRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	area, err := squareRepo.GetArea(areaRecForm.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully get square area", "area": area})
}

//Get area square ...
func (u SquareController) GetPerimeter(c *gin.Context) {
	var areaRecForm form.AreaSquareForm

	if err := c.ShouldBindJSON(&areaRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	area, err := squareRepo.GetPerimeter(areaRecForm.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully get square perimeter", "area": area})
}
