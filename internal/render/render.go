package render

import (
	"main/internal/entity/base"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Draw3DObject(data base.BasicStaticModel) {
	rl.DrawModelEx(data.Model, data.Position, data.RotationAxis, data.RotationAngle, data.Scale, data.Tint)
}

func CleanUp(data base.BasicStaticModel) {
	rl.UnloadModel(data.Model)
	rl.UnloadTexture(data.Texture)
}
