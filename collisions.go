package main

func CheckCollisionBoxToBox(box Box, other Box) bool {
	return box.position.X < other.position.X+other.size.X &&
		box.position.X+box.size.X > other.position.X &&
		box.position.Y < other.position.Y+other.size.Y &&
		box.position.Y+box.size.Y > other.position.Y &&
		box.position.Z < other.position.Z+other.size.Z &&
		box.position.Z+box.size.Z > other.position.Z
}

func (box Box) FacesTouching(other Box) []string {
	var touchingFaces []string

	if box.position.X < other.position.X+other.size.X && box.position.X+box.size.X > other.position.X {
		if box.position.Y < other.position.Y+other.size.Y && box.position.Y+box.size.Y > other.position.Y {
			if box.position.Z+box.size.Z == other.position.Z {
				touchingFaces = append(touchingFaces, "Top")
			}
			if box.position.Z == other.position.Z+other.size.Z {
				touchingFaces = append(touchingFaces, "Bottom")
			}
		}
		if box.position.Z < other.position.Z+other.size.Z && box.position.Z+box.size.Z > other.position.Z {
			if box.position.Y+box.size.Y == other.position.Y {
				touchingFaces = append(touchingFaces, "Front")
			}
			if box.position.Y == other.position.Y+other.size.Y {
				touchingFaces = append(touchingFaces, "Back")
			}
		}
	}
	if box.position.Y < other.position.Y+other.size.Y && box.position.Y+box.size.Y > other.position.Y {
		if box.position.Z < other.position.Z+other.size.Z && box.position.Z+box.size.Z > other.position.Z {
			if box.position.X+box.size.X == other.position.X {
				touchingFaces = append(touchingFaces, "Right")
			}
			if box.position.X == other.position.X+other.size.X {
				touchingFaces = append(touchingFaces, "Left")
			}
		}
	}

	return touchingFaces
}

func (box Box) CollidesWith(other Box) (bool, BoxFace) {
	if box.position.X < other.position.X+other.size.X &&
		box.position.X+box.size.X > other.position.X &&
		box.position.Y < other.position.Y+other.size.Y &&
		box.position.Y+box.size.Y > other.position.Y &&
		box.position.Z < other.position.Z+other.size.Z &&
		box.position.Z+box.size.Z > other.position.Z {

		// Calculate the overlap in each dimension
		overlapX := min(box.position.X+box.size.X, other.position.X+other.size.X) - max(box.position.X, other.position.X)
		overlapY := min(box.position.Y+box.size.Y, other.position.Y+other.size.Y) - max(box.position.Y, other.position.Y)
		overlapZ := min(box.position.Z+box.size.Z, other.position.Z+other.size.Z) - max(box.position.Z, other.position.Z)

		// Determine which face is touching based on the smallest overlap
		if overlapX < overlapY && overlapX < overlapZ {
			if box.position.X < other.position.X {
				return true, LeftFace
			} else {
				return true, RightFace
			}
		} else if overlapY < overlapX && overlapY < overlapZ {
			if box.position.Y < other.position.Y {
				return true, BottomFace
			} else {
				return true, TopFace
			}
		} else {
			if box.position.Z < other.position.Z {
				return true, BackFace
			} else {
				return true, FrontFace
			}
		}
	}

	return false, NoFace
}

func min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
