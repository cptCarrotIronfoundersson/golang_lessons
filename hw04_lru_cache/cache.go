package hw04lrucache

import "fmt"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l lruCache) Set(key Key, value interface{}) bool {
	ls, ok := l.items[key]
	if !ok {
		if len(l.items) == l.capacity {
			l.queue.Remove(l.queue.Back())
			for key := range l.items {
				if l.items[key] == nil {
					delete(l.items, key)
					break
				}
			}
		}
		l.items[key] = l.queue.PushFront(value)
		return false
	}
	ls.Value = value
	l.queue.MoveToFront(ls)
	return true
}

func (l lruCache) Get(key Key) (interface{}, bool) {
	if value, ok := l.items[key]; ok {
		fmt.Println(value.Value)
		l.queue.MoveToFront(l.items[key])
		return value.Value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
