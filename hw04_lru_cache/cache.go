package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

var lock sync.Mutex

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	lock.Lock()
	defer lock.Unlock()
	mapItem, ok := c.items[key]
	if ok {
		mapItem.Value = value
		c.queue.MoveToFront(mapItem)
		return true
	}
	mapItem = c.queue.PushFront(value)

	c.items[key] = mapItem
	if c.queue.Len() > c.capacity {
		itemToRemove := c.queue.Back()
		var keyItemToRemove Key
		for key, value := range c.items {
			if value == itemToRemove {
				keyItemToRemove = key
			}
		}
		delete(c.items, keyItemToRemove)
		c.queue.Remove(itemToRemove)
	}
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	lock.Lock()
	defer lock.Unlock()
	mapItem, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(mapItem)
		return mapItem.Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	lock.Lock()
	defer lock.Unlock()
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
