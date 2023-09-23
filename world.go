package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type World struct {
	game *Game

	player *Player

	entities []Entity
	cubes    []*Cube
}

func (world *World) Init() *World {
	if world.game == nil {
		panic("game is nil")
	}

	world.player = (&Player{}).Init()

	world.entities = make([]Entity, 0)
	world.cubes = make([]*Cube, 0)

	world.entities = append(world.entities, world.player)

	for x := 0; x < 4; x++ {
		for z := 0; z < 4; z++ {
			for y := 0; y < 1; y++ {
				cube := (&Cube{
					position: Vector3New(
						float32(x),
						float32(y),
						float32(z)),
				}).Init()

				world.cubes = append(world.cubes, cube)
			}
		}
	}

	return world
}

func (world *World) Tick() (err error) {
	cameraHitbox := BoxNew(
		world.player.Position(),
		Vector3New(1.0, 1.0, 1.0))

	for _, cube := range world.cubes {
		cubeBox := cube.Box()

		if CheckCollisionBoxToBox(cameraHitbox, cubeBox) {
			cube.color = rl.Red
		} else {
			cube.color = rl.Purple
		}
	}

	return
}

func (world *World) Render() (err error) {
	for _, cube := range world.cubes {
		cubePosition := cube.Position()

		rl.DrawCube(
			cubePosition.ToRaylib(),
			1,
			1,
			1,
			cube.color)
	}

	return
}
