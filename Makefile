.PHONY: build clean

BINARY_NAME=pokedexcli
BUILD_DIR=./bin

build:
	mkdir -p ${BUILD_DIR}
	GOARCH=amd64 GOOS=linux go build -o ${BUILD_DIR}/${BINARY_NAME} main.go
	GOARCH=amd64 GOOS=windows go build -o ${BUILD_DIR}/${BINARY_NAME}.exe main.go

run: build
	${BUILD_DIR}/${BINARY_NAME}

clean:
	go clean
	find ${BUILD_DIR} -name "${BINARY_NAME}-*" -type f -delete
	find . -name "cover.*" -type f -delete

test:
	go test ./... -v -coverprofile cover.out
	go tool cover -html=cover.out -o cover.html
