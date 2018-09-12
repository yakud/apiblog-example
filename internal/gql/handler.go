package gql

import (
	"github.com/gramework/gramework"
	"github.com/graph-gophers/graphql-go"
)

type Handler struct {
	schema *graphql.Schema
}

func (t *Handler) Handler(ctx *gramework.Context) {
	if t.schema == nil {
		ctx.Logger.Error("schema is nil")
		ctx.Err500()
		return
	}

	req, err := ctx.DecodeGQL()
	if err != nil {
		ctx.Logger.Warn("gql request decoding failed")
		ctx.Error("Invalid request", 400)
		return
	}

	if req == nil {
		ctx.Logger.Error("GQL request is nil: invalid content type")
		ctx.Error("Invalid content type", 400)
		return
	}

	if _, err := ctx.Encode(t.schema.Exec(ctx.ToContext(), req.Query, req.OperationName, req.Variables)); err != nil {
		ctx.SetStatusCode(415)
	}
}

func NewHandler(schema *graphql.Schema) *Handler {
	return &Handler{
		schema: schema,
	}
}
