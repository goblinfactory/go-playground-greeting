# no linq

```go
    // querying objects in a similar to linq style, in go.
    var cars = []car{
        {1950, "Jay Leno", "buic", "red", 1},
        {1965, "Jay Leno", "chrysler", "blue", 2},
        {2000, "Fred", "bmw", "black", 3},
        {2010, "Dan", "volvo", "red", 4},
    }

    var g = garage(cars)

    fmt.Println("\nVintage\n", g.vintage())
    fmt.Println("\nRed\n", g.color("red"))
    fmt.Println("\nVintage AND red\n", g.vintage().color("red"))
    fmt.Println("\nVintage OR red\n", g.vintage().or(g.color("red")))
```

produces

```log

Vintage
[{1950 Jay Leno buic red 1} {1965 Jay Leno chrysler blue 2}]

Red
[{1950 Jay Leno buic red 1} {2010 Dan volvo red 4}]

Vintage AND red
[{1950 Jay Leno buic red 1}]

Vintage OR red
[{1950 Jay Leno buic red 1} {1965 Jay Leno chrysler blue 2} {2010 Dan volvo red 4}]

```
