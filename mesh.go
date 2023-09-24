package main

/*
#include "raycube.h"
*/
import "C"
import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Mesh struct {
	position Vector3

	cubes []*Cube

	computed bool
	result   rl.Mesh
}

func (mesh *Mesh) Cubes() []*Cube {
	return mesh.cubes
}

func (mesh *Mesh) AddCube(cube *Cube) {
	mesh.cubes = append(mesh.cubes, cube)
}

func (mesh *Mesh) HasCube(cube *Cube) bool {
	for _, v := range mesh.cubes {
		if v == cube {
			return true
		}
	}

	return false
}

func (mesh *Mesh) Init() *Mesh {
	mesh.computed = false

	return mesh
}

func (mesh *Mesh) Compute() (result rl.Mesh) {
	data := make([]*Cube, 0)

	for _, cube := range mesh.Cubes() {
		flag := false

		for _, face := range BoxFaceList() {
			neighbor := cube.Neighbor(face)
			if neighbor == nil {
				flag = true
				break
			}

			if neighbor.Material() == MaterialAir {
				flag = true
				break
			}
		}

		if !flag {
			continue
		}

		data = append(data, cube)
	}

	mesh.cubes = data
	mesh.computed = true

	return
}
