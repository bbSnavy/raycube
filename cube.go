package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

import "C"

type Cube struct {
	chunk *Chunk

	position Vector3

	color color.RGBA

	material Material

	facesEnabled [7]bool
}

func (cube *Cube) Chunk() *Chunk {
	return cube.chunk
}

func (cube *Cube) Position() Vector3 {
	return cube.position.Copy()
}

func (cube *Cube) Box() Box {
	return BoxNew(
		cube.Position(),
		Vector3NewUnit())
}

func (cube *Cube) Material() Material {
	return cube.material
}

func (cube *Cube) Init() *Cube {
	for index := range cube.facesEnabled {
		cube.facesEnabled[index] = false
	}

	return cube
}

func (cube *Cube) Neighbor(face BoxFace) (result *Cube) {
	var (
		chunk *Chunk
		delta Vector3
	)

	result = nil

	chunk = cube.Chunk()
	delta = Vector3NewZero()

	switch face {
	case TopFace:
		delta = Vector3New(0.0, +1.0, 0.0)
		break

	case BottomFace:
		delta = Vector3New(0.0, -1.0, 0.0)
		break

	case FrontFace:
		delta = Vector3New(0.0, 0.0, +1.0)
		break

	case BackFace:
		delta = Vector3New(0.0, 0.0, -1.0)
		break

	case LeftFace:
		delta = Vector3New(+1.0, 0.0, 0.0)
		break

	case RightFace:
		delta = Vector3New(-1.0, 0.0, 0.0)
		break

	default:
		return
	}

	result = chunk.CubeAtPosition(cube.Position().Add(delta))
	if result == nil {
		return
	}

	return
}

func (cube *Cube) Render() (err error) {
	cubeMaterial := cube.Material()
	cubePositionWorld := cube.
		Position().
		Add(cube.Chunk().PositionBase()).
		ToRaylib()

	x, y, z := cubePositionWorld.X, cubePositionWorld.Y, cubePositionWorld.Z
	width, height, length := float32(1.0), float32(1.0), float32(1.0)

	rl.Begin(rl.RL_TRIANGLES)

	rl.Color4ub(
		rl.Red.R,
		rl.Red.G,
		rl.Red.B,
		rl.Red.A)

	for index, value := range cube.facesEnabled {
		face := BoxFace(index)

		if !value {
			continue
		}

		switch face {
		case FrontFace:
			{
				// Front face
				rl.Vertex3f(x-width/2, y-height/2, z+length/2) // Bottom Left
				rl.Vertex3f(x+width/2, y-height/2, z+length/2) // Bottom Right
				rl.Vertex3f(x-width/2, y+height/2, z+length/2) // Top Left

				rl.Vertex3f(x+width/2, y+height/2, z+length/2) // Top Right
				rl.Vertex3f(x-width/2, y+height/2, z+length/2) // Top Left
				rl.Vertex3f(x+width/2, y-height/2, z+length/2) // Bottom Right

				break
			}

		case BackFace:
			{
				// Back face
				rl.Vertex3f(x-width/2, y-height/2, z-length/2) // Bottom Left
				rl.Vertex3f(x-width/2, y+height/2, z-length/2) // Top Left
				rl.Vertex3f(x+width/2, y-height/2, z-length/2) // Bottom Right

				rl.Vertex3f(x+width/2, y+height/2, z-length/2) // Top Right
				rl.Vertex3f(x+width/2, y-height/2, z-length/2) // Bottom Right
				rl.Vertex3f(x-width/2, y+height/2, z-length/2) // Top Left

				break
			}

		case TopFace:
			{
				// Top face
				rl.Vertex3f(x-width/2, y+height/2, z-length/2) // Top Left
				rl.Vertex3f(x-width/2, y+height/2, z+length/2) // Bottom Left
				rl.Vertex3f(x+width/2, y+height/2, z+length/2) // Bottom Right

				rl.Vertex3f(x+width/2, y+height/2, z-length/2) // Top Right
				rl.Vertex3f(x-width/2, y+height/2, z-length/2) // Top Left
				rl.Vertex3f(x+width/2, y+height/2, z+length/2) // Bottom Right

				break
			}

		case BottomFace:
			{
				// Bottom face
				rl.Vertex3f(x-width/2, y-height/2, z-length/2) // Top Left
				rl.Vertex3f(x+width/2, y-height/2, z+length/2) // Bottom Right
				rl.Vertex3f(x-width/2, y-height/2, z+length/2) // Bottom Left

				rl.Vertex3f(x+width/2, y-height/2, z-length/2) // Top Right
				rl.Vertex3f(x+width/2, y-height/2, z+length/2) // Bottom Right
				rl.Vertex3f(x-width/2, y-height/2, z-length/2) // Top Left

				break
			}

		case RightFace:
			{
				// Right face
				rl.Vertex3f(x+width/2, y-height/2, z-length/2) // Bottom Right
				rl.Vertex3f(x+width/2, y+height/2, z-length/2) // Top Right
				rl.Vertex3f(x+width/2, y+height/2, z+length/2) // Top Left

				rl.Vertex3f(x+width/2, y-height/2, z+length/2) // Bottom Left
				rl.Vertex3f(x+width/2, y-height/2, z-length/2) // Bottom Right
				rl.Vertex3f(x+width/2, y+height/2, z+length/2) // Top Left

				break
			}

		case LeftFace:
			{
				// Left face
				rl.Vertex3f(x-width/2, y-height/2, z-length/2) // Bottom Right
				rl.Vertex3f(x-width/2, y+height/2, z+length/2) // Top Left
				rl.Vertex3f(x-width/2, y+height/2, z-length/2) // Top Right

				rl.Vertex3f(x-width/2, y-height/2, z+length/2) // Bottom Left
				rl.Vertex3f(x-width/2, y+height/2, z+length/2) // Top Left
				rl.Vertex3f(x-width/2, y-height/2, z-length/2) // Bottom Right

				break
			}
		}

	}

	rl.End()

	//rl.DrawCubeWires(
	//	cubePositionWorld,
	//	1.0,
	//	1.0,
	//	1.0,
	//	rl.Red)

	_ = cubeMaterial

	return
}
