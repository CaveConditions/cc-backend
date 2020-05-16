# CaveConditions Backend

This is WIP for https://www.caveconditions.com. This part will become the backend service based on Go, GRPC and Postgres

### migrate the database

$ vim Makefile

- update postgres database source
```
export POSTGRESQL_URL='postgres://postgres:123456@192.168.0.3:5432/cave_conditions?sslmode=disable'
```

### prepare the config.yml

```
server:
  listen_addr: "0.0.0.0:8000" # the grpc listen address

postgres:
  user: "postgres" # the username for postgres database
  passwd: "123456" # the passwd for postgres database
  host: "192.168.0.3" # the host for postgres database
  port: "5432" # the port for postgres database
  name: "cave_conditions" # the postgres database name
```

### compile the server

$ make

### start the grpc server

$ ./CaveConditions serve

### AddCave
➜  CaveConditions git:(master) ✗ ./CaveConditions AddCave -h
Add a new cave entity

Usage:
  CaveConditions AddCave [flags]

Flags:
      --address string   CaveConditions server's address (default "localhost:8000")
      --country string   The country name of cave (default "USA")
  -h, --help             help for AddCave
      --lat float32      The latitude of cave (default 30.28608)
      --lng float32      The longitude of cave (default -84.34147)
      --region string    The region name of cave (default "Florida")
      --title string     The title of cave (default "Leon Sinks")

$ ./CaveConditions AddCave --title  "Park Sinks"

### GetCave

➜  CaveConditions git:(master) ✗ ./CaveConditions GetCave -h
Get a special cave entity

Usage:
  CaveConditions GetCave [flags]

Flags:
      --address string   CaveConditions server's address (default "localhost:8000")
  -h, --help             help for GetCave
      --title string     The title of cave (default "Leon Sinks")

$ ./CaveConditions GetCave --title  "Park Sinks"