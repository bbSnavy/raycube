package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
)

type Cube struct {
	position Vector3

	color color.RGBA

	model  rl.Model
	loaded bool
}

func (cube *Cube) Position() Vector3 {
	return cube.position.Copy()
}

func (cube *Cube) Box() Box {
	return BoxNew(
		cube.Position(),
		Vector3NewUnit())
}

func (cube *Cube) Model() rl.Model {
	return cube.model
}

func (cube *Cube) Init() *Cube {
	cube.loaded = false

	return cube
}
