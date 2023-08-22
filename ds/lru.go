package ds

import "container/list"

type LRUCache struct {
	maxSize int32
	cache   map[string]*list.Element
	items   *list.List
}

type item struct {
	key   string
	value interface{}
}

func NewLRUCache(maxSize int32) *LRUCache {
	return &LRUCache{
		maxSize: maxSize,
		cache:   make(map[string]*list.Element),
		items:   list.New(),
	}
}

func (c LRUCache) Get(key string) interface{} {
	if elm, ok := c.cache[key]; ok {
		c.items.MoveToFront(elm)
		return elm.Value.(*item).value
	}
	return -1
}
func (c LRUCache) Put(key string, value interface{}) {
	if elm, ok := c.cache[key]; ok {
		c.items.MoveToFront(elm)
		elm.Value.(*item).value = value
		return
	}
	if len(c.cache) == int(c.maxSize) {
		last := c.items.Back()
		delete(c.cache, last.Value.(*item).key)
		c.items.Remove(last)
	}

	i := &item{key, value}
	elm := c.items.PushFront(i)
	c.cache[key] = elm
}
