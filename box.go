package main

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
