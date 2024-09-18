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

func gdspxOnEngineStart(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnEngineStart")
	callbacks.OnEngineStart()
	return nil
}
func gdspxOnEngineUpdate(this js.Value, args []js.Value) interface{} {
	callbacks.OnEngineUpdate(0.016)
	return nil
}
func gdspxOnEngineFixedUpdate(this js.Value, args []js.Value) interface{} {
	callbacks.OnEngineFixedUpdate(0.016)
	return nil
}
func gdspxOnEngineDestroy(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnEngineDestroy")
	callbacks.OnEngineDestroy()
	return nil
}
func gdspxOnSceneSpriteInstantiated(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSceneSpriteInstantiated")
	//callbacks.OnSceneSpriteInstantiated()
	return nil
}
func gdspxOnSpriteReady(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteReady")
	//callbacks.OnSpriteReady()
	return nil
}
func gdspxOnSpriteUpdated(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteUpdated")
	//callbacks.OnSpriteUpdated()
	return nil
}
func gdspxOnSpriteFixedUpdated(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteFixedUpdated")
	//callbacks.OnSpriteFixedUpdated()
	return nil
}
func gdspxOnSpriteDestroyed(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteDestroyed")
	//callbacks.OnSpriteDestroyed()
	return nil
}
func gdspxOnSpriteFramesSetChanged(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteFramesSetChanged")
	//callbacks.OnSpriteFramesSetChanged()
	return nil
}
func gdspxOnSpriteAnimationChanged(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteAnimationChanged")
	//callbacks.OnSpriteAnimationChanged()
	return nil
}
func gdspxOnSpriteFrameChanged(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteFrameChanged")
	//callbacks.OnSpriteFrameChanged()
	return nil
}
func gdspxOnSpriteAnimationLooped(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteAnimationLooped")
	//callbacks.OnSpriteAnimationLooped()
	return nil
}
func gdspxOnSpriteAnimationFinished(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteAnimationFinished")
	//callbacks.OnSpriteAnimationFinished()
	return nil
}
func gdspxOnSpriteVfxFinished(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteVfxFinished")
	//callbacks.OnSpriteVfxFinished()
	return nil
}
func gdspxOnSpriteScreenExited(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteScreenExited")
	//callbacks.OnSpriteScreenExited()
	return nil
}
func gdspxOnSpriteScreenEntered(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnSpriteScreenEntered")
	//callbacks.OnSpriteScreenEntered()
	return nil
}
func gdspxOnMousePressed(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnMousePressed")
	//callbacks.OnMousePressed()
	return nil
}
func gdspxOnMouseReleased(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnMouseReleased")
	//callbacks.OnMouseReleased()
	return nil
}
func gdspxOnKeyPressed(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnKeyPressed")
	//callbacks.OnKeyPressed()
	return nil
}
func gdspxOnKeyReleased(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnKeyReleased")
	//callbacks.OnKeyReleased()
	return nil
}
func gdspxOnActionPressed(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnActionPressed")
	//callbacks.OnActionPressed()
	return nil
}
func gdspxOnActionJustPressed(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnActionJustPressed")
	//callbacks.OnActionJustPressed()
	return nil
}
func gdspxOnActionJustReleased(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnActionJustReleased")
	//callbacks.OnActionJustReleased()
	return nil
}
func gdspxOnAxisChanged(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnAxisChanged")
	//callbacks.OnAxisChanged()
	return nil
}
func gdspxOnCollisionEnter(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnCollisionEnter")
	//callbacks.OnCollisionEnter()
	return nil
}
func gdspxOnCollisionStay(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnCollisionStay")
	//callbacks.OnCollisionStay()
	return nil
}
func gdspxOnCollisionExit(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnCollisionExit")
	//callbacks.OnCollisionExit()
	return nil
}
func gdspxOnTriggerEnter(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnTriggerEnter")
	//callbacks.OnTriggerEnter()
	return nil
}
func gdspxOnTriggerStay(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnTriggerStay")
	//callbacks.OnTriggerStay()
	return nil
}
func gdspxOnTriggerExit(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnTriggerExit")
	//callbacks.OnTriggerExit()
	return nil
}
func gdspxOnUiReady(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiReady")
	//callbacks.OnUiReady()
	return nil
}
func gdspxOnUiUpdated(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiUpdated")
	//callbacks.OnUiUpdated()
	return nil
}
func gdspxOnUiDestroyed(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiDestroyed")
	//callbacks.OnUiDestroyed()
	return nil
}
func gdspxOnUiPressed(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiPressed")
	//callbacks.OnUiPressed()
	return nil
}
func gdspxOnUiReleased(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiReleased")
	//callbacks.OnUiReleased()
	return nil
}
func gdspxOnUiHovered(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiHovered")
	//callbacks.OnUiHovered()
	return nil
}
func gdspxOnUiClicked(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiClicked")
	//callbacks.OnUiClicked()
	return nil
}
func gdspxOnUiToggle(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiToggle")
	//callbacks.OnUiToggle()
	return nil
}
func gdspxOnUiTextChanged(this js.Value, args []js.Value) interface{} {
	println("GDExtensionSpxCallbackOnUiTextChanged")
	//callbacks.OnUiTextChanged()
	return nil
}
func resiterFuncPtr2Js() {
	js.Global().Set("gdspx_on_engine_start", js.FuncOf(gdspxOnEngineStart))
	js.Global().Set("gdspx_on_engine_update", js.FuncOf(gdspxOnEngineUpdate))
	js.Global().Set("gdspx_on_engine_fixed_update", js.FuncOf(gdspxOnEngineFixedUpdate))
	js.Global().Set("gdspx_on_engine_destroy", js.FuncOf(gdspxOnEngineDestroy))
	js.Global().Set("gdspx_on_scene_sprite_instantiated", js.FuncOf(gdspxOnSceneSpriteInstantiated))
	js.Global().Set("gdspx_on_sprite_ready", js.FuncOf(gdspxOnSpriteReady))
	js.Global().Set("gdspx_on_sprite_updated", js.FuncOf(gdspxOnSpriteUpdated))
	js.Global().Set("gdspx_on_sprite_fixed_updated", js.FuncOf(gdspxOnSpriteFixedUpdated))
	js.Global().Set("gdspx_on_sprite_destroyed", js.FuncOf(gdspxOnSpriteDestroyed))
	js.Global().Set("gdspx_on_sprite_frames_set_changed", js.FuncOf(gdspxOnSpriteFramesSetChanged))
	js.Global().Set("gdspx_on_sprite_animation_changed", js.FuncOf(gdspxOnSpriteAnimationChanged))
	js.Global().Set("gdspx_on_sprite_frame_changed", js.FuncOf(gdspxOnSpriteFrameChanged))
	js.Global().Set("gdspx_on_sprite_animation_looped", js.FuncOf(gdspxOnSpriteAnimationLooped))
	js.Global().Set("gdspx_on_sprite_animation_finished", js.FuncOf(gdspxOnSpriteAnimationFinished))
	js.Global().Set("gdspx_on_sprite_vfx_finished", js.FuncOf(gdspxOnSpriteVfxFinished))
	js.Global().Set("gdspx_on_sprite_screen_exited", js.FuncOf(gdspxOnSpriteScreenExited))
	js.Global().Set("gdspx_on_sprite_screen_entered", js.FuncOf(gdspxOnSpriteScreenEntered))
	js.Global().Set("gdspx_on_mouse_pressed", js.FuncOf(gdspxOnMousePressed))
	js.Global().Set("gdspx_on_mouse_released", js.FuncOf(gdspxOnMouseReleased))
	js.Global().Set("gdspx_on_key_pressed", js.FuncOf(gdspxOnKeyPressed))
	js.Global().Set("gdspx_on_key_released", js.FuncOf(gdspxOnKeyReleased))
	js.Global().Set("gdspx_on_action_pressed", js.FuncOf(gdspxOnActionPressed))
	js.Global().Set("gdspx_on_action_just_pressed", js.FuncOf(gdspxOnActionJustPressed))
	js.Global().Set("gdspx_on_action_just_released", js.FuncOf(gdspxOnActionJustReleased))
	js.Global().Set("gdspx_on_axis_changed", js.FuncOf(gdspxOnAxisChanged))
	js.Global().Set("gdspx_on_collision_enter", js.FuncOf(gdspxOnCollisionEnter))
	js.Global().Set("gdspx_on_collision_stay", js.FuncOf(gdspxOnCollisionStay))
	js.Global().Set("gdspx_on_collision_exit", js.FuncOf(gdspxOnCollisionExit))
	js.Global().Set("gdspx_on_trigger_enter", js.FuncOf(gdspxOnTriggerEnter))
	js.Global().Set("gdspx_on_trigger_stay", js.FuncOf(gdspxOnTriggerStay))
	js.Global().Set("gdspx_on_trigger_exit", js.FuncOf(gdspxOnTriggerExit))
	js.Global().Set("gdspx_on_ui_ready", js.FuncOf(gdspxOnUiReady))
	js.Global().Set("gdspx_on_ui_updated", js.FuncOf(gdspxOnUiUpdated))
	js.Global().Set("gdspx_on_ui_destroyed", js.FuncOf(gdspxOnUiDestroyed))
	js.Global().Set("gdspx_on_ui_pressed", js.FuncOf(gdspxOnUiPressed))
	js.Global().Set("gdspx_on_ui_released", js.FuncOf(gdspxOnUiReleased))
	js.Global().Set("gdspx_on_ui_hovered", js.FuncOf(gdspxOnUiHovered))
	js.Global().Set("gdspx_on_ui_clicked", js.FuncOf(gdspxOnUiClicked))
	js.Global().Set("gdspx_on_ui_toggle", js.FuncOf(gdspxOnUiToggle))
	js.Global().Set("gdspx_on_ui_text_changed", js.FuncOf(gdspxOnUiTextChanged))
}
