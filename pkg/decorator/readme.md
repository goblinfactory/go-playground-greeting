# decorator

Decorator example written in Go, with equivalent C# code

```go
    type DecoratedAdd func(a int, b int) int

    // this is such a quick way to implement a decorator
    func (f DecoratedAdd) Add(a int, b int) int {
        fmt.Println("before")
        defer fmt.Println("after")
        r := f(a, b)
        fmt.Println("result", r)
        return r
    }

    func Test_howToUseDecorator(t *testing.T) {
        ad := DecoratedAdd(AddNums)
        r := ad.Add(1, 3)
        if r != 4 {
            t.Error("expected 4")
        }
    }

    func AddNums(a int, b int) int {
        return a + b
    }
```

Same code in C#

[see it live in jsfiddle](https://dotnetfiddle.net/9bRx4e)

```csharp
using System;

public class Program
{
	public static void Main()
	{
		var decoratedLogin = Decorate(Program.Login);
		var result = decoratedLogin(10, "123");
		Console.WriteLine(result);
	}

	public static Func<int,string,bool> Decorate(Func<int,string,bool> function) {
		return (a, b)=>  {
			Console.WriteLine("before");
			try {
				return function(a, b);
			}
			finally {
				Console.WriteLine("after");
			}
		};
	}

	public static bool Login(int id, string pw)  {
		Console.WriteLine("logging in");
		return (id == 10 && pw == "123") ;
	}

	public static bool IsActive(int id, string sessionId)  {
		Console.WriteLine("checking session");
		return (id == 10 && sessionId == "456") ;
	}


}
```
