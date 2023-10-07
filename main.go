package main

import (
	_ "git.sr.ht/~mna/zzterm"
	_ "github.com/pkg/term"

	"projects/twpsx/guppy/tiles"
	"projects/twpsx/guppy/tiles/cursor"
	"projects/twpsx/guppy/tiles/term" // tterm
)

func main() {
	// t, err := term.Open("/dev/tty", term.RawMode)
	// if err != nil {
	// 	panic(err)
	// }
	// defer t.Restore()

	// input := zzterm.NewInput()

	root, err := tiles.NewRoot()
	if err != nil {
		panic(err)
	}
	term.Clear()
	root.NewChild(false, false)
	root.NewChild(false, false)
	root.NewChild(true, false)
	// root.Children[0].DrawBorder()
	// root.Children[2].DrawBorder()
	root.Children[0].NewChild(true, false)
	root.Children[0].NewChild(true, false)
	root.DrawBorder()
	root.Children[0].Children[0].DrawBorder()
	root.Children[2].DrawBorder()
	cursor.MoveDownBeginning(1)

	// mainLoop:
	//
	//	for {
	//		k, err := input.ReadKey(t)
	//		if err != nil {
	//			panic(err)
	//		}
	//
	//		switch k.Type() {
	//		case zzterm.KeyRune:
	//			if k.Rune() == 'a' {
	//				cursor.MoveTo(20, 30)
	//				continue
	//			}
	//			if k.Rune() == 'c' {
	//				tterm.Clear()
	//				continue
	//			}
	//			if k.Rune() == '0' {
	//				cursor.MoveTo(0, 0)
	//				continue
	//			}
	//			print(string(k.Rune()))
	//		case zzterm.KeyEnter:
	//			cursor.MoveDownBeginning(1)
	//		case zzterm.KeyUp:
	//			cursor.MoveUp(1)
	//		case zzterm.KeyRight:
	//			cursor.MoveRight(1)
	//		case zzterm.KeyDown:
	//			cursor.MoveDown(1)
	//		case zzterm.KeyLeft:
	//			cursor.MoveLeft(1)
	//		case zzterm.KeyESC, zzterm.KeyCtrlC:
	//			break mainLoop
	//		default:
	//			x, y, err := tterm.GetSize()
	//			if err != nil {
	//				panic(err)
	//			}
	//			fmt.Printf("Window Size: %dx%d\033[1E", x, y)
	//		}
	//	}
}
