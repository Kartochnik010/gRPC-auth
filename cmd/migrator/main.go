package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/golang-migrate/migrate"
)

func main() {
	var storagePath, migrationPath, migrationsTable string

	flag.StringVar(&storagePath, "storage_path", "", "path to storage")
	flag.StringVar(&migrationPath, "migrations path", "", "path to migrations")
	flag.StringVar(&migrationsTable, "migration table", "", "specify name of table to migrate to")
	flag.Parse()
	if storagePath == "" {
		panic("Storage path is required")
	}
	if migrationPath == "" {
		panic("Migration path is required")
	}

	m, err := migrate.New(
		"file://"+migrationPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationsTable),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migration to apply")
			return
		}
		panic(err)
	}
	fmt.Println("migrations successfylly applied")
}
