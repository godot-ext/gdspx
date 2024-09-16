package webffi

import (
	"godot-ext/gdspx/pkg/engine"
	"syscall/js"
)

var (
	callbacks engine.CallbackInfo
)

func Link() bool {
	js.Global().Set("goWasmInit", js.FuncOf(goWasmInit))
	return true
}

func goWasmInit(this js.Value, args []js.Value) interface{} {
	println("Go wasm init succ!")
	resiterFuncPtr2Js()
	return js.ValueOf(nil)
}

func BindCallback(info engine.CallbackInfo) {
	callbacks = info
}

func dlsymGD(funcName string) js.Value {
	return js.Global().Get(funcName)
}
