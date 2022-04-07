package initInstance

import (
	"fmt"
	"sync"
)

type InitSingleton struct {
}

var lock = &sync.Mutex{}
var instance *InitSingleton

func init() {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		fmt.Println("Create new instance")
		instance = &InitSingleton{}
	} else {
		fmt.Println("Instance already created")
	}
}
func GetInstance() *InitSingleton {
	if instance != nil {
		fmt.Println("Instance already created")
		return instance
	}

	panic("instance is empty")
}
