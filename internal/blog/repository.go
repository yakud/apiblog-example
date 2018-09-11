package blog

import (
	"fmt"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Repository struct {
	db *pg.DB
}

func (t *Repository) CreateTable() error {
	opt := &orm.CreateTableOptions{
		Temp: false,
	}

	if err := t.db.CreateTable(&Post{}, opt); err != nil {
		return fmt.Errorf("create table posts: %s", err.Error())
	}

	return nil
}

func (t *Repository) DropTable() error {
	if err := t.db.DropTable(&Post{}, &orm.DropTableOptions{}); err != nil {
		return fmt.Errorf("drop table posts: %s", err.Error())
	}

	return nil
}

func (t *Repository) Create(post *Post) (ID, error) {
	if err := t.db.Insert(post); err != nil {
		return 0, fmt.Errorf("create post: %s", err.Error())
	}

	return post.Id, nil
}

func (t *Repository) Delete(post *Post) error {
	if err := t.db.Delete(post); err != nil {
		return fmt.Errorf("delete post: %s", err.Error())
	}

	return nil
}

func (t *Repository) Update(post *Post) error {
	if err := t.db.Update(post); err != nil {
		return fmt.Errorf("delete post: %s", err.Error())
	}

	return nil
}

func (t *Repository) GetAll() ([]Post, error) {
	var posts []Post

	if err := t.db.Model(&posts).Select(); err != nil {
		return nil, fmt.Errorf("get all post: %s", err.Error())
	}

	return posts, nil
}

func (t *Repository) GetById(id ID, post *Post) error {
	post.Id = id

	if err := t.db.Select(post); err != nil {
		return fmt.Errorf("get post by id %d: %s", id, err)
	}

	return nil

}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{
		db: db,
	}
}
