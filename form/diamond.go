package form

type CreateDiamondForm struct {
	Name      string `form:"name" json:"name" binding:"required"`
	Side      int    `form:"side" json:"side" binding:"required"`
	Diagonal1 int    `form:"diagonal1" json:"diagonal1" binding:"required"`
	Diagonal2 int    `form:"diagonal2" json:"diagonal2" binding:"required"`
}

type GetListDiamondForm struct {
	Limit  int `form:"limit" json:"limit"`
	Offset int `form:"offset" json:"offset"`
}

type UpdateDiamondForm struct {
	Name      string `form:"name" json:"name" binding:"required"`
	Diagonal1 int    `form:"diagonal1" json:"diagonal1"`
	Diagonal2 int    `form:"diagonal2" json:"diagonal2"`
	Side      int    `form:"side" json:"side"`
}

type DeleteDiamondForm struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type AreaDiamondForm struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type PerimeterDiamondForm struct {
	Name string `form:"name" json:"name" binding:"required"`
}
