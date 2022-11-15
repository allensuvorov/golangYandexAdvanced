package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cache struct {
	m  map[string]int
	mu sync.Mutex
}

func NewCache() *Cache {
	c := Cache{
		m: make(map[string]int),
	}
	return &c
}

func (c *Cache) Set(k string, v int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[k] = v
}

func (c *Cache) Get(k string) (int, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	v, ok := c.m[k]
	return v, ok
}

type DB struct{}

func NewDB() *DB {
	db := DB{}
	return &db
}

func (db *DB) SelectUsernames() []string {
	return []string{"gosha@yandex.ru", "sasha@yandex.ru", "test@yandex.ru"}
}

func (db *DB) SelectLikeCount(username string) int {
	return rand.Intn(1000)
}

type Handler struct {
	db              *DB
	unameLikesCache *Cache
}

func NewHandler(db *DB, unamelikesCache *Cache) *Handler {
	h := Handler{
		db:              db,
		unameLikesCache: unamelikesCache,
	}
	return &h
}

func (h *Handler) GetLikes(username string) int {
	count, ok := h.unameLikesCache.Get(username)
	if !ok {
		count = h.db.SelectLikeCount(username)
		h.unameLikesCache.Set(username, count)

	}
	return count
}

func main() {
	db := NewDB()
	cache := NewCache()
	handler := NewHandler(db, cache)

	usernames := db.SelectUsernames()
	for try := 0; try < 3; try++ {
		for _, u := range usernames {
			go func(u string) {
				fmt.Printf("У пользователя %s %d лайков\n", u, handler.GetLikes(u))
			}(u)
		}
	}
	time.Sleep(1 * time.Second)
}
