package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Vector3 struct {
	X float32
	Y float32
	Z float32
}

func Vector3New(x, y, z float32) Vector3 {
	return Vector3{
		x, y, z,
	}
}

func (vector Vector3) Copy() Vector3 {
	return Vector3{
		X: vector.X,
		Y: vector.Y,
		Z: vector.Z,
	}
}

func (vector Vector3) ToRaylib() rl.Vector3 {
	return rl.Vector3{
		X: vector.X,
		Y: vector.Y,
		Z: vector.Z,
	}
}
