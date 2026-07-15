.PHONY: default run test build clean
default: run
#Variables
APP_NAME = goopportunities
#Commands
run:
	go run main.go
test:
	go test ./...
build:
	go build -o $(APP_NAME) main.go
clean:
	rm -rf $(APP_NAME)
	rm -rf ./docs
docs:
	swag init -g main.go
	go doc -all