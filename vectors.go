package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
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

func (vector Vector3) Mul(v Vector3) Vector3 {
	return Vector3New(
		vector.X*v.X,
		vector.Y*v.Y,
		vector.Z*v.Z)
}

func (vector Vector3) Div(v Vector3) Vector3 {
	return Vector3New(
		DivZero(vector.X, v.X),
		DivZero(vector.Y, v.Y),
		DivZero(vector.Z, v.Z))
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

func (vector Vector3) Inv() Vector3 {
	return Vector3New(
		-vector.X,
		-vector.Y,
		-vector.Z)
}

func (vector Vector3) Length() float64 {
	return math.Sqrt(float64(
		vector.X*vector.X + vector.Y*vector.Y + vector.Z*vector.Z))
}

func (vector Vector3) Normalize() Vector3 {
	length := float32(vector.Length())

	return Vector3New(
		DivZero(vector.X, length),
		DivZero(vector.Y, length),
		DivZero(vector.Z, length))
}

func (vector Vector3) DistanceSquare(v Vector3) float64 {
	x := v.X - vector.X
	y := v.Y - vector.Y
	z := v.Z - vector.Z

	return float64(x*x + y*y + z*z)
}

func (vector Vector3) Distance(v Vector3) float64 {
	return math.Sqrt(vector.DistanceSquare(v))
}

func DegreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180.0)
}

func Vector3UnitVector(rotation Vector3) (result Vector3) {
	a := DegreesToRadians(float64(rotation.Y)) // yaw
	b := DegreesToRadians(float64(rotation.Z)) // pitch

	x := math.Cos(a) * math.Cos(b)
	y := math.Sin(b)
	z := math.Sin(a) * math.Cos(b)

	return Vector3New(
		float32(x),
		float32(y),
		float32(z))
}
