package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ueryooo/gqlgen-template/src/db"
)

const seedDir = "/db/seeds/"

func main() {
	db, err := db.MysqlInit()
	if err != nil {
		panic(err)
	}

	currentDir, _ := os.Getwd()
	files, err := os.ReadDir(currentDir + seedDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		bytes, err := ioutil.ReadFile(currentDir + seedDir + file.Name())
		if err != nil {
			continue
		}
		_, err = db.Exec(string(bytes))
		if err != nil {
			fmt.Println("error:", err)
		}
	}
	fmt.Println("======== Data Seeding Succeed ! ========")
}
