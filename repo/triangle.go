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

type TriangleRepo struct{}

/* Triangle */
func (r TriangleRepo) CreateTriangle(createTriangleForm form.CreateTriangleForm) (triangle model.Triangle, err error) {
	var countTriangle int64
	if err := db.PgClient.Table("triangle").Model(&model.User{}).Where("name", createTriangleForm.Name).Count(&countTriangle).Error; err != nil {
		return triangle, err
	}
	if countTriangle > 0 {
		return triangle, errors.New("triangle already exists")
	}

	triangle.Name = createTriangleForm.Name
	triangle.Side1 = createTriangleForm.Side1
	triangle.Side2 = createTriangleForm.Side2
	triangle.Base = createTriangleForm.Base
	triangle.Height = createTriangleForm.Height

	result := db.PgClient.Table("triangle").Create(&triangle)
	if result.Error != nil {
		return triangle, errors.New("something went wrong, please try again later: " + result.Error.Error())
	}

	/* Calculate Area and Perimeter and save redis. Help user can get as fast as posible */

	perimeter := utils.CalculatePerimeterTriangle(createTriangleForm.Side1, createTriangleForm.Side2, createTriangleForm.Base)
	area := utils.CalculateAreaTriangle(createTriangleForm.Base, createTriangleForm.Height)

	err = updateCacheInfo(triangle.Name, constant.Triangle, perimeter, area)
	if err != nil {
		return triangle, err
	}
	/* */
	return triangle, nil
}

func (r TriangleRepo) GetListTriangle(listTriangleForm form.GetListTriangleForm) (triangle []model.Triangle, err error) {

	if err := db.PgClient.Table("triangle").Model(&model.Triangle{}).Offset(listTriangleForm.Offset).Limit(listTriangleForm.Limit).Find(&triangle).Error; err != nil {
		return triangle, errors.New("something went wrong, please try again later: " + err.Error())
	}

	return triangle, nil
}

func (r TriangleRepo) UpdateTriangle(updateTriangleForm form.UpdateTriangleForm) (triangle model.Triangle, err error) {

	if err := db.PgClient.Table("triangle").Model(&model.Triangle{}).Where("name", updateTriangleForm.Name).Updates(updateTriangleForm).Error; err != nil {
		return triangle, errors.New("something went wrong, please try again later: " + err.Error())
	}

	if err := db.PgClient.Table("triangle").Model(&model.Triangle{}).Where("name", updateTriangleForm.Name).First(&triangle).Error; err != nil {
		return triangle, err
	}

	/* Calculate Area and Perimeter and save redis. Help user can get as fast as posible */
	perimeter := utils.CalculatePerimeterTriangle(triangle.Side1, triangle.Side2, triangle.Base)
	area := utils.CalculateAreaTriangle(triangle.Base, triangle.Height)

	err = updateCacheInfo(triangle.Name, constant.Triangle, perimeter, area)
	if err != nil {
		logger.Logger.Error("Fail to update cache info: " + err.Error())
		return triangle, err
	}
	/* */

	return triangle, nil
}

func (r TriangleRepo) DeleteTriangle(deleteTriangleForm form.DeleteTriangleForm) (triangle model.Triangle, err error) {
	if err := db.PgClient.Table("triangle").Model(&model.Triangle{}).Where("name", deleteTriangleForm.Name).Delete(&triangle).Error; err != nil {
		return triangle, errors.New("something went wrong, please try again later: " + err.Error())
	}
	deleteCacheInfo(deleteTriangleForm.Name, constant.Triangle)
	return triangle, nil
}

func (r TriangleRepo) GetArea(name string) (area int, err error) {
	areaStr, err := db.RedisClient.Get(context.Background(), utils.BuildShapeKey(constant.Triangle, name, constant.Area)).Result()
	if err != nil {
		return 0, err
	}
	area, _ = strconv.Atoi(areaStr)
	return area, nil
}

func (r TriangleRepo) GetPerimeter(name string) (perimeter int, err error) {
	perimeterStr, err := db.RedisClient.Get(context.Background(), utils.BuildShapeKey(constant.Triangle, name, constant.Perimeter)).Result()
	if err != nil {
		return 0, err
	}
	perimeter, _ = strconv.Atoi(perimeterStr)
	return perimeter, nil
}
