package db

import (
	"fmt"
	"tasoryx/config"
	"tasoryx/pkg/adapters/storage/types"
	"tasoryx/pkg/logger"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPSQLConnection(cfg config.Database) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Get().Panic("could not connect to database : " + err.Error())
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Get().Panic("could not get sql.DB from gorm.DB: " + err.Error())
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	if err = migration(db); err != nil {
		logger.Get().Panic("could not migrate data : " + err.Error())
	}
	logger.Get().Info("database migration completed successfully")
	return db
}

func migration(db *gorm.DB) error {
	return db.AutoMigrate(types.Task{}, types.User{})
}
