package main

var (
	GAME *Game
)

func init() {
	GAME = (&Game{}).Init()
}

func main() {
	var (
		err error
	)

	err = GAME.Start()
	if err != nil {
		panic(err)
	}
}
