package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
)

type Game struct {
	settings *Settings

	world *World

	active bool

	camera *rl.Camera3D
}

func (game *Game) Init() *Game {
	game.settings = (&Settings{}).Init()

	game.camera = game.InitCamera()
	game.world = game.InitWorld()

	return game
}

func (game *Game) InitCamera() (camera *rl.Camera3D) {
	camera = &rl.Camera3D{}
	camera.Position = rl.Vector3{Y: 2.0, Z: 4.0}
	camera.Target = rl.Vector3{Y: 2.0}
	camera.Up = rl.Vector3{Y: 1.0}
	camera.Fovy = 70.0
	camera.Projection = rl.CameraPerspective

	return
}

func (game *Game) InitWorld() (world *World) {
	world = &World{}

	world.Init()

	return
}

func (game *Game) Start() (err error) {
	game.StartWindow()
	defer game.CloseWindow()

	rl.SetTargetFPS(300)
	rl.DisableCursor()

	game.active = true

	for game.IsRunning() {
		game.Tick()
	}

	return
}

func (game *Game) StartWindow() {
	rl.InitWindow(
		game.settings.GameWindowWidth,
		game.settings.GameWindowHeight,
		game.settings.GameWindowName)
}

func (game *Game) CloseWindow() {
	rl.CloseWindow()
}

func (game *Game) IsRunning() bool {
	return !rl.WindowShouldClose() && game.active
}

func (game *Game) Tick() {
	var (
		err error
	)

	err = game.ProcessInputs()
	if err != nil {
		log.Println("failed at processing game inputs:", err)
		return
	}

	err = game.Render()
	if err != nil {
		log.Println("failed at rendering game: err", err)
		return
	}
}

func (game *Game) ProcessInputs() (err error) {
	if rl.IsKeyPressed(rl.KeyEscape) {
		game.active = false
	}

	if rl.IsKeyDown(rl.KeySpace) {
		game.camera.Position = rl.Vector3Add(
			game.camera.Position,
			rl.Vector3{
				Y: 0.05,
			})

		game.camera.Target.Y += 0.05
	}

	if rl.IsKeyDown(rl.KeyLeftShift) {
		game.camera.Position = rl.Vector3Add(
			game.camera.Position,
			rl.Vector3{
				Y: -0.05,
			})

		game.camera.Target.Y -= 0.05
	}

	if rl.IsKeyPressed(rl.KeyF2) {
		rl.TakeScreenshot("frame.png")
	}

	return
}

func (game *Game) Render() (err error) {
	rl.UpdateCamera(game.camera, rl.CameraThirdPerson)

	rl.BeginDrawing()

	rl.ClearBackground(rl.SkyBlue)

	{
		rl.BeginMode3D(*game.camera)

		err = game.RenderCamera()
		if err != nil {
			return
		}

		err = game.world.Render()
		if err != nil {
			return
		}

		rl.EndMode3D()
	}

	rl.DrawFPS(16, 16)

	rl.EndDrawing()

	return
}

func (game *Game) RenderCamera() (err error) {
	rl.DrawCube(
		game.camera.Target,
		1.0,
		2.0,
		1.0,
		rl.Black)

	return
}
