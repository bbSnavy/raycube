package main

import (
	"image/color"
)

type Cube struct {
	chunk *Chunk

	position Vector3

	color color.RGBA

	material Material
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
