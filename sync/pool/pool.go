package main

import (
	"fmt"
	"sync"
)

func main()  {
	pool()
}

func pool()  {
	var strPool = sync.Pool{
		New: func() interface{} {
			return []string{}
		},
	}

	strPool.Put("bar")

	strPool.Put("baz")
	strPool.Put("foo")
	strPool.Put("doo")

	item := strPool.Get()
	fmt.Println(item)
	strPool.Put(item)
	fmt.Println(strPool.Get())
	fmt.Println(strPool.Get())
	fmt.Println(strPool.Get())
	fmt.Println(strPool.Get())
}