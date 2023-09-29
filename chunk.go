package main

import (
	"os"
	"time"

	"github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ChunkSize = 16
)

type Chunk struct {
	world *World

	position Vector3

	cubes [4096]*Cube

	meshSnapshot []*Mesh
	meshComputed bool
}

func (chunk *Chunk) CubeAtIndex(index int) *Cube {
	if index < 0 || index >= len(chunk.cubes) {
		return nil
	}

	return chunk.cubes[index]
}

func (chunk *Chunk) CubeAtPosition(vector Vector3) *Cube {
	if vector.X < 0 || vector.Y < 0 || vector.Z < 0 {
		return nil
	}

	if vector.X >= ChunkSize || vector.Y >= ChunkSize || vector.Z >= ChunkSize {
		return nil
	}

	for _, cube := range chunk.cubes {
		if cube.Position().Equals(vector) {
			return cube
		}
	}

	return nil
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

var (
	noise *perlin.Perlin
)

func init() {
	noise = perlin.NewPerlin(2, 2, 3, time.Now().UnixMilli())
}

func (chunk *Chunk) Init() *Chunk {
	if chunk.world == nil {
		panic("chunk world is nil")
	}

	index := 0
	for x := 0; x < 16; x++ {
		for y := 0; y < 16; y++ {
			for z := 0; z < 16; z++ {
				materials := []Material{
					MaterialAir,
					MaterialStone,
				}

				cube := (&Cube{
					chunk:    chunk,
					position: Vector3New(float32(x), float32(y), float32(z)),
				}).Init()

				posAbs := cube.PositionWorld()
				posVal := noise.Noise3D(
					float64(posAbs.X)/10,
					float64(posAbs.Y)/10,
					float64(posAbs.Z)/10) * 10.0

				_ = materials
				_ = posVal
				cube.material = MaterialStone

				chunk.cubes[index] = cube

				index++
			}
		}
	}

	chunk.meshComputed = false

	return chunk
}

func (chunk *Chunk) MeshGenerate() {
	var (
		mesh *Mesh
	)

	result := make([]*Mesh, 0)

	mesh = (&Mesh{
		position: chunk.PositionBase(),
	}).Init()

	mesh.cubes = chunk.Cubes()[:]
	mesh.Compute()

	result = append(result, mesh)

	chunk.meshComputed = true
	chunk.meshSnapshot = result
}

var (
	model        rl.Model
	modelTexture rl.Texture2D
	modelLoaded  = false
)

func (chunk *Chunk) Render() (err error) {
	if !modelLoaded {
		modelTexture = rl.LoadTexture(os.Getenv("TEXTURE"))

		rl.SetTextureWrap(modelTexture, rl.RL_TEXTURE_WRAP_REPEAT)

		mesh := rl.GenMeshCube(1.0, 1.0, 1.0)
		model = rl.LoadModelFromMesh(mesh)
		model.Materials.GetMap(rl.MapDiffuse).Texture = modelTexture

		modelLoaded = true
	}

	if !chunk.meshComputed {
		go chunk.MeshGenerate()

		return
	}

	for _, mesh := range chunk.meshSnapshot {
		if !mesh.computed {
			continue
		}

		for _, cube := range mesh.Cubes() {
			err = chunk.RenderCube(cube)
			if err != nil {
				return
			}
		}
	}

	return
}

func (chunk *Chunk) RenderCube(cube *Cube) (err error) {
	rl.Begin(rl.RL_TRIANGLES)

	err = cube.Render()
	if err != nil {
		return
	}

	rl.End()

	return
}
