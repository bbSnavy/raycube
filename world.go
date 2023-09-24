package main

type World struct {
	game *Game

	player *Player

	entities []Entity
	chunks   []*Chunk
}

func (world *World) Player() *Player {
	return world.player
}

func (world *World) Init() *World {
	if world.game == nil {
		panic("game is nil")
	}

	world.player = (&Player{}).Init()

	world.entities = make([]Entity, 0)
	world.chunks = make([]*Chunk, 0)

	world.entities = append(world.entities, world.player)

	for x := 0; x < 2; x++ {
		for z := 0; z < 2; z++ {
			for y := 0; y < 1; y++ {
				chunk := (&Chunk{
					position: Vector3New(float32(x), float32(y), float32(z)),

					world: world,
				}).Init()

				world.chunks = append(world.chunks, chunk)
			}
		}
	}

	return world
}

func (world *World) Tick() (err error) {
	for _, entity := range world.entities {
		entity.Tick(world)
	}

	return
}

func (world *World) Render() (err error) {
	for _, chunk := range world.chunks {
		err = chunk.Render()
		if err != nil {
			return
		}
	}

	return
}
