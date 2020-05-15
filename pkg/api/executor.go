package api

import (
	"CaveConditions/pkg/config"
	"CaveConditions/pkg/db"
)

// apiExecutor can run the GRPC API methods
type apiExecutor struct {
	UnimplementedCaveConditionsServer

	settings *config.Config // settings for the server
	handler  *db.Handler    // handler for db operation
}

// newAPIExecutor returns the api methods executor
func newAPIExecutor(settings *config.Config) *apiExecutor {
	e := &apiExecutor{
		settings: settings,
	}

	// init the database handler
	e.handler = db.NewHandler(settings)

	return e
}

// close the database connections and other resources
func (e *apiExecutor) close() {
	if e.handler != nil {
		e.handler.Close()
	}
}
