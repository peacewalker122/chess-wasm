//go:build js && wasm
// +build js,wasm

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"syscall/js"

	"github.com/peacewalker122/chess-wasm/go/engine"
	"github.com/rs/zerolog"
)

func main() {
	fmt.Println("Go WebAssembly Initialized")

	// Create channel to keep main function running
	ch := make(chan struct{}, 1)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	game := &engine.Game{}

	// Example: Set up a JavaScript function
	js.Global().Set("greetFromGo", js.FuncOf(func(this js.Value, args []js.Value) any {
		message := "Hello from Go WebAssembly!"
		return message
	}))
	js.Global().Set("buildBoard", js.FuncOf(func(this js.Value, args []js.Value) any {
		_true := true
		val := engine.CreateBoard(&_true)
		game.Board = val

		result, err := json.Marshal(val)
		if err != nil {
			return js.ValueOf(err.Error())
		}

		// Convert the byte slice to a string and return it
		// JavaScript can then parse this string using JSON.parse()
		return js.ValueOf(string(result))
	}))
	js.Global().Set("getBoard", js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Printf("Argument is: %+v\n", this)

		result, err := json.Marshal(game.Board)
		if err != nil {
			return js.ValueOf(err.Error())
		}

		// Convert the byte slice to a string and return it
		// JavaScript can then parse this string using JSON.parse()
		return js.ValueOf(string(result))
	}))
	js.Global().Set("startMove", js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Printf("Argument is: %+v\n", this)

		if len(args) < 2 {
			return js.ValueOf(errors.New("minimum args aren't achieved"))
		}

		arg1 := args[0]
		arg2 := args[1]

		board, err := game.CreateMove(arg1.String(), arg2.String())
		if err != nil {
			return js.ValueOf(err)
		}

		result, err := json.Marshal(board)
		if err != nil {
			return js.ValueOf(err.Error())
		}

		return js.ValueOf(string(result))
	}))

	// Keep the program running
	<-ch
}
