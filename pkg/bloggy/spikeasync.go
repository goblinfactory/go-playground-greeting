package bloggy

import (
	"fmt"
	"math/rand"
	"time"
)

const invalidOperation = -1

func downloadFile(name string) chan int32 {
	r := make(chan int32)

	go func(name string) {
		defer close(r)
		switch name {
		case "a.txt":
			time.Sleep(time.Second * 3)
			fmt.Println("downloaded " + name)
			r <- rand.Int31n(100)
		default:
			r <- invalidOperation
		}
	}(name)

	return r
}

// Run simple async test
func Run() {
	fmt.Println("downloading 3 files")
	aCh, bCh, cCh :=
		downloadFile("a.txt"),
		downloadFile("b.txt"),
		downloadFile("c.txt")

	// wait for all dowloads
	// write results
	fmt.Printf("%d,%d,%d\n",
		<-aCh,
		<-bCh,
		<-cCh,
	)
}
