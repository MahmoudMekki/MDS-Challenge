package migration

import (
	"github.com/MahmoudMekki/MDS-task/database"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/rs/zerolog/log"
)

func RunMigration() {
	log.Info().Msg("Migration is started")
	var system []*gormigrate.Migration
	system = append(system, listProduct()...)
	system = append(system, listOrder()...)
	dbConn, err := database.GetDatabaseConnection()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	migration := gormigrate.New(dbConn, gormigrate.DefaultOptions, system)
	if err = migration.Migrate(); err != nil {
		log.Fatal().Msg("Couldn't run the migration")
	}
	log.Info().Msg("Migration is done successfully")
}
