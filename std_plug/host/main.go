package main

import (
	"plugin"
	"std_plug/shared"
)

func main() {

	p, err := plugin.Open("../plugin/plugin.so")
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}
	*v.(*shared.SharedSymbol) = 7
	f_cast := *f.(*shared.SharedFunc) // prints "Hello, number 7"
	f_cast()
}
