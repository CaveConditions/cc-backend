package db

import (
	"CaveConditions/pkg/config"
	"fmt"

	"github.com/jinzhu/gorm"
)

const (
	// DefaultMaxOpenConns - max open connections for database
	DefaultMaxOpenConns = 20
	// DefaultMaxIdleConns - max idle connections for database
	DefaultMaxIdleConns = 5
)

// Handler for the db operations
type Handler struct {
	settings *config.Config // database connection config
	db       *gorm.DB       // postgres database object
}

// NewHandler returns a new database operation handler
func NewHandler(settings *config.Config) *Handler {
	s := &Handler{settings: settings}

	// init the database connections
	source := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		settings.DB.Host,
		settings.DB.Port,
		settings.DB.User,
		settings.DB.Name,
		settings.DB.Passwd,
	)
	db, err := gorm.Open("postgres", source)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxOpenConns(DefaultMaxOpenConns)
	db.DB().SetMaxIdleConns(DefaultMaxIdleConns)
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}
	// init the db for handler
	s.db = db

	return s
}

// Close make sure all the database connections are released
func (s *Handler) Close() {
	if s.db != nil {
		s.db.Close()
	}
}
