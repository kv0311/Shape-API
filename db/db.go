package db

import (
	"fmt"
	"shape-api/config"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PgClient *gorm.DB

func InitPG() error {
	port, _ := strconv.Atoi(config.GetConfig("pg.port"))
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname = %s sslmode=disable",
		config.GetConfig("pg.host"),
		port,
		config.GetConfig("pg.user"),
		config.GetConfig("pg.password"),
		config.GetConfig("pg.name"))
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	PgClient = db
	return err
}
