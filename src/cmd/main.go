package main

import (
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/ueryooo/gqlgen-template/src/db"
	"github.com/ueryooo/gqlgen-template/src/graph"
)

func main() {
	if db, err := db.MysqlInit(); err != nil {
		panic(err)
	} else {
		driver, _ := mysql.WithInstance(db.DB, &mysql.Config{})
		m, _ := migrate.NewWithDatabaseInstance(
			"file://db/migrations",
			"mysql",
			driver,
		)

		fmt.Println("======== Migrations Start ! ========")
		err := m.Up()
		if err != nil && err != migrate.ErrNoChange {
			fmt.Println("======== Migrations Failed ! ========")
		} else {
			fmt.Println("======== Migrations Succeed ! ========")
		}

		graph.ServerInit()
	}

}
