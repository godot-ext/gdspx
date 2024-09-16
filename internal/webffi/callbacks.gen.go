/*
------------------------------------------------------------------------------
//   This code was generated by template ffi_wrapper.go.tmpl.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "ffi_wrapper.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------
*/
package webffi

import (
	"syscall/js"
)

func onSpxCallbackOnEngineStart(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnEngineStart")
	callbacks.OnEngineStart()
	return nil
}
func onSpxCallbackOnEngineUpdate(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnEngineUpdate")
	callbacks.OnEngineUpdate(0.016)
	return nil
}
func onSpxCallbackOnEngineFixedUpdate(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnEngineFixedUpdate")
	callbacks.OnEngineFixedUpdate(0.016)
	return nil
}
func onSpxCallbackOnEngineDestroy(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnEngineDestroy")
	callbacks.OnEngineDestroy()
	return nil
}
func onSpxCallbackOnSceneSpriteInstantiated(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSceneSpriteInstantiated")
	return nil
}
func onSpxCallbackOnSpriteReady(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteReady")
	return nil
}
func onSpxCallbackOnSpriteUpdated(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteUpdated")
	return nil
}
func onSpxCallbackOnSpriteFixedUpdated(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteFixedUpdated")
	return nil
}
func onSpxCallbackOnSpriteDestroyed(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteDestroyed")
	return nil
}
func onSpxCallbackOnSpriteFramesSetChanged(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteFramesSetChanged")
	return nil
}
func onSpxCallbackOnSpriteAnimationChanged(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteAnimationChanged")
	return nil
}
func onSpxCallbackOnSpriteFrameChanged(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteFrameChanged")
	return nil
}
func onSpxCallbackOnSpriteAnimationLooped(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteAnimationLooped")
	return nil
}
func onSpxCallbackOnSpriteAnimationFinished(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteAnimationFinished")
	return nil
}
func onSpxCallbackOnSpriteVfxFinished(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteVfxFinished")
	return nil
}
func onSpxCallbackOnSpriteScreenExited(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteScreenExited")
	return nil
}
func onSpxCallbackOnSpriteScreenEntered(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteScreenEntered")
	return nil
}
func onSpxCallbackOnMousePressed(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnMousePressed")
	return nil
}
func onSpxCallbackOnMouseReleased(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnMouseReleased")
	return nil
}
func onSpxCallbackOnKeyPressed(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnKeyPressed")
	return nil
}
func onSpxCallbackOnKeyReleased(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnKeyReleased")
	return nil
}
func onSpxCallbackOnActionPressed(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnActionPressed")
	return nil
}
func onSpxCallbackOnActionJustPressed(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnActionJustPressed")
	return nil
}
func onSpxCallbackOnActionJustReleased(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnActionJustReleased")
	return nil
}
func onSpxCallbackOnAxisChanged(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnAxisChanged")
	return nil
}
func onSpxCallbackOnCollisionEnter(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnCollisionEnter")
	return nil
}
func onSpxCallbackOnCollisionStay(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnCollisionStay")
	return nil
}
func onSpxCallbackOnCollisionExit(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnCollisionExit")
	return nil
}
func onSpxCallbackOnTriggerEnter(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnTriggerEnter")
	return nil
}
func onSpxCallbackOnTriggerStay(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnTriggerStay")
	return nil
}
func onSpxCallbackOnTriggerExit(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnTriggerExit")
	return nil
}
func onSpxCallbackOnUiReady(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiReady")
	return nil
}
func onSpxCallbackOnUiUpdated(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiUpdated")
	return nil
}
func onSpxCallbackOnUiDestroyed(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiDestroyed")
	return nil
}
func onSpxCallbackOnUiPressed(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiPressed")
	return nil
}
func onSpxCallbackOnUiReleased(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiReleased")
	return nil
}
func onSpxCallbackOnUiHovered(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiHovered")
	return nil
}
func onSpxCallbackOnUiClicked(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiClicked")
	return nil
}
func onSpxCallbackOnUiToggle(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiToggle")
	return nil
}
func onSpxCallbackOnUiTextChanged(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiTextChanged")
	return nil
}
func resiterFuncPtr2Js() {
	js.Global().Set("on_spx_callback_on_engine_start", js.FuncOf(onSpxCallbackOnEngineStart))
	js.Global().Set("on_spx_callback_on_engine_update", js.FuncOf(onSpxCallbackOnEngineUpdate))
	js.Global().Set("on_spx_callback_on_engine_fixed_update", js.FuncOf(onSpxCallbackOnEngineFixedUpdate))
	js.Global().Set("on_spx_callback_on_engine_destroy", js.FuncOf(onSpxCallbackOnEngineDestroy))
	js.Global().Set("on_spx_callback_on_scene_sprite_instantiated", js.FuncOf(onSpxCallbackOnSceneSpriteInstantiated))
	js.Global().Set("on_spx_callback_on_sprite_ready", js.FuncOf(onSpxCallbackOnSpriteReady))
	js.Global().Set("on_spx_callback_on_sprite_updated", js.FuncOf(onSpxCallbackOnSpriteUpdated))
	js.Global().Set("on_spx_callback_on_sprite_fixed_updated", js.FuncOf(onSpxCallbackOnSpriteFixedUpdated))
	js.Global().Set("on_spx_callback_on_sprite_destroyed", js.FuncOf(onSpxCallbackOnSpriteDestroyed))
	js.Global().Set("on_spx_callback_on_sprite_frames_set_changed", js.FuncOf(onSpxCallbackOnSpriteFramesSetChanged))
	js.Global().Set("on_spx_callback_on_sprite_animation_changed", js.FuncOf(onSpxCallbackOnSpriteAnimationChanged))
	js.Global().Set("on_spx_callback_on_sprite_frame_changed", js.FuncOf(onSpxCallbackOnSpriteFrameChanged))
	js.Global().Set("on_spx_callback_on_sprite_animation_looped", js.FuncOf(onSpxCallbackOnSpriteAnimationLooped))
	js.Global().Set("on_spx_callback_on_sprite_animation_finished", js.FuncOf(onSpxCallbackOnSpriteAnimationFinished))
	js.Global().Set("on_spx_callback_on_sprite_vfx_finished", js.FuncOf(onSpxCallbackOnSpriteVfxFinished))
	js.Global().Set("on_spx_callback_on_sprite_screen_exited", js.FuncOf(onSpxCallbackOnSpriteScreenExited))
	js.Global().Set("on_spx_callback_on_sprite_screen_entered", js.FuncOf(onSpxCallbackOnSpriteScreenEntered))
	js.Global().Set("on_spx_callback_on_mouse_pressed", js.FuncOf(onSpxCallbackOnMousePressed))
	js.Global().Set("on_spx_callback_on_mouse_released", js.FuncOf(onSpxCallbackOnMouseReleased))
	js.Global().Set("on_spx_callback_on_key_pressed", js.FuncOf(onSpxCallbackOnKeyPressed))
	js.Global().Set("on_spx_callback_on_key_released", js.FuncOf(onSpxCallbackOnKeyReleased))
	js.Global().Set("on_spx_callback_on_action_pressed", js.FuncOf(onSpxCallbackOnActionPressed))
	js.Global().Set("on_spx_callback_on_action_just_pressed", js.FuncOf(onSpxCallbackOnActionJustPressed))
	js.Global().Set("on_spx_callback_on_action_just_released", js.FuncOf(onSpxCallbackOnActionJustReleased))
	js.Global().Set("on_spx_callback_on_axis_changed", js.FuncOf(onSpxCallbackOnAxisChanged))
	js.Global().Set("on_spx_callback_on_collision_enter", js.FuncOf(onSpxCallbackOnCollisionEnter))
	js.Global().Set("on_spx_callback_on_collision_stay", js.FuncOf(onSpxCallbackOnCollisionStay))
	js.Global().Set("on_spx_callback_on_collision_exit", js.FuncOf(onSpxCallbackOnCollisionExit))
	js.Global().Set("on_spx_callback_on_trigger_enter", js.FuncOf(onSpxCallbackOnTriggerEnter))
	js.Global().Set("on_spx_callback_on_trigger_stay", js.FuncOf(onSpxCallbackOnTriggerStay))
	js.Global().Set("on_spx_callback_on_trigger_exit", js.FuncOf(onSpxCallbackOnTriggerExit))
	js.Global().Set("on_spx_callback_on_ui_ready", js.FuncOf(onSpxCallbackOnUiReady))
	js.Global().Set("on_spx_callback_on_ui_updated", js.FuncOf(onSpxCallbackOnUiUpdated))
	js.Global().Set("on_spx_callback_on_ui_destroyed", js.FuncOf(onSpxCallbackOnUiDestroyed))
	js.Global().Set("on_spx_callback_on_ui_pressed", js.FuncOf(onSpxCallbackOnUiPressed))
	js.Global().Set("on_spx_callback_on_ui_released", js.FuncOf(onSpxCallbackOnUiReleased))
	js.Global().Set("on_spx_callback_on_ui_hovered", js.FuncOf(onSpxCallbackOnUiHovered))
	js.Global().Set("on_spx_callback_on_ui_clicked", js.FuncOf(onSpxCallbackOnUiClicked))
	js.Global().Set("on_spx_callback_on_ui_toggle", js.FuncOf(onSpxCallbackOnUiToggle))
	js.Global().Set("on_spx_callback_on_ui_text_changed", js.FuncOf(onSpxCallbackOnUiTextChanged))
}
