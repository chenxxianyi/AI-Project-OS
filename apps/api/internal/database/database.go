package database

import (
	"fmt"

	"github.com/ai-project-os/api/internal/config"
	"github.com/ai-project-os/api/internal/models"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init(cfg *config.Config) error {
	var logLevel gormlogger.LogLevel
	switch cfg.LogLevel {
	case "debug":
		logLevel = gormlogger.Info
	default:
		logLevel = gormlogger.Warn
	}

	db, err := gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{
		Logger: gormlogger.Default.LogMode(logLevel),
	})
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	if err := db.AutoMigrate(
		&models.User{},
		&models.Project{},
		&models.Prompt{},
		&models.PromptVersion{},
		&models.ProjectDoc{},
		&models.Template{},
		&models.AIGeneration{},
	); err != nil {
		return fmt.Errorf("failed to auto migrate: %w", err)
	}

	DB = db
	zap.L().Info("Database connected and migrated successfully")
	return nil
}

func GetDB() *gorm.DB {
	return DB
}
