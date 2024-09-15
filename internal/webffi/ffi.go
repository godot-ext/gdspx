package webffi

import (
	"godot-ext/gdspx/pkg/engine"
	"syscall/js"
)

var (
	callbacks engine.CallbackInfo
)

func Link() bool {
	return true
}

func BindCallback(info engine.CallbackInfo) {
	callbacks = info
}

func dlsymGD(funcName string) js.Value {
	return js.Global().Get(funcName)
}
