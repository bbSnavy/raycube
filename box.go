package main

import rl "github.com/gen2brain/raylib-go/raylib"

type BoxFace int

const (
	NoFace BoxFace = iota
	TopFace
	BottomFace
	FrontFace
	BackFace
	LeftFace
	RightFace
)

func BoxFaceList() []BoxFace {
	return []BoxFace{
		TopFace,
		BottomFace,
		FrontFace,
		BackFace,
		LeftFace,
		RightFace,
	}
}

func (face BoxFace) Opposite() BoxFace {
	switch face {
	case TopFace:
		return BottomFace
	case BottomFace:
		return TopFace
	case FrontFace:
		return BackFace
	case BackFace:
		return FrontFace
	case LeftFace:
		return RightFace
	case RightFace:
		return LeftFace
	default:
		return NoFace
	}
}

type Box struct {
	position Vector3
	size     Vector3
}

func BoxNew(position, size Vector3) Box {
	return Box{
		position: position,
		size:     size,
	}
}

func (box Box) ToRaylib() rl.BoundingBox {
	return rl.NewBoundingBox(
		box.Min().ToRaylib(),
		box.Max().ToRaylib())
}

func (box Box) PosX() float32 {
	return box.position.X
}

func (box Box) PosY() float32 {
	return box.position.Y
}

func (box Box) PosZ() float32 {
	return box.position.Z
}

func (box Box) SizeX() float32 {
	return box.size.X
}

func (box Box) SizeY() float32 {
	return box.size.Y
}

func (box Box) SizeZ() float32 {
	return box.size.Z
}

func (box Box) Min() Vector3 {
	return box.position
}

func (box Box) Max() Vector3 {
	return box.position.Add(box.size)
}

func (box Box) VectorBoxIntersect(vector Vector3) bool {
	if vector.X >= box.PosX() && vector.X <= (box.PosX()+box.SizeX()) &&
		vector.Y >= box.PosY() && vector.Y <= (box.PosY()+box.SizeY()) &&
		vector.Z >= box.PosZ() && vector.Z <= (box.PosZ()+box.SizeZ()) {
		return true
	}

	return false
}
