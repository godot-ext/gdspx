/*------------------------------------------------------------------------------
//   This code was generated by template ffi_gdextension_interface.go.tmpl.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "ffi_gdextension_interface.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/
package ffi

//revive:disable

var (
	api GDExtensionInterface
)

type GDExtensionInterface struct {
	// All of the GDExtension interface functions.
	SpxAudioPlayAudio GDExtensionSpxAudioPlayAudio
	SpxAudioSetAudioVolume GDExtensionSpxAudioSetAudioVolume
	SpxAudioGetAudioVolume GDExtensionSpxAudioGetAudioVolume
	SpxAudioIsMusicPlaying GDExtensionSpxAudioIsMusicPlaying
	SpxAudioPlayMusic GDExtensionSpxAudioPlayMusic
	SpxAudioSetMusicVolume GDExtensionSpxAudioSetMusicVolume
	SpxAudioGetMusicVolume GDExtensionSpxAudioGetMusicVolume
	SpxAudioPauseMusic GDExtensionSpxAudioPauseMusic
	SpxAudioResumeMusic GDExtensionSpxAudioResumeMusic
	SpxAudioGetMusicTimer GDExtensionSpxAudioGetMusicTimer
	SpxAudioSetMusicTimer GDExtensionSpxAudioSetMusicTimer
	SpxCameraGetCameraPosition GDExtensionSpxCameraGetCameraPosition
	SpxCameraSetCameraPosition GDExtensionSpxCameraSetCameraPosition
	SpxCameraGetCameraZoom GDExtensionSpxCameraGetCameraZoom
	SpxCameraSetCameraZoom GDExtensionSpxCameraSetCameraZoom
	SpxCameraGetViewportRect GDExtensionSpxCameraGetViewportRect
	SpxInputGetMousePos GDExtensionSpxInputGetMousePos
	SpxInputGetKey GDExtensionSpxInputGetKey
	SpxInputGetMouseState GDExtensionSpxInputGetMouseState
	SpxInputGetKeyState GDExtensionSpxInputGetKeyState
	SpxInputGetAxis GDExtensionSpxInputGetAxis
	SpxInputIsActionPressed GDExtensionSpxInputIsActionPressed
	SpxInputIsActionJustPressed GDExtensionSpxInputIsActionJustPressed
	SpxInputIsActionJustReleased GDExtensionSpxInputIsActionJustReleased
	SpxPhysicRaycast GDExtensionSpxPhysicRaycast
	SpxSceneChangeSceneToFile GDExtensionSpxSceneChangeSceneToFile
	SpxSceneReloadCurrentScene GDExtensionSpxSceneReloadCurrentScene
	SpxSceneUnloadCurrentScene GDExtensionSpxSceneUnloadCurrentScene
	SpxSpriteSetDontDestroyOnLoad GDExtensionSpxSpriteSetDontDestroyOnLoad
	SpxSpriteSetProcess GDExtensionSpxSpriteSetProcess
	SpxSpriteSetPhysicProcess GDExtensionSpxSpriteSetPhysicProcess
	SpxSpriteSetChildPosition GDExtensionSpxSpriteSetChildPosition
	SpxSpriteGetChildPosition GDExtensionSpxSpriteGetChildPosition
	SpxSpriteSetChildRotation GDExtensionSpxSpriteSetChildRotation
	SpxSpriteGetChildRotation GDExtensionSpxSpriteGetChildRotation
	SpxSpriteSetChildScale GDExtensionSpxSpriteSetChildScale
	SpxSpriteGetChildScale GDExtensionSpxSpriteGetChildScale
	SpxSpriteCreateSprite GDExtensionSpxSpriteCreateSprite
	SpxSpriteCloneSprite GDExtensionSpxSpriteCloneSprite
	SpxSpriteDestroySprite GDExtensionSpxSpriteDestroySprite
	SpxSpriteIsSpriteAlive GDExtensionSpxSpriteIsSpriteAlive
	SpxSpriteSetPosition GDExtensionSpxSpriteSetPosition
	SpxSpriteGetPosition GDExtensionSpxSpriteGetPosition
	SpxSpriteSetRotation GDExtensionSpxSpriteSetRotation
	SpxSpriteGetRotation GDExtensionSpxSpriteGetRotation
	SpxSpriteSetScale GDExtensionSpxSpriteSetScale
	SpxSpriteGetScale GDExtensionSpxSpriteGetScale
	SpxSpriteSetColor GDExtensionSpxSpriteSetColor
	SpxSpriteGetColor GDExtensionSpxSpriteGetColor
	SpxSpriteSetTexture GDExtensionSpxSpriteSetTexture
	SpxSpriteGetTexture GDExtensionSpxSpriteGetTexture
	SpxSpriteSetVisible GDExtensionSpxSpriteSetVisible
	SpxSpriteGetVisible GDExtensionSpxSpriteGetVisible
	SpxSpriteGetZIndex GDExtensionSpxSpriteGetZIndex
	SpxSpriteSetZIndex GDExtensionSpxSpriteSetZIndex
	SpxSpritePlayAnim GDExtensionSpxSpritePlayAnim
	SpxSpritePlayBackwardsAnim GDExtensionSpxSpritePlayBackwardsAnim
	SpxSpritePauseAnim GDExtensionSpxSpritePauseAnim
	SpxSpriteStopAnim GDExtensionSpxSpriteStopAnim
	SpxSpriteIsPlayingAnim GDExtensionSpxSpriteIsPlayingAnim
	SpxSpriteSetAnim GDExtensionSpxSpriteSetAnim
	SpxSpriteGetAnim GDExtensionSpxSpriteGetAnim
	SpxSpriteSetAnimFrame GDExtensionSpxSpriteSetAnimFrame
	SpxSpriteGetAnimFrame GDExtensionSpxSpriteGetAnimFrame
	SpxSpriteSetAnimSpeedScale GDExtensionSpxSpriteSetAnimSpeedScale
	SpxSpriteGetAnimSpeedScale GDExtensionSpxSpriteGetAnimSpeedScale
	SpxSpriteGetAnimPlayingSpeed GDExtensionSpxSpriteGetAnimPlayingSpeed
	SpxSpriteSetAnimCentered GDExtensionSpxSpriteSetAnimCentered
	SpxSpriteIsAnimCentered GDExtensionSpxSpriteIsAnimCentered
	SpxSpriteSetAnimOffset GDExtensionSpxSpriteSetAnimOffset
	SpxSpriteGetAnimOffset GDExtensionSpxSpriteGetAnimOffset
	SpxSpriteSetAnimFlipH GDExtensionSpxSpriteSetAnimFlipH
	SpxSpriteIsAnimFlippedH GDExtensionSpxSpriteIsAnimFlippedH
	SpxSpriteSetAnimFlipV GDExtensionSpxSpriteSetAnimFlipV
	SpxSpriteIsAnimFlippedV GDExtensionSpxSpriteIsAnimFlippedV
	SpxSpriteSetVelocity GDExtensionSpxSpriteSetVelocity
	SpxSpriteGetVelocity GDExtensionSpxSpriteGetVelocity
	SpxSpriteIsOnFloor GDExtensionSpxSpriteIsOnFloor
	SpxSpriteIsOnFloorOnly GDExtensionSpxSpriteIsOnFloorOnly
	SpxSpriteIsOnWall GDExtensionSpxSpriteIsOnWall
	SpxSpriteIsOnWallOnly GDExtensionSpxSpriteIsOnWallOnly
	SpxSpriteIsOnCeiling GDExtensionSpxSpriteIsOnCeiling
	SpxSpriteIsOnCeilingOnly GDExtensionSpxSpriteIsOnCeilingOnly
	SpxSpriteGetLastMotion GDExtensionSpxSpriteGetLastMotion
	SpxSpriteGetPositionDelta GDExtensionSpxSpriteGetPositionDelta
	SpxSpriteGetFloorNormal GDExtensionSpxSpriteGetFloorNormal
	SpxSpriteGetWallNormal GDExtensionSpxSpriteGetWallNormal
	SpxSpriteGetRealVelocity GDExtensionSpxSpriteGetRealVelocity
	SpxSpriteMoveAndSlide GDExtensionSpxSpriteMoveAndSlide
	SpxSpriteSetGravity GDExtensionSpxSpriteSetGravity
	SpxSpriteGetGravity GDExtensionSpxSpriteGetGravity
	SpxSpriteSetMass GDExtensionSpxSpriteSetMass
	SpxSpriteGetMass GDExtensionSpxSpriteGetMass
	SpxSpriteAddForce GDExtensionSpxSpriteAddForce
	SpxSpriteAddImpulse GDExtensionSpxSpriteAddImpulse
	SpxSpriteSetCollisionLayer GDExtensionSpxSpriteSetCollisionLayer
	SpxSpriteGetCollisionLayer GDExtensionSpxSpriteGetCollisionLayer
	SpxSpriteSetCollisionMask GDExtensionSpxSpriteSetCollisionMask
	SpxSpriteGetCollisionMask GDExtensionSpxSpriteGetCollisionMask
	SpxSpriteSetTriggerLayer GDExtensionSpxSpriteSetTriggerLayer
	SpxSpriteGetTriggerLayer GDExtensionSpxSpriteGetTriggerLayer
	SpxSpriteSetTriggerMask GDExtensionSpxSpriteSetTriggerMask
	SpxSpriteGetTriggerMask GDExtensionSpxSpriteGetTriggerMask
	SpxSpriteSetColliderRect GDExtensionSpxSpriteSetColliderRect
	SpxSpriteSetColliderCircle GDExtensionSpxSpriteSetColliderCircle
	SpxSpriteSetColliderCapsule GDExtensionSpxSpriteSetColliderCapsule
	SpxSpriteSetCollisionEnabled GDExtensionSpxSpriteSetCollisionEnabled
	SpxSpriteIsCollisionEnabled GDExtensionSpxSpriteIsCollisionEnabled
	SpxSpriteSetTriggerRect GDExtensionSpxSpriteSetTriggerRect
	SpxSpriteSetTriggerCircle GDExtensionSpxSpriteSetTriggerCircle
	SpxSpriteSetTriggerCapsule GDExtensionSpxSpriteSetTriggerCapsule
	SpxSpriteSetTriggerEnabled GDExtensionSpxSpriteSetTriggerEnabled
	SpxSpriteIsTriggerEnabled GDExtensionSpxSpriteIsTriggerEnabled
	SpxUiCreateNode GDExtensionSpxUiCreateNode
	SpxUiCreateButton GDExtensionSpxUiCreateButton
	SpxUiCreateLabel GDExtensionSpxUiCreateLabel
	SpxUiCreateImage GDExtensionSpxUiCreateImage
	SpxUiCreateToggle GDExtensionSpxUiCreateToggle
	SpxUiCreateSlider GDExtensionSpxUiCreateSlider
	SpxUiCreateInput GDExtensionSpxUiCreateInput
	SpxUiDestroyNode GDExtensionSpxUiDestroyNode
	SpxUiGetType GDExtensionSpxUiGetType
	SpxUiSetText GDExtensionSpxUiSetText
	SpxUiGetText GDExtensionSpxUiGetText
	SpxUiSetTexture GDExtensionSpxUiSetTexture
	SpxUiGetTexture GDExtensionSpxUiGetTexture
	SpxUiSetColor GDExtensionSpxUiSetColor
	SpxUiGetColor GDExtensionSpxUiGetColor
	SpxUiSetFontSize GDExtensionSpxUiSetFontSize
	SpxUiGetFontSize GDExtensionSpxUiGetFontSize
	SpxUiSetVisible GDExtensionSpxUiSetVisible
	SpxUiGetVisible GDExtensionSpxUiGetVisible
	SpxUiSetInteractable GDExtensionSpxUiSetInteractable
	SpxUiGetInteractable GDExtensionSpxUiGetInteractable
	SpxUiSetRect GDExtensionSpxUiSetRect
	SpxUiGetRect GDExtensionSpxUiGetRect
	}

func (x *GDExtensionInterface) loadProcAddresses() {
	x.SpxAudioPlayAudio = (GDExtensionSpxAudioPlayAudio)(dlsymGD("spx_audio_play_audio"))
	x.SpxAudioSetAudioVolume = (GDExtensionSpxAudioSetAudioVolume)(dlsymGD("spx_audio_set_audio_volume"))
	x.SpxAudioGetAudioVolume = (GDExtensionSpxAudioGetAudioVolume)(dlsymGD("spx_audio_get_audio_volume"))
	x.SpxAudioIsMusicPlaying = (GDExtensionSpxAudioIsMusicPlaying)(dlsymGD("spx_audio_is_music_playing"))
	x.SpxAudioPlayMusic = (GDExtensionSpxAudioPlayMusic)(dlsymGD("spx_audio_play_music"))
	x.SpxAudioSetMusicVolume = (GDExtensionSpxAudioSetMusicVolume)(dlsymGD("spx_audio_set_music_volume"))
	x.SpxAudioGetMusicVolume = (GDExtensionSpxAudioGetMusicVolume)(dlsymGD("spx_audio_get_music_volume"))
	x.SpxAudioPauseMusic = (GDExtensionSpxAudioPauseMusic)(dlsymGD("spx_audio_pause_music"))
	x.SpxAudioResumeMusic = (GDExtensionSpxAudioResumeMusic)(dlsymGD("spx_audio_resume_music"))
	x.SpxAudioGetMusicTimer = (GDExtensionSpxAudioGetMusicTimer)(dlsymGD("spx_audio_get_music_timer"))
	x.SpxAudioSetMusicTimer = (GDExtensionSpxAudioSetMusicTimer)(dlsymGD("spx_audio_set_music_timer"))
	x.SpxCameraGetCameraPosition = (GDExtensionSpxCameraGetCameraPosition)(dlsymGD("spx_camera_get_camera_position"))
	x.SpxCameraSetCameraPosition = (GDExtensionSpxCameraSetCameraPosition)(dlsymGD("spx_camera_set_camera_position"))
	x.SpxCameraGetCameraZoom = (GDExtensionSpxCameraGetCameraZoom)(dlsymGD("spx_camera_get_camera_zoom"))
	x.SpxCameraSetCameraZoom = (GDExtensionSpxCameraSetCameraZoom)(dlsymGD("spx_camera_set_camera_zoom"))
	x.SpxCameraGetViewportRect = (GDExtensionSpxCameraGetViewportRect)(dlsymGD("spx_camera_get_viewport_rect"))
	x.SpxInputGetMousePos = (GDExtensionSpxInputGetMousePos)(dlsymGD("spx_input_get_mouse_pos"))
	x.SpxInputGetKey = (GDExtensionSpxInputGetKey)(dlsymGD("spx_input_get_key"))
	x.SpxInputGetMouseState = (GDExtensionSpxInputGetMouseState)(dlsymGD("spx_input_get_mouse_state"))
	x.SpxInputGetKeyState = (GDExtensionSpxInputGetKeyState)(dlsymGD("spx_input_get_key_state"))
	x.SpxInputGetAxis = (GDExtensionSpxInputGetAxis)(dlsymGD("spx_input_get_axis"))
	x.SpxInputIsActionPressed = (GDExtensionSpxInputIsActionPressed)(dlsymGD("spx_input_is_action_pressed"))
	x.SpxInputIsActionJustPressed = (GDExtensionSpxInputIsActionJustPressed)(dlsymGD("spx_input_is_action_just_pressed"))
	x.SpxInputIsActionJustReleased = (GDExtensionSpxInputIsActionJustReleased)(dlsymGD("spx_input_is_action_just_released"))
	x.SpxPhysicRaycast = (GDExtensionSpxPhysicRaycast)(dlsymGD("spx_physic_raycast"))
	x.SpxSceneChangeSceneToFile = (GDExtensionSpxSceneChangeSceneToFile)(dlsymGD("spx_scene_change_scene_to_file"))
	x.SpxSceneReloadCurrentScene = (GDExtensionSpxSceneReloadCurrentScene)(dlsymGD("spx_scene_reload_current_scene"))
	x.SpxSceneUnloadCurrentScene = (GDExtensionSpxSceneUnloadCurrentScene)(dlsymGD("spx_scene_unload_current_scene"))
	x.SpxSpriteSetDontDestroyOnLoad = (GDExtensionSpxSpriteSetDontDestroyOnLoad)(dlsymGD("spx_sprite_set_dont_destroy_on_load"))
	x.SpxSpriteSetProcess = (GDExtensionSpxSpriteSetProcess)(dlsymGD("spx_sprite_set_process"))
	x.SpxSpriteSetPhysicProcess = (GDExtensionSpxSpriteSetPhysicProcess)(dlsymGD("spx_sprite_set_physic_process"))
	x.SpxSpriteSetChildPosition = (GDExtensionSpxSpriteSetChildPosition)(dlsymGD("spx_sprite_set_child_position"))
	x.SpxSpriteGetChildPosition = (GDExtensionSpxSpriteGetChildPosition)(dlsymGD("spx_sprite_get_child_position"))
	x.SpxSpriteSetChildRotation = (GDExtensionSpxSpriteSetChildRotation)(dlsymGD("spx_sprite_set_child_rotation"))
	x.SpxSpriteGetChildRotation = (GDExtensionSpxSpriteGetChildRotation)(dlsymGD("spx_sprite_get_child_rotation"))
	x.SpxSpriteSetChildScale = (GDExtensionSpxSpriteSetChildScale)(dlsymGD("spx_sprite_set_child_scale"))
	x.SpxSpriteGetChildScale = (GDExtensionSpxSpriteGetChildScale)(dlsymGD("spx_sprite_get_child_scale"))
	x.SpxSpriteCreateSprite = (GDExtensionSpxSpriteCreateSprite)(dlsymGD("spx_sprite_create_sprite"))
	x.SpxSpriteCloneSprite = (GDExtensionSpxSpriteCloneSprite)(dlsymGD("spx_sprite_clone_sprite"))
	x.SpxSpriteDestroySprite = (GDExtensionSpxSpriteDestroySprite)(dlsymGD("spx_sprite_destroy_sprite"))
	x.SpxSpriteIsSpriteAlive = (GDExtensionSpxSpriteIsSpriteAlive)(dlsymGD("spx_sprite_is_sprite_alive"))
	x.SpxSpriteSetPosition = (GDExtensionSpxSpriteSetPosition)(dlsymGD("spx_sprite_set_position"))
	x.SpxSpriteGetPosition = (GDExtensionSpxSpriteGetPosition)(dlsymGD("spx_sprite_get_position"))
	x.SpxSpriteSetRotation = (GDExtensionSpxSpriteSetRotation)(dlsymGD("spx_sprite_set_rotation"))
	x.SpxSpriteGetRotation = (GDExtensionSpxSpriteGetRotation)(dlsymGD("spx_sprite_get_rotation"))
	x.SpxSpriteSetScale = (GDExtensionSpxSpriteSetScale)(dlsymGD("spx_sprite_set_scale"))
	x.SpxSpriteGetScale = (GDExtensionSpxSpriteGetScale)(dlsymGD("spx_sprite_get_scale"))
	x.SpxSpriteSetColor = (GDExtensionSpxSpriteSetColor)(dlsymGD("spx_sprite_set_color"))
	x.SpxSpriteGetColor = (GDExtensionSpxSpriteGetColor)(dlsymGD("spx_sprite_get_color"))
	x.SpxSpriteSetTexture = (GDExtensionSpxSpriteSetTexture)(dlsymGD("spx_sprite_set_texture"))
	x.SpxSpriteGetTexture = (GDExtensionSpxSpriteGetTexture)(dlsymGD("spx_sprite_get_texture"))
	x.SpxSpriteSetVisible = (GDExtensionSpxSpriteSetVisible)(dlsymGD("spx_sprite_set_visible"))
	x.SpxSpriteGetVisible = (GDExtensionSpxSpriteGetVisible)(dlsymGD("spx_sprite_get_visible"))
	x.SpxSpriteGetZIndex = (GDExtensionSpxSpriteGetZIndex)(dlsymGD("spx_sprite_get_z_index"))
	x.SpxSpriteSetZIndex = (GDExtensionSpxSpriteSetZIndex)(dlsymGD("spx_sprite_set_z_index"))
	x.SpxSpritePlayAnim = (GDExtensionSpxSpritePlayAnim)(dlsymGD("spx_sprite_play_anim"))
	x.SpxSpritePlayBackwardsAnim = (GDExtensionSpxSpritePlayBackwardsAnim)(dlsymGD("spx_sprite_play_backwards_anim"))
	x.SpxSpritePauseAnim = (GDExtensionSpxSpritePauseAnim)(dlsymGD("spx_sprite_pause_anim"))
	x.SpxSpriteStopAnim = (GDExtensionSpxSpriteStopAnim)(dlsymGD("spx_sprite_stop_anim"))
	x.SpxSpriteIsPlayingAnim = (GDExtensionSpxSpriteIsPlayingAnim)(dlsymGD("spx_sprite_is_playing_anim"))
	x.SpxSpriteSetAnim = (GDExtensionSpxSpriteSetAnim)(dlsymGD("spx_sprite_set_anim"))
	x.SpxSpriteGetAnim = (GDExtensionSpxSpriteGetAnim)(dlsymGD("spx_sprite_get_anim"))
	x.SpxSpriteSetAnimFrame = (GDExtensionSpxSpriteSetAnimFrame)(dlsymGD("spx_sprite_set_anim_frame"))
	x.SpxSpriteGetAnimFrame = (GDExtensionSpxSpriteGetAnimFrame)(dlsymGD("spx_sprite_get_anim_frame"))
	x.SpxSpriteSetAnimSpeedScale = (GDExtensionSpxSpriteSetAnimSpeedScale)(dlsymGD("spx_sprite_set_anim_speed_scale"))
	x.SpxSpriteGetAnimSpeedScale = (GDExtensionSpxSpriteGetAnimSpeedScale)(dlsymGD("spx_sprite_get_anim_speed_scale"))
	x.SpxSpriteGetAnimPlayingSpeed = (GDExtensionSpxSpriteGetAnimPlayingSpeed)(dlsymGD("spx_sprite_get_anim_playing_speed"))
	x.SpxSpriteSetAnimCentered = (GDExtensionSpxSpriteSetAnimCentered)(dlsymGD("spx_sprite_set_anim_centered"))
	x.SpxSpriteIsAnimCentered = (GDExtensionSpxSpriteIsAnimCentered)(dlsymGD("spx_sprite_is_anim_centered"))
	x.SpxSpriteSetAnimOffset = (GDExtensionSpxSpriteSetAnimOffset)(dlsymGD("spx_sprite_set_anim_offset"))
	x.SpxSpriteGetAnimOffset = (GDExtensionSpxSpriteGetAnimOffset)(dlsymGD("spx_sprite_get_anim_offset"))
	x.SpxSpriteSetAnimFlipH = (GDExtensionSpxSpriteSetAnimFlipH)(dlsymGD("spx_sprite_set_anim_flip_h"))
	x.SpxSpriteIsAnimFlippedH = (GDExtensionSpxSpriteIsAnimFlippedH)(dlsymGD("spx_sprite_is_anim_flipped_h"))
	x.SpxSpriteSetAnimFlipV = (GDExtensionSpxSpriteSetAnimFlipV)(dlsymGD("spx_sprite_set_anim_flip_v"))
	x.SpxSpriteIsAnimFlippedV = (GDExtensionSpxSpriteIsAnimFlippedV)(dlsymGD("spx_sprite_is_anim_flipped_v"))
	x.SpxSpriteSetVelocity = (GDExtensionSpxSpriteSetVelocity)(dlsymGD("spx_sprite_set_velocity"))
	x.SpxSpriteGetVelocity = (GDExtensionSpxSpriteGetVelocity)(dlsymGD("spx_sprite_get_velocity"))
	x.SpxSpriteIsOnFloor = (GDExtensionSpxSpriteIsOnFloor)(dlsymGD("spx_sprite_is_on_floor"))
	x.SpxSpriteIsOnFloorOnly = (GDExtensionSpxSpriteIsOnFloorOnly)(dlsymGD("spx_sprite_is_on_floor_only"))
	x.SpxSpriteIsOnWall = (GDExtensionSpxSpriteIsOnWall)(dlsymGD("spx_sprite_is_on_wall"))
	x.SpxSpriteIsOnWallOnly = (GDExtensionSpxSpriteIsOnWallOnly)(dlsymGD("spx_sprite_is_on_wall_only"))
	x.SpxSpriteIsOnCeiling = (GDExtensionSpxSpriteIsOnCeiling)(dlsymGD("spx_sprite_is_on_ceiling"))
	x.SpxSpriteIsOnCeilingOnly = (GDExtensionSpxSpriteIsOnCeilingOnly)(dlsymGD("spx_sprite_is_on_ceiling_only"))
	x.SpxSpriteGetLastMotion = (GDExtensionSpxSpriteGetLastMotion)(dlsymGD("spx_sprite_get_last_motion"))
	x.SpxSpriteGetPositionDelta = (GDExtensionSpxSpriteGetPositionDelta)(dlsymGD("spx_sprite_get_position_delta"))
	x.SpxSpriteGetFloorNormal = (GDExtensionSpxSpriteGetFloorNormal)(dlsymGD("spx_sprite_get_floor_normal"))
	x.SpxSpriteGetWallNormal = (GDExtensionSpxSpriteGetWallNormal)(dlsymGD("spx_sprite_get_wall_normal"))
	x.SpxSpriteGetRealVelocity = (GDExtensionSpxSpriteGetRealVelocity)(dlsymGD("spx_sprite_get_real_velocity"))
	x.SpxSpriteMoveAndSlide = (GDExtensionSpxSpriteMoveAndSlide)(dlsymGD("spx_sprite_move_and_slide"))
	x.SpxSpriteSetGravity = (GDExtensionSpxSpriteSetGravity)(dlsymGD("spx_sprite_set_gravity"))
	x.SpxSpriteGetGravity = (GDExtensionSpxSpriteGetGravity)(dlsymGD("spx_sprite_get_gravity"))
	x.SpxSpriteSetMass = (GDExtensionSpxSpriteSetMass)(dlsymGD("spx_sprite_set_mass"))
	x.SpxSpriteGetMass = (GDExtensionSpxSpriteGetMass)(dlsymGD("spx_sprite_get_mass"))
	x.SpxSpriteAddForce = (GDExtensionSpxSpriteAddForce)(dlsymGD("spx_sprite_add_force"))
	x.SpxSpriteAddImpulse = (GDExtensionSpxSpriteAddImpulse)(dlsymGD("spx_sprite_add_impulse"))
	x.SpxSpriteSetCollisionLayer = (GDExtensionSpxSpriteSetCollisionLayer)(dlsymGD("spx_sprite_set_collision_layer"))
	x.SpxSpriteGetCollisionLayer = (GDExtensionSpxSpriteGetCollisionLayer)(dlsymGD("spx_sprite_get_collision_layer"))
	x.SpxSpriteSetCollisionMask = (GDExtensionSpxSpriteSetCollisionMask)(dlsymGD("spx_sprite_set_collision_mask"))
	x.SpxSpriteGetCollisionMask = (GDExtensionSpxSpriteGetCollisionMask)(dlsymGD("spx_sprite_get_collision_mask"))
	x.SpxSpriteSetTriggerLayer = (GDExtensionSpxSpriteSetTriggerLayer)(dlsymGD("spx_sprite_set_trigger_layer"))
	x.SpxSpriteGetTriggerLayer = (GDExtensionSpxSpriteGetTriggerLayer)(dlsymGD("spx_sprite_get_trigger_layer"))
	x.SpxSpriteSetTriggerMask = (GDExtensionSpxSpriteSetTriggerMask)(dlsymGD("spx_sprite_set_trigger_mask"))
	x.SpxSpriteGetTriggerMask = (GDExtensionSpxSpriteGetTriggerMask)(dlsymGD("spx_sprite_get_trigger_mask"))
	x.SpxSpriteSetColliderRect = (GDExtensionSpxSpriteSetColliderRect)(dlsymGD("spx_sprite_set_collider_rect"))
	x.SpxSpriteSetColliderCircle = (GDExtensionSpxSpriteSetColliderCircle)(dlsymGD("spx_sprite_set_collider_circle"))
	x.SpxSpriteSetColliderCapsule = (GDExtensionSpxSpriteSetColliderCapsule)(dlsymGD("spx_sprite_set_collider_capsule"))
	x.SpxSpriteSetCollisionEnabled = (GDExtensionSpxSpriteSetCollisionEnabled)(dlsymGD("spx_sprite_set_collision_enabled"))
	x.SpxSpriteIsCollisionEnabled = (GDExtensionSpxSpriteIsCollisionEnabled)(dlsymGD("spx_sprite_is_collision_enabled"))
	x.SpxSpriteSetTriggerRect = (GDExtensionSpxSpriteSetTriggerRect)(dlsymGD("spx_sprite_set_trigger_rect"))
	x.SpxSpriteSetTriggerCircle = (GDExtensionSpxSpriteSetTriggerCircle)(dlsymGD("spx_sprite_set_trigger_circle"))
	x.SpxSpriteSetTriggerCapsule = (GDExtensionSpxSpriteSetTriggerCapsule)(dlsymGD("spx_sprite_set_trigger_capsule"))
	x.SpxSpriteSetTriggerEnabled = (GDExtensionSpxSpriteSetTriggerEnabled)(dlsymGD("spx_sprite_set_trigger_enabled"))
	x.SpxSpriteIsTriggerEnabled = (GDExtensionSpxSpriteIsTriggerEnabled)(dlsymGD("spx_sprite_is_trigger_enabled"))
	x.SpxUiCreateNode = (GDExtensionSpxUiCreateNode)(dlsymGD("spx_ui_create_node"))
	x.SpxUiCreateButton = (GDExtensionSpxUiCreateButton)(dlsymGD("spx_ui_create_button"))
	x.SpxUiCreateLabel = (GDExtensionSpxUiCreateLabel)(dlsymGD("spx_ui_create_label"))
	x.SpxUiCreateImage = (GDExtensionSpxUiCreateImage)(dlsymGD("spx_ui_create_image"))
	x.SpxUiCreateToggle = (GDExtensionSpxUiCreateToggle)(dlsymGD("spx_ui_create_toggle"))
	x.SpxUiCreateSlider = (GDExtensionSpxUiCreateSlider)(dlsymGD("spx_ui_create_slider"))
	x.SpxUiCreateInput = (GDExtensionSpxUiCreateInput)(dlsymGD("spx_ui_create_input"))
	x.SpxUiDestroyNode = (GDExtensionSpxUiDestroyNode)(dlsymGD("spx_ui_destroy_node"))
	x.SpxUiGetType = (GDExtensionSpxUiGetType)(dlsymGD("spx_ui_get_type"))
	x.SpxUiSetText = (GDExtensionSpxUiSetText)(dlsymGD("spx_ui_set_text"))
	x.SpxUiGetText = (GDExtensionSpxUiGetText)(dlsymGD("spx_ui_get_text"))
	x.SpxUiSetTexture = (GDExtensionSpxUiSetTexture)(dlsymGD("spx_ui_set_texture"))
	x.SpxUiGetTexture = (GDExtensionSpxUiGetTexture)(dlsymGD("spx_ui_get_texture"))
	x.SpxUiSetColor = (GDExtensionSpxUiSetColor)(dlsymGD("spx_ui_set_color"))
	x.SpxUiGetColor = (GDExtensionSpxUiGetColor)(dlsymGD("spx_ui_get_color"))
	x.SpxUiSetFontSize = (GDExtensionSpxUiSetFontSize)(dlsymGD("spx_ui_set_font_size"))
	x.SpxUiGetFontSize = (GDExtensionSpxUiGetFontSize)(dlsymGD("spx_ui_get_font_size"))
	x.SpxUiSetVisible = (GDExtensionSpxUiSetVisible)(dlsymGD("spx_ui_set_visible"))
	x.SpxUiGetVisible = (GDExtensionSpxUiGetVisible)(dlsymGD("spx_ui_get_visible"))
	x.SpxUiSetInteractable = (GDExtensionSpxUiSetInteractable)(dlsymGD("spx_ui_set_interactable"))
	x.SpxUiGetInteractable = (GDExtensionSpxUiGetInteractable)(dlsymGD("spx_ui_get_interactable"))
	x.SpxUiSetRect = (GDExtensionSpxUiSetRect)(dlsymGD("spx_ui_set_rect"))
	x.SpxUiGetRect = (GDExtensionSpxUiGetRect)(dlsymGD("spx_ui_get_rect"))
	}
