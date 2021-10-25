package tinykb

import (
	"image/color"
)

// Keyboard defines the interface needed for tinykb
type Keyboard interface {
	Display()
	Redraw(col, row int, selected bool)
	KeyEvent(key Key)
	GetKey() Key
	Layer(index int)
}

// RGB565ToRGBA converts a uint16 used in the display to color.RGBA
func RGB565ToRGBA(c uint16) color.RGBA {
	return color.RGBA{
		R: uint8((c & 0xF800) >> 8),
		G: uint8((c & 0x07E0) >> 3),
		B: uint8((c & 0x001F) << 3),
		A: 0xFF,
	}
}

// RGBATo565 converts a color.RGBA to uint16 used in the display
func RGBATo565(c color.RGBA) uint16 {
	r, g, b, _ := c.RGBA()
	return uint16((r & 0xF800) +
		((g & 0xFC00) >> 5) +
		((b & 0xF800) >> 11))
}
