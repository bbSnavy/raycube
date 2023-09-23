package main

func CheckCollisionBoxToBox(box Box, other Box) bool {
	return box.position.X < other.position.X+other.size.X &&
		box.position.X+box.size.X > other.position.X &&
		box.position.Y < other.position.Y+other.size.Y &&
		box.position.Y+box.size.Y > other.position.Y &&
		box.position.Z < other.position.Z+other.size.Z &&
		box.position.Z+box.size.Z > other.position.Z
}
