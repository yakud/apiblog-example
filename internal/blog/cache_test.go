package blog

import (
	"log"
	"testing"

	"github.com/go-redis/redis"
	redis2 "github.com/yakud/apiblog-example/internal/redis"
)

func TestCache_SetGet(t *testing.T) {
	c := getCache()

	post1 := &Post{Id: 1, Name: "a"}
	if err := c.Set(post1); err != nil {
		t.Error(err)
	}

	post2 := &Post{Id: post1.Id}
	if err := c.Get(post2); err != nil {
		t.Error(err)
	}

	if post1.Name != post2.Name {
		t.Error("not equal cache")
	}

	post3 := &Post{Id: 2, Name: "b"}
	if err := c.Set(post3); err != nil {
		t.Error(err)
	}

	post4 := &Post{Id: post3.Id}
	if err := c.Get(post4); err != nil {
		t.Error(err)
	}

	if post3.Name != post4.Name {
		t.Error("not equal cache")
	}
}

func TestCache_Delete(t *testing.T) {
	c := getCache()

	post1 := &Post{Id: 1, Name: "a"}
	if err := c.Set(post1); err != nil {
		t.Error(err)
	}

	post2 := &Post{Id: post1.Id}
	if err := c.Get(post2); err != nil {
		t.Error(err)
	}

	if err := c.Delete(post2); err != nil {
		t.Error(err)
	}

	post3 := &Post{Id: post2.Id}
	if err := c.Get(post3); err == nil {
		t.Error("expected error")
	}

	if err := c.Delete(post2); err != nil {
		t.Error(err)
	}
}

func getCache() *Cache {
	rep := NewCache(getRedis())
	rep.DropAll()

	return rep
}

func getRedis() *redis.Client {
	db, err := redis2.NewConnection(nil)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
