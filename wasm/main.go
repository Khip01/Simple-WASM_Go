package main

import (
	"fmt"
	"syscall/js"
)

var htmlString = "<h4>Halo, salam dari Go!</h4>"

func GetHtml() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return htmlString
	})
}

func main() {

	ch := make(chan chan struct{}, 0)

	fmt.Println("Hello WASM From GO!")

	js.Global().Set("getHtml", GetHtml())

	<- ch
}
