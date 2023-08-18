package main

import (
	"image/color"
	"math"
	"time"
)

const (
	ContinentRadius   float64 = 4
	Scale             float64 = 0.03
	noiseMultiplier   float64 = 4.0
	heightMultiplier1 int     = 100
	heightMultiplier2 int     = 50
)

var BackgroundColor = color.RGBA{10, 10, 20, 255}
var intervalOffset = 0

func colorTransition(altitude float64) color.Color {
	t := float64(time.Now().UnixNano()) / 1e9 * 0.5

	red := uint8((math.Sin(t)*127 + 128) * altitude)
	green := uint8((math.Sin(t+2*math.Pi/3)*127 + 128) * altitude)
	blue := uint8((math.Sin(t+4*math.Pi/3)*127 + 128) * altitude)
	return color.RGBA{red, green, blue, 255}
}

func drawRadar1() {
	game.Terrain.Fill(color.Black)
	const radarInterval int = 15

	// Calculate altitude normalization bounds
	altMax := (ContinentRadius + noiseMultiplier)

	for y := intervalOffset; y < ScreenHeight; y += radarInterval {
		for x := 0; x < ScreenWidth; x++ {
			dx := XOFF(x)
			dy := YOFF(y)
			altitude := ContinentRadius - math.Sqrt(dx*dx+dy*dy)
			altitude += game.Altitudes.Noise2D(dx, dy) * noiseMultiplier
			if altitude < 0 {
				altitude = 0.0
			}
			// Normalize altitude
			altitude = (altitude) / (altMax)

			var clr uint8 = uint8(255 * altitude)
			var height int = y - int(altitude*float64(heightMultiplier1))
			game.Terrain.Set(x, height, color.RGBA{clr, clr, clr, 255})
		}
	}

	if intervalOffset >= radarInterval {
		intervalOffset = 0
	} else {
		intervalOffset++
	}
}

func drawRadar2() {
	game.Terrain.Fill(color.Black)
	const radarInterval int = 10

	// Calculate altitude normalization bounds
	altMax := (ContinentRadius + noiseMultiplier)

	for y := 0; y < ScreenHeight; y += radarInterval {
		for x := 0; x < ScreenWidth; x++ {
			dx := XOFF(x)
			dy := YOFF(y)
			altitude := ContinentRadius - math.Sqrt(dx*dx+dy*dy)
			altitude += game.Altitudes.Noise2D(dx, (float64(y+intervalOffset)-game.Y)*Scale) * noiseMultiplier
			if altitude < 0 {
				altitude = 0.0
			}
			// Normalize altitude
			altitude = (altitude) / (altMax)

			// var clr uint8 = uint8(255 * altitude)
			var height int = y - int(altitude*float64(heightMultiplier2))
			game.Terrain.Set(x, height, colorTransition(altitude))
		}
	}
	intervalOffset++
}
