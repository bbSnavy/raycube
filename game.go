package main

import (
	"log"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
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

	if rl.IsKeyPressed(rl.KeyF11) {
		rl.ToggleFullscreen()
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		game.Player().AddPosition(Vector3New(0.0, 1.0, 0.0))
	}

	if rl.IsKeyPressed(rl.KeyLeftShift) {
		game.Player().AddPosition(Vector3New(0.0, -1.0, 0.0))
	}

	if rl.IsKeyDown(rl.KeyRight) {
		game.Player().AddRotation(Vector3New(0.0, 1.0, 0.0))
	}

	if rl.IsKeyDown(rl.KeyLeft) {
		game.Player().AddRotation(Vector3New(0.0, -1.0, 0.0))
	}

	if rl.IsKeyDown(rl.KeyUp) {
		game.Player().AddRotation(Vector3New(0.0, 0.0, 1.0))
	}

	if rl.IsKeyDown(rl.KeyDown) {
		game.Player().AddRotation(Vector3New(0.0, 0.0, -1.0))
	}

	mouseDelta := rl.GetMouseDelta()
	mouseDeltaAbs := []float32{
		float32(math.Abs(float64(mouseDelta.X))),
		float32(math.Abs(float64(mouseDelta.Y))),
	}
	if mouseDelta.X < 0 {
		game.Player().AddRotation(Vector3New(0.0, -mouseDeltaAbs[0], 0.0))
	}
	if mouseDelta.X > 0 {
		game.Player().AddRotation(Vector3New(0.0, +mouseDeltaAbs[0], 0.0))
	}
	if mouseDelta.Y < 0 {
		game.Player().AddRotation(Vector3New(0.0, 0.0, +mouseDeltaAbs[1]))
	}
	if mouseDelta.Y > 0 {
		game.Player().AddRotation(Vector3New(0.0, 0.0, -mouseDeltaAbs[1]))
	}

	return
}

func (game *Game) ProcessCamera() (err error) {
	player := game.Player()

	game.camera.Position = player.Position().ToRaylib()
	game.camera.Target = Vector3UnitVector(player.Rotation()).
		Normalize().
		Add(player.Position()).
		ToRaylib()

	rl.DrawCube(
		game.camera.Target,
		0.010,
		0.010,
		0.010,
		rl.Black)

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

	rl.DrawFPS(16, 16)

	rl.EndDrawing()

	return
}
