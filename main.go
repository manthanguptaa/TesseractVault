package main

import (
	"fmt"
	"log"
	"sync"
)

type Setter[K comparable, V any] interface {
	Set(K, V) error
}

type Getter[K comparable, V any] interface {
	Get(K) (V, error)
}

type KVStore[K comparable, V any] struct {
	data map[K]V
	mu   sync.RWMutex
}

func InitialiseNewKVStore[K comparable, V any]() *KVStore[K, V] {
	return &KVStore[K, V]{
		data: make(map[K]V),
	}
}

func (s *KVStore[K, V]) Set(key K, value V) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = value

	return nil
}

func (s *KVStore[K, V]) Get(key K) (V, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, ok := s.data[key]

	if !ok {
		return value, fmt.Errorf("the key (%v) doesn't exist", key)
	}
	return value, nil
}

func StoreThings(s Setter[string, int]) error {
	return s.Set("foo", 1)
}

func main() {
	store := InitialiseNewKVStore[string, string]()

	if err := store.Set("foo", "bar"); err != nil {
		log.Fatal(err)
	}

	value, err := store.Get("foo")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)

}
