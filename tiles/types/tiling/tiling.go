package tiling

import (
	"errors"
	"fmt"

	"projects/twpsx/guppy/tiles"

	"projects/twpsx/guppy/tiles/draw"
	"projects/twpsx/guppy/tiles/term"

	"github.com/google/uuid"
)

type TilingTile struct {
	Left    *TilingTile
	Right   *TilingTile
	Content *tiles.Tile
	id      string
}

func (t *TilingTile) PrintInformation() {
	fmt.Println("Tile", t.id, "IsFocused", t.Content.IsFocused, "pos", t.Content.PosX, t.Content.PosY, "size", t.Content.SizeX, t.Content.SizeY)
}

func NewRoot() (*TilingTile, error) {
	sx, sy, err := term.GetSize()
	if err != nil {
		return nil, err
	}
	return &TilingTile{
		id:    uuid.NewString(),
		Left:  nil,
		Right: nil,
		Content: &tiles.Tile{
			IsFocused: false,
			SizeX:     sx,
			SizeY:     sy,
			PosX:      0,
			PosY:      0,
		},
	}, nil
}

func (t *TilingTile) RemoveChild(root *TilingTile) error {
	parent, err := findParent(root, t.id)
	if err != nil {
		return err
	}
	parent.Left = nil
	parent.Right = nil
	parent.Content.IsFocused = true
	return nil
}

func clearFocused(root *TilingTile) {
	root.Content.IsFocused = false
	if root.Left != nil {
		clearFocused(root.Left)
		clearFocused(root.Right)
	}
}

func (t *TilingTile) NewChild(root *TilingTile, vSplit bool) error {
	if t.Left != nil && t.Right != nil {
		return errors.New("can only create children on leaves")
	}
	clearFocused(root)
	var sx, sy int
	var px, py int
	if vSplit {
		sx = t.Content.SizeX / 2
		sy = t.Content.SizeY
		if t.Content.SizeX%2 == 1 {
			sx += 1
		}
		px = t.Content.PosX + sx
		py = t.Content.PosY
	} else {
		sx = t.Content.SizeX
		sy = t.Content.SizeY / 2
		if t.Content.SizeY%2 == 1 {
			sy += 1
		}
		px = t.Content.PosX
		py = t.Content.PosY + sy
	}
	t.Left = &TilingTile{
		id:    uuid.NewString(),
		Left:  nil,
		Right: nil,
		Content: &tiles.Tile{
			IsFocused: false,
			SizeX:     sx,
			SizeY:     sy,
			PosX:      t.Content.PosX,
			PosY:      t.Content.PosY,
		},
	}
	t.Right = &TilingTile{
		id:    uuid.NewString(),
		Left:  nil,
		Right: nil,
		Content: &tiles.Tile{
			IsFocused: true,
			SizeX:     sx,
			SizeY:     sy,
			PosX:      px,
			PosY:      py,
		},
	}
	return nil
}

func GetLeaves(root *TilingTile) []*TilingTile {
	var leaves []*TilingTile
	if root == nil {
		return leaves
	}
	if root.Left == nil {
		leaves = append(leaves, root)
	}
	leaves = append(leaves, GetLeaves(root.Left)...)
	leaves = append(leaves, GetLeaves(root.Right)...)
	return leaves
}

func findParent(root *TilingTile, id string) (*TilingTile, error) {
	if root == nil {
		return nil, errors.New("invalid root")
	}
	if root.Left == nil {
		return nil, errors.New("no parent node has a child with provided id")
	}
	if root.Left.id == id || root.Right.id == id {
		return root, nil
	}
	ls, err := findParent(root.Left, id)
	if err == nil {
		return ls, nil
	}
	rs, err := findParent(root.Right, id)
	if err == nil {
		return rs, nil
	}
	return nil, errors.New("no parent node has a child with provided id")
}

func FindFocused(root *TilingTile) (*TilingTile, error) {
	if root.Content.IsFocused {
		return root, nil
	}
	if root.Left == nil {
		return nil, errors.New("no focused Tile found")
	}
	ls, err := FindFocused(root.Left)
	if err == nil {
		return ls, nil
	}
	return FindFocused(root.Right)
}

func DrawBorders(root *TilingTile) {
	leaves := GetLeaves(root)
	for _, l := range leaves {
		draw.DrawBorder(l.Content)
	}
	fc, err := FindFocused(root)
	if err != nil {
		return
	}
	draw.DrawBorder(fc.Content)
}
