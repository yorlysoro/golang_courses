package main

// Lissajous generates GIF animations of random Lissajous figures in multiple colors.

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.RGBA{0xFF, 0xFF, 0xFF, 0xFF}, // background (white)
	color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // red
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // green
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // blue
	color.RGBA{0xFF, 0xA5, 0x00, 0xFF}, // orange
	color.RGBA{0x80, 0x00, 0x80, 0xFF}, // purple
}

func main() {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	lissajous(os.Stdout, rnd)
}

func lissajous(out io.Writer, rnd *rand.Rand) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rnd.Float64()*3.0 + 0.5 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	palCount := len(palette) - 1 // exclude background
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// pick a color index based on position/phase for variety
			h := (x + y + 2) / 4                       // normalize to [0,1]
			idx := uint8(1 + int(h*float64(palCount))) // 1..len(palette)-1
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), idx)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	_ = gif.EncodeAll(out, &anim) // ignore encoding errors
}
