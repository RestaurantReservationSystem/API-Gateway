CURRENT_DIR =$(shell pwd)

mod:
	go mod init composition_service
	go mod tidy
	go mod vendor
tidy:
	go mod tidy
	go mod vendor

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o docs




