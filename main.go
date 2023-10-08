package main

import (
	"fmt"

	"git.sr.ht/~mna/zzterm"
	"github.com/pkg/term"

	"projects/twpsx/guppy/tiles"
	"projects/twpsx/guppy/tiles/cursor"
	tterm "projects/twpsx/guppy/tiles/term" // tterm
)

func main() {
	t, err := term.Open("/dev/tty", term.RawMode)
	if err != nil {
		panic(err)
	}
	defer t.Restore()

	root, err := tiles.NewRoot()
	if err != nil {
		panic(err)
	}
	tterm.Clear()
	cursor.MoveDownBeginning(1)

	input := zzterm.NewInput()

mainLoop:
	for {
		k, err := input.ReadKey(t)
		if err != nil {
			panic(err)
		}

		switch k.Type() {
		case zzterm.KeyRune:
			if k.Rune() == 'c' {
				tterm.Clear()
				continue
			}
			print(string(k.Rune()))
		case zzterm.KeyEnter:
			cursor.MoveDownBeginning(1)
		case zzterm.KeyUp:
			cursor.MoveUp(1)
		case zzterm.KeyRight:
			root.NewChild(true, true)
			tterm.Clear()
			for _, c := range root.Children {
				c.DrawBorder()
			}

		case zzterm.KeyLeft:
			cursor.MoveLeft(1)
		case zzterm.KeyESC, zzterm.KeyCtrlC:
			break mainLoop
		default:
			x, y, err := tterm.GetSize()
			if err != nil {
				panic(err)
			}
			fmt.Printf("Window Size: %dx%d\033[1E", x, y)
		}
	}
}
