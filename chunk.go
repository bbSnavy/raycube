package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"os"
)

type Chunk struct {
	world *World

	position Vector3

	cubes [4096]*Cube
}

func (chunk *Chunk) CubeAtIndex(index int) *Cube {
	if index < 0 || index >= len(chunk.cubes) {
		return nil
	}

	return chunk.cubes[index]
}

func (chunk *Chunk) Cubes() []*Cube {
	return chunk.cubes[:]
}

func (chunk *Chunk) Position() Vector3 {
	return chunk.position
}

func (chunk *Chunk) PositionBase() Vector3 {
	return chunk.position.Mul(Vector3New(16.0, 16.0, 16.0))
}

func (chunk *Chunk) Init() *Chunk {
	if chunk.world == nil {
		panic("chunk world is nil")
	}

	index := 0
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			for z := 0; z < 16; z++ {
				chunk.cubes[index] = (&Cube{
					chunk:    chunk,
					position: Vector3New(float32(x), float32(y), float32(z)),
				}).Init()

				index++
			}
		}
	}

	return chunk
}

var (
	model       rl.Model
	modelLoaded = false
)

func (chunk *Chunk) Render() (err error) {
	if !modelLoaded {
		texture := rl.LoadTexture(os.Getenv("TEXTURE"))
		mesh := rl.GenMeshCube(1.0, 1.0, 1.0)
		model = rl.LoadModelFromMesh(mesh)
		model.Materials.GetMap(rl.MapDiffuse).Texture = texture

		modelLoaded = true
	}

	for _, cube := range chunk.cubes {
		rl.DrawModel(
			model,
			cube.
				Position().
				Add(chunk.PositionBase()).
				ToRaylib(),
			1.0,
			rl.White)
	}

	return
}
