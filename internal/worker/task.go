package worker

import (
	"sync"

	"github.com/yakud/apiblog-example/internal/blog"
)

type TaskType byte

const (
	Create TaskType = iota
	Delete
	Update
	GetAll
	Get
	IncrementViewsNumber
)

type Task struct {
	Post *blog.Post
	Type TaskType
	C    AnswerChan
}

func (t *Task) Clear() {
	t.Post = nil
	//Type.C = nil
	t.Type = 0
}

// Task pool
type TaskPool struct {
	pool *sync.Pool
}

func (t *TaskPool) Acquire() *Task {
	return t.pool.Get().(*Task)
}

func (t *TaskPool) Put(task *Task) {
	task.Clear()
	t.pool.Put(task)
}

func NewTaskPool() *TaskPool {
	p := &TaskPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return &Task{
					C: make(AnswerChan),
				}
			},
		},
	}

	return p
}
