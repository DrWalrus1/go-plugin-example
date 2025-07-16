package main

import (
	"fmt"
	"std_plug/shared"
)

var V shared.SharedSymbol
var F shared.SharedFunc = func() { fmt.Printf("Hello, number %d\n", V) }
