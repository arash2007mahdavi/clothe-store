package database

import (
	"fmt"
	"store/src/configs"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbClient *gorm.DB

func InitDB(cfg *configs.Config) error{
	var err error

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Tehran",
		cfg.Postgres.Host, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DbName, cfg.Postgres.Port, cfg.Postgres.SslMode,
	)

	DbClient, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	database, err := DbClient.DB()
	if err != nil {
		return err
	}

	err = database.Ping()
	if err != nil {
		return err
	}
	database.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	database.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	database.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	return nil
}

func GetDB() *gorm.DB {
	return DbClient
}

func CloseDB() {
	con, _ := DbClient.DB()
	con.Close()
}