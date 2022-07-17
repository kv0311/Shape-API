package form

type CreateDiamondForm struct {
	Name      string `form:"name" json:"name" binding:"required"`
	Side      int    `form:"side" json:"side" binding:"required"`
	Diagonal1 int    `form:"diagonal1" json:"diagonal1" binding:"required"`
	Diagonal2 int    `form:"diagonal2" json:"diagonal2" binding:"required"`
}
