# queue

threadsafe queue for golang

## Installation

```
go get github.com/Greyh4t/queue
```

## example

```Go
package main

import (
	"fmt"

	"github.com/Greyh4t/queue"
)

func main() {
	q := queue.New()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			q.PushBack(i)
		} else {
			q.PushFront(i)
		}
	}
	fmt.Println("len1:", q.Len())

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(q.Back())
		} else {
			fmt.Println(q.Front())
		}
	}
	fmt.Println("len2:", q.Len())

	q.PushBackList([]interface{}{"a", "b", "c"})
	fmt.Println("len3:", q.Len())

	q.PushFrontList([]interface{}{"d", "e", "f"})
	fmt.Println("len4:", q.Len())

	q.Init()
	fmt.Println("len5:", q.Len())
}
```
