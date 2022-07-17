package form

type CreateRectangleForm struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Length int    `form:"length" json:"length" binding:"required"`
	Width  int    `form:"width" json:"width" binding:"required"`
}

type GetListRectangleForm struct {
	Limit  int `form:"limit" json:"limit"`
	Offset int `form:"offset" json:"offset"`
}

type UpdateRectangleForm struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Length int    `form:"length" json:"length"`
	Width  int    `form:"width" json:"width"`
}

type DeleteRectangleForm struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type AreaRectangleForm struct {
	Name string `form:"name" json:"name" binding:"required"`
}

type PerimeterRectangleForm struct {
	Name string `form:"name" json:"name" binding:"required"`
}
