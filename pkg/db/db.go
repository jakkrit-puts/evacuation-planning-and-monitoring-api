package db

import (
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jakkrit-puts/evacuation-planning-and-monitoring-api/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.EnvConfig, DBMigrator func(db *gorm.DB) error) *gorm.DB {

	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Bangkok`, config.DBHost, config.DBUsername, config.DBPassword, config.DBName, config.DBPort, config.DBSSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// DryRun: false,
	})

	if err != nil {
		log.Fatalf("Connect Database Fail !!!: %e", err)
	}

	log.Info("Connected to the database...")

	if err := DBMigrator(db); err != nil {
		log.Fatalf("Unable to migrate tables !!!: %e", err)
	}

	return db
}
