package api

import (
	"context"
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/caveconditions/cc-backend/pkg/config"
	"github.com/caveconditions/cc-backend/pkg/db"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	settings := &config.Config{
		DB: config.DBStruct{
			Host:   "127.0.0.1",
			Port:   "5432",
			Name:   "chris",
			User:   "postgres",
			Passwd: "123456",
		},
	}
	e := newAPIExecutor(settings)
	e.handler.DeleteCaves()

	RegisterCaveConditionsServer(s, e)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func resetDb() {
	settings := &config.Config{
		DB: config.DBStruct{
			Host:   "127.0.0.1",
			Port:   "5432",
			Name:   "chris",
			User:   "postgres",
			Passwd: "123456",
		},
	}

	db.NewHandler(settings).DeleteCaves()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestAddCave(t *testing.T) {
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
			resetDb()
			request := AddCaveRequest{
				Cave: &Cave{
					Title: "Test Cave",
				},
			}

			_, err := client.AddCave(ctx, &request)
			So(err, ShouldBeNil)

			grequest := GetCaveRequest{
				Title: "Test Cave",
			}

			reply, _ := client.GetCave(context.Background(), &grequest)
			So(reply.Cave.Title, ShouldEqual, "Test Cave")
		})

		Convey("Fail if cave already exists", func() {
			resetDb()
			request := AddCaveRequest{
				Cave: &Cave{
					Title: "Test Cave",
				},
			}

			client.AddCave(ctx, &request)
			_, err := client.AddCave(ctx, &request)
			So(err, ShouldNotBeNil)

		})
	})

	Convey("For latitude", t, func() {
		resetDb()

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
		resetDb()

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
