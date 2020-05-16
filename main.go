package main

import (
	"github.com/caveconditions/cc-backend/cmd"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	cmd.Execute()
}
