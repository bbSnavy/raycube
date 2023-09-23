package main

import (
	"github.com/google/uuid"
	"log"
)

type Player struct {
	id uuid.UUID

	target Vector3

	EntityBase
}

func (player *Player) CameraOffset() Vector3 {
	return Vector3New(0.00, 1.62, 0.00)
}

func (player *Player) Target() Vector3 {
	return Vector3UnitVector(player.Rotation())
}

func (player *Player) Box() Box {
	return BoxNew(
		player.position,
		Vector3New(1.0, 2.0, 1.0))
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

func (player *Player) Tick(world *World) {
	var (
		acceleration Vector3
		velocity     Vector3
		position     Vector3
	)

	acceleration = player.Acceleration()
	velocity = player.Velocity()
	position = player.Position()

	velocity = velocity.Mul(Vector3New(0.90, 0.90, 0.90))
	velocity = velocity.Add(acceleration.Mul(Vector3New(0.10, 0.10, 0.10)))
	velocity = velocity.Add(Vector3New(0.00, -0.005, 0.00))

	acceleration = acceleration.Mul(Vector3New(0.90, 0.90, 0.90))

	player.SetAcceleration(acceleration)
	player.SetVelocity(velocity)

	for _, cube := range world.cubes {
		if CheckCollisionBoxToBox(cube.Box(), player.Box()) {
			log.Println("collision!", velocity)

			if velocity.Y < 0.0 {
				velocity.Y = 0.0
			}
			break
		}
	}

	position = position.Add(velocity)

	player.SetPosition(position)
}