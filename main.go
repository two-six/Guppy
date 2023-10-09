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
	term.Clear()
	tiling.DrawBorders(root)
	cursor.MoveTo(0, 10)
}
