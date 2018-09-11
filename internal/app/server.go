package app

import (
	"github.com/gramework/gramework"
	"github.com/yakud/apiblog-example/internal/blog"
	"github.com/yakud/apiblog-example/internal/gql"
	"github.com/yakud/apiblog-example/internal/pg"
	"github.com/yakud/apiblog-example/internal/redis"
)

type query struct{}

func (_ *query) Hello() string { return "Hello, world!" }

type Server struct {
}

func (t *Server) Run(config *Config) error {
	// init pg
	pgdb, err := pg.NewConnection(config.PGOptions)
	if err != nil {
		return err
	}
	defer pgdb.Close()

	// init redis
	redisdb, err := redis.NewConnection(config.RedisOptions)
	if err != nil {
		return err
	}
	defer redisdb.Close()

	// init blog repository
	postsRepository := blog.NewRepository(pgdb)
	if err := postsRepository.DropTable(); err != nil {
		return err
	}

	if err := postsRepository.CreateTable(); err != nil {
		return err
	}

	// init blog cache
	postsCache := blog.NewCache(redisdb)
	if err := postsCache.DropAll(); err != nil {
		return err
	}

	// init server
	gr := gramework.New()

	// parse graphql schema
	schema, err := gql.FileMustParseSchema(config.GQLSchemaFile, &query{})
	if err != nil {
		return err
	}

	gr.POST("/", gql.NewHandler(schema))

	return gr.ListenAndServe(config.ServerAddr)
}

func NewServer() *Server {
	return &Server{}
}
