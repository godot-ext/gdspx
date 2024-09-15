package webffi

import (
	. "godot-ext/gdspx/pkg/engine"
	"syscall/js"
)

func ToGdString(object string) js.Value {
	return js.ValueOf(object)
}

// ToObject converts Object to JS object
func ToGdObj(object Object) js.Value {
	return js.ValueOf(int64(object))
}

// ToGdVec2 converts Vec2 to JS object
func ToGdVec2(vec Vec2) js.Value {
	vec2Js := js.Global().Get("Object").New()
	vec2Js.Set("X", vec.X)
	vec2Js.Set("Y", vec.Y)
	return vec2Js
}

// ToGdVec3 converts Vec3 to JS object
func ToGdVec3(vec Vec3) js.Value {
	vec3Js := js.Global().Get("Object").New()
	vec3Js.Set("X", vec.X)
	vec3Js.Set("Y", vec.Y)
	vec3Js.Set("Z", vec.Z)
	return vec3Js
}

// ToGdVec4 converts Vec4 to JS object
func ToGdVec4(vec Vec4) js.Value {
	vec4Js := js.Global().Get("Object").New()
	vec4Js.Set("X", vec.X)
	vec4Js.Set("Y", vec.Y)
	vec4Js.Set("Z", vec.Z)
	vec4Js.Set("W", vec.W)
	return vec4Js
}

// ToGdColor converts Color to JS object
func ToGdColor(color Color) js.Value {
	colorJs := js.Global().Get("Object").New()
	colorJs.Set("R", color.R)
	colorJs.Set("G", color.G)
	colorJs.Set("B", color.B)
	colorJs.Set("A", color.A)
	return colorJs
}

// ToGdRect2 converts Rect2 to JS object
func ToGdRect2(rect Rect2) js.Value {
	rectJs := js.Global().Get("Object").New()
	rectJs.Set("Center", ToGdVec2(rect.Center))
	rectJs.Set("Size", ToGdVec2(rect.Size))
	return rectJs
}

// ToGdBool converts bool to JS object
func ToGdBool(val bool) js.Value {
	return js.ValueOf(val)
}

// ToGdInt converts int64 to JS object
func ToGdInt(val int64) js.Value {
	return js.ValueOf(val)
}

// ToGdFloat converts float32 to JS object
func ToGdFloat(val float32) js.Value {
	return js.ValueOf(val)
}

// ToObj converts JS object to Object (int64)
func ToString(object js.Value) string {
	return object.String()
}

// ToObj converts JS object to Object (int64)
func ToObj(object js.Value) Object {
	value := int64(object.Int())
	return Object(value)
}

// ToObj converts JS object to Object (int64)
func ToObject(object js.Value) Object {
	value := int64(object.Int())
	return Object(value)
}

// ToVec2 converts JS object to Vec2
func ToVec2(vec js.Value) Vec2 {
	return Vec2{
		X: float32(vec.Get("X").Float()),
		Y: float32(vec.Get("Y").Float()),
	}
}

// ToVec3 converts JS object to Vec3
func ToVec3(vec js.Value) Vec3 {
	return Vec3{
		X: float32(vec.Get("X").Float()),
		Y: float32(vec.Get("Y").Float()),
		Z: float32(vec.Get("Z").Float()),
	}
}

// ToVec4 converts JS object to Vec4
func ToVec4(vec js.Value) Vec4 {
	return Vec4{
		X: float32(vec.Get("X").Float()),
		Y: float32(vec.Get("Y").Float()),
		Z: float32(vec.Get("Z").Float()),
		W: float32(vec.Get("W").Float()),
	}
}

// ToColor converts JS object to Color
func ToColor(color js.Value) Color {
	return Color{
		R: float32(color.Get("R").Float()),
		G: float32(color.Get("G").Float()),
		B: float32(color.Get("B").Float()),
		A: float32(color.Get("A").Float()),
	}
}

// ToRect2 converts JS object to Rect2
func ToRect2(rect js.Value) Rect2 {
	return Rect2{
		Center: ToVec2(rect.Get("Center")),
		Size:   ToVec2(rect.Get("Size")),
	}
}

// ToBool converts JS object to bool
func ToBool(val js.Value) bool {
	return val.Bool()
}

// ToInt converts JS object to int64
func ToInt(val js.Value) int64 {
	return int64(val.Int())
}

// ToFloat converts JS object to float32
func ToFloat(val js.Value) float32 {
	return float32(val.Float())
}

// ToFloat converts JS object to float32
func ToFloat32(val js.Value) float32 {
	return float32(val.Float())
}

// ToFloat converts JS object to float32
func ToInt64(val js.Value) int64 {
	return int64(val.Int())
}
