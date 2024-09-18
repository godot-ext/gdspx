package webffi

import (
	"fmt"
	. "godot-ext/gdspx/pkg/engine"
	"syscall/js"
)

// ToObject converts Object to JS object
func JsFromGdObj(val Object) js.Value {
	return JsFromGdInt(int64(val))
}

// ToGdInt converts int64 to JS object
func JsFromGdInt(val int64) js.Value {
	vec2Js := js.Global().Get("Object").New()

	low := uint32(val & 0xFFFFFFFF)
	high := uint32((val >> 32) & 0xFFFFFFFF)
	vec2Js.Set("low", low)
	vec2Js.Set("high", high)
	return vec2Js
}

// ToObj converts JS object to Object (int64)
func JsToGdObject(val js.Value) Object {
	return Object(JsToGdInt(val))
}

func JsToGdObj(val js.Value) int64 {
	return JsToGdInt(val)
}

// ToInt converts JS object to int64
func JsToGdInt(val js.Value) int64 {
	low := uint32(val.Get("low").Int())
	high := uint32(val.Get("high").Int())

	int64Value := int64(high)<<32 | int64(low)
	fmt.Printf("Received int64: %d\n", int64Value)
	return int64Value
}

func JsFromGdString(object string) js.Value {
	return js.ValueOf(object)
}

// ToGdVec2 converts Vec2 to JS object
func JsFromGdVec2(vec Vec2) js.Value {
	vec2Js := js.Global().Get("Object").New()
	vec2Js.Set("x", vec.X)
	vec2Js.Set("y", vec.Y)
	return vec2Js
}

// ToGdVec3 converts Vec3 to JS object
func JsFromGdVec3(vec Vec3) js.Value {
	vec3Js := js.Global().Get("Object").New()
	vec3Js.Set("x", vec.X)
	vec3Js.Set("y", vec.Y)
	vec3Js.Set("z", vec.Z)
	return vec3Js
}

// ToGdVec4 converts Vec4 to JS object
func JsFromGdVec4(vec Vec4) js.Value {
	vec4Js := js.Global().Get("Object").New()
	vec4Js.Set("x", vec.X)
	vec4Js.Set("y", vec.Y)
	vec4Js.Set("z", vec.Z)
	vec4Js.Set("w", vec.W)
	return vec4Js
}

// ToGdColor converts Color to JS object
func JsFromGdColor(color Color) js.Value {
	colorJs := js.Global().Get("Object").New()
	colorJs.Set("r", color.R)
	colorJs.Set("g", color.G)
	colorJs.Set("b", color.B)
	colorJs.Set("a", color.A)
	return colorJs
}

// ToGdRect2 converts Rect2 to JS object
func JsFromGdRect2(rect Rect2) js.Value {
	rectJs := js.Global().Get("Object").New()
	rectJs.Set("center", JsFromGdVec2(rect.Center))
	rectJs.Set("size", JsFromGdVec2(rect.Size))
	return rectJs
}

// ToGdBool converts bool to JS object
func JsFromGdBool(val bool) js.Value {
	return js.ValueOf(val)
}

// ToGdFloat converts float32 to JS object
func JsFromGdFloat(val float32) js.Value {
	return js.ValueOf(val)
}

// ToObj converts JS object to Object (int64)
func JsToGdString(object js.Value) string {
	return object.String()
}

// ToVec2 converts JS object to Vec2
func JsToGdVec2(vec js.Value) Vec2 {
	return Vec2{
		X: float32(vec.Get("x").Float()),
		Y: float32(vec.Get("y").Float()),
	}
}

// ToVec3 converts JS object to Vec3
func JsToGdVec3(vec js.Value) Vec3 {
	return Vec3{
		X: float32(vec.Get("x").Float()),
		Y: float32(vec.Get("y").Float()),
		Z: float32(vec.Get("z").Float()),
	}
}

// ToVec4 converts JS object to Vec4
func JsToGdVec4(vec js.Value) Vec4 {
	return Vec4{
		X: float32(vec.Get("x").Float()),
		Y: float32(vec.Get("y").Float()),
		Z: float32(vec.Get("z").Float()),
		W: float32(vec.Get("w").Float()),
	}
}

// ToColor converts JS object to Color
func JsToGdColor(color js.Value) Color {
	return Color{
		R: float32(color.Get("r").Float()),
		G: float32(color.Get("g").Float()),
		B: float32(color.Get("b").Float()),
		A: float32(color.Get("a").Float()),
	}
}

// ToRect2 converts JS object to Rect2
func JsToGdRect2(rect js.Value) Rect2 {
	return Rect2{
		Center: JsToGdVec2(rect.Get("center")),
		Size:   JsToGdVec2(rect.Get("size")),
	}
}

// ToBool converts JS object to bool
func JsToGdBool(val js.Value) bool {
	return val.Bool()
}

// ToFloat converts JS object to float32
func JsToGdFloat(val js.Value) float32 {
	return float32(val.Float())
}

// ToFloat converts JS object to float32
func JsToGdFloat32(val js.Value) float32 {
	return float32(val.Float())
}

// ToFloat converts JS object to float32
func JsToGdInt64(val js.Value) int64 {
	return int64(val.Int())
}
