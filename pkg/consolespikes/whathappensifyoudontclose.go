package consolespikes

import "time"

// WhatHappensIfYouDontCloseTerminal tests what happens to a full screen terminal if you don't dispose(close) the (t)erminal properly when finished
func WhatHappensIfYouDontCloseTerminal() {
	left, right, _, _, _, _ := SplitLeftRight("server", "requests")
	left.Write("this demo pausing for 3 seconds then closes. What happens is there are goroutines that will be left running that will capture mouse events and convert them to ansi codes that you will see in the main console. Move your mouse over the console to see the result.")
	right.Green("When you have finished you will need to close your terminal to kill the running go-routines.")
	time.Sleep(3 * time.Second)
}

// CorrectWayToCloseTerminal ...
func CorrectWayToCloseTerminal() {
	left, right, wg, _, cancel, _ := SplitLeftRight("server", "requests")

	left.Write("this demo pausing for 3 seconds then closes. What happens is there are goroutines that will be left running that will capture mouse events and convert them to ansi codes that you will see in the main console. Move your mouse over the console to see the result.")
	right.Write("right window here")
	time.Sleep(3 * time.Second)

	cancel()
	wg.Wait()
}
