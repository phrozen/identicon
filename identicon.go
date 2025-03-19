package identicon

import (
	"crypto/md5"
	"image"
	"image/color"
	"iter"

	"golang.org/x/image/draw"
)

// GitHub generates a Github-like identicon from the given data.
// The data is hashed using MD5 and the resulting hash is used to
// generate a 12x12 image (1px margin and 2x2 pixel 5x5 matrix).
// Algorithm Ported from: github.com/dgraham/identicon
func GitHub(identifier string) *image.Paletted {
	src := md5.Sum([]byte(identifier))
	hue, sat, lum := githubHSL(src)
	background := color.RGBA{R: 240, G: 240, B: 240, A: 255} // Default Background
	foreground := HSLToRGB(hue, 0.65-sat, 0.75-lum)          // Calculated Foreground

	palette := []color.Color{background, foreground}
	img := image.NewPaletted(image.Rect(0, 0, 12, 12), palette)

	for i, v := range Nibbler(src[0:8]) {
		if v%2 == 0 || i >= 15 {
			continue
		}
		x := 2 * (i / 5)
		y := (i%5)*2 + 1
		set2x2Idx(img, 5-x, y, 1) // Left
		set2x2Idx(img, 5+x, y, 1) // Right
	}
	return img
}

// GitHubAlternate generates a Github-like identicon from the given data
// using a dark background color scheme and a more saturated foreground.
// The data is hashed using MD5 and the resulting hash is used to generate
// a 16x16 image (1px margin and 2x2 pixel 7x7 matrix), which better aligns
// with the hash size and image size, making it more recognizable.
func GitHubAlternate(identifier string) *image.Paletted {
	src := md5.Sum([]byte(identifier))
	hue, sat, lum := githubHSL(src)
	background := color.RGBA{R: 32, G: 32, B: 32, A: 255}
	foreground := HSLToRGB(hue, 0.75-sat, 0.85-lum)

	palette := []color.Color{background, foreground}
	img := image.NewPaletted(image.Rect(0, 0, 16, 16), palette)

	for i, v := range Nibbler(src[0:14]) {
		if v%2 == 0 {
			continue
		}
		x := 2 * (i / 7)
		y := (i%7)*2 + 1
		set2x2Idx(img, 7-x, y, 1) // Left
		set2x2Idx(img, 7+x, y, 1) // Right
	}
	return img
}

// Nibbler is an iterator that yields the index and the
// nibbles of a byte slice.
func Nibbler(items []byte) iter.Seq2[int, byte] {
	return func(yield func(int, byte) bool) {
		for i, item := range items {
			if !yield(i*2, item>>4) {
				return
			}
			if !yield(i*2+1, item&0x0F) {
				return
			}
		}
	}
}

// Sets a 2x2 block of pixels with the given palette index
// at the given coordinates on a paletted image.
func set2x2Idx(img *image.Paletted, x, y int, idx byte) {
	for j := 0; j < 2; j++ {
		for i := 0; i < 2; i++ {
			img.Pix[(y+j)*img.Stride+(x+i)] = idx
		}
	}
}

// ScalePaletted scales a paletted image to the given size using the
// nearest neighbor algorithm.
func ScalePaletted(src *image.Paletted, width, height int) *image.Paletted {
	dst := image.NewPaletted(image.Rect(0, 0, width, height), src.Palette)
	draw.NearestNeighbor.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)
	return dst
}
