package form

type CreateTriangleForm struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Side1  int    `form:"side1" json:"side1" binding:"required"`
	Side2  int    `form:"side2" json:"side2" binding:"required"`
	Height int    `form:"height" json:"height" binding:"required"`
	Base   int    `form:"base" json:"base" binding:"required"`
}
