# random things about go 

Random thoughts about go, from 4 C# developers - initial thoughts after first impressions with some learning.

*   `9.11.21 [ @Alan ]`

    *I like that even if your code has the incorrect number of arguments to a function, you can always press F12 from VSCode and go straight to the implementation. In C# if you have the wrong arguments, the compiler wont go anywhere since there's no matching overloaded function. Go can only ever have 1 definition for a function, so F12 always knows where to go irrespective of what parameters you are calling the func with.*

*   `19.11.21 [ @T ]` 

    *I think part of the reason we've been enjoying go is that it's pretty much just notepad and a terminal window (albeit via VSCode or Goland). It's back to basics and it's fun!*

*   `26.11.21 [ @Alan ]`

    *One of my biggest takeaways (this week) is that go supports you to trivially easily seperate your concerns. In C# you would need to put everything you want to make private to a concern private inside a class. In Go everything is private inside a package and you make a concious choice what to make part of your packages API. You think about a solution or a project as being made up of building blocks of functionality "packages", and each package is a simple folder. To get this level of seperation in .NET you literally have to move all the code to new project, or move all the code inside a class as a private methods and classes of a class. This is totally non idiomatic, and even if you decided to do this, would require that you then either put all your code inside a single file, or use partial classes.*

    *actually it's more of a case that go trivially gives you a solid IRON CURTAIN between your concerns... you CANT mess with anything that's not deliberately exported as part of a defined API, ...and that sounds honerous, but really what this means, is at some point your decide what funcs get a capital letter, that's it!*

*   `28.11.21 [ @Alan ]`

    *Go has a really clever (pragmatic) way of doing mostly non nullable types in the language. structs (go's classes and nested properties that are structs) cannot be nil. Also, while maps can be set to nil, the language sets it to a special empty nil. so it takes no memory, but will return default zero value for any requests to retrieve keys, without throwing exceptions.* 

    *e.g. if I have type `person struct { name string, brother person }` and try to do `p := person { "me", nil }` <-- I will get compile time error, brother cannot be nil, nice red squiggly. You can't write code that will set it to nil either. e.g. func getBrother() brother {} ... could never return nil.*

    *while this is possible in C# is really been shoehorned into the language and is messy and projects can be really hard to make non nillable.*