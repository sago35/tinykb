package keyboard14x4

import (
	"github.com/sago35/tinykb"
	"tinygo.org/x/drivers"
)

const (
	sz = 22
)

// Keyboard14x4 is the definition of a keyboard.
type Keyboard14x4 struct {
	Layout [2][4][14]tinykb.Key
	Disp   drivers.Displayer
	Column int
	Row    int
	X      int16
	Y      int16
	layer  int
}

// Keyboard14x4Layout defines the key layout for Keyboard14x4.
var Keyboard14x4Layout = [2][4][14]tinykb.Key{
	{
		{
			'`', '1', '2', '3', '4', '5', '6', '7', '8', '9', '0', '-', '=',
			tinykb.KeyBackspace,
		},
		{
			'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p', '[', ']', '\\',
			tinykb.KeyClose,
		},
		{
			'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l', ';', '\'',
			tinykb.KeyReturn,
			tinykb.KeyUp,
			tinykb.KeyShift,
		},
		{
			'z', 'x', 'c', 'v', 'b', 'n', 'm', ',', '.', '/', ' ',
			tinykb.KeyLeft,
			tinykb.KeyDown,
			tinykb.KeyRight,
		},
	},
	{
		{
			'~', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '_', '+',
			tinykb.KeyBackspace,
		},
		{
			'Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O', 'P', '{', '}', '|',
			tinykb.KeyTab,
		},
		{
			'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', ':', '"',
			tinykb.KeyReturn,
			tinykb.KeyUp,
			tinykb.KeyShiftRelease,
		},
		{
			'Z', 'X', 'C', 'V', 'B', 'N', 'M', '<', '>', '?', ' ',
			tinykb.KeyLeft,
			tinykb.KeyDown,
			tinykb.KeyRight,
		},
	},
}

// New creates and returns a Keyboard.
func New(display drivers.Displayer, x, y int16) *Keyboard14x4 {
	return &Keyboard14x4{
		Disp:   display,
		Layout: Keyboard14x4Layout,
		layer:  0,
		X:      x,
		Y:      y,
	}
}

func (k *Keyboard14x4) Display() {
	for row := 0; row < len(k.Layout[k.layer]); row++ {
		for col := 0; col < len(k.Layout[k.layer][0]); col++ {
			if col == k.Column && row == k.Row {
				k.Redraw(col, row, true)
			} else {
				k.Redraw(col, row, false)
			}
		}
	}
}

func (k *Keyboard14x4) Redraw(col, row int, selected bool) {
	btn := k.Layout[k.layer][row][col]
	xxx := k.X + (sz+1)*(int16(col)+0) + sz/2*0
	yyy := k.Y + ((sz + 1) * int16(row))
	if selected {
		btn.DisplaySelected()
	} else {
		btn.Display()
	}
	buf := btn.GetBuffer()
	for yy := int16(0); yy < sz; yy++ {
		for xx := int16(0); xx < sz; xx++ {
			k.Disp.SetPixel(xxx+xx, yyy+yy, tinykb.RGB565ToRGBA(buf[xx+yy*sz]))
		}
	}
}

func (k *Keyboard14x4) GetKey() tinykb.Key {
	return k.Layout[k.layer][k.Row][k.Column]
}

func (k *Keyboard14x4) Layer(index int) {
	k.layer = index
	k.Display()
}

func (k *Keyboard14x4) KeyEvent(key tinykb.Key) {
	col := k.Column
	row := k.Row
	switch key {
	case tinykb.KeyRight:
		col = (col + len(k.Layout[k.layer][0]) + 1) % len(k.Layout[k.layer][0])
	case tinykb.KeyLeft:
		col = (col + len(k.Layout[k.layer][0]) - 1) % len(k.Layout[k.layer][0])
	case tinykb.KeyUp:
		row = (row + len(k.Layout[k.layer]) - 1) % len(k.Layout[k.layer])
	case tinykb.KeyDown:
		row = (row + len(k.Layout[k.layer]) + 1) % len(k.Layout[k.layer])
	default:
		return
	}
	k.Redraw(k.Column, k.Row, false)
	k.Column = col
	k.Row = row
	k.Redraw(col, row, true)
}
