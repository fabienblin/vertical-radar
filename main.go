package main

import "github.com/hajimehoshi/ebiten/v2"

var game *GameInstance

func init() {
	game = initGame()
}

func main() {
	// GenerateTerrain()
	// Run the game loop
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

type Position struct {
	x, y int
	altitude float64
}

func XOFF(x int) float64 {
	return (float64(x) - game.X) * Scale
}

func YOFF(y int) float64 {
	return (float64(y) - game.Y) * Scale
}
