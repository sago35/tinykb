package main

import (
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/sago35/tinykb"
	"github.com/sago35/tinykb/keyboard14x4"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freemono"
)

var (
	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
)

func main() {
	display.FillScreen(black)

	for {
		str, err := run()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%q\n", str)

		time.Sleep(500 * time.Millisecond)
	}
}

func run() (string, error) {
	var kb tinykb.Keyboard
	kb = keyboard14x4.New(display, 0, 150)
	kb.Display()

	str := ""
	needsRedraw := true
	prevKey := uint16(0xFFFF)
	guard := time.Now()
	for {
		key := GetPressedKey()

		if prevKey != key || time.Now().Sub(guard) > 500*time.Millisecond {
			if prevKey == key {
				guard = time.Now().Add(-1 * 400 * time.Millisecond)
			} else {
				guard = time.Now()
			}
			switch key {
			case tinykb.KeyRight, tinykb.KeyLeft, tinykb.KeyUp, tinykb.KeyDown:
				kb.KeyEvent(tinykb.Key(key))
			case tinykb.KeyReturn:
				k := kb.GetKey()
				//fmt.Printf("%s\n", k)
				switch k {
				case tinykb.KeyShift:
					kb.Layer(1)
				case tinykb.KeyShiftRelease:
					kb.Layer(0)
				case tinykb.KeyReturn:
					str += "\n"
				case tinykb.KeyBackspace:
					if len(str) > 0 {
						str = str[:len(str)-1]
					}
				case tinykb.KeyClose:
					return str, nil
				default:
					str += fmt.Sprintf("%c", k)
				}
				needsRedraw = true
			case ' ':
				str += fmt.Sprintf("%c", key)
				needsRedraw = true
			case tinykb.KeyBackspace:
				if len(str) > 0 {
					str = str[:len(str)-1]
					needsRedraw = true
				}
			default:
			}

			if needsRedraw {
				display.FillRectangle(0, 0, 320, 150, black)
				tinyfont.WriteLine(display, &freemono.Regular9pt7b, 0, 15, str, white)
				needsRedraw = false
			}
		}

		prevKey = key
	}
}
