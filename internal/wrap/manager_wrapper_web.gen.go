/*
------------------------------------------------------------------------------
//   This code was generated by template manager_wrapper.go.tmpl.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "manager_wrapper.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------
*/
package wrap

import (
	"fmt"
	. "godot-ext/gdspx/internal/webffi"
	. "godot-ext/gdspx/pkg/engine"
	"reflect"
)

func BindMgr(mgrs []IManager) {
	for _, mgr := range mgrs {
		switch v := mgr.(type) {
		case IAudioMgr:
			AudioMgr = v

		case ICameraMgr:
			CameraMgr = v

		case IInputMgr:
			InputMgr = v

		case IPhysicMgr:
			PhysicMgr = v

		case ISceneMgr:
			SceneMgr = v

		case ISpriteMgr:
			SpriteMgr = v

		case IUiMgr:
			UiMgr = v

		default:
			panic(fmt.Sprintf("engine init error : unknown manager type %s", reflect.TypeOf(mgr).String()))
		}
	}
}

type audioMgr struct {
	baseMgr
}
type cameraMgr struct {
	baseMgr
}
type inputMgr struct {
	baseMgr
}
type physicMgr struct {
	baseMgr
}
type sceneMgr struct {
	baseMgr
}
type spriteMgr struct {
	baseMgr
}
type uiMgr struct {
	baseMgr
}

func createMgrs() []IManager {
	addManager(&audioMgr{})
	addManager(&cameraMgr{})
	addManager(&inputMgr{})
	addManager(&physicMgr{})
	addManager(&sceneMgr{})
	addManager(&spriteMgr{})
	addManager(&uiMgr{})
	return mgrs
}

// call gdextension interface functions

func (pself *audioMgr) PlayAudio(path string) {
	arg0 := ToGdString(path)
	API.SpxAudioPlayAudio.Invoke(arg0)
}
func (pself *audioMgr) SetAudioVolume(volume float32) {
	arg0 := ToGdFloat(volume)
	API.SpxAudioSetAudioVolume.Invoke(arg0)
}
func (pself *audioMgr) GetAudioVolume() float32 {
	_retValue := API.SpxAudioGetAudioVolume.Invoke()
	return ToFloat32(_retValue)
}
func (pself *audioMgr) IsMusicPlaying() bool {
	_retValue := API.SpxAudioIsMusicPlaying.Invoke()
	return ToBool(_retValue)
}
func (pself *audioMgr) PlayMusic(path string) {
	arg0 := ToGdString(path)
	API.SpxAudioPlayMusic.Invoke(arg0)
}
func (pself *audioMgr) SetMusicVolume(volume float32) {
	arg0 := ToGdFloat(volume)
	API.SpxAudioSetMusicVolume.Invoke(arg0)
}
func (pself *audioMgr) GetMusicVolume() float32 {
	_retValue := API.SpxAudioGetMusicVolume.Invoke()
	return ToFloat32(_retValue)
}
func (pself *audioMgr) PauseMusic() {
	API.SpxAudioPauseMusic.Invoke()
}
func (pself *audioMgr) ResumeMusic() {
	API.SpxAudioResumeMusic.Invoke()
}
func (pself *audioMgr) GetMusicTimer() float32 {
	_retValue := API.SpxAudioGetMusicTimer.Invoke()
	return ToFloat32(_retValue)
}
func (pself *audioMgr) SetMusicTimer(time float32) {
	arg0 := ToGdFloat(time)
	API.SpxAudioSetMusicTimer.Invoke(arg0)
}
func (pself *cameraMgr) GetCameraPosition() Vec2 {
	_retValue := API.SpxCameraGetCameraPosition.Invoke()
	return ToVec2(_retValue)
}
func (pself *cameraMgr) SetCameraPosition(position Vec2) {
	arg0 := ToGdVec2(position)
	API.SpxCameraSetCameraPosition.Invoke(arg0)
}
func (pself *cameraMgr) GetCameraZoom() Vec2 {
	_retValue := API.SpxCameraGetCameraZoom.Invoke()
	return ToVec2(_retValue)
}
func (pself *cameraMgr) SetCameraZoom(size Vec2) {
	arg0 := ToGdVec2(size)
	API.SpxCameraSetCameraZoom.Invoke(arg0)
}
func (pself *cameraMgr) GetViewportRect() Rect2 {
	_retValue := API.SpxCameraGetViewportRect.Invoke()
	return ToRect2(_retValue)
}
func (pself *inputMgr) GetMousePos() Vec2 {
	_retValue := API.SpxInputGetMousePos.Invoke()
	return ToVec2(_retValue)
}
func (pself *inputMgr) GetKey(key int64) bool {
	arg0 := ToGdInt(key)
	_retValue := API.SpxInputGetKey.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *inputMgr) GetMouseState(mouse_id int64) bool {
	arg0 := ToGdInt(mouse_id)
	_retValue := API.SpxInputGetMouseState.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *inputMgr) GetKeyState(key int64) int64 {
	arg0 := ToGdInt(key)
	_retValue := API.SpxInputGetKeyState.Invoke(arg0)
	return ToInt64(_retValue)
}
func (pself *inputMgr) GetAxis(neg_action string, pos_action string) float32 {
	arg0 := ToGdString(neg_action)
	arg1 := ToGdString(pos_action)
	_retValue := API.SpxInputGetAxis.Invoke(arg0, arg1)
	return ToFloat32(_retValue)
}
func (pself *inputMgr) IsActionPressed(action string) bool {
	arg0 := ToGdString(action)
	_retValue := API.SpxInputIsActionPressed.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *inputMgr) IsActionJustPressed(action string) bool {
	arg0 := ToGdString(action)
	_retValue := API.SpxInputIsActionJustPressed.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *inputMgr) IsActionJustReleased(action string) bool {
	arg0 := ToGdString(action)
	_retValue := API.SpxInputIsActionJustReleased.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *physicMgr) Raycast(from Vec2, to Vec2, collision_mask int64) Object {
	arg0 := ToGdVec2(from)
	arg1 := ToGdVec2(to)
	arg2 := ToGdInt(collision_mask)
	_retValue := API.SpxPhysicRaycast.Invoke(arg0, arg1, arg2)
	return ToObject(_retValue)
}
func (pself *physicMgr) CheckCollision(from Vec2, to Vec2, collision_mask int64, collide_with_areas bool, collide_with_bodies bool) bool {
	arg0 := ToGdVec2(from)
	arg1 := ToGdVec2(to)
	arg2 := ToGdInt(collision_mask)
	arg3 := ToGdBool(collide_with_areas)
	arg4 := ToGdBool(collide_with_bodies)
	_retValue := API.SpxPhysicCheckCollision.Invoke(arg0, arg1, arg2, arg3, arg4)
	return ToBool(_retValue)
}
func (pself *sceneMgr) ChangeSceneToFile(path string) {
	arg0 := ToGdString(path)
	API.SpxSceneChangeSceneToFile.Invoke(arg0)
}
func (pself *sceneMgr) ReloadCurrentScene() int64 {
	_retValue := API.SpxSceneReloadCurrentScene.Invoke()
	return ToInt64(_retValue)
}
func (pself *sceneMgr) UnloadCurrentScene() {
	API.SpxSceneUnloadCurrentScene.Invoke()
}
func (pself *spriteMgr) SetDontDestroyOnLoad(obj Object) {
	arg0 := ToGdObj(obj)
	API.SpxSpriteSetDontDestroyOnLoad.Invoke(arg0)
}
func (pself *spriteMgr) SetProcess(obj Object, is_on bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(is_on)
	API.SpxSpriteSetProcess.Invoke(arg0, arg1)
}
func (pself *spriteMgr) SetPhysicProcess(obj Object, is_on bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(is_on)
	API.SpxSpriteSetPhysicProcess.Invoke(arg0, arg1)
}
func (pself *spriteMgr) SetChildPosition(obj Object, path string, pos Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdString(path)
	arg2 := ToGdVec2(pos)
	API.SpxSpriteSetChildPosition.Invoke(arg0, arg1, arg2)
}
func (pself *spriteMgr) GetChildPosition(obj Object, path string) Vec2 {
	arg0 := ToGdObj(obj)
	arg1 := ToGdString(path)
	_retValue := API.SpxSpriteGetChildPosition.Invoke(arg0, arg1)
	return ToVec2(_retValue)
}
func (pself *spriteMgr) SetChildRotation(obj Object, path string, rot float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdString(path)
	arg2 := ToGdFloat(rot)
	API.SpxSpriteSetChildRotation.Invoke(arg0, arg1, arg2)
}
func (pself *spriteMgr) GetChildRotation(obj Object, path string) float32 {
	arg0 := ToGdObj(obj)
	arg1 := ToGdString(path)
	_retValue := API.SpxSpriteGetChildRotation.Invoke(arg0, arg1)
	return ToFloat32(_retValue)
}
func (pself *spriteMgr) SetChildScale(obj Object, path string, scale Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdString(path)
	arg2 := ToGdVec2(scale)
	API.SpxSpriteSetChildScale.Invoke(arg0, arg1, arg2)
}
func (pself *spriteMgr) GetChildScale(obj Object, path string) Vec2 {
	arg0 := ToGdObj(obj)
	arg1 := ToGdString(path)
	_retValue := API.SpxSpriteGetChildScale.Invoke(arg0, arg1)
	return ToVec2(_retValue)
}
func (pself *spriteMgr) CreateSprite(path string) Object {
	arg0 := ToGdString(path)
	_retValue := API.SpxSpriteCreateSprite.Invoke(arg0)
	return ToObject(_retValue)
}
func (pself *spriteMgr) CloneSprite(obj Object) Object {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteCloneSprite.Invoke(arg0)
	return ToObject(_retValue)
}
func (pself *spriteMgr) DestroySprite(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteDestroySprite.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *spriteMgr) IsSpriteAlive(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteIsSpriteAlive.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *spriteMgr) SetPosition(obj Object, pos Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(pos)
	API.SpxSpriteSetPosition.Invoke(arg0, arg1)
}
func (pself *spriteMgr) GetPosition(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetPosition.Invoke(arg0)
	return ToVec2(_retValue)
}
func (pself *spriteMgr) SetRotation(obj Object, rot float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdFloat(rot)
	API.SpxSpriteSetRotation.Invoke(arg0, arg1)
}
func (pself *spriteMgr) GetRotation(obj Object) float32 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetRotation.Invoke(arg0)
	return ToFloat32(_retValue)
}
func (pself *spriteMgr) SetScale(obj Object, scale Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(scale)
	API.SpxSpriteSetScale.Invoke(arg0, arg1)
}
func (pself *spriteMgr) GetScale(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetScale.Invoke(arg0)
	return ToVec2(_retValue)
}
func (pself *spriteMgr) SetColor(obj Object, color Color) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdColor(color)
	API.SpxSpriteSetColor.Invoke(arg0, arg1)
}
func (pself *spriteMgr) GetColor(obj Object) Color {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetColor.Invoke(arg0)
	return ToColor(_retValue)
}
func (pself *spriteMgr) SetTexture(obj Object, path string) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdString(path)
	API.SpxSpriteSetTexture.Invoke(arg0, arg1)
}
func (pself *spriteMgr) GetTexture(obj Object) string {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetTexture.Invoke(arg0)
	return ToString(_retValue)
}
func (pself *spriteMgr) SetVisible(obj Object, visible bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(visible)
	API.SpxSpriteSetVisible.Invoke(arg0, arg1)
}
func (pself *spriteMgr) GetVisible(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetVisible.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *spriteMgr) GetZIndex(obj Object) int64 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetZIndex.Invoke(arg0)
	return ToInt64(_retValue)
}
func (pself *spriteMgr) SetZIndex(obj Object, z int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(z)
	API.SpxSpriteSetZIndex.Invoke(arg0, arg1)
}
func (pself *spriteMgr) PlayAnim(obj Object, p_name string, p_custom_scale float32, p_from_end bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdString(p_name)
	arg2 := ToGdFloat(p_custom_scale)
	arg3 := ToGdBool(p_from_end)
	API.SpxSpritePlayAnim.Invoke(arg0, arg1, arg2, arg3)
}
func (pself *spriteMgr) PlayBackwardsAnim(obj Object, p_name string) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdString(p_name)
	API.SpxSpritePlayBackwardsAnim.Invoke(arg0, arg1)
}
func (pself *spriteMgr) PauseAnim(obj Object) {
	arg0 := ToGdObj(obj)
	API.SpxSpritePauseAnim.Invoke(arg0)
}
func (pself *spriteMgr) StopAnim(obj Object) {
	arg0 := ToGdObj(obj)
	API.SpxSpriteStopAnim.Invoke(arg0)
}
func (pself *spriteMgr) IsPlayingAnim(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteIsPlayingAnim.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *spriteMgr) SetAnim(obj Object, p_name string) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdString(p_name)
	API.SpxSpriteSetAnim.Invoke(arg0, arg1)
}
func (pself *spriteMgr) GetAnim(obj Object) string {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetAnim.Invoke(arg0)
	return ToString(_retValue)
}
func (pself *spriteMgr) SetAnimFrame(obj Object, p_frame int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(p_frame)
	API.SpxSpriteSetAnimFrame.Invoke(arg0, arg1)
}
func (pself *spriteMgr) GetAnimFrame(obj Object) int64 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetAnimFrame.Invoke(arg0)
	return ToInt64(_retValue)
}
func (pself *spriteMgr) SetAnimSpeedScale(obj Object, p_speed_scale float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdFloat(p_speed_scale)
	API.SpxSpriteSetAnimSpeedScale.Invoke(arg0, arg1)
}
func (pself *spriteMgr) GetAnimSpeedScale(obj Object) float32 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetAnimSpeedScale.Invoke(arg0)
	return ToFloat32(_retValue)
}
func (pself *spriteMgr) GetAnimPlayingSpeed(obj Object) float32 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetAnimPlayingSpeed.Invoke(arg0)
	return ToFloat32(_retValue)
}
func (pself *spriteMgr) SetAnimCentered(obj Object, p_center bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(p_center)
	API.SpxSpriteSetAnimCentered.Invoke(arg0, arg1)
}
func (pself *spriteMgr) IsAnimCentered(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteIsAnimCentered.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *spriteMgr) SetAnimOffset(obj Object, p_offset Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(p_offset)
	API.SpxSpriteSetAnimOffset.Invoke(arg0, arg1)
}
func (pself *spriteMgr) GetAnimOffset(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetAnimOffset.Invoke(arg0)
	return ToVec2(_retValue)
}
func (pself *spriteMgr) SetAnimFlipH(obj Object, p_flip bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(p_flip)
	API.SpxSpriteSetAnimFlipH.Invoke(arg0, arg1)
}
func (pself *spriteMgr) IsAnimFlippedH(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteIsAnimFlippedH.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *spriteMgr) SetAnimFlipV(obj Object, p_flip bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(p_flip)
	API.SpxSpriteSetAnimFlipV.Invoke(arg0, arg1)
}
func (pself *spriteMgr) IsAnimFlippedV(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteIsAnimFlippedV.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *spriteMgr) SetVelocity(obj Object, velocity Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(velocity)
	API.SpxSpriteSetVelocity.Invoke(arg0, arg1)
}
func (pself *spriteMgr) GetVelocity(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetVelocity.Invoke(arg0)
	return ToVec2(_retValue)
}
func (pself *spriteMgr) IsOnFloor(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteIsOnFloor.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *spriteMgr) IsOnFloorOnly(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteIsOnFloorOnly.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *spriteMgr) IsOnWall(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteIsOnWall.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *spriteMgr) IsOnWallOnly(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteIsOnWallOnly.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *spriteMgr) IsOnCeiling(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteIsOnCeiling.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *spriteMgr) IsOnCeilingOnly(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteIsOnCeilingOnly.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *spriteMgr) GetLastMotion(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetLastMotion.Invoke(arg0)
	return ToVec2(_retValue)
}
func (pself *spriteMgr) GetPositionDelta(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetPositionDelta.Invoke(arg0)
	return ToVec2(_retValue)
}
func (pself *spriteMgr) GetFloorNormal(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetFloorNormal.Invoke(arg0)
	return ToVec2(_retValue)
}
func (pself *spriteMgr) GetWallNormal(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetWallNormal.Invoke(arg0)
	return ToVec2(_retValue)
}
func (pself *spriteMgr) GetRealVelocity(obj Object) Vec2 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetRealVelocity.Invoke(arg0)
	return ToVec2(_retValue)
}
func (pself *spriteMgr) MoveAndSlide(obj Object) {
	arg0 := ToGdObj(obj)
	API.SpxSpriteMoveAndSlide.Invoke(arg0)
}
func (pself *spriteMgr) SetGravity(obj Object, gravity float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdFloat(gravity)
	API.SpxSpriteSetGravity.Invoke(arg0, arg1)
}
func (pself *spriteMgr) GetGravity(obj Object) float32 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetGravity.Invoke(arg0)
	return ToFloat32(_retValue)
}
func (pself *spriteMgr) SetMass(obj Object, mass float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdFloat(mass)
	API.SpxSpriteSetMass.Invoke(arg0, arg1)
}
func (pself *spriteMgr) GetMass(obj Object) float32 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetMass.Invoke(arg0)
	return ToFloat32(_retValue)
}
func (pself *spriteMgr) AddForce(obj Object, force Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(force)
	API.SpxSpriteAddForce.Invoke(arg0, arg1)
}
func (pself *spriteMgr) AddImpulse(obj Object, impulse Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(impulse)
	API.SpxSpriteAddImpulse.Invoke(arg0, arg1)
}
func (pself *spriteMgr) SetCollisionLayer(obj Object, layer int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(layer)
	API.SpxSpriteSetCollisionLayer.Invoke(arg0, arg1)
}
func (pself *spriteMgr) GetCollisionLayer(obj Object) int64 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetCollisionLayer.Invoke(arg0)
	return ToInt64(_retValue)
}
func (pself *spriteMgr) SetCollisionMask(obj Object, mask int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(mask)
	API.SpxSpriteSetCollisionMask.Invoke(arg0, arg1)
}
func (pself *spriteMgr) GetCollisionMask(obj Object) int64 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetCollisionMask.Invoke(arg0)
	return ToInt64(_retValue)
}
func (pself *spriteMgr) SetTriggerLayer(obj Object, layer int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(layer)
	API.SpxSpriteSetTriggerLayer.Invoke(arg0, arg1)
}
func (pself *spriteMgr) GetTriggerLayer(obj Object) int64 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetTriggerLayer.Invoke(arg0)
	return ToInt64(_retValue)
}
func (pself *spriteMgr) SetTriggerMask(obj Object, mask int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(mask)
	API.SpxSpriteSetTriggerMask.Invoke(arg0, arg1)
}
func (pself *spriteMgr) GetTriggerMask(obj Object) int64 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteGetTriggerMask.Invoke(arg0)
	return ToInt64(_retValue)
}
func (pself *spriteMgr) SetColliderRect(obj Object, center Vec2, size Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(center)
	arg2 := ToGdVec2(size)
	API.SpxSpriteSetColliderRect.Invoke(arg0, arg1, arg2)
}
func (pself *spriteMgr) SetColliderCircle(obj Object, center Vec2, radius float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(center)
	arg2 := ToGdFloat(radius)
	API.SpxSpriteSetColliderCircle.Invoke(arg0, arg1, arg2)
}
func (pself *spriteMgr) SetColliderCapsule(obj Object, center Vec2, size Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(center)
	arg2 := ToGdVec2(size)
	API.SpxSpriteSetColliderCapsule.Invoke(arg0, arg1, arg2)
}
func (pself *spriteMgr) SetCollisionEnabled(obj Object, enabled bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(enabled)
	API.SpxSpriteSetCollisionEnabled.Invoke(arg0, arg1)
}
func (pself *spriteMgr) IsCollisionEnabled(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteIsCollisionEnabled.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *spriteMgr) SetTriggerRect(obj Object, center Vec2, size Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(center)
	arg2 := ToGdVec2(size)
	API.SpxSpriteSetTriggerRect.Invoke(arg0, arg1, arg2)
}
func (pself *spriteMgr) SetTriggerCircle(obj Object, center Vec2, radius float32) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(center)
	arg2 := ToGdFloat(radius)
	API.SpxSpriteSetTriggerCircle.Invoke(arg0, arg1, arg2)
}
func (pself *spriteMgr) SetTriggerCapsule(obj Object, center Vec2, size Vec2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdVec2(center)
	arg2 := ToGdVec2(size)
	API.SpxSpriteSetTriggerCapsule.Invoke(arg0, arg1, arg2)
}
func (pself *spriteMgr) SetTriggerEnabled(obj Object, trigger bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(trigger)
	API.SpxSpriteSetTriggerEnabled.Invoke(arg0, arg1)
}
func (pself *spriteMgr) IsTriggerEnabled(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxSpriteIsTriggerEnabled.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *uiMgr) CreateNode(path string) Object {
	arg0 := ToGdString(path)
	_retValue := API.SpxUiCreateNode.Invoke(arg0)
	return ToObject(_retValue)
}
func (pself *uiMgr) CreateButton(path string, text string) Object {
	arg0 := ToGdString(path)
	arg1 := ToGdString(text)
	_retValue := API.SpxUiCreateButton.Invoke(arg0, arg1)
	return ToObject(_retValue)
}
func (pself *uiMgr) CreateLabel(path string, text string) Object {
	arg0 := ToGdString(path)
	arg1 := ToGdString(text)
	_retValue := API.SpxUiCreateLabel.Invoke(arg0, arg1)
	return ToObject(_retValue)
}
func (pself *uiMgr) CreateImage(path string) Object {
	arg0 := ToGdString(path)
	_retValue := API.SpxUiCreateImage.Invoke(arg0)
	return ToObject(_retValue)
}
func (pself *uiMgr) CreateToggle(path string, value bool) Object {
	arg0 := ToGdString(path)
	arg1 := ToGdBool(value)
	_retValue := API.SpxUiCreateToggle.Invoke(arg0, arg1)
	return ToObject(_retValue)
}
func (pself *uiMgr) CreateSlider(path string, value float32) Object {
	arg0 := ToGdString(path)
	arg1 := ToGdFloat(value)
	_retValue := API.SpxUiCreateSlider.Invoke(arg0, arg1)
	return ToObject(_retValue)
}
func (pself *uiMgr) CreateInput(path string, text string) Object {
	arg0 := ToGdString(path)
	arg1 := ToGdString(text)
	_retValue := API.SpxUiCreateInput.Invoke(arg0, arg1)
	return ToObject(_retValue)
}
func (pself *uiMgr) DestroyNode(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxUiDestroyNode.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *uiMgr) GetType(obj Object) int64 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxUiGetType.Invoke(arg0)
	return ToInt64(_retValue)
}
func (pself *uiMgr) SetText(obj Object, text string) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdString(text)
	API.SpxUiSetText.Invoke(arg0, arg1)
}
func (pself *uiMgr) GetText(obj Object) string {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxUiGetText.Invoke(arg0)
	return ToString(_retValue)
}
func (pself *uiMgr) SetTexture(obj Object, path string) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdString(path)
	API.SpxUiSetTexture.Invoke(arg0, arg1)
}
func (pself *uiMgr) GetTexture(obj Object) string {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxUiGetTexture.Invoke(arg0)
	return ToString(_retValue)
}
func (pself *uiMgr) SetColor(obj Object, color Color) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdColor(color)
	API.SpxUiSetColor.Invoke(arg0, arg1)
}
func (pself *uiMgr) GetColor(obj Object) Color {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxUiGetColor.Invoke(arg0)
	return ToColor(_retValue)
}
func (pself *uiMgr) SetFontSize(obj Object, size int64) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdInt(size)
	API.SpxUiSetFontSize.Invoke(arg0, arg1)
}
func (pself *uiMgr) GetFontSize(obj Object) int64 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxUiGetFontSize.Invoke(arg0)
	return ToInt64(_retValue)
}
func (pself *uiMgr) SetVisible(obj Object, visible bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(visible)
	API.SpxUiSetVisible.Invoke(arg0, arg1)
}
func (pself *uiMgr) GetVisible(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxUiGetVisible.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *uiMgr) SetInteractable(obj Object, interactable bool) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdBool(interactable)
	API.SpxUiSetInteractable.Invoke(arg0, arg1)
}
func (pself *uiMgr) GetInteractable(obj Object) bool {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxUiGetInteractable.Invoke(arg0)
	return ToBool(_retValue)
}
func (pself *uiMgr) SetRect(obj Object, rect Rect2) {
	arg0 := ToGdObj(obj)
	arg1 := ToGdRect2(rect)
	API.SpxUiSetRect.Invoke(arg0, arg1)
}
func (pself *uiMgr) GetRect(obj Object) Rect2 {
	arg0 := ToGdObj(obj)
	_retValue := API.SpxUiGetRect.Invoke(arg0)
	return ToRect2(_retValue)
}