NAME:=apiblog-example
MAINTAINER:="Aleksei Kiselev <yakudgm@gmail.com>"
DESCRIPTION:="Api Blog example"

APP_SERVER_PATH=$(GOPATH)/src/github.com/yakud/apiblog-example/cmd/apiblog-server

all: $(NAME)

up:
	docker-compose up -d

down:
	docker-compose stop

install:
	go install $(APP_SERVER_PATH)

run:
	$(GOPATH)/bin/apiblog-server