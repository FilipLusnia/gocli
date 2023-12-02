package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {

	cache := NewCache(time.Microsecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddGetCache(t *testing.T) {
	testKey := "someKey"
	testVal := "someValue"

	cache := NewCache(time.Microsecond)
	cache.Add(testKey, []byte(testVal))
	result, ok := cache.Get(testKey)
	if !ok {
		t.Errorf("key \"%v\" not found", testKey)
	}
	if string(result) != testVal {
		t.Error("value doesn't match")
	}
}

func TestReap(t *testing.T) {
	testKey := "someKey"
	testVal := "someValue"

	interval := time.Microsecond * 10
	cache := NewCache(interval)

	cache.Add(testKey, []byte(testVal))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(testKey)
	if !ok {
		t.Errorf("%s shouldn't have been reaped", testKey)
	}
}
