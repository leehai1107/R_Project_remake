package base

import (
	"main/internal/global"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type BasicStaticModel struct {
	Model         rl.Model
	Texture       rl.Texture2D
	RotationAxis  rl.Vector3
	RotationAngle float32
	Position      rl.Vector3
	Scale         rl.Vector3
	Tint          rl.Color
}

func NewBasicStaticModel() *BasicStaticModel {
	return &BasicStaticModel{
		RotationAxis: global.VECTOR_UP,
	}
}

func (b *BasicStaticModel) SetPosition(x, y, z float32) {
	b.Position.X = x
	b.Position.Y = y
	b.Position.Z = z
}

func (b *BasicStaticModel) SetScale(s rl.Vector3) {
	b.Scale = s
}

func (b *BasicStaticModel) SetTint(c rl.Color) {
	b.Tint = c
}

func (b *BasicStaticModel) SetModel(m rl.Model) {
	b.Model = m
}

func (b *BasicStaticModel) SetTexture(t rl.Texture2D) {
	b.Texture = t
}

func (b *BasicStaticModel) SetRotationAxis(r rl.Vector3) {
	b.RotationAxis = r
}

func (b *BasicStaticModel) SetRotationAngle(a float32) {
	b.RotationAngle = a
}

func (b *BasicStaticModel) GetModel() rl.Model {
	return b.Model
}

func (b *BasicStaticModel) GetPosition() rl.Vector3 {
	return b.Position
}

func (b *BasicStaticModel) GetScale() rl.Vector3 {
	return b.Scale
}

func (b *BasicStaticModel) GetTint() rl.Color {
	return b.Tint
}

func (b *BasicStaticModel) GetTexture() rl.Texture2D {
	return b.Texture
}

func (b *BasicStaticModel) GetRotationAxis() rl.Vector3 {
	return b.RotationAxis
}

func (b *BasicStaticModel) GetRotationAngle() float32 {
	return b.RotationAngle
}
