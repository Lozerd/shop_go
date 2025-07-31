package postgres

import (
	"fmt"

	"github.com/lozerd/shop_go/domain/product"
	"github.com/lozerd/shop_go/infrastructure/logging"
	"github.com/lozerd/shop_go/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var logger = logging.NewLogger("postgres", true)

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
		logger.Error().Err(err).Msg("failed to connect to database")
		return nil, err
	}
	return db, nil
}

func Ping() {
	db, err := GetDB()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to ping database")
	}

	sqlDB, err := db.DB()
	defer sqlDB.Close()

	if err != nil {
		logger.Fatal().Err(err).Msg("failed to ping database")
	}

	if err := sqlDB.Ping(); err != nil {
		logger.Fatal().Err(err).Msg("failed to ping database")
	}
}

func Migrate() {
	db, err := GetDB()
	if err != nil {
		logger.Fatal().Err(err).Msg("database migration couln't establish connection")
	}

	err = db.AutoMigrate(&product.Product{})
	if err != nil {
		logger.Fatal().Err(err).Msg("database migration failed")
	}
	logger.Info().Msg("database migration done")
}
