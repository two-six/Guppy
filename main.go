package main

import (
	"projects/twpsx/guppy/tiles"
	"projects/twpsx/guppy/tiles/cursor"
	"projects/twpsx/guppy/tiles/draw"
	"projects/twpsx/guppy/tiles/term"
	"projects/twpsx/guppy/tiles/tiling"
)

func main() {
	floating := tiles.Tile{
		IsFocused: true,
		PosX:      10,
		PosY:      10,
		SizeX:     20,
		SizeY:     10,
	}
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
	for {
		newSize, err := tiling.RefreshSize(root)
		if err != nil {
			panic(err)
		}
		if newSize {
			term.Clear()
			tiling.DrawBorders(root)
			draw.DrawBorder(&floating)
			cursor.MoveTo(30, 30)
			// printAllInformation(root)
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
