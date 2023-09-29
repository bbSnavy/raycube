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

func (game *Game) Player() *Player {
	return game.world.player
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

	world.game = game

	world.Init()

	return
}

func (game *Game) Start() (err error) {
	rl.SetConfigFlags(rl.FlagWindowResizable)

	game.StartWindow()
	defer game.CloseWindow()

	rl.SetTargetFPS(120)
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

	err = game.world.Tick()
	if err != nil {
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
		return
	}

	if rl.IsKeyPressed(rl.KeyF2) {
		rl.TakeScreenshot("frame.png")
	}

	if rl.IsKeyPressed(rl.KeyF) {
		rl.ToggleFullscreen()
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		game.Player().AddVelocity(Vector3New(0.0, 0.5, 0.0))
	}

	if rl.IsKeyDown(rl.KeyLeftShift) {
		game.Player().AddAcceleration(Vector3New(0.0, -0.05, 0.0))
	}

	if rl.IsKeyDown(rl.KeyW) {
		game.Player().
			AddAcceleration(game.Player().Target().Normalize().Mul(Vector3New(0.03, 0.00, 0.03)))
	}

	if rl.IsKeyDown(rl.KeyS) {
		game.Player().
			AddAcceleration(game.Player().Target().Normalize().Mul(Vector3New(0.03, 0.00, 0.03)).Inv())
	}

	mouseDelta := rl.GetMouseDelta()
	game.Player().AddRotation(Vector3New(0.0, mouseDelta.X*0.25, mouseDelta.Y*0.25))

	return
}

func (game *Game) ProcessCamera() (err error) {
	player := game.Player()

	game.camera.Position = player.Position().Add(player.CameraOffset()).ToRaylib()
	game.camera.Target = player.TargetRelative().Add(player.CameraOffset()).ToRaylib()

	return
}

func (game *Game) Render() (err error) {
	rl.BeginDrawing()

	rl.ClearBackground(rl.SkyBlue)

	{
		rl.BeginMode3D(*game.camera)

		err = game.ProcessCamera()
		if err != nil {
			return
		}

		err = game.world.Render()
		if err != nil {
			return
		}

		rl.EndMode3D()
	}

	{
		rl.DrawRectangle(
			int32(rl.GetRenderWidth())/2-4,
			int32(rl.GetRenderHeight())/2-4,
			8, 8,
			rl.RayWhite)
	}

	rl.DrawFPS(16, 16)

	rl.EndDrawing()

	return
}
