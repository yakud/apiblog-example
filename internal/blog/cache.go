package blog

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

const RedisKeyPrefix = "post:"
const CacheExpirationDefault = time.Hour

type Cache struct {
	CacheExpiration time.Duration

	db *redis.Client
}

func (t *Cache) Set(post *Post) error {
	data, err := t.Marshal(post)
	if err != nil {
		return fmt.Errorf("post marshal: %s", err.Error())
	}

	err = t.db.Set(t.makeKey(post), data, t.CacheExpiration).Err()
	if err != nil {
		return fmt.Errorf("post cache set: %s", err.Error())
	}

	return err
}

func (t *Cache) Get(post *Post) error {
	data, err := t.db.Get(t.makeKey(post)).Result()
	if err != nil {
		return fmt.Errorf("post cache get: %s", err.Error())
	}

	if data == "" {
		return fmt.Errorf("post cache get empty data %d: %s", post.Id, err.Error())
	}

	if err := t.Unmarshal([]byte(data), post); err != nil {
		return fmt.Errorf("post unmarshal: %s", err.Error())
	}

	return nil
}

func (t *Cache) Delete(post *Post) error {
	if err := t.db.Del(t.makeKey(post)).Err(); err != nil {
		return fmt.Errorf("post cache del: %s", err.Error())
	}

	return nil
}

func (t *Cache) DropAll() error {
	return t.db.FlushAll().Err()
}

func (t *Cache) Marshal(post *Post) ([]byte, error) {
	return json.Marshal(post)
}
func (t *Cache) Unmarshal(data []byte, post *Post) error {
	return json.Unmarshal(data, post)
}

func (t *Cache) makeKey(post *Post) string {
	return RedisKeyPrefix + post.Id.String()
}

func NewCache(db *redis.Client) *Cache {
	return &Cache{
		db: db,

		CacheExpiration: CacheExpirationDefault,
	}
}
