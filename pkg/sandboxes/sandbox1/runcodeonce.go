package sandbox1

import (
	"fmt"
	"sync"
)

type someSingleton struct {
	Cnt  int
	Code string
}

// instantiates a default 'once' same as var once = sync.Once {}
var once sync.Once
var counter someSingleton

// DemoRunOnce ..
func DemoRunOnce() {
	fmt.Println("starting")
	doIt()
	doIt()
	doIt()
}

func doIt() {
	// once handles mutexes etc.
	once.Do(func() {
		counter = expensiveInitialisation()
	})
	counter.inc()
	fmt.Println("counter", counter.Cnt)
}

func (c *someSingleton) inc() {
	c.Cnt++
}

func expensiveInitialisation() someSingleton {
	return someSingleton{0, "CODE01"}
}
