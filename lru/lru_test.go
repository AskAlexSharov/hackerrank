package lru

import (
	"strconv"
	"testing"
)

func TestLRU2(t *testing.T) {
	lru := NewLRU(2)
	lru.Set("1", 1)
	lru.Set("3", 3)
	lru.Set("5", 5)

	expectEvicted(t, lru, "1")
	expectAbsent(t, lru, "2")
	expectEqual(t, lru, "3", 3)
	lru.Set("7", 7)
	expectEvicted(t, lru, "5")
}

func TestLRU1(t *testing.T) {
	lru := NewLRU(1)
	lru.Set("1", 1)
	lru.Set("3", 3)
	lru.Set("5", 5)

	expectEvicted(t, lru, "1")
	expectAbsent(t, lru, "2")
	expectEvicted(t, lru, "3")
	expectEqual(t, lru, "5", 5)
	lru.Set("7", 7)
	expectEvicted(t, lru, "5")
	expectEqual(t, lru, "7", 7)
}

func TestLRUMuchSetGet(t *testing.T) {
	lru := NewLRU(1000)
	go lru.Set("1", 1)
	go lru.Set("1", 1)
	go lru.Set("1", 1)
	go lru.Set("1", 1)
	go lru.Set("1", 1)
	go lru.Set("1", 1)
	lru.Set("1", 1)
	expectEqual(t, lru, "1", 1)
	expectEqual(t, lru, "1", 1)
	expectEqual(t, lru, "1", 1)
	expectEqual(t, lru, "1", 1)
	expectEqual(t, lru, "1", 1)
	expectEqual(t, lru, "1", 1)
	expectAbsent(t, lru, "2")
}

func expectAbsent(t *testing.T, lru *lru, key string) {
	if _, ok := lru.Get(key); ok {
		t.Errorf("Expected key %v to be absend", key)
	}
}

func expectEvicted(t *testing.T, lru *lru, key string) {
	if _, ok := lru.Get(key); ok {
		t.Errorf("Expected key %v to be evicted", key)
	}
}

func expectEqual(t *testing.T, lru *lru, key string, expect int) {
	if v, ok := lru.Get(key); !ok {
		t.Errorf("Expected to find element by key %v", key)
	} else if v != expect {
		t.Errorf("Wrong value of key %v: %v, expected %v", key, v, expect)
	}
}

func BenchmarkLru(b *testing.B) {
	lru := NewLRU(2)
	for i := 0; i < b.N; i++ {
		s := strconv.Itoa(i)
		lru.Set(s, i)
		lru.Set(s, i)
		lru.Set(s, i)
		lru.Get(s)
		lru.Get(s)
		lru.Get(s)
		lru.Set(s, i)
		lru.Get(s)
	}
}
