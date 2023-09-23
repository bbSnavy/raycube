package main

import (
	"github.com/google/uuid"
)

type Player struct {
	id uuid.UUID

	target Vector3

	EntityBase
}

func (player *Player) Target() Vector3 {
	return Vector3UnitVector(player.Rotation())
}

func (player *Player) TargetRelative() Vector3 {
	return player.
		Target().
		Normalize().
		Add(player.Position())
}

func (player *Player) Init() *Player {
	player.id = uuid.New()

	player.EntityBase.Init()

	player.SetPosition(Vector3New(2.0, 8.0, 2.0))

	return player
}

func (player *Player) Tick() {
	var (
		acceleration Vector3
		velocity     Vector3
	)

	acceleration = player.Acceleration()
	velocity = player.Velocity()

	velocity = velocity.Mul(Vector3New(0.90, 0.90, 0.90))
	velocity = velocity.Add(acceleration.Mul(Vector3New(0.10, 0.10, 0.10)))

	acceleration = acceleration.Mul(Vector3New(0.90, 0.90, 0.90))

	player.SetAcceleration(acceleration)
	player.SetVelocity(velocity)
	player.AddPosition(velocity)
}
