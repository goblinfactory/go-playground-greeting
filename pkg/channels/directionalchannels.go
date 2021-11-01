package channels

// CreateWriteChan returns a channel you can write to. Caller is responsible for closing channel. In any case this is almost always a sign of a bad design.
func CreateWriteChan() chan<- int {
	ch := make(chan int)
	return ch
}

// CreateReadChan returns a channel you can read from. these two examples show you how to constrain
func CreateReadChan() <-chan int {
	ch := make(chan int)
	return ch
}

// references
// https://blog.gopheracademy.com/advent-2019/directional-channels/
