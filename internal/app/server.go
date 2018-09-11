package app

import (
	"github.com/gramework/gramework"
	"github.com/yakud/apiblog-example/internal/gql"
)

type query struct{}

func (_ *query) Hello() string { return "Hello, world!" }

type Server struct {
}

func (t *Server) Run(config *Config) error {
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
