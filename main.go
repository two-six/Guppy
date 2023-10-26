package main

import (
	"strings"

	pkgterm "github.com/pkg/term"

	"git.sr.ht/~mna/zzterm"

	"projects/twpsx/guppy/tiles/cursor"
	"projects/twpsx/guppy/tiles/term"
	"projects/twpsx/guppy/tiles/tiling"
	"projects/twpsx/guppy/typing"
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
	leaf, err := tiling.FindFocused(root)
	if err != nil {
		panic(err)
	}
	writer := typing.New(leaf.Content.PosX, leaf.Content.PosY, leaf.Content.SizeX, leaf.Content.SizeY)
	go refreshScreen(root, writer)
	for {
		k, err := input.ReadKey(t)
		if err != nil {
			panic(err)
		}
		switch k.Type() {
		case zzterm.KeyLeft:
			tiling.SwitchFocus(root, true)
			focused, err := tiling.FindFocused(root)
			if err != nil {
				panic(err)
			}
			refreshWriter(focused, writer)
			term.Clear()
			tiling.DrawBorders(root)
			printWriter(writer)
			cursor.MoveTo(writer.PosX+writer.CursorPosX, writer.PosY+writer.CursorPosY+1)
		case zzterm.KeyRight:
			tiling.SwitchFocus(root, false)
			focused, err := tiling.FindFocused(root)
			if err != nil {
				panic(err)
			}
			refreshWriter(focused, writer)
			term.Clear()
			tiling.DrawBorders(root)
			printWriter(writer)
			cursor.MoveTo(writer.PosX+writer.CursorPosX, writer.PosY+writer.CursorPosY+1)
		case zzterm.KeyUp:
			leaf, err := tiling.FindFocused(root)
			if err != nil {
				panic(err)
			}
			leaf.NewChild(root, true)
			focused, err := tiling.FindFocused(root)
			if err != nil {
				panic(err)
			}
			refreshWriter(focused, writer)
			term.Clear()
			tiling.DrawBorders(root)
			printWriter(writer)
			cursor.MoveTo(writer.PosX+writer.CursorPosX, writer.PosY+writer.CursorPosY+1)
		case zzterm.KeyDown:
			leaf, err := tiling.FindFocused(root)
			if err != nil {
				panic(err)
			}
			leaf.NewChild(root, false)
			focused, err := tiling.FindFocused(root)
			if err != nil {
				panic(err)
			}
			refreshWriter(focused, writer)
			term.Clear()
			tiling.DrawBorders(root)
			printWriter(writer)
			cursor.MoveTo(writer.PosX+writer.CursorPosX, writer.PosY+writer.CursorPosY+1)
		case zzterm.KeyDelete:
			leaf, err := tiling.FindFocused(root)
			if err != nil {
				panic(err)
			}
			leaf.RemoveChild(root)
			leaves := tiling.GetLeaves(root)
			leaves[0].Content.IsFocused = true
			refreshWriter(leaves[0], writer)
			printWriter(writer)
			cursor.MoveTo(writer.PosX+writer.CursorPosX, writer.PosY+writer.CursorPosY+1)
		case zzterm.KeyRune:
			switch k.Rune() {
			case 'x':
				writer.RemoveLastCharacter()
			default:
				writer.Write(string(k.Rune()))
			}
			printWriter(writer)
			cursor.MoveTo(writer.PosX+writer.CursorPosX, writer.PosY+writer.CursorPosY+1)
		case zzterm.KeyEnter:
			writer.InsertNewline()
			printWriter(writer)
			cursor.MoveTo(writer.PosX+writer.CursorPosX, writer.PosY+writer.CursorPosY+1)
		case zzterm.KeyESC, zzterm.KeyCtrlC:
			return
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

func refreshScreen(root *tiling.TilingTile, writer *typing.TypingArea) {
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
			cursor.MoveTo(writer.PosX, writer.PosY)
			writer.Print()
		}
	}
}

func printWriter(writer *typing.TypingArea) {
	str := writer.Print()
	cursor.MoveTo(writer.PosX, writer.PosY)
	for i, line := range strings.Split(str, "\n") {
		println(line)
		cursor.MoveTo(writer.PosX, writer.PosY+i+2)
	}
}

func refreshWriter(leaf *tiling.TilingTile, writer *typing.TypingArea) {
	writer.PosX = leaf.Content.PosX
	writer.PosY = leaf.Content.PosY
	writer.SizeX = leaf.Content.SizeX
	writer.SizeY = leaf.Content.SizeY
	writer.AlignToSize(0)
}
