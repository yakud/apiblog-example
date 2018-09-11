NAME:=apiblog-example
MAINTAINER:="Aleksei Kiselev <yakudgm@gmail.com>"
DESCRIPTION:="Api Blog example"

APP_SERVER_PATH=$(GOPATH)/src/github.com/yakud/apiblog-example/cmd/apiblog-server

all: $(NAME)

up:
	cd ./docker && \
	docker-compose up

install:
	go install $(APP_SERVER_PATH)

run:
	go run $(APP_SERVER_PATH)/server.go