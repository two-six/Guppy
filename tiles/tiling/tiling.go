package tiling

import (
	"errors"
	"fmt"
	"math"

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

func RefreshSize(root *TilingTile) (bool, error) {
	sx, sy, err := term.GetSize()
	if err != nil {
		return false, err
	}
	if sx == root.Content.SizeX && sy == root.Content.SizeY {
		return false, nil
	}
	root.Content.SizeX = sx
	root.Content.SizeY = sy
	return true, refreshSizes(root)
}

func refreshSizes(parent *TilingTile) error {
	if parent.Left == nil {
		return nil
	}
	vSplit, err := isVSplit(parent)
	if err != nil {
		return err
	}
	refreshChildrenSize(parent, vSplit)
	refreshChildrenPos(parent, vSplit)
	if err = refreshSizes(parent.Left); err != nil {
		return err
	}
	return refreshSizes(parent.Right)
}

func refreshChildrenSize(parent *TilingTile, vSplit bool) {
	if vSplit {
		parent.Left.Content.SizeY = parent.Content.SizeY
		parent.Right.Content.SizeY = parent.Content.SizeY
		oldSX := parent.Left.Content.SizeX + parent.Right.Content.SizeX
		leftPercentage := math.Floor(float64(parent.Left.Content.SizeX) / float64(oldSX) * 100)
		rightPercentage := math.Floor(float64(parent.Right.Content.SizeX) / float64(oldSX) * 100)
		if leftPercentage > rightPercentage {
			parent.Left.Content.SizeX = int(math.Floor(float64(parent.Content.SizeX) / 100 * leftPercentage))
			parent.Right.Content.SizeX = parent.Content.SizeX - parent.Left.Content.SizeX
		} else {
			parent.Right.Content.SizeX = int(math.Floor(float64(parent.Content.SizeX) / 100 * leftPercentage))
			parent.Left.Content.SizeX = parent.Content.SizeX - parent.Right.Content.SizeX

		}
	} else {
		parent.Left.Content.SizeX = parent.Content.SizeX
		parent.Right.Content.SizeX = parent.Content.SizeX
		oldSY := parent.Left.Content.SizeY + parent.Right.Content.SizeY
		leftPercentage := math.Floor(float64(parent.Left.Content.SizeY) / float64(oldSY) * 100)
		rightPercentage := math.Floor(float64(parent.Right.Content.SizeY) / float64(oldSY) * 100)
		if leftPercentage > rightPercentage {
			parent.Left.Content.SizeY = int(math.Floor(float64(parent.Content.SizeY) / 100 * leftPercentage))
			parent.Right.Content.SizeY = parent.Content.SizeY - parent.Left.Content.SizeY
		} else {
			parent.Right.Content.SizeY = int(math.Floor(float64(parent.Content.SizeY) / 100 * leftPercentage))
			parent.Left.Content.SizeY = parent.Content.SizeY - parent.Right.Content.SizeY

		}
	}
}

func refreshChildrenPos(parent *TilingTile, vSplit bool) {
	if vSplit {
		parent.Left.Content.PosX = parent.Content.PosX
		parent.Right.Content.PosX = parent.Content.PosX + parent.Left.Content.SizeX
	} else {
		parent.Left.Content.PosY = parent.Content.PosY
		parent.Right.Content.PosY = parent.Content.PosY + parent.Left.Content.SizeY
	}
}

func isVSplit(parent *TilingTile) (bool, error) {
	if parent.Left == nil {
		return false, errors.New("parent has no children")
	}
	return parent.Left.Content.PosX != parent.Right.Content.PosX, nil
}

func clearFocused(root *TilingTile) {
	root.Content.IsFocused = false
	if root.Left != nil {
		clearFocused(root.Left)
		clearFocused(root.Right)
	}
}

func calculateProportionsOfNewChild(t *TilingTile, vSplit bool) (int, int, int, int) {
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
	return sx, sy, px, py
}

func (t *TilingTile) NewChild(root *TilingTile, vSplit bool) error {
	if t.Left != nil && t.Right != nil {
		return errors.New("can only create children on leaves")
	}
	sx, sy, px, py := calculateProportionsOfNewChild(t, vSplit)
	clearFocused(root)
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
	if ls, err := findParent(root.Left, id); err == nil {
		return ls, nil
	} else if rs, err := findParent(root.Right, id); err == nil {
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
	if ls, err := FindFocused(root.Left); err == nil {
		return ls, nil
	}
	return FindFocused(root.Right)
}

func DrawBorders(root *TilingTile) {
	leaves := GetLeaves(root)
	for _, l := range leaves {
		draw.DrawBorder(l.Content)
	}
	if fc, err := FindFocused(root); err == nil {
		draw.DrawBorder(fc.Content)
	}
}

func (t *TilingTile) Resize(root *TilingTile, n int) error {
	parent, err := findParent(root, t.id)
	if err != nil {
		return err
	}
	vSplit, err := isVSplit(parent)
	if err != nil {
		return err
	}
	var choosen, other *TilingTile
	if parent.Left.id == t.id {
		choosen = parent.Left
		other = parent.Right
	} else {
		choosen = parent.Right
		other = parent.Left
	}
	if vSplit {
		choosen.Content.SizeX += n
		if choosen.Content.SizeX > parent.Content.SizeX {
			choosen.Content.SizeX = parent.Content.SizeX - 1
		} else if choosen.Content.SizeX == 0 {
			choosen.Content.SizeX = 1
		}
		other.Content.SizeX = parent.Content.SizeX - choosen.Content.SizeX
		parent.Right.Content.PosX = parent.Content.PosX + parent.Left.Content.SizeX
	} else {
		choosen.Content.SizeY += n
		if choosen.Content.SizeY > parent.Content.SizeY {
			choosen.Content.SizeY = parent.Content.SizeY - 1
		} else if choosen.Content.SizeY <= 0 {
			choosen.Content.SizeY = 1
		}
		other.Content.SizeY = parent.Content.SizeY - choosen.Content.SizeY
		parent.Right.Content.PosY = parent.Content.PosY + parent.Left.Content.SizeY
	}
	refreshSizes(choosen)
	refreshSizes(other)
	RefreshSize(root)

	return nil
}
