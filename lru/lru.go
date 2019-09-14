package lru

import (
	"sync"
)

type node struct {
	next *node
	prev *node
	key  string
	val  int
}

type lru struct {
	pointers    map[string]*node
	head        *node
	tail        *node
	maxSize     int
	currentSize int
	lock        *sync.RWMutex
	pool        sync.Pool
}

func NewLRU(maxSize int) *lru {
	return &lru{
		pointers: make(map[string]*node, maxSize),
		maxSize:  maxSize,
		lock:     &sync.RWMutex{},
		pool: sync.Pool{
			New: func() interface{} {
				return &node{}
			},
		},
	}
}

func (this *lru) Get(key string) (int, bool) {
	this.lock.RLock()
	defer this.lock.RUnlock()

	n, ok := this.pointers[key]
	if !ok {
		return 0, false
	}

	this.touch(n)
	return n.val, true
}

func (this *lru) Set(key string, v int) {
	this.lock.Lock()
	defer this.lock.Unlock()

	n, ok := this.pointers[key]
	if ok {
		n.val = v
		this.cutFromMiddleOfList(n)
	} else {
		n = this.pool.Get().(*node)
		n.key = key
		n.val = v
		n.next = nil
		n.prev = nil

		this.pointers[key] = n
		this.currentSize += 1
	}

	this.touch(n)
	this.evict()
}

func (this *lru) cutFromMiddleOfList(n *node) {
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
}

// Touch means make touched node new head of list
func (this *lru) touch(n *node) {
	if n == this.tail {
		this.tail = n.prev
	}

	n.next = this.head
	n.prev = nil
	if this.head != nil {
		this.head.prev = n
	}

	this.head = n

	if this.tail == nil {
		this.tail = this.head
	}
}

func (this *lru) cutTail() {
	if this.tail == nil {
		return
	}

	tmp := this.tail
	this.tail = this.tail.prev
	if this.tail != nil {
		this.tail.next = nil
	}
	this.currentSize -= 1
	this.pool.Put(tmp)
}

func (this *lru) evict() {
	if this.currentSize <= this.maxSize {
		return
	}

	delete(this.pointers, this.tail.key)
	this.cutTail()
}
