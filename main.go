package main

import (
	"os"

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
	// input := zzterm.NewInput()
	// shouldRefreshChan := make(chan bool, 1)
	shouldRefresh := false
	for {
		newSize, err := tiling.RefreshSize(root)
		if err != nil {
			panic(err)
		}
		if newSize || shouldRefresh {
			term.Clear()
			tiling.DrawBorders(root)
			cursor.MoveTo(30, 30)
			// printAllInformation(root)
		}
		// go readKeys(input, root, t, shouldRefreshChan)
		// select {
		// case msg := <-shouldRefreshChan:
		// 	shouldRefresh = msg
		// default:
		// 	shouldRefresh = false
		// }
	}
}

func printAllInformation(root *tiling.TilingTile) {
	root.PrintInformation()
	if root.Left != nil {
		printAllInformation(root.Left)
		printAllInformation(root.Right)

	}
}

func readKeys(input *zzterm.Input, root *tiling.TilingTile, t *pkgterm.Term, shouldRefreshChan chan<- bool) {
	for {
		k, err := input.ReadKey(t)
		if err != nil {
			panic(err)
		}
		switch k.Type() {
		case zzterm.KeyLeft:
			tiling.SwitchFocus(root, true)
			shouldRefreshChan <- true
		case zzterm.KeyRight:
			tiling.SwitchFocus(root, false)
			shouldRefreshChan <- true
		case zzterm.KeyESC, zzterm.KeyCtrlC:
			os.Exit(0)
		}
	}
}
