package db

import (
	"go-server/pkg/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBHandler struct {
	gDB *gorm.DB
}

func NewAndConnectGorm(dsn string) (*DBHandler, error) {
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	gormDB.AutoMigrate(&model.Board{})

	dbHandler := &DBHandler{
		gDB: gormDB,
	}

	return dbHandler, err
}
