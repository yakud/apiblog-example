package app

import (
	"context"

	"sync"

	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/gramework/gramework"
	"github.com/gramework/gramework/graphiql"
	"github.com/yakud/apiblog-example/internal/blog"
	"github.com/yakud/apiblog-example/internal/gql"
	"github.com/yakud/apiblog-example/internal/pg"
	"github.com/yakud/apiblog-example/internal/redis"
	"github.com/yakud/apiblog-example/internal/worker"
)

type Server struct {
}

func (t *Server) Run(config *Config) error {
	// Init workers
	ctx, cancelWorkers := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	workersPool := worker.NewWorkersPool()
	if b, err := t.newBlogInstance(config); err == nil {
		workersPool.Run(b, ctx, wg)
	} else {
		return err
	}

	// parse graphql schema
	schema, err := gql.FileMustParseSchema(
		config.GQLSchemaFile,
		gql.NewResolver(workersPool),
	)
	if err != nil {
		return err
	}

	// init server
	gr := gramework.New(func(app *gramework.App) {
		app.Logger = &log.Logger{
			Level:   log.ErrorLevel,
			Handler: cli.New(os.Stdout),
		}
	})

	gr.POST("/graphql", gql.NewHandler(schema))
	gr.GET("/", graphiql.Handler)

	if err := gr.ListenAndServe(config.ServerAddr); err != nil {
		return err
	}

	cancelWorkers()
	wg.Wait()

	return nil
}

func (t *Server) newBlogInstance(config *Config) (*blog.Instance, error) {
	// init pg
	pgdb, err := pg.NewConnection(config.PGOptions)
	if err != nil {
		return nil, err
	}

	// init redis
	redisdb, err := redis.NewConnection(config.RedisOptions)
	if err != nil {
		return nil, err
	}

	// init blog repository
	postsRepository := blog.NewRepository(pgdb)
	postsRepository.DropTable()

	if err := postsRepository.CreateTable(); err != nil {
		return nil, err
	}

	// init blog cache
	postsCache := blog.NewCache(redisdb)
	postsCache.DropAll()

	instance := blog.NewInstance(
		postsRepository,
		postsCache,
	)

	return instance, nil
}

func NewServer() *Server {
	return &Server{}
}
