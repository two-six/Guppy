package main

import (
	"projects/twpsx/guppy/tiles/cursor"
	"projects/twpsx/guppy/tiles/term"
	"projects/twpsx/guppy/tiles/types/tiling"
)

func main() {
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
	term.Clear()
	tiling.DrawBorders(root)
	cursor.MoveTo(0, 10)
	for {
		newSize, err := tiling.RefreshSize(root)
		if err != nil {
			panic(err)
		}
		if newSize {
			term.Clear()
			tiling.DrawBorders(root)
			cursor.MoveTo(0, 10)
			printAllInformation(root)
		}
	}
}

func printAllInformation(root *tiling.TilingTile) {
	root.PrintInformation()
	if root.Left != nil {
		printAllInformation(root.Left)
		printAllInformation(root.Right)

	}
}
