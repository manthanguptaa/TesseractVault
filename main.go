package main

import (
	"fmt"
	"sync"
)

type Setter[K comparable, V any] interface {
	Set(K, V) error
}

type Getter[K comparable, V any] interface {
	Get(K) (V, error)
}

type Updater[K comparable, V any] interface {
	Update(K, V) error
}

type Deleter[K comparable, V any] interface {
	Delete(K) (V, error)
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

// checks if the given key is present in the key value store
func (s *KVStore[K, V]) Exists(key K) bool {
	_, ok := s.data[key]
	return ok
}

func (s *KVStore[K, V]) Update(key K, value V) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.Exists(key) {
		return fmt.Errorf("the key (%v) doesn't exist", key)
	}

	s.data[key] = value
	return nil
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

func (s *KVStore[K, V]) Delete(key K) (V, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	value, ok := s.data[key]

	if !ok {
		return value, fmt.Errorf("the key (%v) doesn't exist", key)
	}

	delete(s.data, key)

	return value, nil
}

func StoreThings(s Setter[string, int]) error {
	return s.Set("foo", 1)
}

func main() {
	// store := InitialiseNewKVStore[string, string]()

}
