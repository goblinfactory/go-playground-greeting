package channels

// import "fmt"

// func ServiceStatusChannel() {

// 	// const STOP = "$STOP$"
// 	ch := make(chan string)

// 	go PingPong(ch)

// 	for _, m := range []string{"ping", "poop", "pop", "ping", "boink"} {
// 		ch <- m
// 	}

// 	close(ch)
// 	// how to block until all finished? like an Task.WhenAll(PingPong)
// 	fin := <-ch

// 	fmt.Println(fin)

// }

// func PingPong(ch chan string) {
// 	cnt := 0
// 	pings := 0
// 	for m := range ch {
// 		cnt++
// 		r, p := ping(m)
// 		if p {
// 			pings++
// 		}
// 		fmt.Println(m, r)
// 	}
// 	ch <- fmt.Sprintf("Finished. There were %d pings out of %d messages.", pings, cnt)
// }

// func ping(m string) (string, bool) {
// 	if m == "ping" {
// 		return "pong", true
// 	} else {
// 		return "N/A", false
// 	}
// }
