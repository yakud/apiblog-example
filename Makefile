NAME:=apiblog-example
MAINTAINER:="Aleksei Kiselev <yakudgm@gmail.com>"
DESCRIPTION:="Api Blog example"

all: $(NAME)

up:
	docker-compose up

down:
	docker-compose stop

install:
	go get -v ./... && \
	ls -lh $(GOPROJECT) && \
	go build -o $(GOPATH)/bin/apiblog-server github.com/yakud/apiblog-example/cmd/apiblog-server

run:
	$(GOPATH)/bin/apiblog-server