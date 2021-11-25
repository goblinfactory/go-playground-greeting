package functionsandinterfaces

import (
	"net/http"
	"testing"
)

func TestPassingFuncThatSatisfiesinterface(t *testing.T) {

	handler := func(w http.ResponseWriter, r *http.Request) {}

	demo1(http.HandlerFunc(handler))
	demo1(http.HandlerFunc(demo2))
	demo3(http.HandlerFunc(demo2))
}

func demo1(h http.Handler) {

}

func demo2(rw http.ResponseWriter, r *http.Request) {

}

func demo3(h http.HandlerFunc) {

}

func TestMyOwn(t *testing.T) {
	handler := func(a int, b int) {}
	demo4(AdderFunc(handler))
	demo4(AdderFunc(demo5))
	demo6(AdderFunc(demo5))
}

type Adder interface {
	Add(int, int)
}

// this type 'alias' is created so that it can be used as a receiver onto which I can create
// an extension method.
type AdderFunc func(int, int)

// I need an extension method on func(int,int) that adds so that I can use it!
// but can't use func(int,int) as a receiver, I have to extend an actual type hence 'AdderFunc' type above
func (f AdderFunc) Add(a int, b int) {
	f(a, b)
}

func demo4(h Adder) {

}

func demo5(a int, b int) {

}

func demo6(f AdderFunc) {

}
