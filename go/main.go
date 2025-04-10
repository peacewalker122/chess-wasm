//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/peacewalker122/chess-wasm/go/engine"
)

func main() {
	fmt.Println("Go WebAssembly Initialized")

	// Create channel to keep main function running
	ch := make(chan struct{}, 1)

	// Example: Set up a JavaScript function
	js.Global().Set("greetFromGo", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		message := "Hello from Go WebAssembly!"
		return message
	}))
	js.Global().Set("buildBoard", js.FuncOf(func(this js.Value, args []js.Value) any {
		var param *bool

		if this.IsNull() {
			_true := true
			param = &_true
		} else {
			val := this.Bool()
			param = &val
		}

		val := engine.CreateBoard(param)

		result, err := json.Marshal(val)
		if err != nil {
			return js.ValueOf(err)
		}

		return js.ValueOf(result)
	}))

	// Keep the program running
	<-ch
}
