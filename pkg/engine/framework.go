package engine

import (
	"reflect"
)

var (
	Id2Sprites         = make(map[Object]ISpriter)
	Id2UiNodes         = make(map[Object]IUiNode)
	TimeSinceGameStart = float32(0)
	name2SpriteType    = make(map[string]reflect.Type)
)

func isNodeExist(id Object) bool {
	if _, ok := Id2UiNodes[id]; ok {
		return true
	}
	if _, ok := Id2Sprites[id]; ok {
		return true
	}
	return false
}

func getPrefabPath(name string) string {
	assetName := name
	return "res://assets/prefabs/" + assetName + ".tscn"
}

func getUiPath(name string) string {
	assetName := name
	return "res://assets/ui/" + assetName + ".tscn"
}

func InternalInitEngine() {
	initKeyCode()
}

func InternalUpdateEngine(delta float32) {
	updateTimers(delta)
	updateTweens(delta)
}

func ClearAllSprites() {
	for _, sprite := range Id2Sprites {
		sprite.Destroy()
	}
	Id2Sprites = make(map[Object]ISpriter)
}

func RegisterSpriteType[T any]() {
	tType := reflect.TypeOf((*T)(nil)).Elem()
	name := tType.Name()
	name2SpriteType[name] = tType
}

func GetSprite(id Object) ISpriter {
	if sprite, ok := Id2Sprites[id]; ok {
		return sprite
	}
	return nil
}

func BindSceneInstantiatedSprite(id Object, type_name string) {
	if t, ok := name2SpriteType[type_name]; ok {
		createSprite(t, id)
	} else {
		println("BindSceneInstantiatedSprite: type not found", type_name)
	}
}
func CreateSprite[T any]() *T {
	tType := reflect.TypeOf((*T)(nil)).Elem()
	name := tType.Name()
	id := SpriteMgr.CreateSprite(getPrefabPath(name))
	spriteValue := createSprite(tType, id)
	return spriteValue.Addr().Interface().(*T)
}
func CreateUI[T any](prefabName string) *T {
	tType := reflect.TypeOf((*T)(nil)).Elem()
	name := tType.Name()
	if prefabName != "" {
		name = prefabName
	}
	nodeValue := reflect.New(tType).Elem()
	id := UiMgr.CreateNode(getUiPath(name))
	node := nodeValue.Addr().Interface().(IUiNode)
	node.SetId(id)
	node.onCreate()
	Id2UiNodes[id] = node
	node.OnStart()
	return nodeValue.Addr().Interface().(*T)
}

func createSprite(tType reflect.Type, id Object) reflect.Value {
	spriteValue := reflect.New(tType).Elem()
	sprite := spriteValue.Addr().Interface().(ISpriter)
	sprite.SetId(id)
	sprite.onCreate()
	Id2Sprites[id] = sprite
	sprite.OnStart()
	return spriteValue
}
