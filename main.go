package main

import "log"

func main() {
	game := (&Game{}).Init()
	log.Println(game.Start())
}
