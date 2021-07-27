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

	pool1 := []string{"foo", "bar"}
	strPool.Put(pool1)
	item := strPool.Get().([]string)
	fmt.Println(item)
	item = append(item, "baz")
	strPool.Put(item)
	strPool.Put([]string{"some", "values"})
	fmt.Println(strPool.Get())
	fmt.Println(strPool.Get())
	fmt.Println(strPool.Get())
}