package main

import "log"

func main() {
	game := (&Game{}).Init()
	log.Println(game.Start())

	//rl.InitWindow(1280, 720, "raycube")
	//defer rl.CloseWindow()
	//
	//rl.SetTargetFPS(300)
	//rl.DisableCursor()
	//
	//for !rl.WindowShouldClose() {
	//	if rl.IsKeyPressed(rl.KeyEscape) {
	//		break
	//	}
	//
	//	rl.UpdateCamera(camera, rl.CameraFirstPerson)
	//
	//	rl.BeginDrawing()
	//	rl.BeginMode3D(*camera)
	//
	//	rl.ClearBackground(rl.SkyBlue)
	//
	//	rl.DrawCube(
	//		rl.Vector3{X: 4.0, Z: 4.0, Y: 2.0},
	//		4,
	//		4,
	//		4,
	//		rl.RayWhite)
	//
	//	rl.DrawCube(
	//		camera.Position,
	//		1,
	//		1,
	//		1,
	//		rl.Red)
	//
	//	rl.EndMode3D()
	//
	//	rl.EndDrawing()
	//}
}
