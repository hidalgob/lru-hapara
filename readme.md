lru-hapara
==========

Design and implement a data structure for Least Recently Used (LRU) cache. (https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU) It should support the following operations: get, put and delete
get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reaches its capacity, it should invalidate the least recently used item before inserting a new item.
delete(key) - Remove the value of the key if the key exists in the cache. Return the value if the key exists in the cache, otherwise return -1.
Please use basic data structures where possible. We are interested to see how you would implement the algorithm.  

Constraints:  
1 <= capacity <= 1000  
0 <= key <= 103  
0 <= value <= 105  


Tech test done for Golang Developer role at Hapara  
Submission for Bruno Hidalgo Lemergas


Testing
=======

For the simple purpose of the test, just run `main()` at `lru.go` file
Corner case tests are commented so no `panic` is throw.

References
=======

Links used for the completion of this test    

https://en.wikipedia.org/wiki/Cache_replacement_policies#LRU  
https://go.dev/doc/tutorial/getting-started  
https://go.dev/doc/  
https://www.w3schools.com/go/go_syntax.php  
https://www.golang-book.com/books/intro/6  
https://blog.learngoprogramming.com/learn-go-lang-variables-visual-tutorial-and-ebook-9a061d29babe  
https://github.com/hashicorp/golang-lru  
https://gobyexample.com/structs  
https://go.dev/tour/moretypes/15  
https://go.dev/tour/moretypes/11  
https://go.dev/blog/maps  
https://pkg.go.dev/container/list  
https://go.dev/play/  
https://stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang  
https://go.dev/tour/flowcontrol/1  
https://gobyexample.com/panic  
