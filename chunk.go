package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"os"
	"time"
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
					material: []Material{MaterialAir, MaterialStone}[1],
				}).Init()

				index++
			}
		}
	}

	chunk.meshComputed = false

	return chunk
}

func (chunk *Chunk) MeshGenerate() (result []*Mesh, err error) {
	var (
		mesh  *Mesh
		start *Cube
	)

	result = make([]*Mesh, 0)

	start = nil
	for _, cube := range chunk.Cubes() {
		if cube.Material() == MaterialAir {
			continue
		}

		start = cube
		break
	}

	if start == nil {
		return
	}

	mesh = (&Mesh{
		position: chunk.PositionBase(),
	}).Init()

	var (
		fn func(cube *Cube, mesh *Mesh)
	)

	timeStart := time.Now().UnixMicro()

	fn = func(cube *Cube, mesh *Mesh) {
		if mesh.HasCube(cube) {
			return
		}

		//cube.material = MaterialGrass

		mesh.AddCube(cube)

		for _, face := range BoxFaceList() {
			//time.Sleep(50 * time.Millisecond)

			neighbor := cube.Neighbor(face)
			if neighbor == nil {
				continue
			}

			if neighbor.Material() == MaterialAir {
				continue
			}

			fn(neighbor, mesh)
		}
	}

	fn(start, mesh)

	timeEnd := time.Now().UnixMicro()

	log.Println(float64(timeEnd-timeStart) / 1000.0)

	log.Println("mesh length->", len(mesh.Cubes()))
	log.Println("mesh computed->", mesh.Compute())

	result = append(result, mesh)

	return
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
		chunk.meshSnapshot, err = chunk.MeshGenerate()
		if err != nil {
			return
		}

		chunk.meshComputed = true
	}

	for _, mesh := range chunk.meshSnapshot {
		if !mesh.computed {
			panic("mesh has not been computed")
		}

		for _, cube := range mesh.Cubes() {
			err = chunk.RenderCube(cube)
			if err != nil {
				return
			}
		}

		//meshResult := mesh.result
		//
		//meshModel := rl.LoadModelFromMesh(meshResult)
		//
		//meshModel.Materials.GetMap(rl.MapDiffuse).Texture = modelTexture
		//
		//rl.DrawModel(
		//	meshModel,
		//	mesh.position.
		//		Add(Vector3New(0.0, 16.0, 0.0)).
		//		ToRaylib(),
		//	1.0,
		//	rl.White)
	}

	//for _, cube := range chunk.cubes {
	//
	//}

	return
}

func (chunk *Chunk) RenderCube(cube *Cube) (err error) {
	cubeMaterial := cube.Material()

	switch cubeMaterial {
	case MaterialAir:
		break

	case MaterialStone:
		rl.DrawModel(
			model,
			cube.
				Position().
				Add(chunk.PositionBase()).
				ToRaylib(),
			1.0,
			rl.White)

		break

	case MaterialGrass:
		rl.DrawCube(
			cube.
				Position().
				Add(chunk.PositionBase()).
				ToRaylib(),
			1.0,
			1.0,
			1.0,
			rl.DarkGreen)

		break
	}

	return
}
