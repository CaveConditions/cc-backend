package api

import (
	"CaveConditions/pkg/config"
	"CaveConditions/pkg/db"
	"CaveConditions/pkg/db/model"
	"CaveConditions/pkg/log"
	context "context"
	"time"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// AddCave - check if the cave is existed, if not, add the cave to database
func (e *apiExecutor) AddCave(ctx context.Context, request *AddCaveRequest) (*AddCaveReply, error) {
	if request.Cave.Title == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Cave title is empty")
	}
	// check if the cave is existed
	cave, err := e.handler.FindCaveByUK(request.Cave.Title)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Find cave: %v", err)
	}
	if cave != nil {
		return nil, status.Errorf(codes.AlreadyExists, "Cave is existed")
	}

	cave = &model.Cave{
		Title:       request.Cave.Title,
		CountryName: request.Cave.CountryName,
		RegionName:  request.Cave.RegionName,
		Longitude:   request.Cave.Longitude,
		Latitude:    request.Cave.Latitude,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	// insert the cave to database
	if err := e.handler.CreateCave(cave); err != nil {
		log.Errorf("Create cave: %v", err)
		return nil, status.Errorf(codes.Internal, "Create cave: %v", err)
	}

	return &AddCaveReply{}, nil
}

// GetCave - query the cave by the unique key 'title'
func (e *apiExecutor) GetCave(ctx context.Context, request *GetCaveRequest) (*GetCaveReply, error) {
	if request.Title == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Cave title is empty")
	}

	// check if the cave is existed
	cave, err := e.handler.FindCaveByUK(request.Title)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Find cave: %v", err)
	}
	if cave == nil {
		return nil, status.Errorf(codes.AlreadyExists, "Cave is not found")
	}

	return &GetCaveReply{
		Cave: &Cave{
			Title:       cave.Title,
			CountryName: cave.CountryName,
			RegionName:  cave.RegionName,
			Longitude:   cave.Longitude,
			Latitude:    cave.Latitude,
		},
	}, nil
}
