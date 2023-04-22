//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/abhisekp/webassembly/pkg/json"
)

func jsonWrapper() js.Func {
	result := map[string]any{}
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			result["error"] = "Invalid no of arguments passed"
			return result
		}
		jsDoc := js.Global().Get("document")
		if !jsDoc.Truthy() {
			result["error"] = "document not found"
			return result
		}

		jsonOutputTextArea := jsDoc.Call("getElementById", "jsonoutput")
		if !jsonOutputTextArea.Truthy() {
			result["error"] = "jsonoutput not found"
			return result
		}

		inputJSON := args[0].String()
		fmt.Printf("input %s\n", inputJSON)
		pretty, err := json.PrettyJson(inputJSON)
		if err != nil {
			fmt.Printf("unable to convert to json %s\n", err)
			result["error"] = err.Error()
			return result
		}
		jsonOutputTextArea.Set("value", pretty)
		return nil
	})
	return jsonFunc
}
