package gql

import (
	"strconv"

	"github.com/graph-gophers/graphql-go"
	"github.com/yakud/apiblog-example/internal/blog"
)

type Resolver struct {
	blog     *blog.Blog
	postPool *blog.PostPool
}

// get(id: ID!): Post
func (t *Resolver) Get(args struct{ ID graphql.ID }) (*PostResolver, error) {
	p := &blog.Post{
		Id: blog.ID(t.graphIDToInt(args.ID)),
	}

	if err := t.blog.Get(p); err != nil {
		return nil, err
	}

	return &PostResolver{Post: p}, nil
}

// getAll: [Post]!
func (t *Resolver) GetAll() ([]*PostResolver, error) {
	posts, err := t.blog.GetAll()
	if err != nil {
		return nil, err
	}

	postsResolvers := make([]*PostResolver, len(posts))
	for i, post := range posts {
		p := post // copy
		postsResolvers[i] = &PostResolver{Post: &p}
	}

	return postsResolvers, nil
}

// create(name: String, shortDescr: String, preview: String, content: String, uri: String): Post
func (t *Resolver) Create(args struct {
	Name       *string
	ShortDescr *string
	Preview    *string
	Content    *string
	Uri        *string
}) (*PostResolver, error) {
	p := t.postPool.Acquire()
	defer t.postPool.Put(p)

	if args.Name != nil {
		p.Name = *args.Name
	}
	if args.ShortDescr != nil {
		p.ShortDescr = *args.ShortDescr
	}
	if args.Preview != nil {
		p.Preview = *args.Preview
	}
	if args.Content != nil {
		p.Content = *args.Content
	}
	if args.Uri != nil {
		p.URI = *args.Uri
	}

	if _, err := t.blog.Create(p); err != nil {
		return nil, err
	}

	return &PostResolver{Post: p}, nil
}

// update(id: ID!, name: String, shortDescr: String, preview: String, content: String, uri: String): Post
func (t *Resolver) Update(args struct {
	ID         graphql.ID
	Name       *string
	ShortDescr *string
	Preview    *string
	Content    *string
	Uri        *string
}) (*PostResolver, error) {
	p := &blog.Post{
		Id: blog.ID(t.graphIDToInt(args.ID)),
	}

	if err := t.blog.Get(p); err != nil {
		return nil, err
	}

	if args.Name != nil {
		p.Name = *args.Name
	}
	if args.ShortDescr != nil {
		p.ShortDescr = *args.ShortDescr
	}
	if args.Preview != nil {
		p.Preview = *args.Preview
	}
	if args.Content != nil {
		p.Content = *args.Content
	}
	if args.Uri != nil {
		p.URI = *args.Uri
	}

	if err := t.blog.Update(p); err != nil {
		return nil, err
	}

	return &PostResolver{Post: p}, nil
}

// incrementViewsNumber(id: ID!): Post
func (t *Resolver) IncrementViewsNumber(args struct{ ID graphql.ID }) (*PostResolver, error) {
	p := &blog.Post{
		Id: blog.ID(t.graphIDToInt(args.ID)),
	}

	if err := t.blog.IncrementViewsNumber(p); err != nil {
		return nil, err
	}

	return &PostResolver{Post: p}, nil
}

// delete(id: ID!): Boolean!
func (t *Resolver) Delete(args struct{ ID graphql.ID }) (bool, error) {
	p := &blog.Post{
		Id: blog.ID(t.graphIDToInt(args.ID)),
	}

	if err := t.blog.Delete(p); err != nil {
		return false, err
	}

	return true, nil
}

func (t *Resolver) graphIDToInt(ID graphql.ID) int {
	i, _ := strconv.Atoi(string(ID))
	return i
}

func NewResolver(b *blog.Blog) *Resolver {
	return &Resolver{
		blog:     b,
		postPool: blog.NewPostPool(),
	}
}
