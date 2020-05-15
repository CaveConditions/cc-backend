package cmd

import (
	"CaveConditions/pkg/api"
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func init() {
	// command parameters for AddCave subcommand
	AddCaveCmd.PersistentFlags().StringVar(&address, "address", "localhost:8000", "CaveConditions server's address")
	AddCaveCmd.PersistentFlags().StringVar(&title, "title", "Leon Sinks", "The title of cave")
	AddCaveCmd.PersistentFlags().StringVar(&countryName, "country", "USA", "The country name of cave")
	AddCaveCmd.PersistentFlags().StringVar(&regionName, "region", "Florida", "The region name of cave")
	AddCaveCmd.PersistentFlags().Float32Var(&longitude, "lng", -84.34147, "The longitude of cave")
	AddCaveCmd.PersistentFlags().Float32Var(&latitude, "lat", 30.28608, "The latitude of cave")
}

func init() {
	RootCmd.AddCommand(AddCaveCmd)
}

// new a grpc client with address
func newClient(address string) (api.CaveConditionsClient, error) {
	// create the grpc connection
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	// create the client with connection
	return api.NewCaveConditionsClient(conn), nil
}

// AddCaveCmd subcommand
var AddCaveCmd = &cobra.Command{
	Use:   "AddCave",
	Short: "Add a new cave entity",
	Run: func(cmd *cobra.Command, args []string) {
		// create the grpc client
		client, err := newClient(address)
		if err != nil {
			fmt.Printf("New client: %v\n", err)
			return
		}

		request := &api.AddCaveRequest{
			Cave: &api.Cave{
				Title:       title,
				CountryName: countryName,
				RegionName:  regionName,
				Longitude:   longitude,
				Latitude:    latitude,
			},
		}
		// call the AddCave interface
		if _, err := client.AddCave(context.Background(), request); err != nil {
			fmt.Printf("Add cave: %v\n", err)
			return
		}
		fmt.Println("Add a new cave entity success")
	},
}
