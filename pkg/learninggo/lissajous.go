package learninggo

// lissajous animated gif generator from Learning Go, An Idiomatic Approach...
// plus small changes.

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	color.RGBA{0x77, 0x77, 0x77, 0x77},
}

const (
	whiteIndex      = 0 // first color in palette
	blackIndex      = 1 // next color in palette
	backgroundIndex = 2
)

// LissajousFromArgs creates a random Lissajous animated gif. Try 5, 50, 64, 8, 0.001
func LissajousFromArgs(args []string) {
	path := args[0]
	cycles, _ := strconv.ParseFloat(args[1], 64)
	size, _ := strconv.ParseFloat(args[2], 64)
	nframes, _ := strconv.ParseFloat(args[3], 64)
	delay, _ := strconv.ParseFloat(args[4], 64)
	res, _ := strconv.ParseFloat(args[5], 64)
	Lissajous(path, cycles, size, nframes, delay, res)
}

// Lissajous creates a random Lissajous animated gif. Try 5, 0.001, 50, 64, 8
func Lissajous(path string, cycles, res, size, nframes, delay float64) {

	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()
	freq := rand.Float64() * 2.0
	anim := gif.GIF{LoopCount: int(nframes)}
	phase := 0.0
	j := 0
	for i := 0; i < int(nframes); i++ {
		rect := image.Rect(0, 0, int(2*size)+1, int(2*size)+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			x2 := math.Tan(t)
			y := math.Sin(t*freq + phase)
			j++
			if j%2 == 1 {
				img.SetColorIndex(int(size+(x*size+0.5)), int(size+(y*size+0.5)), blackIndex)
			} else {
				img.SetColorIndex(int(size+(x*size+0.5)+1), int(size+(y*size+0.5)-1), backgroundIndex)
				img.SetColorIndex(int(size+(x2*size+0.5)+1), int(size+(y*size+0.5)-1), backgroundIndex)
			}

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, int(delay))
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(file, &anim)
}
