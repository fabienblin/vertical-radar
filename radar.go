package main

import (
	"image/color"
	"math"
)

const (
	ContinentRadius   float64 = 4
	Scale             float64 = 0.03
	noiseMultiplier   float64 = 4.0
	heightMultiplier1 int     = 100
	heightMultiplier2 int     = 50
)

var intervalOffset = 0

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

			var clr uint8 = uint8(255 * altitude)
			var height int = y - int(altitude*float64(heightMultiplier2))
			game.Terrain.Set(x, height, color.RGBA{clr, clr / 2, clr, 255})
		}
	}
	intervalOffset++
}
