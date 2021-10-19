package di

// Greeter interface
type Greeter interface {
	Greet(name string)
}

// provider

// Cat consumer dependant on Greeter
type Cat struct {
	Name    string
	greeter Greeter
}

// SayHello ..
func (c Cat) SayHello(say string) {
	c.greeter.Greet(say + " from " + c.Name)
}
