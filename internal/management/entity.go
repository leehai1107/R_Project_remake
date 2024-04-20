package management

import rl "github.com/gen2brain/raylib-go/raylib"

var entities map[int32]any

func Init() {
	entities = make(map[int32]any)
}

func CreateEntity(objectID int32, data any) {
	entities[objectID] = data
	rl.TraceLog(rl.LogInfo, "Entities created!")
}

func DeleteEntity(objectID int32) {
	delete(entities, objectID)
	rl.TraceLog(rl.LogInfo, "Entities deleted!")

}

func GetEntity(objectId int32) any {
	return entities[objectId]
}
