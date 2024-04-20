package cmd

import (
	"main/internal/camera"
	"main/internal/entity"
	"main/internal/management"
	"main/internal/render"
	"main/internal/windows"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Execute interface {
	Game()
}

type execute struct {
	windows windows.Windows
	camera  camera.Camera
}

func Run() Execute {
	return &execute{
		windows: windows.NewWindows(),
		camera:  camera.NewCamera3D(),
	}
}

func (e *execute) Game() {
	e.initialize()

	e.process()

	defer e.cleanup()
}

func (e *execute) initialize() {
	e.windows.Init()
	management.Init()
}

func (e *execute) cleanup() {
	e.windows.Close()
}

func (e *execute) process() {
	tree := entity.NewTree("assets/obj/tree.obj", "assets/texture/tree.png", rl.NewVector3(0, 0, 0), rl.NewVector3(1, 1, 1))
	management.CreateEntity(1, tree)
	for !rl.WindowShouldClose() {
		e.camera.UpdateCamera()
		rl.BeginDrawing()

		rl.BeginMode3D(e.camera.GetCamera())

		rl.ClearBackground(rl.White)

		rl.DrawGrid(100, 1)

		render.Draw3DObject(management.GetEntity(1).(*entity.Tree).Model)

		rl.EndMode3D()

		rl.DrawFPS(10, 10)
		rl.EndDrawing()
	}
	render.CleanUp(management.GetEntity(1).(*entity.Tree).Model)
	management.DeleteEntity(1)

}
