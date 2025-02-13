package bootstrap

import (
	"valorx-auth/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitiatePostgreSQL(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.Postgre.URL))
	if err != nil {
		return db, err
	}

	psqlDB, err := db.DB()
	if err != nil {
		return db, err
	}

	psqlDB.SetMaxIdleConns(cfg.Postgre.MaxIdleConn)
	psqlDB.SetMaxOpenConns(cfg.Postgre.MaxOpenConn)

	return db, nil
}
