package blog

import "sync"

type PostPool struct {
	pool *sync.Pool
}

func (t *PostPool) Acquire() *Post {
	return t.pool.Get().(*Post)
}

func (t *PostPool) Put(post *Post) {
	post.Clear()
	t.pool.Put(post)
}

func NewPostPool() *PostPool {
	p := &PostPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return &Post{}
			},
		},
	}

	return p
}
