package repo

import (
	"context"
	"shape-api/constant"
	"shape-api/db"
	"shape-api/utils"
	"strconv"
	"time"
)

func updateCacheInfo(name string, shapeType string, perimeter int, area int) (err error) {
	cacheTimeRedis := time.Hour * 24 * 30

	errSavePerimeter := db.RedisClient.Set(context.Background(), utils.BuildShapeKey(shapeType, name, constant.Perimeter), strconv.Itoa(perimeter), cacheTimeRedis).Err()
	if errSavePerimeter != nil {
		return errSavePerimeter
	}
	errSaveArea := db.RedisClient.Set(context.Background(), utils.BuildShapeKey(shapeType, name, constant.Area), area, cacheTimeRedis).Err()
	if errSaveArea != nil {
		return errSaveArea
	}
	return nil
}

func deleteCacheInfo(name string, shapeType string) {
	db.RedisClient.Del(context.Background(), utils.BuildShapeKey(shapeType, name, constant.Perimeter))
	db.RedisClient.Del(context.Background(), utils.BuildShapeKey(shapeType, name, constant.Area))
}
