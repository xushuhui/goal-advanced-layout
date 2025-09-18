INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
.PHONY: init
init:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/golang/mock/mockgen@latest
	go install github.com/swaggo/swag/cmd/swag@latest



.PHONY: mock
mock:
	mockgen -source=internal/biz/user.go -destination test/mocks/biz/user.go
	mockgen -source=internal/data/user.go -destination test/mocks/data/user.go

.PHONY: test
test:
	go test -coverpkg=./internal/handler,./internal/service,./internal/repository -coverprofile=./coverage.out ./test/server/...
	go tool cover -html=./coverage.out -o coverage.html

.PHONY: build
build:
	go build -ldflags="-s -w" -o ./bin/server ./cmd/server

.PHONY: docker
docker:
	docker build -f deploy/build/Dockerfile --build-arg APP_RELATIVE_PATH=./cmd/job -t 1.1.1.1:5000/demo-job:v1 .
	docker run --rm -i 1.1.1.1:5000/demo-job:v1

.PHONY: swag
swag:
	swag init  -g cmd/server/main.go -o ./docs --parseDependency

.PHONY: config
config:
	protoc --proto_path=./internal \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)
