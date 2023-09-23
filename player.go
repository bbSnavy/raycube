package main

import "github.com/google/uuid"

type Player struct {
	id uuid.UUID

	target Vector3

	EntityBase
}

func (player *Player) Target() Vector3 {
	return player.target
}

func (player *Player) Init() *Player {
	player.id = uuid.New()

	player.EntityBase.Init()

	return player
}
