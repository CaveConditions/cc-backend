## export an env
$ export POSTGRESQL_URL='postgres://postgres:123456@192.168.0.3:5432/cave_conditions?sslmode=disable'

## init the up and down sql for creating table
$ migrate create -ext sql -dir migrations -seq create_caves_table

## up the table
$ migrate -database ${POSTGRESQL_URL} -path migrations up

## down the table
$ migrate -database ${POSTGRESQL_URL} -path migrations down