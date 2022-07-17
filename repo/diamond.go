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

type DiamondRepo struct{}

/* Diamond */
func (r DiamondRepo) CreateDiamond(createDiamondForm form.CreateDiamondForm) (diamond model.Diamond, err error) {
	var countDiamond int64
	if err := db.PgClient.Table("diamond").Model(&model.User{}).Where("name", createDiamondForm.Name).Count(&countDiamond).Error; err != nil {
		return diamond, err
	}
	if countDiamond > 0 {
		return diamond, errors.New("diamond already exists")
	}

	diamond.Name = createDiamondForm.Name
	diamond.Side = createDiamondForm.Side

	result := db.PgClient.Table("diamond").Create(&diamond)
	if result.Error != nil {
		return diamond, errors.New("something went wrong, please try again later: " + result.Error.Error())
	}

	/* Calculate Area and Perimeter and save redis. Help user can get as fast as posible */

	perimeter := utils.CalculatePerimeterDiamond(createDiamondForm.Side)
	area := utils.CalculateAreaDiamond(createDiamondForm.Diagonal1, createDiamondForm.Diagonal2)

	err = updateCacheInfo(diamond.Name, constant.Diamond, perimeter, area)
	if err != nil {
		return diamond, err
	}
	/* */
	return diamond, nil
}

func (r DiamondRepo) GetListDiamond(listDiamondForm form.GetListDiamondForm) (diamond []model.Diamond, err error) {

	if err := db.PgClient.Table("diamond").Model(&model.Diamond{}).Offset(listDiamondForm.Offset).Limit(listDiamondForm.Limit).Find(&diamond).Error; err != nil {
		return diamond, errors.New("something went wrong, please try again later: " + err.Error())
	}

	return diamond, nil
}

func (r DiamondRepo) UpdateDiamond(updateDiamondForm form.UpdateDiamondForm) (diamond model.Diamond, err error) {

	if err := db.PgClient.Table("diamond").Model(&model.Diamond{}).Where("name", updateDiamondForm.Name).Updates(updateDiamondForm).Error; err != nil {
		return diamond, errors.New("something went wrong, please try again later: " + err.Error())
	}

	if err := db.PgClient.Table("diamond").Model(&model.Diamond{}).Where("name", updateDiamondForm.Name).First(&diamond).Error; err != nil {
		return diamond, err
	}

	/* Calculate Area and Perimeter and save redis. Help user can get as fast as posible */
	perimeter := utils.CalculatePerimeterDiamond(diamond.Side)
	area := utils.CalculateAreaDiamond(diamond.Diagonal1, diamond.Diagonal2)

	err = updateCacheInfo(diamond.Name, constant.Diamond, perimeter, area)
	if err != nil {
		logger.Logger.Error("Fail to update cache info: " + err.Error())
		return diamond, err
	}
	/* */

	return diamond, nil
}

func (r DiamondRepo) DeleteDiamond(deleteDiamondForm form.DeleteDiamondForm) (diamond model.Diamond, err error) {
	if err := db.PgClient.Table("diamond").Model(&model.Diamond{}).Where("name", deleteDiamondForm.Name).Delete(&diamond).Error; err != nil {
		return diamond, errors.New("something went wrong, please try again later: " + err.Error())
	}
	deleteCacheInfo(deleteDiamondForm.Name, constant.Diamond)
	return diamond, nil
}

func (r DiamondRepo) GetArea(name string) (area int, err error) {
	areaStr, err := db.RedisClient.Get(context.Background(), utils.BuildShapeKey(constant.Diamond, name, constant.Area)).Result()
	if err != nil {
		return 0, err
	}
	area, _ = strconv.Atoi(areaStr)
	return area, nil
}

func (r DiamondRepo) GetPerimeter(name string) (perimeter int, err error) {
	perimeterStr, err := db.RedisClient.Get(context.Background(), utils.BuildShapeKey(constant.Diamond, name, constant.Perimeter)).Result()
	if err != nil {
		return 0, err
	}
	perimeter, _ = strconv.Atoi(perimeterStr)
	return perimeter, nil
}
