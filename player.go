package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
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

	player.SetPosition(Vector3New(2.0, 80.0, 2.0))

	return player
}

func (player *Player) Tick(world *World) {
	player.TickMovement(world)
	player.TickRay(world)
}

func (player *Player) TickRay(world *World) {
	var (
		ray Vector3
	)

	ray = player.
		Target().
		Mul(Vector3New(10.0, 10.0, 10.0))

	rayRL := rl.NewRay(
		player.Position().
			Add(player.CameraOffset()).
			Add(Vector3New(0.50, 0.00, 0.50)).
			ToRaylib(),
		ray.
			ToRaylib())

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		for _, chunk := range world.chunks {
			flag := false

			for _, cube := range chunk.Cubes() {
				if cube.Material() == MaterialAir {
					continue
				}

				if !rl.GetRayCollisionBox(rayRL, cube.Box().ToRaylib()).Hit {
					continue
				}

				cube.material = MaterialAir
				flag = true
			}

			if flag {
				go chunk.MeshGenerate()
			}
		}
	}

	return
}

func (player *Player) TickMovement(world *World) {
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

	acceleration = acceleration.Mul(Vector3New(0.90, 0.90, 0.90))

	player.SetAcceleration(acceleration)
	player.SetVelocity(velocity)

	playerBox := player.Box()
	playerBox.position = position.Add(velocity)

	for _, chunk := range world.chunks {
		for _, cube := range chunk.Cubes() {
			if cube.Material() == MaterialAir {
				continue
			}

			if cube.PositionWorld().Distance(player.Position()) > 8.0 {
				continue
			}

			a, b := cube.Box().CollidesWith(playerBox)
			if a {
				switch b {
				case NoFace:
					break

				case TopFace:
					velocity.Y = 0.0

				case BottomFace:
					velocity.Y = 0.0

				case LeftFace:
					velocity.X = 0.0

				case RightFace:
					velocity.X = 0.0

				case FrontFace:
					velocity.Z = 0.0

				case BackFace:
					velocity.Z = 0.0

				default:
					{
						log.Println("unhandled face:", b)
						break
					}
				}
			}

		}
	}

	position = position.Add(velocity)

	player.SetPosition(position)
}
