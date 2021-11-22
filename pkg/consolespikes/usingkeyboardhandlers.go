package consolespikes

import "fmt"

// import "fmt"

// SpikeUsingkeyboardHandlers ...
func SpikeUsingkeyboardHandlers() {

	fmt.Println("starting")
	left, right, wg, _, _, h := SplitLeftRight("left", "right")
	h.Handlers['a'] = func() { right.Write("left\n") }
	h.Handlers['s'] = func() { right.Write("down\n") }
	h.Handlers['d'] = func() { right.Write("right\n") }
	h.Handlers['w'] = func() { right.Write("up\n") }

	left.Write("hello world!")
	wg.Wait()
}
