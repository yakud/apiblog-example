package app

import (
	"github.com/gramework/gramework"
	"github.com/yakud/apiblog-example/internal/blog"
	"github.com/yakud/apiblog-example/internal/gql"
	"github.com/yakud/apiblog-example/internal/pg"
	"github.com/yakud/apiblog-example/internal/redis"
)

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
	postsRepository.DropTable()

	if err := postsRepository.CreateTable(); err != nil {
		return err
	}

	// init blog cache
	postsCache := blog.NewCache(redisdb)
	postsCache.DropAll()

	blogCompose := blog.NewBlog(
		postsRepository,
		postsCache,
	)

	// init server
	gr := gramework.New()

	// parse graphql schema
	schema, err := gql.FileMustParseSchema(
		config.GQLSchemaFile,
		gql.NewResolver(blogCompose),
	)
	if err != nil {
		return err
	}

	gr.POST("/graphql", gql.NewHandler(schema))
	gr.GET("/", func(ctx *gramework.Context) error {
		ctx.SetBody([]byte(indexPage))
		ctx.HTML()

		return nil
	})

	return gr.ListenAndServe(config.ServerAddr)
}

func NewServer() *Server {
	return &Server{}
}
