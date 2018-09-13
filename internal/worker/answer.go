package worker

import (
	"sync"

	"github.com/yakud/apiblog-example/internal/blog"
)

type Answer struct {
	Id    blog.ID
	Err   error
	Posts []blog.Post
}

type AnswerChan chan *Answer

func (t *Answer) Clear() {
	t.Id = 0
	t.Err = nil
	t.Posts = nil
}

// Task pool
type AnswerPool struct {
	pool *sync.Pool
}

func (t *AnswerPool) Acquire() *Answer {
	return t.pool.Get().(*Answer)
}

func (t *AnswerPool) Put(answer *Answer) {
	answer.Clear()
	t.pool.Put(answer)
}

func NewAnswerPool() *AnswerPool {
	p := &AnswerPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return &Answer{}
			},
		},
	}

	return p
}
