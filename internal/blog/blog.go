package blog

import "sync"

type Blog interface {
	Create(*Post) (ID, error)
	Delete(*Post) error
	Update(*Post) error
	GetAll() ([]Post, error)
	Get(post *Post) error
	IncrementViewsNumber(post *Post) error
}

// Блокировка при обновлении строк на уровне структуры
// todo: maybe блокировать только конкретную запись
type Instance struct {
	cache      *Cache
	repository *Repository

	updateMutex *sync.RWMutex
}

func (t *Instance) Create(post *Post) (ID, error) {
	t.updateMutex.Lock()
	defer t.updateMutex.Unlock()

	if _, err := t.repository.Create(post); err != nil {
		return 0, err
	}

	if err := t.cache.Set(post); err != nil {
		return 0, err
	}

	return post.Id, nil
}

func (t *Instance) Delete(post *Post) error {
	t.updateMutex.Lock()
	defer t.updateMutex.Unlock()

	if err := t.repository.Delete(post); err != nil {
		return err
	}

	if err := t.cache.Delete(post); err != nil {
		return err
	}

	return nil
}

func (t *Instance) Update(post *Post) error {
	t.updateMutex.Lock()
	defer t.updateMutex.Unlock()

	if err := t.repository.Update(post); err != nil {
		return err
	}

	if err := t.cache.Set(post); err != nil {
		return err
	}

	return nil
}

func (t *Instance) GetAll() ([]Post, error) {
	t.updateMutex.RLock()
	defer t.updateMutex.RUnlock()

	return t.repository.GetAll()
}

func (t *Instance) Get(post *Post) error {
	t.updateMutex.RLock()
	defer t.updateMutex.RUnlock()

	if err := t.cache.Get(post); err != nil {
		if err := t.repository.Get(post); err != nil {
			// not found
			return err
		}

		// Write lock
		t.updateMutex.Lock()
		defer t.updateMutex.Unlock()

		// cache miss, set
		if err := t.cache.Set(post); err != nil {
			return err
		}
	}

	// found in cache
	return nil
}

func (t *Instance) IncrementViewsNumber(post *Post) error {
	t.updateMutex.Lock()
	defer t.updateMutex.Unlock()

	if err := t.repository.IncrementViewsNumber(post); err != nil {
		return err
	}

	if err := t.cache.Set(post); err != nil {
		return err
	}

	return nil
}

func NewInstance(rep *Repository, cache *Cache) *Instance {
	return &Instance{
		cache:       cache,
		repository:  rep,
		updateMutex: &sync.RWMutex{},
	}
}
