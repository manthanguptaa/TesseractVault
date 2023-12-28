package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetAndGet(t *testing.T) {
	store := InitialiseNewKVStore[string, int]()
	key := "myKey"
	value := 42

	err := store.Set(key, value)
	assert.NoError(t, err)

	retrievedValue, err := store.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, value, retrievedValue)
}

func TestUpdate(t *testing.T) {
	store := InitialiseNewKVStore[string, string]()
	key := "name"
	initialValue := "Alice"
	updatedValue := "Bob"

	err := store.Set(key, initialValue)
	assert.NoError(t, err)

	err = store.Update(key, updatedValue)
	assert.NoError(t, err)

	retrievedValue, err := store.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, updatedValue, retrievedValue)
}

func TestDelete(t *testing.T) {
	store := InitialiseNewKVStore[int, string]()
	key := 123
	value := "hello"

	err := store.Set(key, value)
	assert.NoError(t, err)

	deletedValue, err := store.Delete(key)
	assert.NoError(t, err)
	assert.Equal(t, value, deletedValue)

	_, err = store.Get(key)
	assert.Error(t, err) // Key should no longer exist
}

func TestUpdateNonExistentKey(t *testing.T) {
	store := InitialiseNewKVStore[string, int]()
	key := "nonExistentKey"
	value := 55

	err := store.Update(key, value)
	assert.Error(t, err)
}

func TestDeleteNonExistentKey(t *testing.T) {
	store := InitialiseNewKVStore[string, string]()
	key := "missingKey"

	_, err := store.Delete(key)
	assert.Error(t, err)
}

func TestExists(t *testing.T) {
	store := InitialiseNewKVStore[string, string]()
	key1 := "existingKey"
	key2 := "missingKey"

	store.Set(key1, "value")

	assert.True(t, store.Exists(key1))
	assert.False(t, store.Exists(key2))
}
