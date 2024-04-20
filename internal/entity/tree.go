package entity

import (
	"main/internal/entity/base"
	"main/internal/global"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Tree struct {
	Model base.BasicStaticModel
}

func NewTree(modelPath string, texturePath string, posistion rl.Vector3, scale rl.Vector3) *Tree {
	model := rl.LoadModel(modelPath)
	texture := rl.LoadTexture(texturePath)
	rl.SetMaterialTexture(model.Materials, rl.MapDiffuse, texture)
	return &Tree{
		Model: base.BasicStaticModel{
			Model:         model,
			Texture:       texture,
			Position:      posistion,
			Scale:         scale,
			RotationAxis:  global.VECTOR_UP,
			RotationAngle: 0,
			Tint:          rl.White,
		},
	}
}
