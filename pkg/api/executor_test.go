package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/caveconditions/cc-backend/pkg/config"
	"github.com/caveconditions/cc-backend/pkg/db"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func registerServer() sqlmock.Sqlmock {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	settings := &config.Config{}
	e := &apiExecutor{
		settings: settings,
	}

	// init the database handler
	dbb, mock, _ := sqlmock.New()
	dbh, _ := gorm.Open("postgres", dbb)
	dbh.LogMode(true)
	e.handler = db.NewMockHandler(settings, dbh)

	RegisterCaveConditionsServer(s, e)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	return mock
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestAddCave(t *testing.T) {
	mock := registerServer()
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := NewCaveConditionsClient(conn)

	Convey("For a given title", t, func() {

		Convey("Returns title too long failure on 101 char title", func() {
			request := AddCaveRequest{
				Cave: &Cave{
					Title: "bUyOX8utWFyOBxV5yJenS3oo5IcwiY0wlK6dtti4PlvEQx5KUSBUHKGU9tXniBVJ0wQLIQAHxvSqVS4WyMRggR5owaefKnyPfIId2",
				},
			}

			_, err := client.AddCave(ctx, &request)
			So(err, ShouldNotBeNil)
		})

		Convey("Returns error if title not defined", func() {
			request := AddCaveRequest{
				Cave: &Cave{},
			}

			_, err := client.AddCave(ctx, &request)
			So(err, ShouldNotBeNil)

		})

		Convey("Returns correct title", func() {
			request := AddCaveRequest{
				Cave: &Cave{
					Title: "Test Cave",
				},
			}
			mock.ExpectBegin()
			mock.ExpectQuery(`^INSERT INTO "caves" (.+) RETURNING`).
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), request.Cave.Title, "", "", 0.0, 0.0).
				WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			mock.ExpectCommit()

			_, err := client.AddCave(ctx, &request)
			So(err, ShouldBeNil)

		})

	})

	Convey("For latitude", t, func() {

		Convey("Return error on latitude < -90", func() {
			request := AddCaveRequest{
				Cave: &Cave{
					Title:    "Test Cave",
					Latitude: -100,
				},
			}

			_, err := client.AddCave(ctx, &request)
			So(err, ShouldNotBeNil)
		})

		Convey("Return error on latitude  > 90", func() {
			request := AddCaveRequest{
				Cave: &Cave{
					Title:    "Test Cave",
					Latitude: 100,
				},
			}

			_, err := client.AddCave(ctx, &request)
			So(err, ShouldNotBeNil)
		})

		Convey("Returns error if latitude is set but not longitude", func() {
			request := AddCaveRequest{
				Cave: &Cave{
					Title:    "Test Cave",
					Latitude: 80,
				},
			}

			_, err := client.AddCave(ctx, &request)
			So(err, ShouldNotBeNil)
		})

	})

	Convey("For longitude", t, func() {

		Convey("Return error on longitude < -180", func() {
			request := AddCaveRequest{
				Cave: &Cave{
					Title:     "Test Cave",
					Longitude: -190,
				},
			}

			_, err := client.AddCave(ctx, &request)
			So(err, ShouldNotBeNil)
		})

		Convey("Return error on longitude  > 180", func() {
			request := AddCaveRequest{
				Cave: &Cave{
					Title:     "Test Cave",
					Longitude: 190,
				},
			}

			_, err := client.AddCave(ctx, &request)
			fmt.Println(err)
			So(err, ShouldNotBeNil)
		})

		Convey("Returns error if longitude is set but not latitude", func() {
			request := AddCaveRequest{
				Cave: &Cave{
					Title:     "Test Cave",
					Longitude: 80,
				},
			}

			_, err := client.AddCave(ctx, &request)
			So(err, ShouldNotBeNil)
		})

	})

}
