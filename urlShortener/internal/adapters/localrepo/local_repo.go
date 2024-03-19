package localrepo

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"urlShortener/urlShortener/internal/entity"
)

type URLRepo interface {
	CreateShortURL(ctx context.Context, u *entity.URLShortener) (*entity.URLShortener, error)
	GetShortURL(ctx context.Context,u *entity.URLShortener) (*entity.URLShortener, error)
}

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type repository struct {
	hashTable map[string]string
	originURL map[string]bool
	mu        sync.RWMutex
	count     int
}

func New() *repository {
	return &repository{
		hashTable: make(map[string]string),
		originURL: make(map[string]bool),
		mu:        sync.RWMutex{}}
}

func (r *repository) CreateShortURL(ctx context.Context,u *entity.URLShortener) (*entity.URLShortener, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.originURL[u.URL]; ok {
		return nil, fmt.Errorf("url: %s - already exist", u.URL)
	}
	r.count++
	u.ShortURL = base62Encode(r.count)

	r.hashTable[u.ShortURL] = u.URL
	r.originURL[u.URL] = true

	return u, nil
}

func (r *repository) GetShortURL(ctx context.Context,id int) (*entity.URLShortener, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.hashTable[base62Encode(id)]; !ok {
		return nil, errors.New("not found")
	}

	url := &entity.URLShortener{
		URL: r.hashTable[base62Encode(id)],
	}
	return url, nil
}
func base62Encode(num int) string {
	encoded := ""
	for num > 0 {
		remainder := num % 62
		num /= 62
		encoded = string(base62Chars[remainder]) + encoded
	}
	return encoded
}
