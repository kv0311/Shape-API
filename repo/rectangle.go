package repo

import (
	"context"
	"errors"
	"shape-api/constant"
	"shape-api/db"
	"shape-api/form"
	"shape-api/logger"
	"shape-api/model"
	"shape-api/utils"
	"strconv"
)

type RectangelRepo struct{}

/* Rectangle */
func (r RectangelRepo) CreateRectangle(createRecForm form.CreateRectangleForm) (rectangle model.Rectangle, err error) {
	var countRectangle int64
	if err := db.PgClient.Table("rectangle").Model(&model.User{}).Where("name", createRecForm.Name).Count(&countRectangle).Error; err != nil {
		return rectangle, err
	}
	if countRectangle > 0 {
		return rectangle, errors.New("rectangle already exists")
	}

	rectangle.Name = createRecForm.Name
	rectangle.Length = createRecForm.Length
	rectangle.Width = createRecForm.Width

	result := db.PgClient.Table("rectangle").Create(&rectangle)
	if result.Error != nil {
		return rectangle, errors.New("something went wrong, please try again later: " + result.Error.Error())
	}

	/* Calculate Area and Perimeter and save redis. Help user can get as fast as posible */

	perimeter := utils.CalculatePerimeterRectangle(createRecForm.Width, createRecForm.Length)
	area := utils.CalculateAreaRectangle(createRecForm.Width, createRecForm.Length)

	err = updateCacheInfo(rectangle.Name, constant.Rectangle, perimeter, area)
	if err != nil {
		return rectangle, err
	}
	/* */
	return rectangle, nil
}

func (r RectangelRepo) GetListRectangle(listRectangleForm form.GetListRectangleForm) (rectangle []model.Rectangle, err error) {

	if err := db.PgClient.Table("rectangle").Model(&model.Rectangle{}).Offset(listRectangleForm.Offset).Limit(listRectangleForm.Limit).Find(&rectangle).Error; err != nil {
		return rectangle, errors.New("something went wrong, please try again later: " + err.Error())
	}

	return rectangle, nil
}

func (r RectangelRepo) UpdateRectangle(updateRectangleForm form.UpdateRectangleForm) (rectangle model.Rectangle, err error) {

	if err := db.PgClient.Table("rectangle").Model(&model.Rectangle{}).Where("name", updateRectangleForm.Name).Updates(updateRectangleForm).Error; err != nil {
		return rectangle, errors.New("something went wrong, please try again later: " + err.Error())
	}

	if err := db.PgClient.Table("rectangle").Model(&model.Rectangle{}).Where("name", updateRectangleForm.Name).First(&rectangle).Error; err != nil {
		return rectangle, err
	}

	/* Calculate Area and Perimeter and save redis. Help user can get as fast as posible */
	perimeter := utils.CalculatePerimeterRectangle(rectangle.Width, rectangle.Length)
	area := utils.CalculateAreaRectangle(rectangle.Width, rectangle.Length)

	err = updateCacheInfo(rectangle.Name, constant.Rectangle, perimeter, area)
	if err != nil {
		logger.Logger.Error("Fail to update cache info: " + err.Error())
		return rectangle, err
	}
	/* */

	return rectangle, nil
}

func (r RectangelRepo) DeleteRectangle(deleteRectangleForm form.DeleteRectangleForm) (rectangle model.Rectangle, err error) {
	if err := db.PgClient.Table("rectangle").Model(&model.Rectangle{}).Where("name", deleteRectangleForm.Name).Delete(&rectangle).Error; err != nil {
		return rectangle, errors.New("something went wrong, please try again later: " + err.Error())
	}
	deleteCacheInfo(deleteRectangleForm.Name, constant.Rectangle)
	return rectangle, nil
}

func (r RectangelRepo) GetArea(name string) (area int, err error) {
	areaStr, err := db.RedisClient.Get(context.Background(), utils.BuildShapeKey(constant.Rectangle, name, constant.Area)).Result()
	if err != nil {
		return 0, err
	}
	area, _ = strconv.Atoi(areaStr)
	return area, nil
}

func (r RectangelRepo) GetPerimeter(name string) (perimeter int, err error) {
	perimeterStr, err := db.RedisClient.Get(context.Background(), utils.BuildShapeKey(constant.Rectangle, name, constant.Perimeter)).Result()
	if err != nil {
		return 0, err
	}
	perimeter, _ = strconv.Atoi(perimeterStr)
	return perimeter, nil
}
