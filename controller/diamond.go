package controller

import (
	"net/http"

	"shape-api/form"
	"shape-api/repo"

	"github.com/gin-gonic/gin"
)

//UserController ...
type DiamondController struct{}

var diamondForm = new(form.CreateDiamondForm)

var diamondRepo = new(repo.DiamondRepo)

//Create diamond ...
func (u DiamondController) CreateDiamond(c *gin.Context) {
	var CreateRecForm form.CreateDiamondForm

	if err := c.ShouldBindJSON(&CreateRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	user, err := diamondRepo.CreateDiamond(CreateRecForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully create diamond", "user": user})
}

//Get area diamond ...
func (u DiamondController) GetListDiamond(c *gin.Context) {
	var listRecForm form.GetListDiamondForm

	if err := c.ShouldBindJSON(&listRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	area, err := diamondRepo.GetListDiamond(listRecForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully get list diamond", "area": area})
}

//Update diamond ...
func (u DiamondController) UpdateDiamond(c *gin.Context) {
	var updateRecForm form.UpdateDiamondForm

	if err := c.ShouldBindJSON(&updateRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	user, err := diamondRepo.UpdateDiamond(updateRecForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully update diamond", "user": user})
}

//Delete diamond ...
func (u DiamondController) DeleteDiamond(c *gin.Context) {
	var deleteRecForm form.DeleteDiamondForm

	if err := c.ShouldBindJSON(&deleteRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	user, err := diamondRepo.DeleteDiamond(deleteRecForm)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully delete diamond", "user": user})
}

//Get area diamond ...
func (u DiamondController) GetArea(c *gin.Context) {
	var areaRecForm form.AreaDiamondForm

	if err := c.ShouldBindJSON(&areaRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	area, err := diamondRepo.GetArea(areaRecForm.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully get diamond area", "area": area})
}

//Get area diamond ...
func (u DiamondController) GetPerimeter(c *gin.Context) {
	var areaRecForm form.AreaDiamondForm

	if err := c.ShouldBindJSON(&areaRecForm); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	area, err := diamondRepo.GetPerimeter(areaRecForm.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully get diamond perimeter", "area": area})
}
