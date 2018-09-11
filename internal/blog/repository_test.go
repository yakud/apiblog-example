package blog

import (
	"log"
	"testing"

	"github.com/go-pg/pg"
	pg2 "github.com/yakud/apiblog-example/internal/pg"
)

func TestRepository_CreateTable(t *testing.T) {
	rep := getRep(false)

	if err := rep.CreateTable(); err != nil {
		t.Error(err)
	}
}

func TestRepository_Create(t *testing.T) {
	rep := getRep(true)

	post := &Post{
		Name: "a",
	}

	id, err := rep.Create(post)
	if err != nil {
		t.Error(err)
	}

	if id != 1 || post.Id != 1 {
		t.Error("id not equal 1")
	}
}

func TestRepository_Get(t *testing.T) {
	rep := getRep(true)

	post := &Post{
		Name: "a",
	}

	if _, err := rep.Create(post); err != nil {
		t.Error(err)
	}

	post2 := &Post{Id: post.Id}
	if err := rep.Get(post2); err != nil {
		t.Error(err)
	}

	if post.Name != post2.Name {
		t.Error("names not equal", post.Name)
	}
}

func TestRepository_Update(t *testing.T) {
	rep := getRep(true)

	post := &Post{
		Name: "a",
	}

	if _, err := rep.Create(post); err != nil {
		t.Error(err)
	}

	// Update fields
	post.Name = "b"
	post.URI = "c"

	if err := rep.Update(post); err != nil {
		t.Error(err)
	}

	post2 := &Post{Id: post.Id}
	if err := rep.Get(post2); err != nil {
		t.Error(err)
	}

	if post.Name != post2.Name || post.URI != post2.URI {
		t.Error("names not equal", post.Name)
	}
}

func TestRepository_Delete(t *testing.T) {
	rep := getRep(true)

	post := &Post{
		Name: "a",
	}

	if _, err := rep.Create(post); err != nil {
		t.Error(err)
	}

	if err := rep.Delete(post); err != nil {
		t.Error(err)
	}

	post2 := &Post{Id: post.Id}
	if err := rep.Get(post2); err == nil {
		t.Error("Expected error")
	}
}

func TestRepository_GetAll(t *testing.T) {
	rep := getRep(true)

	posts1 := []*Post{
		{Name: "a"},
		{Name: "b"},
		{Name: "c", Content: "d"},
	}

	for _, post := range posts1 {
		if _, err := rep.Create(post); err != nil {
			t.Error(err)
		}
	}

	posts2, err := rep.GetAll()
	if err != nil {
		t.Error(err)
	}

	for i, post2 := range posts2 {
		if posts1[i].Name != post2.Name || posts1[i].Content != post2.Content {
			t.Error(posts1[i].Name, "!=", post2.Name)
		}
	}
}

func TestRepository_IncrementViewsNumber(t *testing.T) {
	rep := getRep(true)
	//rep := NewRepository(getPg())

	post := &Post{
		Name:        "a",
		ViewsNumber: 0,
	}

	if _, err := rep.Create(post); err != nil {
		t.Error(err)
	}

	if err := rep.IncrementViewsNumber(post); err != nil {
		t.Error(err)
	}

	if post.ViewsNumber != 1 {
		t.Error("expected", 1)
	}

	if err := rep.IncrementViewsNumber(post); err != nil {
		t.Error(err)
	}

	if post.ViewsNumber != 2 {
		t.Error("expected", 2)
	}

	if err := rep.IncrementViewsNumber(post); err != nil {
		t.Error(err)
	}

	if post.ViewsNumber != 3 {
		t.Error("expected", 3)
	}
}

func getRep(createTable bool) *Repository {
	rep := NewRepository(getPg())

	if err := rep.DropTable(); err != nil {
		log.Fatal(err)
	}

	if createTable {
		if err := rep.CreateTable(); err != nil {
			log.Fatal(err)
		}
	}

	return rep
}

func getPg() *pg.DB {
	db, err := pg2.NewConnection(nil)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
