package dbmigration

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func RunDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("Can not create new migrate instanace:", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}

	log.Println("db migrated successfully!")
}

// func RunDBMigrationClient(migrationURL string, driver string) {
// 	migration, err := migrate.NewWithDatabaseInstance(migrationURL, "postgres", driver)
// 	if err != nil {
// 		log.Fatal("Can not create new migrate instanace:", err)
// 	}

// 	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
// 		log.Fatal("failed to run migrate up:", err)
// 	}

// 	log.Println("db migrated successfully!")
// }
