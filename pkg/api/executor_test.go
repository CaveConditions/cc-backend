package api

import (
	"CaveConditions/pkg/config"
	"context"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// init the database config
func newTestAPIExecutor() *apiExecutor {
	settings := &config.Config{
		DB: config.DBStruct{
			Host:   "192.168.0.4",
			Port:   "5432",
			Name:   "cave_conditions",
			User:   "postgres",
			Passwd: "123456",
		},
	}

	return newAPIExecutor(settings)
}

func TestAddCave(t *testing.T) {
	// init the test executor
	e := newTestAPIExecutor()

	request := AddCaveRequest{
		Cave: &Cave{
			Title:       "Leon Sinks",
			CountryName: "USA",
			RegionName:  "Florida",
			Longitude:   -84.34147,
			Latitude:    30.28608,
		},
	}
	_, err := e.AddCave(context.Background(), &request)
	if err != nil {
		t.Fatalf("AddCave: %v\n", err)
	}
}

func TestGetCave(t *testing.T) {
	// init the test executor
	e := newTestAPIExecutor()

	request := GetCaveRequest{
		Title: "Leon Sinks",
	}
	reply, err := e.GetCave(context.Background(), &request)
	if err != nil {
		t.Fatalf("GetCave: %v\n", err)
	}
	t.Logf("%v\n", reply)
}
