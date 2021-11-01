# Compare Golang channels withC# BlockingCollection & Threads

Don't do this; i.e. don't use channels and functions to create collections. It's the wrong use of concurrecy; This contrived example is shown here so that we can have some insights to compare channels to C# BlockingCollection and Threads.

Typically this might be an expensive operation, calculateExpensiveFoo. I've not used long names in the example below, because the code doesnt fit in side by side when i do! Please use your imagination. 😇

<table style="padding:0px">
<tr>
<th>Go</th>
<th>[C#](https://dotnetfiddle.net/ReDK9q)</th>
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
	ch := make(chan int)
	go func() {
		for i := 0; i < cnt; i++ {
			fmt.Printf("Adding: %d\n", i)
			ch <- i
		}
		close(ch)
	}()
	return ch
}
```

</td>
<td style="vertical-align:top;" >

```csharp

public static void Main() {
	foreach(var n in GenerateNums(10)) {
		Console.WriteLine("num: {0}", n);
	}
}

public static IEnumerable<int> GenerateNums(int cnt) {
	var bc = new BlockingCollection<int>(1);
	Task.Run(()=> {
		for(int i=1; i<cnt; i++) {
			Console.WriteLine("adding: {0}", i);
			bc.Add(i);
		}
		bc.CompleteAdding();
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
    </ul>
</td>
<td>
    <ul>
        <li>Each task, uses a thread.
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
0
1
Adding: 2
Adding: 3
2
3
Adding: 4
Adding: 5
4
5
Adding: 6
Adding: 7
6
7
Adding: 8
Adding: 9
8
9

```

</td>
<td style="vertical-align: top;">

```ruby
adding: 1
adding: 2
adding: 3
num: 1
adding: 4
num: 2
num: 3
adding: 5
num: 4
adding: 6
num: 5
adding: 7
num: 6
adding: 8
num: 7
adding: 9
num: 8
num: 9
adding: 10
num: 10
```

</td>
</tr>
</table>
