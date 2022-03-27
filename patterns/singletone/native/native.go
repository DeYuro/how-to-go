package native

import (
	"fmt"
	"sync"
)

type nativeSingleton struct {
}

var lock = &sync.Mutex{}
var instance *nativeSingleton

func GetInstance() *nativeSingleton {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		fmt.Println("Create new instance")
		instance = &nativeSingleton{}
	} else {
		fmt.Println("Instance already created")
	}

	return instance
}
