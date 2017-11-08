// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
//this seemed to be fastest for input of 32 (hence 32 concurrent) 154 millseconds
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"strconv"
	"time"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	var s int = 1
	if len(os.Args) > 1 {
		chanSize, err := strconv.Atoi(os.Args[1])
		if err != nil || width%chanSize != 0 {
			fmt.Println("Input arg must be a number divisible by 1024")
			os.Exit(1)
		}
		s = chanSize
	}

	//buffered channel
	ch := make(chan bool, s)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	f, err := os.Create("res.png")
	defer f.Close()
	if err != nil {
		fmt.Printf("Unable to create create out file res.png")
		os.Exit(1)
	}
	start := time.Now()
	for i := 0; i < s; i++ {
		lower := i * (height / s)       // 0,512 , if 4 then 0, 256, 512, 768, 1024
		upper := (height / s) * (i + 1) //512,1024, if 4 then 256,512,768,1024
		go func(i int) {
			for py := lower; py < upper; py++ {
				y := float64(py)/height*(ymax-ymin) + ymin
				for px := 0; px < width; px++ {
					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					img.Set(px, py, mandelbrot(z))
				}
			}
			ch <- true
		}(i)
	}

	for i := 0; i < s; i++ {
		_, ok := <-ch
		if !ok {
			break
		}
	}
	png.Encode(f, img) // NOTE: ignoring errors
	elapsed := time.Since(start)
	fmt.Printf("Time is %s", elapsed)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

//!-

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
