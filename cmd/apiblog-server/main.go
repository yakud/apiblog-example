package main

import (
	"log"

	"github.com/yakud/apiblog-example/internal/app"
	"github.com/go-pg/pg"
	"github.com/go-redis/redis"
	"github.com/pkg/profile"
)

func main() {
	config := &app.Config{
		ServerAddr:    "0.0.0.0:8080",
		GQLSchemaFile: "schema/schema.graphql",

		PGOptions: &pg.Options{
			User:     "pgadmin",
			Password: "pgadmin",
			Database: "apiblog",
			Addr:     "postgres:5432",
		},

		RedisOptions: &redis.Options{
			Addr:     "redis:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		},
	}

	defer profile.Start(profile.MemProfile).Stop()

	err := app.NewServer().Run(config)
	if err != nil {
		log.Fatal(err)
	}
}
