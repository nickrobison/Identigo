// Package identicon creates a simple, two-color identicon from a provided byteString
package identicon

import (
	"crypto/md5"
	"encoding/hex"
	"image"
	"image/color"
	"image/draw"
	"math"
	"strings"
)

type colorValues struct {
	red uint8
	green uint8
	blue uint8
	alpha uint8
}

func buildArray(hashString string) []bool {
	fill := make([]bool, 32, 64)

	for i, digit := range hashString {

		if digit > 58 {
			fill[i] = true
		} else {
			fill[i] = false
		}
	}

	return fill
}

// Create function builds the identicon using a standard MD5 hash
// The following foreground and background colors are supported:
//
// black
// navy
// blue
// green
// teal
// lime
// aqua
// maroon
// purple
// olive
// gray
// silver
// red
// fuchsia
// yellow
// white
//
// If an invalid color is specified, the function defaults to a gray icon with green coloring
func Create(bytestring []byte, tilesPerSide int, multiplier int, backcolor string, forecolor string) image.Image {

	colors := map[string]colorValues{
	"black": {0,0,0,255},
	"navy": {0,0,128,255},
	"blue": {0,0,128,255},
	"green": {0,255,0,255},
	"teal": {0,128,128,255},
	"lime": {0,255,0,255},
	"aqua": {0,255,255,255},
	"maroon": {128,0,0,255},
	"purple": {128,0,128,255},
	"olive": {128,128,0,255},
	"gray": {128,128,128,255},
	"silver": {192,192,192,255},
	"red": {255,0,0,255},
	"fuchsia": {255,0,255,255},
	"yellow": {255,255,0,255},
	"white": {255,255,255,255},
}

	// Check to see if colors exists in map
	if _, ok := colors[backcolor]; !ok {
		backcolor = "gray"
	}
	if _, ok := colors[forecolor]; !ok {
		forecolor = "green"
	}

	hash := md5.New()

	hash.Write(bytestring)
	hashString := hex.EncodeToString(hash.Sum([]byte{}))
	hashString = strings.ToUpper(hashString)

	fill := buildArray(hashString)

	m := image.NewRGBA(image.Rect(0, 0, tilesPerSide*multiplier, tilesPerSide*multiplier))
	background := color.RGBA{colors[backcolor].red,colors[backcolor].green,colors[backcolor].blue,colors[backcolor].alpha}
	draw.Draw(m, m.Bounds(), &image.Uniform{background}, image.ZP, draw.Src)

	var col int
	var row float64
	var x, y int

	for i, value := range fill {

		if value == true {

			col = i / tilesPerSide
			row = math.Mod(float64(i), float64(tilesPerSide))

			if col != 0 {
				col = col * multiplier
			}
			if row != 0 {
				row = row * float64(multiplier)
			}

			for i := 0; i < multiplier; i++ {
				x = col + i
				for j := 0; j < multiplier; j++ {
					y = int(row) + j
					m.Set(int(x), int(y), color.RGBA{colors[forecolor].red, colors[forecolor].green, colors[forecolor].blue, colors[forecolor].alpha})

				}
			}
		}
	}
	return m
}
