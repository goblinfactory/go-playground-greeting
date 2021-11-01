package concurrencypatterns

// WowThisIsSubtle shows examples of noobie Go code that can really trip you up. Code examples show using a goRoutine to generate a list
// this is something you shouldnt do, (wrong use of concurrency) but is show here because it's interesting to look at the details
// as well as compare how the similar code can be written in C#
func WowThisIsSubtle() {

}

// so this func will return a goRoutine that will generate 10 numbers

// This is similar to the following C# code
// code below can be accessed at
// https://dotnetfiddle.net/rJdqJV

/*

using System;
using System.Collections.Concurrent;
using System.Threading;
using System.Threading.Tasks;

public class Program
{
	public static void Main()
	{
		Console.WriteLine("hello:{0}", 10);
		var nums = GenerateNums(10);

		foreach(var n in nums.GetConsumingEnumerable()) {
			Console.WriteLine("num: {0}", n);
		}
	}

	public static BlockingCollection<int> GenerateNums(int max) {
		var bc = new BlockingCollection<int>(1);
		// start a task to push 10 ints to the the blocking collection
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



*/

// generates the following output

/*

hello:10
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

*/
