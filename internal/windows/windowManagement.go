package windows

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Windows interface {
	Init()
	Close()
}

type windows struct{}

func NewWindows() Windows {
	return &windows{}
}

func (w *windows) Init() {
	rl.InitWindow(680, 480, "Hello World")
	rl.SetTargetFPS(60)
	rl.SetWindowState(
		rl.FlagWindowResizable,
	)
	rl.SetWindowFocused()
}

func (w *windows) Close() {
	rl.CloseWindow()
}
