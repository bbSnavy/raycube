package main

import "image/color"

type Cube struct {
	position Vector3

	color color.RGBA
}

func (cube *Cube) Position() Vector3 {
	return cube.position.Copy()
}

func (cube *Cube) Box() Box {
	return BoxNew(
		cube.Position(),
		Vector3NewUnit())
}

func (cube *Cube) Init() *Cube {
	return cube
}
