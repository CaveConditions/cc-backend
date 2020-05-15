package model

import "time"

// Cave represents the database model for table 'caves'
type Cave struct {
	Title       string
	CountryName string
	RegionName  string
	Longitude   float32
	Latitude    float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
