package identicon

import (
	"image/color"
	"math"
)

func githubHSL(src [16]byte) (float64, float64, float64) {
	h := uint32(uint16(src[12]&0x0f)<<8 | uint16(src[13]))
	s := uint32(src[14])
	l := uint32(src[15])

	hue := interpolate(float64(h), 0, 4095, 0, 360)
	sat := interpolate(float64(s), 0, 255, 0, 0.2)
	lum := interpolate(float64(l), 0, 255, 0, 0.2)

	return hue, sat, lum
}

func HSLToRGB(h, s, l float64) color.Color {
	if h < 0 || h >= 360 ||
		s < 0 || s > 1 ||
		l < 0 || l > 1 {
		// Values out of range
		return color.Black
	}
	// When 0 ≤ h < 360, 0 ≤ s ≤ 1 and 0 ≤ l ≤ 1:
	C := (1 - math.Abs((2*l)-1)) * s
	X := C * (1 - math.Abs(math.Mod(h/60, 2)-1))
	m := l - (C / 2)
	var Rnot, Gnot, Bnot float64

	switch {
	case 0 <= h && h < 60:
		Rnot, Gnot, Bnot = C, X, 0
	case 60 <= h && h < 120:
		Rnot, Gnot, Bnot = X, C, 0
	case 120 <= h && h < 180:
		Rnot, Gnot, Bnot = 0, C, X
	case 180 <= h && h < 240:
		Rnot, Gnot, Bnot = 0, X, C
	case 240 <= h && h < 300:
		Rnot, Gnot, Bnot = X, 0, C
	case 300 <= h && h < 360:
		Rnot, Gnot, Bnot = C, 0, X
	}

	return color.RGBA{
		R: uint8(math.Round((Rnot + m) * 255)),
		G: uint8(math.Round((Gnot + m) * 255)),
		B: uint8(math.Round((Bnot + m) * 255)),
		A: 255,
	}
}

func interpolate(value, low1, high1, low2, high2 float64) float64 {
	return low2 + (high2-low2)*((value-low1)/(high1-low1))
}
