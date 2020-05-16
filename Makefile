GO=GO111MODULE=on go
GOBUILD=$(GO) build
GOINSTALL=$(GO) install

default: build

all: build install proto db_up db_down

build:
	$(GOBUILD) -o CaveConditions main.go

install:
	$(GOINSTALL)

proto:
	protoc --go_out=plugins=grpc:. ./pkg/api/api.proto

db_up:
	export POSTGRESQL_URL='postgres://postgres:123456@192.168.0.4:5432/cave_conditions?sslmode=disable'
	migrate -database ${POSTGRESQL_URL} -path db/migrations up

db_down:
	export POSTGRESQL_URL='postgres://postgres:123456@192.168.0.4:5432/cave_conditions?sslmode=disable'
	migrate -database ${POSTGRESQL_URL} -path db/migrations down