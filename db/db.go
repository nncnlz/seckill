package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env 文件未找到，使用系统环境变量")
	}
}

func InitTiDBConfig() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local&tls=%s",
		os.Getenv("TIDB_USER"),
		os.Getenv("TIDB_PASSWORD"),
		os.Getenv("TIDB_HOST"),
		os.Getenv("TIDB_PORT"),
		os.Getenv("TIDB_DB_NAME"),
		os.Getenv("USE_SSL"),
	)

}

func OpenDB() (*gorm.DB, error) {
	LoadEnv()
	dsn := InitTiDBConfig()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, fmt.Errorf("failed to open DB: %w", err)
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
