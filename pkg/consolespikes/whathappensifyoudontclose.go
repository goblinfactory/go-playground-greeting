package consolespikes

import "time"

// WhenYouDontCloseAndWaitConsoleDoesNotResetProperly tests what happens to a full screen terminal if you don't dispose(close) the (t)erminal properly when finished
func WhenYouDontCloseAndWaitConsoleDoesNotResetProperly() {

	left, right, _, _, _, _ := SplitLeftRight("server", "requests")
	left.Write("left")
	right.Write("right")
}

// WhatHappensIfYouDontClose2 tests what happens to a full screen terminal if you don't dispose(close) the (t)erminal properly when finished
func WhatHappensIfYouDontClose2() {

	left, right, wg, _, cancel, _ := SplitLeftRight("server", "requests")
	defer func() {
		cancel()
		wg.Wait()
	}()

	left.Write("pausing for 3 seconds then closing")
	right.Write("right window here")
	time.Sleep(3 * time.Second)
}
