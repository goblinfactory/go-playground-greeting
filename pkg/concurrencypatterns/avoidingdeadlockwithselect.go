package concurrencypatterns

/*

// Go's select statement compared to C# tasks

select {
	case ch2 <- v:
	case v2 = <- ch1:
}

C# equivalent would be something somewhat similar to the following
(ignoring for a moment how you do blocked reading and writing to
 a channel in C#)

	Task.WaitAny(
		Task.Run(ch2<-v),
		Task.Run(v2 = <- ch1)
	)

*/
