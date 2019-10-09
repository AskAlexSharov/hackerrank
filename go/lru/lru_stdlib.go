package lru

import (
	"container/list"
	"sync"
)

type value struct {
	key string
	val int
}

type lru_stdlib struct {
	list     *list.List
	pointers map[string]*list.Element
	maxSize  int
	lock     *sync.RWMutex
}

func NewLRUStdlib(maxSize int) *lru_stdlib {
	return &lru_stdlib{
		pointers: map[string]*list.Element{},
		list:     list.New(),
		maxSize:  maxSize,
		lock:     &sync.RWMutex{},
	}
}

func (this *lru_stdlib) Get(key string) (int, bool) {
	this.lock.RLock()
	defer this.lock.RUnlock()

	n, ok := this.pointers[key]
	if !ok {
		return 0, false
	}

	this.list.MoveToFront(n)
	return n.Value.(*value).val, true
}

func (this *lru_stdlib) Set(key string, v int) {
	this.lock.Lock()
	defer this.lock.Unlock()

	n, ok := this.pointers[key]
	if ok {
		n.Value.(*value).val = v
		this.list.MoveToFront(n)
	} else {
		this.pointers[key] = this.list.PushFront(&value{key: key, val: v})
	}

	this.evict()
}

func (this *lru_stdlib) evict() {
	if this.list.Len() <= this.maxSize {
		return
	}

	delete(this.pointers, this.list.Back().Value.(*value).key)
	this.list.Remove(this.list.Back())
}
