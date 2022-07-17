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

type SquareRepo struct{}

/* Square */
func (r SquareRepo) CreateSquare(createSquareForm form.CreateSquareForm) (square model.Square, err error) {
	var countSquare int64
	if err := db.PgClient.Table("square").Model(&model.User{}).Where("name", createSquareForm.Name).Count(&countSquare).Error; err != nil {
		return square, err
	}
	if countSquare > 0 {
		return square, errors.New("square already exists")
	}

	square.Name = createSquareForm.Name
	square.Side = createSquareForm.Side

	result := db.PgClient.Table("square").Create(&square)
	if result.Error != nil {
		return square, errors.New("something went wrong, please try again later: " + result.Error.Error())
	}

	/* Calculate Area and Perimeter and save redis. Help user can get as fast as posible */

	perimeter := utils.CalculatePerimeterSquare(createSquareForm.Side)
	area := utils.CalculateAreaSquare(createSquareForm.Side)

	err = updateCacheInfo(square.Name, constant.Square, perimeter, area)
	if err != nil {
		return square, err
	}
	/* */
	return square, nil
}

func (r SquareRepo) GetListSquare(listSquareForm form.GetListSquareForm) (square []model.Square, err error) {

	if err := db.PgClient.Table("square").Model(&model.Square{}).Offset(listSquareForm.Offset).Limit(listSquareForm.Limit).Find(&square).Error; err != nil {
		return square, errors.New("something went wrong, please try again later: " + err.Error())
	}

	return square, nil
}

func (r SquareRepo) UpdateSquare(updateSquareForm form.UpdateSquareForm) (square model.Square, err error) {

	if err := db.PgClient.Table("square").Model(&model.Square{}).Where("name", updateSquareForm.Name).Updates(updateSquareForm).Error; err != nil {
		return square, errors.New("something went wrong, please try again later: " + err.Error())
	}

	if err := db.PgClient.Table("square").Model(&model.Square{}).Where("name", updateSquareForm.Name).First(&square).Error; err != nil {
		return square, err
	}

	/* Calculate Area and Perimeter and save redis. Help user can get as fast as posible */
	perimeter := utils.CalculatePerimeterSquare(square.Side)
	area := utils.CalculateAreaSquare(square.Side)

	err = updateCacheInfo(square.Name, constant.Square, perimeter, area)
	if err != nil {
		logger.Logger.Error("Fail to update cache info: " + err.Error())
		return square, err
	}
	/* */

	return square, nil
}

func (r SquareRepo) DeleteSquare(deleteSquareForm form.DeleteSquareForm) (square model.Square, err error) {
	if err := db.PgClient.Table("square").Model(&model.Square{}).Where("name", deleteSquareForm.Name).Delete(&square).Error; err != nil {
		return square, errors.New("something went wrong, please try again later: " + err.Error())
	}
	deleteCacheInfo(deleteSquareForm.Name, constant.Square)
	return square, nil
}

func (r SquareRepo) GetArea(name string) (area int, err error) {
	areaStr, err := db.RedisClient.Get(context.Background(), utils.BuildShapeKey(constant.Square, name, constant.Area)).Result()
	if err != nil {
		return 0, err
	}
	area, _ = strconv.Atoi(areaStr)
	return area, nil
}

func (r SquareRepo) GetPerimeter(name string) (perimeter int, err error) {
	perimeterStr, err := db.RedisClient.Get(context.Background(), utils.BuildShapeKey(constant.Square, name, constant.Perimeter)).Result()
	if err != nil {
		return 0, err
	}
	perimeter, _ = strconv.Atoi(perimeterStr)
	return perimeter, nil
}
