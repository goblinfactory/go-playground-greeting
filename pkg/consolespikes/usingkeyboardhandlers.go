package consolespikes

// SpikeUsingkeyboardHandlers ...
func SpikeUsingkeyboardHandlers() {

	left, right, wg, _, _, h := SplitLeftRight("left", "right")
	h.Handlers['a'] = func() { right.Write("left\n") }
	h.Handlers['s'] = func() { right.Write("down\n") }
	h.Handlers['d'] = func() { right.Write("right\n") }
	h.Handlers['w'] = func() { right.Write("up\n") }

	left.GreenLine("keyboard hander demo")
	left.Write("press a,s,d keys, press 'q' to quit.")
	wg.Wait()
}
