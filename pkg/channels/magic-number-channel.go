package channels

import (
	"fmt"
	"math/rand"
	"time"
)

// whats a good example where you need blocking behavior, maybe a pipeline of sorts?

const serviceStoppedNoError = "FIN"

// MagicNumberMessage runs a test involving two goroutines signaling to change behaviour using a dedicated signal channel
func MagicNumberMessage() {

	fmt.Println("MagicNumberMessage")

	// start two nano-services (as go routines, aka threads or C# Task.Run ... )

	r1 := make(chan string)
	r2 := make(chan string)
	pingPong1 := startPingPongService(r1, "A")
	pingPong2 := startPingPongService(r2, "B")

	for _, m := range []string{"ping", "poop", "pop", "boink"} {
		pingPong1 <- m
	}

	for _, m := range []string{"ping", "poop", "pop", "boink"} {
		pingPong2 <- m
	}

	close(pingPong1)
	close(pingPong2)

	// block for a response on the same channel indicating service has finished.
	fmt.Println(<-r1)
	fmt.Println(<-r2)

	fmt.Println("all services should have finished and disposed.")
}

// what was the problem, and what did I do to fix it?
// problem was I was using the same channel for the results, not realising taht the service I was starting was also listening on that channel
// furthermore the channel write is non blocking and I immediately read, which doesnt block because it's a buffered channel...
// so the code exits and I have a leak.
// the fix, is use a dedicated blocking (non buffered) channel for signalling that the service has completed.

func startPingPongService(results chan string, name string) chan string {
	ch := make(chan string, 10)
	go pingPongService(ch, results, name)
	return ch
}

// pingPongService service responds to all pings on the channel by printing "pong" to the console. Send the string zero value "" to tell the service to shut down.
func pingPongService(ch chan string, results chan string, name string) {

	cnt := 0
	pings := 0

	fmt.Println("started", name)
	defer pingPongDispose(name, results)

	for m := range ch {
		cnt++
		ms := time.Duration(rand.Intn(2000)) * time.Millisecond
		time.Sleep(ms)
		r, p := ping(m)
		if p {
			pings++
		}
		fmt.Println(name, m, r)
	}
}

func pingPongDispose(name string, ch chan string) {
	fmt.Println("cleaning up resources, open files etc")
	fmt.Println("...done cleaning up.")
	ch <- fmt.Sprintf("%s,%s", name, serviceStoppedNoError)
}

func ping(m string) (string, bool) {
	if m == "ping" {
		return "pong", true
	}
	return "N/A", false
}
