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

func (cube *Cube) PositionWorld() Vector3 {
	return cube.Chunk().PositionBase().Add(cube.Position())
}

func (cube *Cube) Box() Box {
	return BoxNew(
		cube.PositionWorld(),
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
		delta = Vector3New(-1.0, 0.0, 0.0)
		break

	case RightFace:
		delta = Vector3New(+1.0, 0.0, 0.0)
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

	//switch cubeMaterial {
	//case MaterialStone:
	//	{
	//		rl.Color4ub(
	//			rl.Gray.R,
	//			rl.Gray.G,
	//			rl.Gray.B,
	//			rl.Gray.A)
	//
	//		break
	//	}
	//
	//case MaterialGrass:
	//	{
	//		rl.Color4ub(
	//			rl.DarkGreen.R,
	//			rl.DarkGreen.G,
	//			rl.DarkGreen.B,
	//			rl.DarkGreen.A)
	//
	//		break
	//	}
	//}

	for index, value := range cube.facesEnabled {
		face := BoxFace(index)

		if !value {
			continue
		}

		switch face {
		case FrontFace:
			{
				rl.Color4ub(128, 128, 128, 255)

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
				rl.Color4ub(128, 128, 128, 255)

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
				rl.Color4ub(255, 255, 255, 255)

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
				rl.Color4ub(64, 64, 64, 255)

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
				rl.Color4ub(128, 128, 128, 255)

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
				rl.Color4ub(128, 128, 128, 255)

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
