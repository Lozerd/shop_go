package postgres

import (
	"fmt"

	"github.com/lozerd/shop_go/internal/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDSN() string {
	cfg := config.GetConfig()
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)
}

func GetDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(GetDSN()), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to database")
		return nil, err
	}
	return db, nil
}

func Ping() {
	db, err := GetDB()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to ping database")
	}

	sqlDB, err := db.DB()
	defer sqlDB.Close()

	if err != nil {
		log.Fatal().Err(err).Msg("failed to ping database")
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal().Err(err).Msg("failed to ping database")
	}
}
