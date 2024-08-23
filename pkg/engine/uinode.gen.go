package engine

func (pself *UiNode) GetType() int64 {
	return UiMgr.GetType(pself.Id)
}
func (pself *UiNode) SetText(text string) {
	UiMgr.SetText(pself.Id, text)
}
func (pself *UiNode) GetText() string {
	return UiMgr.GetText(pself.Id)
}
func (pself *UiNode) SetTexture(path string) {
	UiMgr.SetTexture(pself.Id, path)
}
func (pself *UiNode) GetTexture() string {
	return UiMgr.GetTexture(pself.Id)
}
func (pself *UiNode) SetColor(color Color) {
	UiMgr.SetColor(pself.Id, color)
}
func (pself *UiNode) GetColor() Color {
	return UiMgr.GetColor(pself.Id)
}
func (pself *UiNode) SetFontSize(size int64) {
	UiMgr.SetFontSize(pself.Id, size)
}
func (pself *UiNode) GetFontSize() int64 {
	return UiMgr.GetFontSize(pself.Id)
}
func (pself *UiNode) SetVisible(visible bool) {
	UiMgr.SetVisible(pself.Id, visible)
}
func (pself *UiNode) GetVisible() bool {
	return UiMgr.GetVisible(pself.Id)
}
func (pself *UiNode) SetInteractable(interactable bool) {
	UiMgr.SetInteractable(pself.Id, interactable)
}
func (pself *UiNode) GetInteractable() bool {
	return UiMgr.GetInteractable(pself.Id)
}
func (pself *UiNode) SetRect(rect Rect2) {
	UiMgr.SetRect(pself.Id, rect)
}
func (pself *UiNode) GetRect() Rect2 {
	return UiMgr.GetRect(pself.Id)
}
