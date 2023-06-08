package main

import "fmt"

type LRUCache struct {
	capacity int
	cache    map[int]int
	keys     []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]int),
		keys:     make([]int, 0, capacity),
	}
}

func (c *LRUCache) Put(key int, value int) {
	c.cache[key] = value
	c.keys = append(c.keys, key)
}
func (c *LRUCache) Get(key int) int {
	if val, ok := c.cache[key]; ok {
		c.keys = append(c.keys, key)
		return val
	}
	return -1
}

func (c *LRUCache) Delete(key int) int {
	if val, ok := c.cache[key]; ok {
		delete(c.cache, key)
		return val
	}
	return -1
}

func main() {
	c := Constructor(2)
	c.Put(1, 1)
	fmt.Println(c)
	fmt.Println(c.Get(1))
	c.Delete(1)
	fmt.Println("Exiting")
}
