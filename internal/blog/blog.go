package blog

// todo: как дела с конкурентностью???
// todo: тут можем поставить некорректный кэш
// todo: обеспечить блокировку записи на момент обновления
type Blog struct {
	cache      *Cache
	repository *Repository
}

func (t *Blog) Create(post *Post) (ID, error) {
	if _, err := t.repository.Create(post); err != nil {
		return 0, err
	}

	if err := t.cache.Set(post); err != nil {
		return 0, err
	}

	return post.Id, nil
}

func (t *Blog) Delete(post *Post) error {
	if err := t.repository.Delete(post); err != nil {
		return err
	}

	if err := t.cache.Delete(post); err != nil {
		return err
	}

	return nil
}

func (t *Blog) Update(post *Post) error {
	if err := t.repository.Update(post); err != nil {
		return err
	}

	if err := t.cache.Set(post); err != nil {
		return err
	}

	return nil
}

func (t *Blog) GetAll() ([]Post, error) {
	return t.repository.GetAll()
}

func (t *Blog) Get(post *Post) error {
	if err := t.cache.Get(post); err != nil {
		if err := t.repository.Get(post); err != nil {
			// not found
			return err
		}

		// cache miss, set
		if err := t.cache.Set(post); err != nil {
			return err
		}
	}

	// found in cache
	return nil
}

func (t *Blog) IncrementViewsNumber(post *Post) error {
	if err := t.repository.IncrementViewsNumber(post); err != nil {
		return err
	}

	if err := t.cache.Set(post); err != nil {
		return err
	}

	return nil
}

func NewBlog(rep *Repository, cache *Cache) *Blog {
	return &Blog{
		cache:      cache,
		repository: rep,
	}
}
