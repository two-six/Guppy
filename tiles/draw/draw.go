package draw

import (
	"github.com/fatih/color"
	"projects/twpsx/guppy/tiles"
	"projects/twpsx/guppy/tiles/cursor"
)

func DrawBorder(leaf *tiles.Tile) {
	var c color.Color
	if leaf.IsFocused {
		c = *color.New(color.FgHiRed)
	} else {
		c = *color.New(color.FgBlack)
	}
	px, py := leaf.PosX, leaf.PosY
	sx, sy := leaf.SizeX, leaf.SizeY
	cursor.MoveTo(px, py)
	for i := 0; i < sx; i++ {
		c.Print("█")
	}
	cursor.MoveTo(px, py+sy)
	for i := 0; i < sx; i++ {
		c.Print("█")
	}
	cursor.MoveTo(px, py)
	for i := 0; i < sy; i++ {
		c.Print("█")
		cursor.MoveTo(px, py+1+i)
	}
	cursor.MoveTo(px+sx, py)
	for i := 0; i <= sy; i++ {
		c.Print("█")
		cursor.MoveTo(px+sx, py+1+i)
	}
}
