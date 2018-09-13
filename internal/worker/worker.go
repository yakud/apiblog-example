package worker

import (
	"context"

	"sync"

	"github.com/pkg/errors"
	"github.com/yakud/apiblog-example/internal/blog"
)

type WorkersPool struct {
	blog.Blog

	tasks       chan *Task
	tasksPool   *TaskPool
	answersPool *AnswerPool
}

func (t *WorkersPool) Run(blog *blog.Instance, ctx context.Context, wg *sync.WaitGroup) {
	ctx.Value().
		wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case task, ok := <-t.tasks:
				if !ok {
					return
				}

				answer := t.answersPool.Acquire()

				switch task.Type {
				case Create:
					answer.Id, answer.Err = blog.Create(task.Post)
				case Delete:
					answer.Err = blog.Delete(task.Post)
				case Update:
					answer.Err = blog.Update(task.Post)
				case GetAll:
					answer.Posts, answer.Err = blog.GetAll()
				case Get:
					answer.Err = blog.Get(task.Post)
				default:
					answer.Err = errors.New("undefined task type")
				}

				task.C <- answer

			case <-ctx.Done():
				return
			}
		}
	}()
}

func (t *WorkersPool) Create(post *blog.Post) (blog.ID, error) {
	answer := t.sendAndWaitAnswer(post, Create)
	defer t.answersPool.Put(answer)

	return answer.Id, answer.Err
}

func (t *WorkersPool) Delete(post *blog.Post) error {
	answer := t.sendAndWaitAnswer(post, Delete)
	defer t.answersPool.Put(answer)

	return answer.Err
}

func (t *WorkersPool) Update(post *blog.Post) error {
	answer := t.sendAndWaitAnswer(post, Update)
	defer t.answersPool.Put(answer)

	return answer.Err
}

func (t *WorkersPool) GetAll() ([]blog.Post, error) {
	answer := t.sendAndWaitAnswer(nil, GetAll)
	defer t.answersPool.Put(answer)

	return answer.Posts, answer.Err
}

func (t *WorkersPool) Get(post *blog.Post) error {
	answer := t.sendAndWaitAnswer(post, Get)
	defer t.answersPool.Put(answer)

	return answer.Err
}

func (t *WorkersPool) IncrementViewsNumber(post *blog.Post) error {
	answer := t.sendAndWaitAnswer(post, IncrementViewsNumber)
	defer t.answersPool.Put(answer)

	return answer.Err
}

func (t *WorkersPool) sendAndWaitAnswer(post *blog.Post, taskType TaskType) *Answer {
	// Create task
	task := t.tasksPool.Acquire()
	defer t.tasksPool.Put(task)

	task.Post = post
	task.Type = taskType

	// Send to worker
	t.tasks <- task

	// Waiting for answer
	return <-task.C
}

func NewWorkersPool() *WorkersPool {
	return &WorkersPool{
		tasksPool:   NewTaskPool(),
		answersPool: NewAnswerPool(),
		tasks:       make(chan *Task),
	}
}
