package main

import (
	"fmt"
	"sync"
)

/*

Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает
каждый из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

Плюсы:
- Горячая замена алгоритмов на лету.
- Изолирует код и данные алгоритмов от остальных классов.
- Уход от наследования к делегированию.
- Реализует принцип открытости/закрытости.

Минусы:
- Усложняет программу за счёт дополнительных классов.
- Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.

*/

func main() {
	lru := &LRU[string, int]{}

	cache := NewCache[string, int](lru, 3)

	cache.Set("a", 1)
	cache.Set("b", 2)
	cache.Set("c", 3)
	cache.Set("d", 4)
}

type Cache[K comparable, V any] struct {
	eviction Eviction[K, V]
	size     int
	capacity int
	mu       sync.Mutex
	storage  map[K]V
}

func NewCache[K comparable, V any](eviction Eviction[K, V], capacity int) *Cache[K, V] {
	return &Cache[K, V]{
		eviction: eviction,
		size:     0,
		capacity: capacity,
		storage:  make(map[K]V, capacity),
	}
}

func (cache *Cache[K, V]) Set(k K, v V) {
	if cache.size >= cache.capacity {
		cache.Evict()
	}

	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.storage[k] = v
	cache.size++
}

func (cache *Cache[K, V]) Evict() {
	cache.eviction.Evict(cache)
	cache.size--
}

func (cache *Cache[K, V]) Get(k K) V {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	return cache.storage[k]
}

type Eviction[K comparable, V any] interface {
	Evict(cache *Cache[K, V])
}

type LRU[K comparable, V any] struct {
}

func (LRU *LRU[K, V]) Evict(cache *Cache[K, V]) {
	fmt.Println("[LRU] evict")
}

type FIFO[K comparable, V any] struct {
}

func (FIFO *FIFO[K, V]) Evict(cache *Cache[K, V]) {
	fmt.Println("[FIFO] evict")
}

type LFU[K comparable, V any] struct {
}

func (LFU *LFU[K, V]) Evict(cache *Cache[K, V]) {
	fmt.Println("[LFU] evict")
}
