package main

type Cube struct {
	position Vector3
}

func (cube *Cube) Position() Vector3 {
	return cube.position.Copy()
}

func (cube *Cube) Init() *Cube {
	return cube
}
