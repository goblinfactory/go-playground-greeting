# Go select compared to various C# alternatives

PLACEHOLDER PAGE (still need to work through this)

Considering the following go code, what would be the closest equivalent idiomatic C# ?

```go

done := make(chan struct{})

select {
    case ch1 <- someFunc()
    case <- ch2
    case <- done
}

```

Code above needs to be made self explanatory with the smallest real word scenario.

**Possible scenarios to consider?**

-   Returning first from a cluster of nodes to respond. E.g. during a mission critical operation, e.g. rocket launch, you can't afford for 1 server to hang for a few seconds during garbage collection. Not that you'd possibly (?) use a GC language for this type of automation, but it's a useful analogy.
