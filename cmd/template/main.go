package main

import (
	. "godot-ext/gdspx/pkg/engine"
	"godot-ext/gdspx/pkg/gdspx"
)

func main() {
	gdspx.LinkEngine(EngineCallbackInfo{
		OnEngineStart:   onStart,
		OnEngineUpdate:  onUpdate,
		OnEngineDestroy: onDestory,
	})
}

func onStart() {
	println("onEngineStart")
	obj := NewSprite("")
	obj2 := NewSprite("")
	obj.SetPosition(Vec2{100, 100})
	obj2.SetPosition(Vec2{110, 110})
	obj.SetTexture("res://icon.png")
	obj2.SetTexture("res://icon.png")
}

func onUpdate(delta float32) {
	//println("onEngineUpdate")
}

func onDestory() {
	println("onEngineDestory")
}
