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
			switch k.Rune() {
			case 'c':
				tterm.Clear()
			case 'w':
				err := root.NextFocus()
				if err != nil {
					panic(err)
				}
				tterm.Clear()
				root.DrawCanBeFocusedTiles()
			case 'r':
				err := root.PrevFocus()
				if err != nil {
					panic(err)
				}
				tterm.Clear()
				root.DrawCanBeFocusedTiles()
			case 'a':
				focused, err := root.FindFocused()
				if err != nil {
					panic(err)
				}
				root.ClearFocus()
				err = focused.ChangeSplitDirection(true)
				if err != nil {
					panic(err)
				}
				focused.CanBeFocused = false
				focused.NewChild(false, true)
				focused.NewChild(false, true)
				focused.Children[1].IsFocused = true
				tterm.Clear()
				root.DrawCanBeFocusedTiles()
				cursor.MoveTo(0, 0)
				focused.Children[0].Information()
				cursor.MoveTo(0, 10)
				focused.Children[1].Information()
			case 's':
				focused, err := root.FindFocused()
				if err != nil {
					panic(err)
				}
				root.ClearFocus()
				err = focused.ChangeSplitDirection(false)
				if err != nil {
					panic(err)
				}
				focused.CanBeFocused = false
				focused.NewChild(false, true)
				focused.NewChild(false, true)
				focused.Children[1].IsFocused = true
				tterm.Clear()
				root.DrawCanBeFocusedTiles()
				cursor.MoveTo(0, 0)
				focused.Children[0].Information()
				cursor.MoveTo(0, 10)
				focused.Children[1].Information()

			default:
				print(string(k.Rune()))
			}
		case zzterm.KeyEnter:
			cursor.MoveDownBeginning(1)
		case zzterm.KeyUp:
			cursor.MoveUp(1)
		case zzterm.KeyRight:
			cursor.MoveRight(1)
		case zzterm.KeyLeft:
			cursor.MoveLeft(1)
		case zzterm.KeyESC, zzterm.KeyCtrlC:
			break mainLoop
		case zzterm.KeyDown:
			cursor.MoveDown(1)

		default:
			x, y, err := tterm.GetSize()
			if err != nil {
				panic(err)
			}
			fmt.Printf("Window Size: %dx%d\033[1E", x, y)
		}
	}
}
