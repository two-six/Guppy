package tiles

import (
	"errors"
	"fmt"

	"projects/twpsx/guppy/tiles/term"

	"github.com/google/uuid"
)

type Tile struct {
	Left         *Tile
	Right        *Tile
	id           string
	IsFocused    bool
	sizeX, sizeY int
	posX, posY   int
}

func (t *Tile) PrintInformation() {
	fmt.Println("Tile", t.id, "IsFocused", t.IsFocused, "pos", t.posX, t.posY, "size", t.sizeX, t.sizeY)
}

func NewRoot() (*Tile, error) {
	x, y, err := term.GetSize()
	if err != nil {
		return nil, err
	}
	return &Tile{
		id:        uuid.NewString(),
		Left:      nil,
		Right:     nil,
		IsFocused: true,
		sizeX:     x,
		sizeY:     y,
		posX:      0,
		posY:      0,
	}, nil
}

func (t *Tile) RemoveChild(root *Tile) error {
	parent, err := findParent(root, t.id)
	if err != nil {
		return err
	}
	parent.Left = nil
	parent.Right = nil
	parent.IsFocused = true
	return nil
}

func clearFocused(root *Tile) {
	root.IsFocused = false
	if root.Left != nil {
		clearFocused(root.Left)
		clearFocused(root.Right)
	}
}

func (t *Tile) NewChild(root *Tile, vSplit bool) error {
	if t.Left != nil && t.Right != nil {
		return errors.New("can only create children on leaves")
	}
	clearFocused(root)
	var sx, sy int
	var px, py int
	if vSplit {
		sx = t.sizeX / 2
		sy = t.sizeY
		if t.sizeX%2 == 1 {
			sx += 1
		}
		px = t.posX + sx
		py = t.posY
	} else {
		sx = t.sizeX
		sy = t.sizeY / 2
		if t.sizeY%2 == 1 {
			sy += 1
		}
		px = t.posX
		py = t.posY + sy
	}
	t.Left = &Tile{
		id:        uuid.NewString(),
		Left:      nil,
		Right:     nil,
		IsFocused: false,
		sizeX:     sx,
		sizeY:     sy,
		posX:      t.posX,
		posY:      t.posY,
	}
	t.Right = &Tile{
		id:        uuid.NewString(),
		Left:      nil,
		Right:     nil,
		IsFocused: true,
		sizeX:     sx,
		sizeY:     sy,
		posX:      px,
		posY:      py,
	}
	return nil
}

func GetLeaves(root *Tile) []*Tile {
	var leaves []*Tile
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

func findParent(root *Tile, id string) (*Tile, error) {
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

func FindFocused(root *Tile) (*Tile, error) {
	if root.IsFocused {
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

func (t *Tile) GetPosition() (int, int) {
	return t.posX, t.posY
}

func (t *Tile) GetSize() (int, int) {
	return t.sizeX, t.sizeY
}
