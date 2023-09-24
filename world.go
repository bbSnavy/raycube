package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
)

type World struct {
	game *Game

	player *Player

	entities []Entity
	cubes    []*Cube
}

func (world *World) Player() *Player {
	return world.player
}

func (world *World) Init() *World {
	if world.game == nil {
		panic("game is nil")
	}

	world.player = (&Player{}).Init()

	world.entities = make([]Entity, 0)
	world.cubes = make([]*Cube, 0)

	world.entities = append(world.entities, world.player)

	for x := 0; x < 64; x++ {
		for z := 0; z < 64; z++ {
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
	for _, entity := range world.entities {
		entity.Tick(world)
	}

	return
}

var (
	model       rl.Model
	modelLoaded = false
)

func (world *World) Render() (err error) {
	if !modelLoaded {
		texture := rl.LoadTexture(os.Getenv("TEXTURE"))
		mesh := rl.GenMeshCube(1.0, 1.0, 1.0)
		model = rl.LoadModelFromMesh(mesh)
		model.Materials.GetMap(rl.MapDiffuse).Texture = texture

		modelLoaded = true
	}

	for _, cube := range world.cubes {
		cubePosition := cube.Position()

		//if !cube.loaded {
		//
		//	cube.model = model
		//	cube.loaded = true
		//}

		rl.DrawModel(
			model,
			cubePosition.ToRaylib(),
			1.0,
			rl.White)
	}

	return
}
