package form

type CreateSquareForm struct {
	Name string `form:"name" json:"name" binding:"required"`
	Side int    `form:"side" json:"side" binding:"required"`
}

type GetListSquareForm struct {
	Limit  int `form:"limit" json:"limit"`
	Offset int `form:"offset" json:"offset"`
}

type UpdateSquareForm struct {
	Name string `form:"name" json:"name" binding:"required"`
	Side int    `form:"side" json:"side"`
}

type DeleteSquareForm struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type AreaSquareForm struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type PerimeterSquareForm struct {
	Name string `form:"name" json:"name" binding:"required"`
}
