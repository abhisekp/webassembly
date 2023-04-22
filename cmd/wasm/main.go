//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	js.Global().Set("formatJSON", jsonWrapper())
	<-make(chan bool)
}
