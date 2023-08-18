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
var intervalOffsetY = 0
var intervalOffsetX = 0

func getAltitude(x int, y int, altMax float64) float64 {
	dx := XOFF(x)
	dy := YOFF(y)
	altitude := ContinentRadius - math.Sqrt(dx*dx+dy*dy)
	altitude += game.Altitudes.Noise2D(dx, (float64(y+intervalOffsetY)-game.Y)*Scale) * noiseMultiplier
	if altitude < 0 {
		altitude = 0.0
	}
	// Normalize altitude
	altitude = (altitude) / (altMax)

	return altitude
}

func colorTransition(altitude float64) color.Color {
	t := float64(time.Now().UnixNano()) / 1e9 * 0.5

	red := uint8((math.Sin(t)*127 + 128) * altitude)
	green := uint8((math.Sin(t+2*math.Pi/3)*127 + 128) * altitude)
	blue := uint8((math.Sin(t+4*math.Pi/3)*127 + 128) * altitude)
	return color.RGBA{red, green, blue, 255}
}

func drawLine(p1, p2 Position, clr color.Color) {
	dx := p2.x - p1.x
	dy := p2.y - p1.y

	steep := false
	if math.Abs(float64(dy)) > math.Abs(float64(dx)) {
		p1.x, p1.y = p1.y, p1.x
		p2.x, p2.y = p2.y, p2.x
		steep = true
	}

	if p1.x > p2.x {
		p1.x, p2.x = p2.x, p1.x
		p1.y, p2.y = p2.y, p1.y
	}

	dx = p2.x - p1.x
	dy = p2.y - p1.y
	yIncrement := 1
	if dy < 0 {
		yIncrement = -1
		dy = -dy
	}

	// Bresenham's line algorithm
	decision := 2*dy - dx
	y := p1.y
	for x := p1.x; x <= p2.x; x++ {
		if steep {
			game.Image.Set(y, x, clr)
		} else {
			game.Image.Set(x, y, clr)
		}
		if decision > 0 {
			y += yIncrement
			decision -= 2 * dx
		}
		decision += 2 * dy
	}
}

func abs(dy int) {
	panic("unimplemented")
}

func drawRadarLines() {
	game.Image.Fill(color.Black)
	const radarInterval int = 10

	// Calculate altitude normalization bounds
	altMax := (ContinentRadius + noiseMultiplier)

	for y := 0; y < ScreenHeight; y += radarInterval {
		for x := 0; x < ScreenWidth; x++ {
			altitude := getAltitude(x, y, altMax)

			var height int = y - int(altitude*float64(heightMultiplier2))
			game.Image.Set(x, height, colorTransition(altitude))
		}
	}
	intervalOffsetY++
}

func drawRadarDots() {
	game.Image.Fill(color.Black)
	const radarInterval int = 5
	var xdelta float64 = (float64(radarInterval) * float64(1.8))

	// Calculate altitude normalization bounds
	altMax := (ContinentRadius + noiseMultiplier)

	for y := 0; y < ScreenHeight; y += radarInterval {
		x := 0
		if (y/2)%radarInterval == 0 {
			x += int(xdelta) / 2
		}
		for x < ScreenWidth {
			altitude := getAltitude(x, y, altMax)

			var height int = y - int(altitude*float64(heightMultiplier2))
			clr := colorTransition(altitude)
			game.Image.Set(x, height, clr)
			game.Image.Set(x+1, height, clr)
			game.Image.Set(x, height+1, clr)
			game.Image.Set(x+1, height+1, clr)

			x += int(xdelta)
		}
	}
	intervalOffsetY++
}

func drawRadarTriangles() {
	game.Image.Fill(color.Black)
	const radarInterval int = 5
	var xdelta float64 = (float64(radarInterval) * float64(1.8))
	var prevPos1 *Position = nil

	// Calculate altitude normalization bounds
	altMax := (ContinentRadius + noiseMultiplier)

	for y := 0; y < ScreenHeight; y += radarInterval {
		var prevPos2 *Position = nil
		// var prevPos3 *Position = nil
		var tmpPrevPos2 *Position = nil
		// var tmpPrevPos3 *Position = nil
		x := 0
		if (y/2)%radarInterval == 0 {
			x += int(xdelta) / 2
		}
		for x < ScreenWidth {
			altitude := getAltitude(x, y, altMax)

			// var height int = y - int(altitude*float64(heightMultiplier2))
			// game.Image.Set(x, y, colorTransition(altitude))
			currentPos := Position{x, y}
			if prevPos1 != nil {
				drawLine(currentPos, *prevPos1, colorTransition(altitude))
			}
			if prevPos2 != nil {
				drawLine(currentPos, *prevPos2, colorTransition(altitude))
			} else {
				tmpPrevPos2 = &currentPos
			}
			// if prevPos3 != nil {
			// 	drawLine(currentPos, *prevPos3, colorTransition(altitude))
			// }
			prevPos1 = &currentPos

			x += int(xdelta)
		}
		prevPos2 = tmpPrevPos2
	}
	intervalOffsetY++
}
