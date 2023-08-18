package main

import (
	_ "image/png"

	"github.com/aquilax/go-perlin"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  int     = 400
	ScreenHeight int     = 400
	Octaves      float64 = 6
	Persistence  float64 = 8
	Lacunarity   int32   = 4
	Seed         int64   = 12345
)

var (
	runnerImage *ebiten.Image
)

type GameInstance struct {
	Terrain   *ebiten.Image
	X, Y      float64
	Altitudes *perlin.Perlin
}

func initGame() *GameInstance {
	g := &GameInstance{
		Terrain:   ebiten.NewImage(ScreenWidth, ScreenHeight),
		X:         float64(ScreenWidth) / 2,
		Y:         float64(ScreenHeight) / 2,
		Altitudes: perlin.NewPerlin(Octaves, Persistence, Lacunarity, Seed),
	}

	// Set the game window properties
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetFullscreen(true)

	ebiten.SetWindowTitle("Vertical Radar")
	ebiten.SetTPS(25)

	return g
}

func (g *GameInstance) Update() error {
	return nil
}

func (g *GameInstance) Draw(screen *ebiten.Image) {
	drawRadar3()
	screen.DrawImage(g.Terrain, nil)
}

func (g *GameInstance) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenWidth, ScreenHeight
}
