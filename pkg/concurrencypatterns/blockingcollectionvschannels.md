# Compare Golang channels withC# BlockingCollection & Threads

Don't do this; i.e. don't use channels and functions to create collections. It's the wrong use of concurrecy; This contrived example is shown here so that we can have some insights to compare channels to C# BlockingCollection and Threads. Typically this might be an expensive operation, calculateExpensiveFoo. I've not used long names in the example below, because the code doesnt fit in side by side when i do! Please use your imagination. 😇

This is a draft : Still need to add to the bottom where this pattern goes horribly wrong (deadlocks)and how to fix it (improve) the code in both C# and Go. Real Go advantages start to kick in with select. Need to add that below.

<table style="padding:0px">
<tr>
<th>Go</th>
<th>[C#]</th>
</tr>
<tr>
<td style="vertical-align:top;">

https://play.golang.org/p/F8kAX9Xi81u

```go
func main() {
	for n := range GenerateNums(10) {
		fmt.Println(n)
	}
}

func GenerateNums(cnt int) <-chan int {
	ch := make(chan int, 5)
	go func() {
		defer close(ch)
		for i := 0; i < cnt; i++ {
			fmt.Printf("Adding: %d\n", i)
			ch <- i
		}
	}()
	return ch
}
```

</td>
<td style="vertical-align:top;" >

https://dotnetfiddle.net/kS6jdv

```csharp
public static void Main() {
	foreach(var n in GenerateNums(10)) {
		Console.WriteLine("num: {0}", n);
	}
}

public static IEnumerable<int> GenerateNums(int cnt) {
	var bc = new BlockingCollection<int>(1);
	Task.Run(()=> {
		try {
			for(int i=1; i<cnt; i++) {
				Console.WriteLine("adding: {0}", i);
				bc.Add(i);
			}
		} finally {
			bc.CompleteAdding();
		}
	});
	return bc.GetConsumingEnumerable();
}
```

</td>
</tr>
<tr>
<td>
    <ul>
        <li>Goroutines are much more lightweigh than a thread. 1 Thread per multiple goroutines.
		<li>Goroutines are more optimal use less memory. Not a huge advantage over C# unless you're creating hundreds of collections.
    </ul>
</td>
<td>
    <ul>
        <li>Each task, uses a thread. 
		<li>Thread per task run is not serious in this example, but could become an issue if we needed 1000 collections.
    </ul>    
</td>
</tr>
<tr>
<td colspan=2>
Running the code above produces the following output
</td>
</tr>
<tr>
<td style="vertical-align: top;">

```ruby
Adding: 0
Adding: 1
Adding: 2
Adding: 3
Adding: 4
Adding: 5
Adding: 6
0
1
2
3
4
5
6
Adding: 7
Adding: 8
Adding: 9
7
8
9


```

</td>
<td style="vertical-align: top;">

```ruby
adding: 1
adding: 2
adding: 3
adding: 4
adding: 5
adding: 6
num: 1
num: 2
num: 3
num: 4
num: 5
adding: 7
adding: 8
adding: 9
adding: 10
num: 6
num: 7
num: 8
num: 9
num: 10
```

</td>
</tr>
</table>
