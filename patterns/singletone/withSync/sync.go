package withSync

import (
	"fmt"
	"sync"
)

type syncSingleton struct {
}

var instance *syncSingleton
var once = sync.Once{}

func GetInstance() *syncSingleton {

	if instance == nil {
		once.Do(func() {
			instance = &syncSingleton{}

		})
		fmt.Println("Create new instance")

	} else {
		fmt.Println("Instance already created")
	}

	return instance
}
