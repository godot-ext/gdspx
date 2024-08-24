/**************************************************************************/
/*  gdextension_spx_ext.h                                                 */
/**************************************************************************/
/*                         This file is part of:                          */
/*                             GODOT ENGINE                               */
/*                        https://godotengine.org                         */
/**************************************************************************/
/* Copyright (c) 2014-present Godot Engine contributors (see AUTHORS.md). */
/* Copyright (c) 2007-2014 Juan Linietsky, Ariel Manzur.                  */
/*                                                                        */
/* Permission is hereby granted, free of charge, to any person obtaining  */
/* a copy of this software and associated documentation files (the        */
/* "Software"), to deal in the Software without restriction, including    */
/* without limitation the rights to use, copy, modify, merge, publish,    */
/* distribute, sublicense, and/or sell copies of the Software, and to     */
/* permit persons to whom the Software is furnished to do so, subject to  */
/* the following conditions:                                              */
/*                                                                        */
/* The above copyright notice and this permission notice shall be         */
/* included in all copies or substantial portions of the Software.        */
/*                                                                        */
/* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,        */
/* EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF     */
/* MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. */
/* IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY   */
/* CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,   */
/* TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE      */
/* SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.                 */
/**************************************************************************/

#ifndef GDEXTENSION_SPX_EXT_H
#define GDEXTENSION_SPX_EXT_H

#include "gdextension_interface.h"
#ifndef NOT_GODOT_ENGINE
#include "core/variant/variant.h"
extern void gdextension_spx_setup_interface();
#endif

#ifdef __cplusplus
extern "C" {
#endif

typedef GDExtensionConstStringPtr GdString;
typedef GDExtensionInt GdInt;
typedef GDExtensionInt GdObj;
typedef GDExtensionBool GdBool;
typedef real_t GdFloat;
typedef Vector4 GdVec4;
typedef Vector3 GdVec3;
typedef Vector2 GdVec2;
typedef Color GdColor;
typedef Rect2 GdRect2;


typedef struct {
	// 0 is return value
	// 1-7 are arguments
	GdVec4 Ret;
	GdVec4 Arg0;
	GdVec4 Arg1;
	GdVec4 Arg2;
	GdVec4 Arg3;
	GdVec4 Arg4;
	GdVec4 Arg5;
	GdVec4 Arg6;
	GdVec4 Arg7;
} CallFrameArgs;

typedef void *GDExtensionSpxCallbackInfoPtr;
typedef void (*GDExtensionSpxGlobalRegisterCallbacks)(GDExtensionSpxCallbackInfoPtr callback_ptr);


// callback
typedef void (*GDExtensionSpxCallbackOnEngineStart)();
typedef void (*GDExtensionSpxCallbackOnEngineUpdate)(GdFloat delta);
typedef void (*GDExtensionSpxCallbackOnEngineDestroy)();

typedef void (*GDExtensionSpxCallbackOnSpriteReady)(GdObj obj);
typedef void (*GDExtensionSpxCallbackOnSpriteUpdated)(GdObj obj);
typedef void (*GDExtensionSpxCallbackOnSpriteDestroyed)(GdObj obj);

typedef void (*GDExtensionSpxCallbackOnMousePressed)(GdInt keyid);
typedef void (*GDExtensionSpxCallbackOnMouseReleased)(GdInt keyid);
typedef void (*GDExtensionSpxCallbackOnKeyPressed)(GdInt keyid);
typedef void (*GDExtensionSpxCallbackOnKeyReleased)(GdInt keyid);
typedef void (*GDExtensionSpxCallbackOnActionPressed)(GdString action_name);
typedef void (*GDExtensionSpxCallbackOnActionJustPressed)(GdString action_name);
typedef void (*GDExtensionSpxCallbackOnActionJustReleased)(GdString action_name);
typedef void (*GDExtensionSpxCallbackOnAxisChanged)(GdString action_name, GdFloat value);

typedef void (*GDExtensionSpxCallbackOnCollisionEnter)(GdInt self_id, GdInt other_id);
typedef void (*GDExtensionSpxCallbackOnCollisionStay)(GdInt self_id, GdInt other_id);
typedef void (*GDExtensionSpxCallbackOnCollisionExit)(GdInt self_id, GdInt other_id);
typedef void (*GDExtensionSpxCallbackOnTriggerEnter)(GdInt self_id, GdInt other_id);
typedef void (*GDExtensionSpxCallbackOnTriggerStay)(GdInt self_id, GdInt other_id);
typedef void (*GDExtensionSpxCallbackOnTriggerExit)(GdInt self_id, GdInt other_id);

typedef void (*GDExtensionSpxCallbackOnUIReady)(GdObj obj);
typedef void (*GDExtensionSpxCallbackOnUIUpdated)(GdObj obj);
typedef void (*GDExtensionSpxCallbackOnUIDestroyed)(GdObj obj);

typedef void (*GDExtensionSpxCallbackOnUIPressed)(GdObj obj);
typedef void (*GDExtensionSpxCallbackOnUIReleased)(GdObj obj);
typedef void (*GDExtensionSpxCallbackOnUIHovered)(GdObj obj);
typedef void (*GDExtensionSpxCallbackOnUIClicked)(GdObj obj);
typedef void (*GDExtensionSpxCallbackOnUIToggle)(GdObj obj, GdBool is_on);
typedef void (*GDExtensionSpxCallbackOnUITextChanged)(GdObj obj, GdString text);

// string
typedef void (*GDExtensionSpxStringNewWithLatin1Chars)(GDExtensionUninitializedStringPtr r_dest, const char *p_contents);
typedef void (*GDExtensionSpxStringNewWithUtf8Chars)(GDExtensionUninitializedStringPtr r_dest, const char *p_contents);
typedef void (*GDExtensionSpxStringNewWithLatin1CharsAndLen)(GDExtensionUninitializedStringPtr r_dest, const char *p_contents, GdInt p_size);
typedef void (*GDExtensionSpxStringNewWithUtf8CharsAndLen)(GDExtensionUninitializedStringPtr r_dest, const char *p_contents, GdInt p_size);
typedef GdInt (*GDExtensionSpxStringToLatin1Chars)(GDExtensionConstStringPtr p_self, char *r_text, GdInt p_max_write_length);
typedef GdInt (*GDExtensionSpxStringToUtf8Chars)(GDExtensionConstStringPtr p_self, char *r_text, GdInt p_max_write_length);
// variant
typedef GDExtensionPtrConstructor (*GDExtensionSpxVariantGetPtrConstructor)(GDExtensionVariantType p_type, int32_t p_constructor);
typedef GDExtensionPtrDestructor (*GDExtensionSpxVariantGetPtrDestructor)(GDExtensionVariantType p_type);

typedef struct {
	// engine
	GDExtensionSpxCallbackOnEngineStart func_on_engine_start;
	GDExtensionSpxCallbackOnEngineUpdate func_on_engine_update;
	GDExtensionSpxCallbackOnEngineDestroy func_on_engine_destroy;

	// sprite
	GDExtensionSpxCallbackOnSpriteReady func_on_sprite_ready;
	GDExtensionSpxCallbackOnSpriteUpdated func_on_sprite_updated;
	GDExtensionSpxCallbackOnSpriteDestroyed func_on_sprite_destroyed;

	// input
	GDExtensionSpxCallbackOnMousePressed func_on_mouse_pressed;
	GDExtensionSpxCallbackOnMouseReleased func_on_mouse_released;
	GDExtensionSpxCallbackOnKeyPressed func_on_key_pressed;
	GDExtensionSpxCallbackOnKeyReleased func_on_key_released;
	GDExtensionSpxCallbackOnActionPressed func_on_action_pressed;
	GDExtensionSpxCallbackOnActionJustPressed func_on_action_just_pressed;
	GDExtensionSpxCallbackOnActionJustReleased func_on_action_just_released;
	GDExtensionSpxCallbackOnAxisChanged func_on_axis_changed;

	// physic
	GDExtensionSpxCallbackOnCollisionEnter func_on_collision_enter;
	GDExtensionSpxCallbackOnCollisionStay func_on_collision_stay;
	GDExtensionSpxCallbackOnCollisionExit func_on_collision_exit;
	GDExtensionSpxCallbackOnTriggerEnter func_on_trigger_enter;
	GDExtensionSpxCallbackOnTriggerStay func_on_trigger_stay;
	GDExtensionSpxCallbackOnTriggerExit func_on_trigger_exit;

	// ui
	GDExtensionSpxCallbackOnUIReady func_on_ui_ready;
	GDExtensionSpxCallbackOnUIUpdated func_on_ui_updated;
	GDExtensionSpxCallbackOnUIDestroyed func_on_ui_destroyed;

	GDExtensionSpxCallbackOnUIPressed func_on_ui_pressed;
	GDExtensionSpxCallbackOnUIReleased func_on_ui_released;
	GDExtensionSpxCallbackOnUIHovered func_on_ui_hovered;
	GDExtensionSpxCallbackOnUIClicked func_on_ui_clicked;
	GDExtensionSpxCallbackOnUIToggle func_on_ui_toggle;
	GDExtensionSpxCallbackOnUITextChanged func_on_ui_text_changed;

} SpxCallbackInfo;



// SpxAudio
typedef void (*GDExtensionSpxAudioPlayAudio)(GdString path);
typedef void (*GDExtensionSpxAudioSetAudioVolume)(GdFloat volume);
typedef void (*GDExtensionSpxAudioGetAudioVolume)(GdFloat* ret_value);
typedef void (*GDExtensionSpxAudioIsMusicPlaying)(GdBool* ret_value);
typedef void (*GDExtensionSpxAudioPlayMusic)(GdString path);
typedef void (*GDExtensionSpxAudioSetMusicVolume)(GdFloat volume);
typedef void (*GDExtensionSpxAudioGetMusicVolume)(GdFloat* ret_value);
typedef void (*GDExtensionSpxAudioPauseMusic)();
typedef void (*GDExtensionSpxAudioResumeMusic)();
typedef void (*GDExtensionSpxAudioGetMusicTimer)(GdFloat* ret_value);
typedef void (*GDExtensionSpxAudioSetMusicTimer)(GdFloat time);
// SpxInput
typedef void (*GDExtensionSpxInputGetMousePos)(GdVec2* ret_value);
typedef void (*GDExtensionSpxInputGetKey)(GdInt key, GdBool* ret_value);
typedef void (*GDExtensionSpxInputGetMouseState)(GdInt mouse_id, GdBool* ret_value);
typedef void (*GDExtensionSpxInputGetKeyState)(GdInt key, GdInt* ret_value);
typedef void (*GDExtensionSpxInputGetAxis)(GdString neg_action,GdString pos_action, GdFloat* ret_value);
typedef void (*GDExtensionSpxInputIsActionPressed)(GdString action, GdBool* ret_value);
typedef void (*GDExtensionSpxInputIsActionJustPressed)(GdString action, GdBool* ret_value);
typedef void (*GDExtensionSpxInputIsActionJustReleased)(GdString action, GdBool* ret_value);
// SpxPhysic
typedef void (*GDExtensionSpxPhysicRaycast)(GdVec2 from, GdVec2 to, GdInt collision_mask, GdObj* ret_value);
// SpxSprite
typedef void (*GDExtensionSpxSpriteCreateSprite)(GdString path, GdObj* ret_value);
typedef void (*GDExtensionSpxSpriteCloneSprite)(GdObj obj, GdObj* ret_value);
typedef void (*GDExtensionSpxSpriteDestroySprite)(GdObj obj, GdBool* ret_value);
typedef void (*GDExtensionSpxSpriteIsSpriteAlive)(GdObj obj, GdBool* ret_value);
typedef void (*GDExtensionSpxSpriteSetPosition)(GdObj obj, GdVec2 pos);
typedef void (*GDExtensionSpxSpriteSetRotation)(GdObj obj, GdFloat rot);
typedef void (*GDExtensionSpxSpriteSetScale)(GdObj obj, GdVec2 scale);
typedef void (*GDExtensionSpxSpriteGetPosition)(GdObj obj, GdVec2* ret_value);
typedef void (*GDExtensionSpxSpriteGetRotation)(GdObj obj, GdFloat* ret_value);
typedef void (*GDExtensionSpxSpriteGetScale)(GdObj obj, GdVec2* ret_value);
typedef void (*GDExtensionSpxSpriteSetColor)(GdObj obj, GdColor color);
typedef void (*GDExtensionSpxSpriteGetColor)(GdObj obj, GdColor* ret_value);
typedef void (*GDExtensionSpxSpriteSetTexture)(GdObj obj, GdString path);
typedef void (*GDExtensionSpxSpriteGetTexture)(GdObj obj, GdString* ret_value);
typedef void (*GDExtensionSpxSpriteSetVisible)(GdObj obj, GdBool visible);
typedef void (*GDExtensionSpxSpriteGetVisible)(GdObj obj, GdBool* ret_value);
typedef void (*GDExtensionSpxSpriteGetZIndex)(GdObj obj, GdInt* ret_value);
typedef void (*GDExtensionSpxSpriteSetZIndex)(GdObj obj, GdInt z);
typedef void (*GDExtensionSpxSpritePlayAnim)(GdObj obj, GdString p_name , GdFloat p_custom_scale , GdBool p_from_end );
typedef void (*GDExtensionSpxSpritePlayBackwardsAnim)(GdObj obj,  GdString p_name );
typedef void (*GDExtensionSpxSpritePauseAnim)(GdObj obj);
typedef void (*GDExtensionSpxSpriteStopAnim)(GdObj obj);
typedef void (*GDExtensionSpxSpriteIsPlayingAnim)(GdObj obj, GdBool* ret_value);
typedef void (*GDExtensionSpxSpriteSetAnim)(GdObj obj, GdString p_name);
typedef void (*GDExtensionSpxSpriteGetAnim)(GdObj obj, GdString* ret_value);
typedef void (*GDExtensionSpxSpriteSetAnimFrame)(GdObj obj, GdInt p_frame);
typedef void (*GDExtensionSpxSpriteGetAnimFrame)(GdObj obj, GdInt* ret_value);
typedef void (*GDExtensionSpxSpriteSetAnimSpeedScale)(GdObj obj, GdFloat p_speed_scale);
typedef void (*GDExtensionSpxSpriteGetAnimSpeedScale)(GdObj obj, GdFloat* ret_value);
typedef void (*GDExtensionSpxSpriteGetAnimPlayingSpeed)(GdObj obj, GdFloat* ret_value);
typedef void (*GDExtensionSpxSpriteSetAnimCentered)(GdObj obj, GdBool p_center);
typedef void (*GDExtensionSpxSpriteIsAnimCentered)(GdObj obj, GdBool* ret_value);
typedef void (*GDExtensionSpxSpriteSetAnimOffset)(GdObj obj, GdVec2 p_offset);
typedef void (*GDExtensionSpxSpriteGetAnimOffset)(GdObj obj, GdVec2* ret_value);
typedef void (*GDExtensionSpxSpriteSetAnimFlipH)(GdObj obj, GdBool p_flip);
typedef void (*GDExtensionSpxSpriteIsAnimFlippedH)(GdObj obj, GdBool* ret_value);
typedef void (*GDExtensionSpxSpriteSetAnimFlipV)(GdObj obj, GdBool p_flip);
typedef void (*GDExtensionSpxSpriteIsAnimFlippedV)(GdObj obj, GdBool* ret_value);
typedef void (*GDExtensionSpxSpriteSetVelocity)(GdObj obj, GdVec2 velocity);
typedef void (*GDExtensionSpxSpriteGetVelocity)(GdObj obj, GdVec2* ret_value);
typedef void (*GDExtensionSpxSpriteIsOnFloor)(GdObj obj, GdBool* ret_value);
typedef void (*GDExtensionSpxSpriteIsOnFloorOnly)(GdObj obj, GdBool* ret_value);
typedef void (*GDExtensionSpxSpriteIsOnWall)(GdObj obj, GdBool* ret_value);
typedef void (*GDExtensionSpxSpriteIsOnWallOnly)(GdObj obj, GdBool* ret_value);
typedef void (*GDExtensionSpxSpriteIsOnCeiling)(GdObj obj, GdBool* ret_value);
typedef void (*GDExtensionSpxSpriteIsOnCeilingOnly)(GdObj obj, GdBool* ret_value);
typedef void (*GDExtensionSpxSpriteGetLastMotion)(GdObj obj, GdVec2* ret_value);
typedef void (*GDExtensionSpxSpriteGetPositionDelta)(GdObj obj, GdVec2* ret_value);
typedef void (*GDExtensionSpxSpriteGetFloorNormal)(GdObj obj, GdVec2* ret_value);
typedef void (*GDExtensionSpxSpriteGetWallNormal)(GdObj obj, GdVec2* ret_value);
typedef void (*GDExtensionSpxSpriteGetRealVelocity)(GdObj obj, GdVec2* ret_value);
typedef void (*GDExtensionSpxSpriteMoveAndSlide)(GdObj obj);
typedef void (*GDExtensionSpxSpriteSetGravity)(GdObj obj, GdFloat gravity);
typedef void (*GDExtensionSpxSpriteGetGravity)(GdObj obj, GdFloat* ret_value);
typedef void (*GDExtensionSpxSpriteSetMass)(GdObj obj, GdFloat mass);
typedef void (*GDExtensionSpxSpriteGetMass)(GdObj obj, GdFloat* ret_value);
typedef void (*GDExtensionSpxSpriteAddForce)(GdObj obj, GdVec2 force);
typedef void (*GDExtensionSpxSpriteAddImpulse)(GdObj obj, GdVec2 impulse);
typedef void (*GDExtensionSpxSpriteSetCollisionLayer)(GdObj obj, GdInt layer);
typedef void (*GDExtensionSpxSpriteGetCollisionLayer)(GdObj obj, GdInt* ret_value);
typedef void (*GDExtensionSpxSpriteSetCollisionMask)(GdObj obj, GdInt mask);
typedef void (*GDExtensionSpxSpriteGetCollisionMask)(GdObj obj, GdInt* ret_value);
typedef void (*GDExtensionSpxSpriteSetTriggerLayer)(GdObj obj, GdInt layer);
typedef void (*GDExtensionSpxSpriteGetTriggerLayer)(GdObj obj, GdInt* ret_value);
typedef void (*GDExtensionSpxSpriteSetTriggerMask)(GdObj obj, GdInt mask);
typedef void (*GDExtensionSpxSpriteGetTriggerMask)(GdObj obj, GdInt* ret_value);
typedef void (*GDExtensionSpxSpriteSetColliderRect)(GdObj obj, GdVec2 center, GdVec2 size);
typedef void (*GDExtensionSpxSpriteSetColliderCircle)(GdObj obj, GdVec2 center, GdFloat radius);
typedef void (*GDExtensionSpxSpriteSetColliderCapsule)(GdObj obj, GdVec2 center, GdVec2 size);
typedef void (*GDExtensionSpxSpriteSetCollisionEnabled)(GdObj obj, GdBool enabled);
typedef void (*GDExtensionSpxSpriteIsCollisionEnabled)(GdObj obj, GdBool* ret_value);
typedef void (*GDExtensionSpxSpriteSetTriggerRect)(GdObj obj, GdVec2 center, GdVec2 size);
typedef void (*GDExtensionSpxSpriteSetTriggerCircle)(GdObj obj, GdVec2 center, GdFloat radius);
typedef void (*GDExtensionSpxSpriteSetTriggerCapsule)(GdObj obj, GdVec2 center, GdVec2 size);
typedef void (*GDExtensionSpxSpriteSetTriggerEnabled)(GdObj obj, GdBool trigger);
typedef void (*GDExtensionSpxSpriteIsTriggerEnabled)(GdObj obj, GdBool* ret_value);
// SpxUI
typedef void (*GDExtensionSpxUICreateNode)(GdString path, GdObj* ret_value);
typedef void (*GDExtensionSpxUICreateButton)(GdString path,GdString text, GdObj* ret_value);
typedef void (*GDExtensionSpxUICreateLabel)(GdString path, GdString text, GdObj* ret_value);
typedef void (*GDExtensionSpxUICreateImage)(GdString path, GdObj* ret_value);
typedef void (*GDExtensionSpxUICreateToggle)(GdString path, GdBool value, GdObj* ret_value);
typedef void (*GDExtensionSpxUICreateSlider)(GdString path, GdFloat value, GdObj* ret_value);
typedef void (*GDExtensionSpxUICreateInput)(GdString path, GdString text, GdObj* ret_value);
typedef void (*GDExtensionSpxUIDestroyNode)(GdObj obj, GdBool* ret_value);
typedef void (*GDExtensionSpxUIGetType)(GdObj obj, GdInt* ret_value);
typedef void (*GDExtensionSpxUISetText)(GdObj obj, GdString text);
typedef void (*GDExtensionSpxUIGetText)(GdObj obj, GdString* ret_value);
typedef void (*GDExtensionSpxUISetTexture)(GdObj obj, GdString path);
typedef void (*GDExtensionSpxUIGetTexture)(GdObj obj, GdString* ret_value);
typedef void (*GDExtensionSpxUISetColor)(GdObj obj, GdColor color);
typedef void (*GDExtensionSpxUIGetColor)(GdObj obj, GdColor* ret_value);
typedef void (*GDExtensionSpxUISetFontSize)(GdObj obj, GdInt size);
typedef void (*GDExtensionSpxUIGetFontSize)(GdObj obj, GdInt* ret_value);
typedef void (*GDExtensionSpxUISetVisible)(GdObj obj, GdBool visible);
typedef void (*GDExtensionSpxUIGetVisible)(GdObj obj, GdBool* ret_value);
typedef void (*GDExtensionSpxUISetInteractable)(GdObj obj, GdBool interactable);
typedef void (*GDExtensionSpxUIGetInteractable)(GdObj obj, GdBool* ret_value);
typedef void (*GDExtensionSpxUISetRect)(GdObj obj, GdRect2 rect);
typedef void (*GDExtensionSpxUIGetRect)(GdObj obj, GdRect2* ret_value);


#ifdef __cplusplus
}
#endif

#endif // GDEXTENSION_SPX_EXT_H
