# Compare Golang channels withC# BlockingCollection & Threads

Don't do this; i.e. don't use channels and functions to create collections. It's the wrong use of concurrecy; This contrived example is shown here so that we can have some insights to compare channels to C# BlockingCollection and Threads.

<table>
<tr>
<th>Go</th>
<th>C#</th>
</tr>
<tr>
<td style="vertical-align: top;">

```go
package concurrencypatterns

import "fmt"

func main() {
	for n := range generateNumbers(10) {
		fmt.Println(n)
	}
}

func generateNumbers(cnt int) <-chan int {
	ch := make(chan int)
	// start a goroutine to push 10 ints to the channel
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
<td style="vertical-align: top;" >

```csharp
using System;
using System.Collections.Concurrent;
using System.Threading;
using System.Threading.Tasks;

public class Program
{
	public static void Main()
	{
		var nums = GenerateNums(10);

		foreach(var n in nums.GetConsumingEnumerable()) {
			Console.WriteLine("num: {0}", n);
		}
	}

	public static BlockingCollection<int> GenerateNums(int max) {
		var bc = new BlockingCollection<int>(1);
		// start a task to push 10 ints to the blocking collection
		Task.Run(()=> {
			for(int i=1; i<11; i++) {
				Console.WriteLine("adding: {0}", i);
				bc.Add(i);
			}
			bc.CompleteAdding();
		});
		return bc;
	}
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
