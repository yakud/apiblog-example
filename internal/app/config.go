package app

import (
	"github.com/go-pg/pg"
	"github.com/go-redis/redis"
)

type Config struct {
	ServerAddr string

	GQLSchemaFile string

	PGOptions    *pg.Options
	RedisOptions *redis.Options
}
