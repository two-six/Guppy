package main

import (
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
	input := zzterm.NewInput()
	go refreshScreen(root)
	for {
		k, err := input.ReadKey(t)
		if err != nil {
			panic(err)
		}
		switch k.Type() {
		case zzterm.KeyLeft:
			tiling.SwitchFocus(root, true)
		case zzterm.KeyRight:
			tiling.SwitchFocus(root, false)
		case zzterm.KeyUp:
			leave, err := tiling.FindFocused(root)
			if err != nil {
				panic(err)
			}
			leave.NewChild(root, true)
		case zzterm.KeyDown:
			leave, err := tiling.FindFocused(root)
			if err != nil {
				panic(err)
			}
			leave.NewChild(root, false)
		case zzterm.KeyDelete:
			leave, err := tiling.FindFocused(root)
			if err != nil {
				panic(err)
			}
			leave.RemoveChild(root)
			leaves := tiling.GetLeaves(root)
			leaves[0].Content.IsFocused = true
		case zzterm.KeyESC, zzterm.KeyCtrlC:
			return
		}
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

func refreshScreen(root *tiling.TilingTile) {
	prevX, prevY := 0, 0
	for {
		x, y, err := term.GetSize()
		if err != nil {
			return
		}
		if x != prevX || y != prevY {
			prevX = x
			prevY = y
			tiling.RefreshSize(root, x, y)
			term.Clear()
			tiling.DrawBorders(root)
		}
	}
}
