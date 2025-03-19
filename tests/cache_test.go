package tests

import (
	"github.com/maevlava/pokedex/internal/pokecache"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCache_AddAndGet(t *testing.T) {
	cache := pokecache.NewCache(2 * time.Second)

	key := "test_key"
	value := []byte("test_value")

	cache.Add(key, value)

	retrieved, found := cache.Get(key)

	assert.True(t, found, "Expected to find value in cache, but it was missing")
	assert.Equal(t, value, retrieved, "Expected value to be equal")
}

func TestCache_Expiration(t *testing.T) {
	cache := pokecache.NewCache(1 * time.Second)

	key := "expire_key"
	value := []byte("expire_value")

	cache.Add(key, value)
	time.Sleep(2 * time.Second) // The test

	_, found := cache.Get(key)

	assert.False(t, found, "Expected to find value in cache, but it was missing")
}
