package timing

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Echo echoes whatever is passed in.
func Echo() {
	fmt.Println(os.Args[1:])
}

// CompareConcatVsJoin test (comparing) speed of concatenating strings in a loop, and using string.join
func CompareConcatVsJoin() {

	src := []string{"mary", "had", "a", "little", "lamb", "whose", "fleece", "was", "white", "as", "snow"}
	cnt := 1000000
	t1 := timeit(src, "concatIt", cnt, concatIt)
	t2 := timeit(src, "joinIt", cnt, joinIt)
	fmt.Printf("%5.2f times faster", t1/t2)
}

type sut func([]string) string

func timeit(src []string, title string, iterations int, sut sut) float64 {
	start := time.Now()
	fmt.Println(title, iterations, "iterations;")
	for i := 0; i < iterations; i++ {
		sut(src)
	}
	fin := time.Now().Sub(start).Milliseconds()
	fmt.Println("---")
	fmt.Println(fin, "ms")
	fmt.Println("---")
	return float64(fin)
}

func concatIt(src []string) string {
	r := src[0]
	for _, w := range src[1:] {
		add := " " + w
		r += add
	}
	return r
}

func joinIt(src []string) string {
	return strings.Join(src, " ")
}
