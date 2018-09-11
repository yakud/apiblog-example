package main

import (
	"log"

	"github.com/yakud/apiblog-example/internal/app"
)

func main() {
	config := &app.Config{
		ServerAddr:    "127.0.0.1:8080",
		GQLSchemaFile: "schema/schema.graphql",
	}

	err := app.NewServer().Run(config)
	if err != nil {
		log.Fatal(err)
	}
}
