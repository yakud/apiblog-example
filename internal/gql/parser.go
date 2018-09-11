package gql

import (
	"io/ioutil"
	"os"

	"github.com/graph-gophers/graphql-go"
)

func FileMustParseSchema(file string, resolver interface{}, opts ...graphql.SchemaOpt) (*graphql.Schema, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	schemaRaw, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return graphql.MustParseSchema(string(schemaRaw), resolver, opts...), nil
}
