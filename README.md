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

	q.PushBack(1, 2, 3)
	fmt.Println("len3:", q.Len())
	q.PushFront(4, 5, 6)
	q.Init()
	fmt.Println("len4:", q.Len())
}
```
