package gql

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/yakud/apiblog-example/internal/blog"
)

type post interface {
	Id() *graphql.ID
	ViewsNumber() *int32
	Name() *string
	ShortDescr() *string
	Preview() *string
	Content() *string
	Uri() *string
}

type PostResolver struct {
	post
	Post *blog.Post
}

func (t PostResolver) Id() *graphql.ID {
	id := graphql.ID(t.Post.Id.String())
	return &id
}

func (t PostResolver) ViewsNumber() *int32 {
	return &t.Post.ViewsNumber
}

func (t PostResolver) Name() *string {
	return &t.Post.Name
}

func (t PostResolver) ShortDescr() *string {
	return &t.Post.ShortDescr
}

func (t PostResolver) Preview() *string {
	return &t.Post.Preview
}

func (t PostResolver) Content() *string {
	return &t.Post.Content
}

func (t PostResolver) Uri() *string {
	return &t.Post.URI
}
