package db

import (
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

type Config struct {
	DSN          string        // data source name.
	MaxOpenConns int           // pool
	MaxIdleConns int           // pool
	MaxIdleTime  time.Duration // connect max idle time.
	MaxLifetime  time.Duration //
}

func NewMysqlConnection(config *Config) (*gorm.DB, error) {
	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      false,
			LogLevel:                  logger.Info,
		},
	)

	db, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{
		Logger:                                   dbLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(config.MaxLifetime)
	sqlDB.SetConnMaxIdleTime(config.MaxIdleTime)
	return db, nil
}

func InitDB(dsn string) error {
	config := &Config{
		DSN:          dsn,
		MaxOpenConns: 50,
		MaxIdleConns: 10,
		MaxIdleTime:  time.Hour,
		MaxLifetime:  time.Hour,
	}
	db, err := NewMysqlConnection(config)
	if err != nil {
		return err
	}
	DB = db
	return nil
}
