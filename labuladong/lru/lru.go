package lru

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	maxEntry int

	ll *list.List

	cache map[int]*list.Element
}

type entry struct {
	key, value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		maxEntry: capacity,
		ll:       list.New(),
		cache:    make(map[int]*list.Element),
	}
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		this.ll.MoveToBack(v)
		return v.Value.((*entry)).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.cache[key]; ok {
		this.ll.MoveToBack(e)
		en := e.Value.(*entry)
		en.value = value
		return
	}
	entry := &entry{key, value}
	ele := this.ll.PushBack(entry)
	this.cache[key] = ele
	if this.Len() > this.maxEntry {
		this.DelOldest()
	}
	fmt.Println(this.cache)
}
func (this *LRUCache) DelOldest() {
	this.removeElement(this.ll.Front())
}
func (this *LRUCache) removeElement(e *list.Element) {
	if e == nil {
		return
	}
	this.ll.Remove(e)
	en := e.Value.(*entry)
	delete(this.cache, en.key)
}
func (this *LRUCache) Len() int {
	return len(this.cache)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
