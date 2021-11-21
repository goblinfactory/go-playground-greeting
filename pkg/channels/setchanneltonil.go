package channels

import (
	"context"
	"sync"
	"time"

	"github.com/goblinfactory/greeting/pkg/consolespikes"
	"github.com/goblinfactory/greeting/pkg/rnd"
)

// DemoNotSettingChannelToNilCausesALotOfWastedCycles demonstrates how NOT setting a channel to nil in a select causes the first channel to close to always read default 0 value
// this shows up as
func DemoNotSettingChannelToNilCausesALotOfWastedCycles() {

	sc, pc, fc, wg, ctx, _ := consolespikes.SplitColumns123("status", "pressure", "fuel")

	sk := consolespikes.NewKonsole(sc) // status konsole
	pk := consolespikes.NewKonsole(pc) // pressure konsole
	fk := consolespikes.NewKonsole(fc) // fuel konsole

	sk.WriteLine("press q to quit.")
	sk.Gray("(each red asterisks  printed simulate 10 million wasted cpu cycles)")
	// I have just picked 10 million as a randomly large number to make you really think about the waste. I haven't load tested this to work out the actual number, can do that later.

	wg.Add(2)
	chf := readRocketStageFuelLevel(ctx, wg)
	chp := readRocketStagePressure(ctx, wg)

Loop: // example of using a label to break out of a for select loop.
	for {
		select {
		case <-ctx.Done():
			break Loop
		case p, ok := <-chp:
			if !ok {
				pk.Red("*")
				time.Sleep(100 * time.Millisecond) // this is here to slow down the wasted CPU cycles enough to visual see them appear.
			} else {
				pk.Green(p, "Mpa ")
			}
		case p, ok := <-chf:
			if !ok {
				fk.Red("*")
				time.Sleep(100 * time.Millisecond) // this is here to slow down the wasted CPU cycles enough to visual see them appear.
			} else {
				fk.Green(p, "kg ")
			}
		}
	}
	wg.Wait()
}

func readRocketStagePressure(ctx context.Context, wg *sync.WaitGroup) chan int {
	nums := []int{352, 235} // only returning 2 values so that we can test what happens when we don't turn off the case in the select.
	return FakeReadIODevice(ctx, wg, nums, 100, 1000)
}

func readRocketStageFuelLevel(ctx context.Context, wg *sync.WaitGroup) chan int {
	nums := []int{100, 99, 70, 66, 55, 40, 33, 11, 5, 3}
	return FakeReadIODevice(ctx, wg, nums, 100, 1000)
}

// FakeReadIODevice starts a goroutine to push fake IO device "readings" to a channel with a random pause between readings between msMin and msMax milliseconds.
func FakeReadIODevice(ctx context.Context, wg *sync.WaitGroup, fakereadings []int, msMin int, msMax int) chan int {
	ch := make(chan int, len(fakereadings))
	go func() {
		defer func() {
			close(ch)
			wg.Done()
		}()
		for n := range fakereadings {
			select {
			case <-ctx.Done():
				return
			default:
				rnd.SleepMinMaxMs(msMin, msMax)
				ch <- n
			}
		}
	}()
	return ch
}

// DemoActuallySettingChannelToNilTurnsOFFTheChannelWithZeroCPUWaste demonstrates how setting a channel to nil in a select causes the channel to be efficiently turned off, avoiding CPU expensive tight polling loops.
func DemoActuallySettingChannelToNilTurnsOFFTheChannelWithZeroCPUWaste() {

	sc, pc, fc, wg, ctx, _ := consolespikes.SplitColumns123("status", "pressure", "fuel")

	sk := consolespikes.NewKonsole(sc) // status konsole
	pk := consolespikes.NewKonsole(pc) // pressure konsole
	fk := consolespikes.NewKonsole(fc) // fuel konsole

	sk.WriteLine("press q to quit.\n")
	sk.Gray("If you're reading this wondering what the big deal is, then please run DemoNotSettingChannelToNilCausesALotOfWastedCycles, let the func run for a few seconds and both fuel and pressure IO simulators will stop streaming data, and look at the wasted cpu cycles represented by the red asterisks;\n\n")

	wg.Add(2)
	chf := readRocketStageFuelLevel(ctx, wg)
	chp := readRocketStagePressure(ctx, wg)

Loop: // example of using a label to break out of a for select loop.
	for {
		select {
		case <-ctx.Done():
			break Loop
		case p, ok := <-chp:
			if !ok {
				pk.Red("*")
				chp = nil // <-- this is IT, this is ALL YOU HAVE TO DO!
			} else {
				pk.Green(p, "Mpa ")
			}
		case p, ok := <-chf:
			if !ok {
				fk.Red("*")
				chf = nil // <-- this is IT, this is ALL YOU HAVE TO DO! Booya, stopped in it's tracks. No wasted CPU cycles.
			} else {
				fk.Green(p, "kg ")
			}
		}
	}
	wg.Wait()
}
