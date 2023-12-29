package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
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

type Storer[K comparable, V any] interface {
	Setter[K, V]
	Getter[K, V]
	Updater[K, V]
	Deleter[K, V]
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

type Server struct {
	Storage    Storer[string, string]
	ListenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		Storage:    InitialiseNewKVStore[string, string](),
		ListenAddr: listenAddr,
	}
}

func (s *Server) handleSet(c echo.Context) error {
	key := c.Param("key")
	value := c.Param("value")
	s.Storage.Set(key, value)
	return c.JSON(http.StatusOK, map[string]string{"msg": "ok"})
}

func (s *Server) handleGet(c echo.Context) error {
	key := c.Param("key")
	value, err := s.Storage.Get(key)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{key: value})
}

func (s *Server) Start() {
	fmt.Printf("HTTP server is running on port %s", s.ListenAddr)

	e := echo.New()

	e.GET("/set/:key/:value", s.handleSet)

	e.GET("/get/:key", s.handleGet)

	e.Start(s.ListenAddr)
}

func main() {
	s := NewServer(":3000")

	s.Start()
}
