## source the .env
$ source migrations/.env

## init the up and down sql for creating table
$ migrate create -ext sql -dir migrations -seq create_caves_table

## up the table
$ migrate -database ${POSTGRESQL_URL} -path migrations up

## down the table
$ migrate -database ${POSTGRESQL_URL} -path migrations down