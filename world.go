package main

import rl "github.com/gen2brain/raylib-go/raylib"

type World struct {
	Cubes []*Cube
}

func (world *World) Init() *World {
	world.Cubes = make([]*Cube, 0)

	for x := 0; x < 4; x++ {
		for z := 0; z < 4; z++ {
			for y := 0; y < 1; y++ {
				cube := (&Cube{
					position: Vector3New(
						float32(x),
						float32(y),
						float32(z)),
				}).Init()

				world.Cubes = append(world.Cubes, cube)
			}
		}
	}

	return world
}

func (world *World) Render() (err error) {
	for _, cube := range world.Cubes {
		cubePosition := cube.Position()

		rl.DrawCube(
			cubePosition.ToRaylib(),
			1,
			1,
			1,
			rl.Purple)
	}

	return
}
