package model

type Rectangle struct {
	Name   string `form:"name" json:"name" binding:"required"` //fullName rule is in validator.go
	Length int    `form:"length" json:"length" binding:"required"`
	Width  int    `form:"width" json:"width" binding:"required"`
}

type Triangle struct {
	Name   string `form:"name" json:"name" binding:"required"` //fullName rule is in validator.go
	Side1  int    `form:"side1" json:"side1" binding:"required"`
	Side2  int    `form:"side2" json:"side2" binding:"required"`
	Height int    `form:"height" json:"height" binding:"required"`
	Base   int    `form:"base" json:"base" binding:"required"`
}

type Square struct {
	Name string `form:"name" json:"name" binding:"required"` //fullName rule is in validator.go
	Side int    `form:"side" json:"side" binding:"required"`
}

type Diamond struct {
	Name      string `form:"name" json:"name" binding:"required"` //fullName rule is in validator.go
	Side      int    `form:"side" json:"side" binding:"required"`
	Diagonal1 int    `form:"diagonal1" json:"diagonal1" binding:"required"`
	Diagonal2 int    `form:"diagonal2" json:"diagonal2" binding:"required"`
}
