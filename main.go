package main

import (
	"CaveConditions/cmd"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	cmd.Execute()
}
