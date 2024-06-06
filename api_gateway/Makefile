CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

exp:
	export DBURL='postgres://postgres:root@localhost:5432/public?sslmode=disable'

mig-up:
	migrate -path migrations -database 'postgres://postgres:root@localhost:5432/public?sslmode=disable' -verbose up

mig-down:
	migrate -path migrations -database ${DBURL} -verbose down


mig-create:
	migrate create -ext sql -dir migrations -seq create_table

mig-insert:
	migrate create -ext sql -dir migrations -seq insert_table

swag-init:
	swag init -g api/api.go -o docs/docs