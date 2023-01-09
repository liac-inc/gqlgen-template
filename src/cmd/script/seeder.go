package main

import (
	"fmt"

	"os"

	"github.com/liac-inc/gqlgen-template/src/db"
)

const seedDir = "/db/seed/"

func main() {
	db, err := db.Connect()
	if err != nil {
		panic(err)
	}

	currentDir, _ := os.Getwd()
	files, err := os.ReadDir(currentDir + seedDir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		bytes, err := os.ReadFile(currentDir + seedDir + file.Name())
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
