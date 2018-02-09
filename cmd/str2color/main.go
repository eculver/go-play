package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"math"
)

const (
	start = "\033["
	reset = "\033[0m"

	fgWhite = 15
	fgBlack = 0

	// Coefficients derived from the ITU-R recommendation BT.709 for
	// determining relative luminance. See https://en.wikipedia.org/wiki/Luma_(video)#Rec._601_luma_versus_Rec._709_luma_coefficients
	// A color's relative luminance is used to determine whether foreground
	// text should be white or black.
	// This SO post was also very helpful: https://stackoverflow.com/questions/3942878/how-to-decide-font-color-in-white-or-black-depending-on-background-color
	rLumaCoef = 0.2126
	gLumaCoef = 0.7152
	bLumaCoef = 0.0722
)

func main() {
	in := []string{
		"foo",
		"bar",
		"baz",
		"qux",
		"¯\\_(ツ)_/¯",
	}

	for _, s := range in {
		bs := md5.Sum([]byte(s))
		r := int(bs[0])
		g := int(bs[1])
		b := int(bs[2])

		// fmt.Printf("%s => rgb(%d,%d,%d)\n", s, int(bs[0]), int(bs[1]), int(bs[2]))
		buf := bytes.NewBufferString(start)

		// colorize as foreground
		fmt.Fprintf(buf, "38;2;%d;%d;%d", r, g, b)
		buf.WriteRune('m')
		buf.WriteString(s)
		buf.WriteString(reset)

		// reset
		buf.WriteString("   |   ")
		buf.WriteString(start)

		// colorize as background + foreground based on background luminance
		fmt.Fprintf(buf, "38;5;%d;", fgColor(r, g, b))
		fmt.Fprintf(buf, "48;2;%d;%d;%d", r, g, b)
		buf.WriteRune('m')
		buf.WriteString(s)
		buf.WriteString(reset)

		// write to stdout
		fmt.Println(buf.String())
	}
}

func lumaNormalize(c float64) float64 {
	if c <= 0.03928 {
		return c / 12.92
	}
	return math.Pow(((c + 0.055) / 1.055), 2.4)
}

func fgColor(r, g, b int) int {
	// These are 8-bit values, have to use sRGB per the W3C recommandation
	// AND have to use normalized sRGB values to calculate relative luminance.
	// See https://www.w3.org/TR/WCAG20/#relativeluminancedef for more info.
	rNorm := lumaNormalize(float64(r) / 255.0)
	gNorm := lumaNormalize(float64(g) / 255.0)
	bNorm := lumaNormalize(float64(b) / 255.0)
	luminance := (rLumaCoef * rNorm) + (gLumaCoef * gNorm) + (bLumaCoef * bNorm)
	// fmt.Printf("luminance: %.2f\n", luminance)
	if luminance > 0.179 {
		return fgBlack
	}
	return fgWhite
}
