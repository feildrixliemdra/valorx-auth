package migration

import (
	"valorx-auth/internal/bootstrap"
	"valorx-auth/pkg/dbmigration"
)

func MigrateDatabase() {
	cfg := bootstrap.NewConfig()

	dbmigration.DatabaseMigration(cfg)
}
