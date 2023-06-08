package main

import "fmt"

type LRUCache struct {
	capacity int
	cache    map[int]int
	keys     []int
}

func Constructor(capacity int) LRUCache {
	if capacity < 1 || capacity > 1000 {
		panic("Capacity must be between 1 and 1000")
	}
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]int),
		keys:     make([]int, 0, capacity),
	}
}

func remove(slice []int, s int) []int {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func (c *LRUCache) Put(key int, value int) {
	if key < 0 || key > 103 {
		panic("Key must be between 0 and 103")
	}
	if value < 0 || value > 105 {
		panic("Value must be between 0 and 105")
	}
	if len(c.keys) >= c.capacity {
		oldestKey := c.keys[0]
		c.keys = c.keys[1:]
		delete(c.cache, oldestKey)
	}
	c.cache[key] = value
	c.keys = append(c.keys, key)
}

func (c *LRUCache) Get(key int) int {
	if key < 0 || key > 103 {
		panic("Key must be between 0 and 103")
	}
	if val, ok := c.cache[key]; ok {
		c.keys = remove(c.keys, key)
		c.keys = append(c.keys, key)
		return val
	}
	return -1
}

func (c *LRUCache) Delete(key int) int {
	if key < 0 || key > 103 {
		panic("Key must be between 0 and 103")
	}
	if val, ok := c.cache[key]; ok {
		c.keys = remove(c.keys, key)
		delete(c.cache, key)
		return val
	}
	return -1
}

func main() {

	// Testing
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	fmt.Println(c.Get(1)) // returns 1
	c.Put(3, 3)           // evicts key 2
	fmt.Println(c)
	fmt.Println(c.Get(2)) // returns -1 (not found)
	c.Put(4, 4)           // evicts key 1
	fmt.Println(c)
	fmt.Println(c.Get(1))    // returns -1 (not found)
	fmt.Println(c.Get(3))    // returns 3
	fmt.Println(c.Get(4))    // returns 4
	fmt.Println(c.Delete(3)) // returns 3
	fmt.Println(c)
	fmt.Println(c.Get(3))     // returns -1 (not found)
	fmt.Println(c.Delete(99)) // returns -1 (not found)
	fmt.Println(c.Get(100))   // returns -1 (not found)

	/*
	   Output of testing the above on debugging session vscode
	   1
	   {2 map[1:1 3:3] [1 3]}
	   -1
	   {2 map[3:3 4:4] [3 4]}
	   -1
	   3
	   4
	   3
	   {2 map[4:4] [3 4]}
	   -1
	   -1
	   -1
	*/

	/*
		Should fail tests

		c2 := Constructor(0) // Capacity is 0. Should panic.
		c2.Put(1, 1)

		c2 := Constructor(1001) // Capacity is 1001. Should panic.
		c2.Put(1, 1)

		c2 := Constructor(2)
		c2.Put(-1, 1) // Key is negative. Should panic.

		c2 := Constructor(2)
		c2.Put(104, 1) // Key is greater than 103. Should panic.

		c2 := Constructor(2)
		c2.Put(1, -1) // Value is negative. Should panic.

		c2 := Constructor(2)
		c2.Put(1, 106) // // Value is greater than 105. Should panic.
	*/
}
