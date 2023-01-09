package main

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/liac-inc/gqlgen-template/src/db"
	"github.com/liac-inc/gqlgen-template/src/graph"
)

func main() {
	if db, err := db.Connect(); err != nil {
		panic(err)
	} else {
		driver, _ := postgres.WithInstance(db, &postgres.Config{})
		m, _ := migrate.NewWithDatabaseInstance(
			"file://db/migration",
			"postgres",
			driver,
		)

		fmt.Println("======== Migrations Start ! ========")
		err := m.Up()
		if err != nil && err != migrate.ErrNoChange {
			fmt.Println("======== Migrations Failed ! ========")
		} else {
			fmt.Println("======== Migrations Succeed ! ========")
		}

		graph.ServerInit(db)
	}
}
