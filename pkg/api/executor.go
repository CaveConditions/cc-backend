package api

import (
	context "context"
	"time"

	"github.com/caveconditions/cc-backend/pkg/log"

	"github.com/caveconditions/cc-backend/pkg/db"

	"github.com/caveconditions/cc-backend/pkg/db/model"

	"github.com/caveconditions/cc-backend/pkg/config"

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

func checkCaveRequest(request *AddCaveRequest) error {
	if len(request.Cave.Title) > 100 {
		return status.Error(codes.InvalidArgument, "Title should be max. 100 characters long")
	}

	if request.Cave.Title == "" {
		return status.Errorf(codes.InvalidArgument, "Cave title is empty")
	}

	if request.Cave.Latitude < -90 {
		return status.Errorf(codes.InvalidArgument, "Latitude has to be between -90 and 90")
	}

	if request.Cave.Latitude > 90 {
		return status.Errorf(codes.InvalidArgument, "Latitude has to be between -90 and 90")
	}

	if request.Cave.Longitude < -180 {
		return status.Errorf(codes.InvalidArgument, "Longitude has to be between -180 and 180")
	}

	if request.Cave.Longitude > 180 {
		return status.Errorf(codes.InvalidArgument, "Longitude has to be between -180 and 180")
	}

	if request.Cave.Latitude != 0 && request.Cave.Longitude == 0 {
		return status.Errorf(codes.InvalidArgument, "If latitude is provided, then longitude has to be provided as well")
	}

	if request.Cave.Longitude != 0 && request.Cave.Latitude == 0 {
		return status.Errorf(codes.InvalidArgument, "If longitude is provided, then latitude has to be provided as well")
	}

	return nil
}

// AddCave - check if the cave is existing, if not, add the cave to database
func (e *apiExecutor) AddCave(ctx context.Context, request *AddCaveRequest) (*AddCaveReply, error) {

	err := checkCaveRequest(request)

	if err != nil {
		return nil, err
	}

	// check if the cave is existing
	cave, err := e.handler.FindCaveByUK(request.Cave.Title)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Find cave: %v", err)
	}
	if cave != nil {
		return nil, status.Errorf(codes.AlreadyExists, "Cave is existing")
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
