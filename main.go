package main

import (
	"time"

	pkgterm "github.com/pkg/term"

	"git.sr.ht/~mna/zzterm"

	"projects/twpsx/guppy/tiles/cursor"
	"projects/twpsx/guppy/tiles/term"
	"projects/twpsx/guppy/tiles/tiling"
)

func main() {
	t, err := pkgterm.Open("/dev/tty", pkgterm.RawMode)
	if err != nil {
		panic(err)
	}
	defer t.Restore()
	defer term.Clear()
	defer cursor.MoveTo(0, 0)
	root, err := tiling.NewRoot()
	if err != nil {
		panic(err)
	}
	root.NewChild(root, true)
	root.Left.NewChild(root, false)
	root.Left.Right.NewChild(root, false)
	root.Right.NewChild(root, true)
	if err = root.Left.Left.Resize(root, 10); err != nil {
		panic(err)
	}
	input := zzterm.NewInput()
	go refreshScreen(time.Second/20, root)
	for {
		k, err := input.ReadKey(t)
		if err != nil {
			panic(err)
		}
		switch k.Type() {
		case zzterm.KeyLeft:
			err := tiling.SwitchFocus(root, true)
			if err != nil {
				panic(err)
			}
		case zzterm.KeyRight:
			err := tiling.SwitchFocus(root, false)
			if err != nil {
				panic(err)
			}
		case zzterm.KeyESC, zzterm.KeyCtrlC:
			return
		}
		tiling.RefreshSize(root)
		term.Clear()
		tiling.DrawBorders(root)

	}
}

func printAllInformation(root *tiling.TilingTile) {
	root.PrintInformation()
	if root.Left != nil {
		printAllInformation(root.Left)
		printAllInformation(root.Right)

	}
}

func refreshScreen(t time.Duration, root *tiling.TilingTile) {
	prevX, prevY := 0, 0
	for {
		x, y, err := term.GetSize()
		if err != nil {
			return
		}
		if x != prevX || y != prevY {
			prevX = x
			prevY = y
			tiling.RefreshSize(root)
			term.Clear()
			tiling.DrawBorders(root)
		}
		time.Sleep(t)
	}
}
