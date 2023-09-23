package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"log"
	"math"
)

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

func Vector3NewZero() Vector3 {
	return Vector3New(0.0, 0.0, 0.0)
}

func Vector3NewUnit() Vector3 {
	return Vector3New(1.0, 1.0, 1.0)
}

func Vector3FromRaylib(v rl.Vector3) Vector3 {
	return Vector3New(
		v.X,
		v.Y,
		v.Z)
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

func (vector Vector3) Add(v Vector3) Vector3 {
	return Vector3New(
		vector.X+v.X,
		vector.Y+v.Y,
		vector.Z+v.Z)
}

func (vector Vector3) Sub(v Vector3) Vector3 {
	return Vector3New(
		vector.X-v.X,
		vector.Y-v.Y,
		vector.Z-v.Z)
}

func (vector Vector3) Mod(v Vector3) Vector3 {
	x := vector.X
	y := vector.Y
	z := vector.Z

	if x >= v.X {
		x = 0.0
	} else if x < 0.0 {
		x = v.X
	}

	if y >= v.Y {
		y = 0.0
	} else if y < 0.0 {
		y = v.Y
	}

	if z >= v.Z {
		z = 0.0
	} else if z < 0.0 {
		z = v.Z
	}

	return Vector3New(x, y, z)
}

func (vector Vector3) Length() float64 {
	return math.Sqrt(float64(
		vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z))
}

func (vector Vector3) Normalize() Vector3 {
	length := float32(vector.Length())

	return Vector3New(
		vector.X/length,
		vector.Y/length,
		vector.Z/length)
}

func DegreesToRadians(degrees float64) float64 {
	//return rl.Deg2rad * degrees

	return degrees * (math.Pi / 180.0)
}

func Vector3EulerToMatrix(rotation Vector3) (result Vector3) {
	phi := DegreesToRadians(float64(rotation.Y))
	theta := DegreesToRadians(float64(rotation.Z))

	log.Println(rotation, []float64{phi, theta})

	x := math.Sin(phi)
	y := math.Cos(theta)
	z := math.Sin(theta) * math.Cos(phi)

	//x := math.Cos(float64(phi)) + math.Cos(float64(theta))
	//y := math.Cos(float64(phi)) + math.Sin(float64(theta))
	//z := math.Cos(float64(phi))

	return Vector3New(
		float32(x),
		float32(y),
		float32(z))

	///*
	//	x = cos(yaw)*cos(pitch)
	//	y = sin(yaw)*cos(pitch)
	//	z = sin(pitch)
	//*/
	//
	//a := rl.Deg2rad * rotation.X // yaw
	//b := rl.Deg2rad * rotation.Y // pitch
	//c := rl.Deg2rad * rotation.Z // roll
	//
	//va := math.Cos(float64(a)) * math.Cos(float64(b))
	//vb := math.Sin(float64(a)) * math.Cos(float64(b))
	//vc := math.Sin(float64(b))
	//
	//_ = c
	//
	//return Vector3New(
	//	float32(va),
	//	float32(vb),
	//	float32(vc))
}
