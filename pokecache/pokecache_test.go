package pockecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 10)

	if cache.cache == nil {
		t.Error("The cache should not be nil")
	}
}

func TestAddToCache(t *testing.T) {
	cache := NewCache(time.Millisecond * 10)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key",
			inputVal: []byte("value"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("value2"),
		},
	}

	for _, cs := range cases {
		cache.Add(cs.inputKey, cs.inputVal)

		actual, ok := cache.Get(cs.inputKey)
		if !ok {
			t.Errorf("The key %v should be in the cache", cs.inputKey)
		}

		if string(actual) != string(cs.inputVal) {
			t.Errorf("The value %v should be in the cache", cs.inputVal)
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Second
	cache := NewCache(interval)

	keyOne := "keyOne"
	keyTwo := "keyTwo"

	cache.Add(keyOne, []byte("value"))
	cache.Add(keyTwo, []byte("value"))

	time.Sleep(interval * 2)

	_, ok := cache.Get(keyOne)
	if ok {
		t.Errorf("The key %v should be reaped", keyOne)
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Second
	cache := NewCache(interval)

	keyOne := "keyOne"
	keyTwo := "keyTwo"

	cache.Add(keyOne, []byte("value"))
	cache.Add(keyTwo, []byte("value"))

	time.Sleep(interval / 2)

	_, ok := cache.Get(keyOne)
	if !ok {
		t.Errorf("The key %v should not be reaped", keyOne)
	}
}
