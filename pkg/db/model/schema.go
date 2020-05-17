package model

import (
	"github.com/jinzhu/gorm"
)

// Cave represents the database model for table 'caves'
type Cave struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100);unique;not null"`
	CountryName string
	RegionName  string
	Longitude   float32
	Latitude    float32
}
