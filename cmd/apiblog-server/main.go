package main

import (
	"log"

	"flag"

	"github.com/go-pg/pg"
	"github.com/go-redis/redis"
	"github.com/yakud/apiblog-example/internal/app"
)

func main() {
	workers := flag.Int("workers", 1, "as int")
	GQLSchemaFile := flag.String("gql-schema", "schema/schema.graphql", "as string path")
	serverAddr := flag.String("addr", "0.0.0.0:8080", "as string path")
	pgAddr := flag.String("pg-addr", "postgres:5432", "as string path")
	redisAddr := flag.String("redis-addr", "redis:6379", "as string path")
	flag.Parse()

	config := &app.Config{
		ServerAddr:    *serverAddr,
		GQLSchemaFile: *GQLSchemaFile,
		Workers:       *workers,

		PGOptions: &pg.Options{
			User:     "pgadmin",
			Password: "pgadmin",
			Database: "apiblog",
			Addr:     *pgAddr,
		},

		RedisOptions: &redis.Options{
			Addr:     *redisAddr,
			Password: "", // no password set
			DB:       0,  // use default DB
		},
	}

	err := app.NewServer().Run(config)
	if err != nil {
		log.Fatal(err)
	}
}
