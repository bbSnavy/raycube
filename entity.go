package main

type Entity interface {
	Position() Vector3
	Velocity() Vector3
	Acceleration() Vector3
	Rotation() Vector3
}

type EntityBase struct {
	position     Vector3
	velocity     Vector3
	acceleration Vector3
	rotation     Vector3
}

func (entity *EntityBase) Position() Vector3 {
	return entity.position
}

func (entity *EntityBase) Velocity() Vector3 {
	return entity.velocity
}

func (entity *EntityBase) Acceleration() Vector3 {
	return entity.acceleration
}

func (entity *EntityBase) Rotation() Vector3 {
	return entity.rotation
}

func (entity *EntityBase) AddPosition(vector Vector3) {
	entity.position = entity.position.Add(vector)
}

func (entity *EntityBase) AddRotation(vector Vector3) {
	entity.rotation = entity.rotation.
		Add(vector).
		Mod(Vector3New(360.0, 360.0, 360.0))
}

func (entity *EntityBase) Init() *EntityBase {
	return entity
}
