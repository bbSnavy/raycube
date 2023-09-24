package main

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

func BoxFaceListTest() []BoxFace {
	return []BoxFace{
		FrontFace,
		BackFace,
		TopFace,
		BottomFace,
		RightFace,
		LeftFace,
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
