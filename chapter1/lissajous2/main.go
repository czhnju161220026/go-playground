package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = [] color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func generate(out io.Writer, cycles int ) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles) *2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t * freq + phase)
			img.SetColorIndex(size + int(x * size + 0.5), size + int(y*size + 0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	_ = gif.EncodeAll(out, &anim)
}


func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err == nil{
		s := r.Form.Get("cycles")
		fmt.Println(s)
		cycles, _ := strconv.Atoi(s)
		generate(w, cycles)
	}
}

func main() {
	http.HandleFunc("/lsj", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}