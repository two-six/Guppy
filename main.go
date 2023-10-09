package main

import (
	"projects/twpsx/guppy/tiles"
	"projects/twpsx/guppy/tiles/cursor"
	"projects/twpsx/guppy/tiles/draw"
	"projects/twpsx/guppy/tiles/term"
)

func main() {
	root, err := tiles.NewRoot()
	if err != nil {
		panic(err)
	}
	root.NewChild(root, true)
	root.Left.NewChild(root, false)
	root.Left.Right.NewChild(root, false)
  root.Right.NewChild(root, true)
	term.Clear()
	draw.DrawBorders(root)
	cursor.MoveTo(0, 10)
}
