package main

import (
	"log"

	"syscall/js"
	"wasm_sample/calc"
	"wasm_sample/image"
)

func main() {
	ch := make(chan struct{})

	grayFn := js.FuncOf(grayScale)
	defer grayFn.Release()

	fibFn := js.FuncOf(fib)
	defer fibFn.Release()

	fibMemFn := js.FuncOf(fibMem)
	defer fibMemFn.Release()

	js.Global().Set("grayScale", grayFn)
	js.Global().Set("fibonacci", fibFn)
	js.Global().Set("fibonacciMemorized", fibMemFn)
	<-ch
}

func grayScale(_ js.Value, args []js.Value) interface{} {
	enc, err := image.GrayScale(args[0].String())
	if err != nil {
		log.Printf("error: %v", err)
	}
	return enc
}

func fib(_ js.Value, args []js.Value) interface{} {
	return calc.Fibonacci(args[0].Int())
}

func fibMem(_ js.Value, args []js.Value) interface{} {
	return calc.FibonacciMemorized(args[0].Int())
}
