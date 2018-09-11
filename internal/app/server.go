package app

import (
	"github.com/gramework/gramework"
	"github.com/yakud/apiblog-example/internal/gql"
	"github.com/yakud/apiblog-example/internal/pg"
)

type query struct{}

func (_ *query) Hello() string { return "Hello, world!" }

type Server struct {
}

func (t *Server) Run(config *Config) error {
	pgdb, err := pg.NewConnection(nil)
	if err != nil {
		return err
	}

	defer pgdb.Close()

	gr := gramework.New()

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
