package tinykb

import (
	"tinygo.org/x/tinyfont"
)

var Regular9pt7b = tinyfont.Font{
	Glyphs: []tinyfont.Glyph{
		///*   */ tinyfont.Glyph{Rune: 32, Width: 0x0, Height: 0x0, XAdvance: 0xb, XOffset: 0, YOffset: 1, Bitmaps: []uint8{}},
		/* 0x020 */ tinyfont.Glyph{Rune: 0x020, Width: 0xb, Height: 0xd, XAdvance: 0xb, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0xC0, 0x1F, 0xFE}},
		/* ! */ tinyfont.Glyph{Rune: 33, Width: 0x2, Height: 0xb, XAdvance: 0xb, XOffset: 4, YOffset: -10, Bitmaps: []uint8{0xaa, 0xa8, 0xc}},
		/* " */ tinyfont.Glyph{Rune: 34, Width: 0x6, Height: 0x5, XAdvance: 0xb, XOffset: 2, YOffset: -10, Bitmaps: []uint8{0xed, 0x24, 0x92, 0x48}},
		/* # */ tinyfont.Glyph{Rune: 35, Width: 0x7, Height: 0xc, XAdvance: 0xb, XOffset: 2, YOffset: -10, Bitmaps: []uint8{0x24, 0x48, 0x91, 0x2f, 0xe4, 0x89, 0x7f, 0x28, 0x51, 0x22, 0x40}},
		/* $ */ tinyfont.Glyph{Rune: 36, Width: 0x8, Height: 0xc, XAdvance: 0xb, XOffset: 1, YOffset: -10, Bitmaps: []uint8{0x8, 0x3e, 0x62, 0x40, 0x30, 0xe, 0x1, 0x81, 0xc3, 0xbe, 0x8, 0x8}},
		/* % */ tinyfont.Glyph{Rune: 37, Width: 0x7, Height: 0xb, XAdvance: 0xb, XOffset: 2, YOffset: -10, Bitmaps: []uint8{0x71, 0x12, 0x23, 0x80, 0x23, 0xb8, 0xe, 0x22, 0x44, 0x70}},
		/* & */ tinyfont.Glyph{Rune: 38, Width: 0x7, Height: 0xa, XAdvance: 0xb, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0x38, 0x81, 0x2, 0x6, 0x1a, 0x65, 0x46, 0xc8, 0xec}},
		/* ' */ tinyfont.Glyph{Rune: 39, Width: 0x3, Height: 0x5, XAdvance: 0xb, XOffset: 4, YOffset: -10, Bitmaps: []uint8{0xe9, 0x24}},
		/* ( */ tinyfont.Glyph{Rune: 40, Width: 0x2, Height: 0xd, XAdvance: 0xb, XOffset: 5, YOffset: -10, Bitmaps: []uint8{0x5a, 0xaa, 0xa9, 0x40}},
		/* ) */ tinyfont.Glyph{Rune: 41, Width: 0x2, Height: 0xd, XAdvance: 0xb, XOffset: 4, YOffset: -10, Bitmaps: []uint8{0xa9, 0x55, 0x5a, 0x80}},
		/* * */ tinyfont.Glyph{Rune: 42, Width: 0x7, Height: 0x7, XAdvance: 0xb, XOffset: 2, YOffset: -10, Bitmaps: []uint8{0x10, 0x22, 0x4b, 0xe3, 0x5, 0x11, 0x0}},
		/* + */ tinyfont.Glyph{Rune: 43, Width: 0x7, Height: 0x7, XAdvance: 0xb, XOffset: 2, YOffset: -8, Bitmaps: []uint8{0x10, 0x20, 0x47, 0xf1, 0x2, 0x4, 0x0}},
		/* , */ tinyfont.Glyph{Rune: 44, Width: 0x3, Height: 0x5, XAdvance: 0xb, XOffset: 2, YOffset: -1, Bitmaps: []uint8{0x6b, 0x48}},
		/* - */ tinyfont.Glyph{Rune: 45, Width: 0x9, Height: 0x1, XAdvance: 0xb, XOffset: 1, YOffset: -5, Bitmaps: []uint8{0xff, 0x0}},
		/* . */ tinyfont.Glyph{Rune: 46, Width: 0x2, Height: 0x2, XAdvance: 0xb, XOffset: 4, YOffset: -1, Bitmaps: []uint8{0xf0}},
		/* / */ tinyfont.Glyph{Rune: 47, Width: 0x7, Height: 0xd, XAdvance: 0xb, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0x2, 0x8, 0x10, 0x60, 0x81, 0x4, 0x8, 0x20, 0x41, 0x2, 0x8, 0x0}},
		/* 0 */ tinyfont.Glyph{Rune: 48, Width: 0x7, Height: 0xb, XAdvance: 0xb, XOffset: 2, YOffset: -10, Bitmaps: []uint8{0x38, 0x8a, 0xc, 0x18, 0x30, 0x60, 0xc1, 0x82, 0x88, 0xe0}},
		/* 1 */ tinyfont.Glyph{Rune: 49, Width: 0x5, Height: 0xb, XAdvance: 0xb, XOffset: 3, YOffset: -10, Bitmaps: []uint8{0x27, 0x28, 0x42, 0x10, 0x84, 0x21, 0x3e}},
		/* 2 */ tinyfont.Glyph{Rune: 50, Width: 0x7, Height: 0xb, XAdvance: 0xb, XOffset: 2, YOffset: -10, Bitmaps: []uint8{0x38, 0x8a, 0x8, 0x10, 0x20, 0x82, 0x8, 0x61, 0x3, 0xf8}},
		/* 3 */ tinyfont.Glyph{Rune: 51, Width: 0x8, Height: 0xb, XAdvance: 0xb, XOffset: 1, YOffset: -10, Bitmaps: []uint8{0x7c, 0x6, 0x2, 0x2, 0x1c, 0x6, 0x1, 0x1, 0x1, 0x42, 0x3c}},
		/* 4 */ tinyfont.Glyph{Rune: 52, Width: 0x6, Height: 0xb, XAdvance: 0xb, XOffset: 3, YOffset: -10, Bitmaps: []uint8{0x18, 0xa2, 0x92, 0x8a, 0x28, 0xbf, 0x8, 0x21, 0xc0}},
		/* 5 */ tinyfont.Glyph{Rune: 53, Width: 0x7, Height: 0xb, XAdvance: 0xb, XOffset: 2, YOffset: -10, Bitmaps: []uint8{0x7c, 0x81, 0x3, 0xe4, 0x40, 0x40, 0x81, 0x3, 0x88, 0xe0}},
		/* 6 */ tinyfont.Glyph{Rune: 54, Width: 0x7, Height: 0xb, XAdvance: 0xb, XOffset: 2, YOffset: -10, Bitmaps: []uint8{0x1e, 0x41, 0x4, 0xb, 0x98, 0xb0, 0xc1, 0xc2, 0x88, 0xe0}},
		/* 7 */ tinyfont.Glyph{Rune: 55, Width: 0x7, Height: 0xb, XAdvance: 0xb, XOffset: 2, YOffset: -10, Bitmaps: []uint8{0xfe, 0x4, 0x8, 0x20, 0x40, 0x82, 0x4, 0x8, 0x20, 0x40}},
		/* 8 */ tinyfont.Glyph{Rune: 56, Width: 0x7, Height: 0xb, XAdvance: 0xb, XOffset: 2, YOffset: -10, Bitmaps: []uint8{0x38, 0x8a, 0xc, 0x14, 0x47, 0x11, 0x41, 0x83, 0x8c, 0xe0}},
		/* 9 */ tinyfont.Glyph{Rune: 57, Width: 0x7, Height: 0xb, XAdvance: 0xb, XOffset: 2, YOffset: -10, Bitmaps: []uint8{0x38, 0x8a, 0x1c, 0x18, 0x68, 0xce, 0x81, 0x4, 0x13, 0xc0}},
		/* : */ tinyfont.Glyph{Rune: 58, Width: 0x2, Height: 0x8, XAdvance: 0xb, XOffset: 4, YOffset: -7, Bitmaps: []uint8{0xf0, 0xf}},
		/* ; */ tinyfont.Glyph{Rune: 59, Width: 0x3, Height: 0xb, XAdvance: 0xb, XOffset: 3, YOffset: -7, Bitmaps: []uint8{0x6c, 0x0, 0xd2, 0xd2, 0x0}},
		/* < */ tinyfont.Glyph{Rune: 60, Width: 0x8, Height: 0x8, XAdvance: 0xb, XOffset: 1, YOffset: -8, Bitmaps: []uint8{0x3, 0x4, 0x18, 0x60, 0x60, 0x18, 0x4, 0x3}},
		/* = */ tinyfont.Glyph{Rune: 61, Width: 0x9, Height: 0x4, XAdvance: 0xb, XOffset: 1, YOffset: -6, Bitmaps: []uint8{0xff, 0x80, 0x0, 0x1f, 0xf0}},
		/* > */ tinyfont.Glyph{Rune: 62, Width: 0x9, Height: 0x8, XAdvance: 0xb, XOffset: 1, YOffset: -8, Bitmaps: []uint8{0x40, 0x18, 0x3, 0x0, 0x60, 0x20, 0x60, 0xc0, 0x80}},
		/* ? */ tinyfont.Glyph{Rune: 63, Width: 0x7, Height: 0xa, XAdvance: 0xb, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0x3d, 0x84, 0x8, 0x30, 0xc2, 0x0, 0x0, 0x0, 0x30}},
		/* @ */ tinyfont.Glyph{Rune: 64, Width: 0x8, Height: 0xc, XAdvance: 0xb, XOffset: 2, YOffset: -10, Bitmaps: []uint8{0x3c, 0x46, 0x82, 0x8e, 0xb2, 0xa2, 0xa2, 0x9f, 0x80, 0x80, 0x40, 0x3c}},
		/* A */ tinyfont.Glyph{Rune: 65, Width: 0xb, Height: 0xa, XAdvance: 0xb, XOffset: 0, YOffset: -9, Bitmaps: []uint8{0x3c, 0x1, 0x40, 0x28, 0x9, 0x1, 0x10, 0x42, 0xf, 0xc1, 0x4, 0x40, 0x9e, 0x3c}},
		/* B */ tinyfont.Glyph{Rune: 66, Width: 0x9, Height: 0xa, XAdvance: 0xb, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xfe, 0x21, 0x90, 0x48, 0x67, 0xe2, 0x9, 0x2, 0x81, 0x41, 0xff, 0x80}},
		/* C */ tinyfont.Glyph{Rune: 67, Width: 0x9, Height: 0xa, XAdvance: 0xb, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0x3e, 0xb0, 0xf0, 0x30, 0x8, 0x4, 0x2, 0x0, 0x80, 0x60, 0x8f, 0x80}},
		/* D */ tinyfont.Glyph{Rune: 68, Width: 0x9, Height: 0xa, XAdvance: 0xb, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xfe, 0x21, 0x90, 0x68, 0x14, 0xa, 0x5, 0x2, 0x83, 0x43, 0x7f, 0x0}},
		/* E */ tinyfont.Glyph{Rune: 69, Width: 0x9, Height: 0xa, XAdvance: 0xb, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xff, 0x20, 0x90, 0x8, 0x87, 0xc2, 0x21, 0x0, 0x81, 0x40, 0xff, 0xc0}},
		/* F */ tinyfont.Glyph{Rune: 70, Width: 0x9, Height: 0xa, XAdvance: 0xb, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xff, 0xa0, 0x50, 0x8, 0x87, 0xc2, 0x21, 0x0, 0x80, 0x40, 0x78, 0x0}},
		/* G */ tinyfont.Glyph{Rune: 71, Width: 0xa, Height: 0xa, XAdvance: 0xb, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0x1e, 0x98, 0x6c, 0xa, 0x0, 0x80, 0x20, 0xf8, 0xb, 0x2, 0x60, 0x87, 0xc0}},
		/* H */ tinyfont.Glyph{Rune: 72, Width: 0x9, Height: 0xa, XAdvance: 0xb, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xe3, 0xa0, 0x90, 0x48, 0x27, 0xf2, 0x9, 0x4, 0x82, 0x41, 0x71, 0xc0}},
		/* I */ tinyfont.Glyph{Rune: 73, Width: 0x5, Height: 0xa, XAdvance: 0xb, XOffset: 3, YOffset: -9, Bitmaps: []uint8{0xf9, 0x8, 0x42, 0x10, 0x84, 0x27, 0xc0}},
		/* J */ tinyfont.Glyph{Rune: 74, Width: 0x8, Height: 0xa, XAdvance: 0xb, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0x1f, 0x2, 0x2, 0x2, 0x2, 0x2, 0x82, 0x82, 0xc6, 0x78}},
		/* K */ tinyfont.Glyph{Rune: 75, Width: 0x9, Height: 0xa, XAdvance: 0xb, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xe3, 0xa1, 0x11, 0x9, 0x5, 0x83, 0x21, 0x8, 0x84, 0x41, 0x70, 0xc0}},
		/* L */ tinyfont.Glyph{Rune: 76, Width: 0x8, Height: 0xa, XAdvance: 0xb, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0xe0, 0x40, 0x40, 0x40, 0x40, 0x40, 0x41, 0x41, 0x41, 0xff}},
		/* M */ tinyfont.Glyph{Rune: 77, Width: 0xb, Height: 0xa, XAdvance: 0xb, XOffset: 0, YOffset: -9, Bitmaps: []uint8{0xe0, 0xec, 0x19, 0x45, 0x28, 0xa4, 0xa4, 0x94, 0x91, 0x12, 0x2, 0x40, 0x5c, 0x1c}},
		/* N */ tinyfont.Glyph{Rune: 78, Width: 0x9, Height: 0xa, XAdvance: 0xb, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xc3, 0xb0, 0x94, 0x4a, 0x24, 0x92, 0x49, 0x14, 0x8a, 0x43, 0x70, 0x80}},
		/* O */ tinyfont.Glyph{Rune: 79, Width: 0x9, Height: 0xa, XAdvance: 0xb, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0x1e, 0x31, 0x90, 0x50, 0x18, 0xc, 0x6, 0x2, 0x82, 0x63, 0xf, 0x0}},
		/* P */ tinyfont.Glyph{Rune: 80, Width: 0x8, Height: 0xa, XAdvance: 0xb, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xfe, 0x43, 0x41, 0x41, 0x42, 0x7c, 0x40, 0x40, 0x40, 0xf0}},
		/* Q */ tinyfont.Glyph{Rune: 81, Width: 0x9, Height: 0xd, XAdvance: 0xb, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0x1c, 0x31, 0x90, 0x50, 0x18, 0xc, 0x6, 0x2, 0x82, 0x63, 0x1f, 0x4, 0x7, 0x92, 0x30}},
		/* R */ tinyfont.Glyph{Rune: 82, Width: 0x9, Height: 0xa, XAdvance: 0xb, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xfe, 0x21, 0x90, 0x48, 0x24, 0x23, 0xe1, 0x10, 0x84, 0x41, 0x70, 0xc0}},
		/* S */ tinyfont.Glyph{Rune: 83, Width: 0x7, Height: 0xa, XAdvance: 0xb, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0x3a, 0xcd, 0xa, 0x3, 0x1, 0x80, 0xc1, 0xc7, 0x78}},
		/* T */ tinyfont.Glyph{Rune: 84, Width: 0x9, Height: 0xa, XAdvance: 0xb, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xff, 0xc4, 0x62, 0x21, 0x0, 0x80, 0x40, 0x20, 0x10, 0x8, 0x1f, 0x0}},
		/* U */ tinyfont.Glyph{Rune: 85, Width: 0x9, Height: 0xa, XAdvance: 0xb, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xe3, 0xa0, 0x90, 0x48, 0x24, 0x12, 0x9, 0x4, 0x82, 0x22, 0xe, 0x0}},
		/* V */ tinyfont.Glyph{Rune: 86, Width: 0xb, Height: 0xa, XAdvance: 0xb, XOffset: 0, YOffset: -9, Bitmaps: []uint8{0xf1, 0xe8, 0x10, 0x82, 0x10, 0x42, 0x10, 0x22, 0x4, 0x80, 0x50, 0xc, 0x0, 0x80}},
		/* W */ tinyfont.Glyph{Rune: 87, Width: 0xb, Height: 0xa, XAdvance: 0xb, XOffset: 0, YOffset: -9, Bitmaps: []uint8{0xf1, 0xe8, 0x9, 0x11, 0x25, 0x44, 0xa8, 0x55, 0xc, 0xa1, 0x8c, 0x31, 0x84, 0x30}},
		/* X */ tinyfont.Glyph{Rune: 88, Width: 0x9, Height: 0xa, XAdvance: 0xb, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xe3, 0xa0, 0x88, 0x82, 0x80, 0x80, 0xc0, 0x90, 0x44, 0x41, 0x71, 0xc0}},
		/* Y */ tinyfont.Glyph{Rune: 89, Width: 0x9, Height: 0xa, XAdvance: 0xb, XOffset: 1, YOffset: -9, Bitmaps: []uint8{0xe3, 0xa0, 0x88, 0x82, 0x81, 0x40, 0x40, 0x20, 0x10, 0x8, 0x1f, 0x0}},
		/* Z */ tinyfont.Glyph{Rune: 90, Width: 0x7, Height: 0xa, XAdvance: 0xb, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0xfd, 0xa, 0x20, 0x81, 0x4, 0x10, 0x21, 0x83, 0xfc}},
		/* [ */ tinyfont.Glyph{Rune: 91, Width: 0x2, Height: 0xd, XAdvance: 0xb, XOffset: 5, YOffset: -10, Bitmaps: []uint8{0xea, 0xaa, 0xaa, 0xc0}},
		/* \ */ tinyfont.Glyph{Rune: 92, Width: 0x7, Height: 0xd, XAdvance: 0xb, XOffset: 2, YOffset: -11, Bitmaps: []uint8{0x80, 0x81, 0x3, 0x2, 0x4, 0x4, 0x8, 0x8, 0x10, 0x10, 0x20, 0x20}},
		/* ] */ tinyfont.Glyph{Rune: 93, Width: 0x2, Height: 0xd, XAdvance: 0xb, XOffset: 4, YOffset: -10, Bitmaps: []uint8{0xd5, 0x55, 0x55, 0xc0}},
		/* ^ */ tinyfont.Glyph{Rune: 94, Width: 0x7, Height: 0x5, XAdvance: 0xb, XOffset: 2, YOffset: -10, Bitmaps: []uint8{0x10, 0x51, 0x22, 0x28, 0x20}},
		/* _ */ tinyfont.Glyph{Rune: 95, Width: 0xb, Height: 0x1, XAdvance: 0xb, XOffset: 0, YOffset: 2, Bitmaps: []uint8{0xff, 0xe0}},
		/* ` */ tinyfont.Glyph{Rune: 96, Width: 0x3, Height: 0x3, XAdvance: 0xb, XOffset: 3, YOffset: -11, Bitmaps: []uint8{0x88, 0x80}},
		/* a */ tinyfont.Glyph{Rune: 97, Width: 0x9, Height: 0x8, XAdvance: 0xb, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0x7e, 0x0, 0x80, 0x47, 0xec, 0x14, 0xa, 0xc, 0xfb}},
		/* b */ tinyfont.Glyph{Rune: 98, Width: 0x9, Height: 0xb, XAdvance: 0xb, XOffset: 1, YOffset: -10, Bitmaps: []uint8{0xc0, 0x20, 0x10, 0xb, 0xc6, 0x12, 0x5, 0x2, 0x81, 0x40, 0xb0, 0xb7, 0x80}},
		/* c */ tinyfont.Glyph{Rune: 99, Width: 0x7, Height: 0x8, XAdvance: 0xb, XOffset: 2, YOffset: -7, Bitmaps: []uint8{0x3a, 0x8e, 0xc, 0x8, 0x10, 0x10, 0x9e}},
		/* d */ tinyfont.Glyph{Rune: 100, Width: 0x9, Height: 0xb, XAdvance: 0xb, XOffset: 1, YOffset: -10, Bitmaps: []uint8{0x3, 0x0, 0x80, 0x47, 0xa4, 0x34, 0xa, 0x5, 0x2, 0x81, 0x21, 0x8f, 0x60}},
		/* e */ tinyfont.Glyph{Rune: 101, Width: 0x8, Height: 0x8, XAdvance: 0xb, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0x3c, 0x43, 0x81, 0xff, 0x80, 0x80, 0x61, 0x3e}},
		/* f */ tinyfont.Glyph{Rune: 102, Width: 0x6, Height: 0xb, XAdvance: 0xb, XOffset: 3, YOffset: -10, Bitmaps: []uint8{0x3d, 0x4, 0x3e, 0x41, 0x4, 0x10, 0x41, 0xf, 0x80}},
		/* g */ tinyfont.Glyph{Rune: 103, Width: 0x9, Height: 0xb, XAdvance: 0xb, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0x3d, 0xa1, 0xa0, 0x50, 0x28, 0x14, 0x9, 0xc, 0x7a, 0x1, 0x1, 0x87, 0x80}},
		/* h */ tinyfont.Glyph{Rune: 104, Width: 0x9, Height: 0xb, XAdvance: 0xb, XOffset: 1, YOffset: -10, Bitmaps: []uint8{0xc0, 0x20, 0x10, 0xb, 0xc6, 0x32, 0x9, 0x4, 0x82, 0x41, 0x20, 0xb8, 0xe0}},
		/* i */ tinyfont.Glyph{Rune: 105, Width: 0x7, Height: 0xa, XAdvance: 0xb, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0x10, 0x1, 0xc0, 0x81, 0x2, 0x4, 0x8, 0x11, 0xfc}},
		/* j */ tinyfont.Glyph{Rune: 106, Width: 0x5, Height: 0xd, XAdvance: 0xb, XOffset: 3, YOffset: -9, Bitmaps: []uint8{0x10, 0x3e, 0x10, 0x84, 0x21, 0x8, 0x42, 0x3f, 0x0}},
		/* k */ tinyfont.Glyph{Rune: 107, Width: 0x8, Height: 0xb, XAdvance: 0xb, XOffset: 2, YOffset: -10, Bitmaps: []uint8{0xc0, 0x40, 0x40, 0x4f, 0x44, 0x58, 0x70, 0x48, 0x44, 0x42, 0xc7}},
		/* l */ tinyfont.Glyph{Rune: 108, Width: 0x7, Height: 0xb, XAdvance: 0xb, XOffset: 2, YOffset: -10, Bitmaps: []uint8{0x70, 0x20, 0x40, 0x81, 0x2, 0x4, 0x8, 0x10, 0x23, 0xf8}},
		/* m */ tinyfont.Glyph{Rune: 109, Width: 0x9, Height: 0x8, XAdvance: 0xb, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xb7, 0x64, 0x62, 0x31, 0x18, 0x8c, 0x46, 0x23, 0x91}},
		/* n */ tinyfont.Glyph{Rune: 110, Width: 0x9, Height: 0x8, XAdvance: 0xb, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0x5e, 0x31, 0x90, 0x48, 0x24, 0x12, 0x9, 0x5, 0xc7}},
		/* o */ tinyfont.Glyph{Rune: 111, Width: 0x9, Height: 0x8, XAdvance: 0xb, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0x3e, 0x31, 0xa0, 0x30, 0x18, 0xc, 0x5, 0x8c, 0x7c}},
		/* p */ tinyfont.Glyph{Rune: 112, Width: 0x9, Height: 0xb, XAdvance: 0xb, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xde, 0x30, 0x90, 0x28, 0x14, 0xa, 0x5, 0x84, 0xbc, 0x40, 0x20, 0x38, 0x0}},
		/* q */ tinyfont.Glyph{Rune: 113, Width: 0x9, Height: 0xb, XAdvance: 0xb, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0x3d, 0xa1, 0xa0, 0x50, 0x28, 0x14, 0x9, 0xc, 0x7a, 0x1, 0x0, 0x80, 0xe0}},
		/* r */ tinyfont.Glyph{Rune: 114, Width: 0x7, Height: 0x8, XAdvance: 0xb, XOffset: 3, YOffset: -7, Bitmaps: []uint8{0xce, 0xa1, 0x82, 0x4, 0x8, 0x10, 0x7c}},
		/* s */ tinyfont.Glyph{Rune: 115, Width: 0x7, Height: 0x8, XAdvance: 0xb, XOffset: 2, YOffset: -7, Bitmaps: []uint8{0x3a, 0x8d, 0xb, 0x80, 0xf0, 0x70, 0xde}},
		/* t */ tinyfont.Glyph{Rune: 116, Width: 0x8, Height: 0xa, XAdvance: 0xb, XOffset: 2, YOffset: -9, Bitmaps: []uint8{0x40, 0x40, 0xfc, 0x40, 0x40, 0x40, 0x40, 0x40, 0x41, 0x3e}},
		/* u */ tinyfont.Glyph{Rune: 117, Width: 0x8, Height: 0x8, XAdvance: 0xb, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xc3, 0x41, 0x41, 0x41, 0x41, 0x41, 0x43, 0x3d}},
		/* v */ tinyfont.Glyph{Rune: 118, Width: 0x9, Height: 0x8, XAdvance: 0xb, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xe3, 0xa0, 0x90, 0x84, 0x42, 0x20, 0xa0, 0x50, 0x10}},
		/* w */ tinyfont.Glyph{Rune: 119, Width: 0x9, Height: 0x8, XAdvance: 0xb, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xe3, 0xc0, 0x92, 0x4b, 0x25, 0x92, 0xa9, 0x98, 0x44}},
		/* x */ tinyfont.Glyph{Rune: 120, Width: 0x9, Height: 0x8, XAdvance: 0xb, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xe3, 0x31, 0x5, 0x1, 0x1, 0x41, 0x11, 0x5, 0xc7}},
		/* y */ tinyfont.Glyph{Rune: 121, Width: 0x9, Height: 0xb, XAdvance: 0xb, XOffset: 1, YOffset: -7, Bitmaps: []uint8{0xe3, 0xa0, 0x90, 0x84, 0x42, 0x40, 0xa0, 0x60, 0x10, 0x10, 0x8, 0x3e, 0x0}},
		/* z */ tinyfont.Glyph{Rune: 122, Width: 0x7, Height: 0x8, XAdvance: 0xb, XOffset: 2, YOffset: -7, Bitmaps: []uint8{0xfd, 0x8, 0x20, 0x82, 0x8, 0x10, 0xbf}},
		/* { */ tinyfont.Glyph{Rune: 123, Width: 0x3, Height: 0xd, XAdvance: 0xb, XOffset: 4, YOffset: -10, Bitmaps: []uint8{0x29, 0x24, 0xa2, 0x49, 0x26}},
		/* | */ tinyfont.Glyph{Rune: 124, Width: 0x1, Height: 0xd, XAdvance: 0xb, XOffset: 5, YOffset: -10, Bitmaps: []uint8{0xff, 0xf8}},
		/* } */ tinyfont.Glyph{Rune: 125, Width: 0x3, Height: 0xd, XAdvance: 0xb, XOffset: 4, YOffset: -10, Bitmaps: []uint8{0x89, 0x24, 0x8a, 0x49, 0x2c}},
		/* ~ */ tinyfont.Glyph{Rune: 126, Width: 0x7, Height: 0x3, XAdvance: 0xb, XOffset: 2, YOffset: -6, Bitmaps: []uint8{0x61, 0x24, 0x30}},
		/* KeyReturn */ tinyfont.Glyph{Rune: 0x101, Width: 0xb, Height: 0xd, XAdvance: 0xb, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x00, 0x00, 0x00, 0x02, 0x00, 0x40, 0x08, 0x41, 0x10, 0x27, 0xFC, 0x40, 0x04, 0x00, 0x00, 0x00, 0x00}},
		/* KeyTab */ tinyfont.Glyph{Rune: 0x102, Width: 0xb, Height: 0xd, XAdvance: 0xb, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x00, 0x00, 0x04, 0x08, 0x81, 0x90, 0x3B, 0xFF, 0xC0, 0xE8, 0x19, 0x02, 0x20, 0x04, 0x00, 0x00, 0x00}},
		/* KeyBackspace */ tinyfont.Glyph{Rune: 0x103, Width: 0xb, Height: 0xd, XAdvance: 0xb, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x0F, 0xE2, 0x04, 0x80, 0xA4, 0x58, 0x53, 0x04, 0x61, 0x4A, 0x45, 0x20, 0x22, 0x04, 0x3F, 0x80, 0x00}},
		/* KeyRight */ tinyfont.Glyph{Rune: 0x106, Width: 0xb, Height: 0xd, XAdvance: 0xb, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x08, 0x00, 0xBF, 0xF8, 0x02, 0x00, 0x80, 0x20, 0x00, 0x00, 0x00}},
		/* KeyLeft */ tinyfont.Glyph{Rune: 0x107, Width: 0xb, Height: 0xd, XAdvance: 0xb, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x00, 0x00, 0x00, 0x00, 0x08, 0x02, 0x00, 0x80, 0x3F, 0xFA, 0x00, 0x20, 0x02, 0x00, 0x00, 0x00, 0x00}},
		/* KeyDown */ tinyfont.Glyph{Rune: 0x108, Width: 0xb, Height: 0xd, XAdvance: 0xb, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x00, 0x00, 0x80, 0x10, 0x02, 0x00, 0x40, 0x08, 0x01, 0x00, 0x20, 0x24, 0x82, 0xA0, 0x38, 0x02, 0x00}},
		/* KeyUp */ tinyfont.Glyph{Rune: 0x109, Width: 0xb, Height: 0xd, XAdvance: 0xb, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x00, 0x00, 0x80, 0x38, 0x0A, 0x82, 0x48, 0x08, 0x01, 0x00, 0x20, 0x04, 0x00, 0x80, 0x10, 0x02, 0x00}},
		/* KeyShift */ tinyfont.Glyph{Rune: 0x1FD, Width: 0xb, Height: 0xd, XAdvance: 0xb, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x00, 0x00, 0x80, 0x38, 0x0F, 0x83, 0xF8, 0xFF, 0x87, 0xC0, 0xF8, 0x1F, 0x00, 0x00, 0x7C, 0x0F, 0x80, 0x00}},
		/* KeyShiftRelease */ tinyfont.Glyph{Rune: 0x1FE, Width: 0xb, Height: 0xd, XAdvance: 0xb, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x00, 0x00, 0x80, 0x28, 0x08, 0x82, 0x08, 0xE3, 0x84, 0x40, 0x88, 0x1F, 0x00, 0x00, 0x7C, 0x0F, 0x80, 0x00}},
		/* KeyClose */ tinyfont.Glyph{Rune: 0x1FF, Width: 0xb, Height: 0xd, XAdvance: 0xb, XOffset: 0, YOffset: -10, Bitmaps: []uint8{0x00, 0x00, 0x03, 0xFF, 0xBF, 0xE7, 0xFC, 0x7F, 0x0F, 0xE0, 0xF8, 0x1F, 0x01, 0xC0, 0x10, 0x02, 0x00}},
	},

	YAdvance: 0x12,
}
