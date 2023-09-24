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
	for _, cube := range mesh.Cubes() {
		for _, face := range BoxFaceListTest() {
			faceIndex := face

			neighbor := cube.Neighbor(face)

			if neighbor == nil {
				cube.facesEnabled[faceIndex] = true
				continue
			}

			if neighbor.Material() == MaterialAir {
				cube.facesEnabled[faceIndex] = true
				continue
			}
		}
	}

	mesh.computed = true

	return
}
