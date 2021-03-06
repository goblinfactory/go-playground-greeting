# Compare Go channels with C#'s BlockingCollection & Threads

Don't do this; i.e. don't use channels and functions to create collections. It's the wrong use of concurrecy; This contrived example is shown here so that we can have some insights to compare channels to C# BlockingCollection and Threads. Typically this might be an expensive operation, calculateExpensiveFoo. I've not used long names in the example below, because the code doesnt fit in side by side when i do! Please use your imagination. 😇

<table style="padding:0px">
<tr>
<th>Go</th>
<th>[C#]</th>
</tr>
<tr>
	<td>
		https://play.golang.org/p/F8kAX9Xi81u
	</td>
	<td>
		https://dotnetfiddle.net/6Cabjc
	</td>
</tr>
<tr>
<td style="vertical-align:top;">

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
<td style="vertical-align:top;">

```csharp
public static void Main() {
	foreach(var n in GenerateNums(10)) {
		Console.WriteLine("num: {0}", n);
	}
}

public static IEnumerable<int> GenerateNums(int cnt) {
	var bc = new BlockingCollection<int>(5);
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
		<li>Goroutines are more optimal use less memory. Not an advantage in trivial examples but is serious when you're storing and creating hundreds of collections.
		<li>defer guarantees the channel closes, and doesnt cause extra indenting that makes code harder to read. Many cases are simply 1 liners.
		<li>Code does <b>not</b> need to be modified when used asychronously.
    </ul>
</td>
<td>
    <ul>
        <li>Each task, uses a thread. 
		<li>Thread per task run is not serious in this example, but could become an issue if we needed 1000 collections.
		<li>return type and code needs to change if you want to call this code using async.
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

This is a draft : Still need to add to the bottom where this pattern goes horribly wrong (deadlocks)and how to fix it (improve) the code in both C# and Go. Real Go advantages start to kick in with select. Need to add that below.

# Compare Go's select vs C#

`select` allows you to block until any of a number of blocking channels `unblocks` with waiting channel data. This is a complex software pattern to solve in other languages and is **imho** one of the core unsung differentiators between go and other languages. I'll attempt to find the best idiomatic C# equivalent and compare the two with real typical use cases for both languages. The examples will be more complex and won't be able to be shown side by side due to limit of column width with github markdown.

TBD.

update: 16:25 Mon, 1 Nov: I have just stumpled across C# Channels for the first time (`System.Threading.Tasks.Channels`). So the comparison above may or may not be relevant, or the best comparison. Going to put a pin in this and come back to this later.

Below are the links that I stumbled across, and not finished reading yet : it's a lot to take in.

-   https://alexyakunin.medium.com/go-vs-c-part-1-goroutines-vs-async-await-ac909c651c11
-   https://www.dotnetcurry.com/dotnetcore/1509/async-dotnetcore-pattern

My mates keep telling me to stop doing side by side comparison, and just focus on learning Go.

Sob, ...I may have to do just that.
