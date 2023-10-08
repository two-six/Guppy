package tiles

import (
	"errors"
	"fmt"

	"github.com/fatih/color"

	"projects/twpsx/guppy/tiles/cursor"
	"projects/twpsx/guppy/tiles/term"
)

type tile struct {
	Parent                  *tile
	Children                []*tile
	posX, posY              int
	sizeX, sizeY            int
	cursorPosX, cursorPosY  int
	CanBeFocused, IsFocused bool
	vSplit                  bool
}

func NewRoot() (tile, error) {
	sX, sY, err := term.GetSize()
	return tile{
		Parent:       nil,
		Children:     make([]*tile, 0),
		posX:         0,
		posY:         0,
		sizeX:        sX,
		sizeY:        sY,
		CanBeFocused: true,
		IsFocused:    true,
		vSplit:       true,
	}, err
}

func (t *tile) NewChild(vSplit, canBeFocused bool) {
	var posX, posY, sizeX, sizeY int
	if len(t.Children) == 0 {
		posX = 0
		posY = 0
		sizeX = t.sizeX
		sizeY = t.sizeY
	} else {
		if t.vSplit {
			childWidth := t.sizeX / (len(t.Children) + 1)
			for i, c := range t.Children {
				c.posX = i * childWidth
				c.sizeX = childWidth
			}
			posX = len(t.Children) * childWidth
			posY = t.Children[0].posY
			sizeX = childWidth
			sizeY = t.Children[0].sizeY
			sizeX += t.sizeX - (childWidth * (len(t.Children) + 1))
		} else {
			childHeight := t.sizeY / (len(t.Children) + 1)
			for i, c := range t.Children {
				c.posY = i * childHeight
				c.sizeY = childHeight
			}
			posY = len(t.Children) * childHeight
			posX = t.Children[0].posX
			sizeY = childHeight
			sizeX = t.Children[0].sizeX
			sizeY += t.sizeY - (childHeight * (len(t.Children) + 1))
		}
	}
	t.Children = append(t.Children, &tile{
		Parent:       t,
		Children:     make([]*tile, 0),
		posX:         posX,
		posY:         posY,
		sizeX:        sizeX,
		sizeY:        sizeY,
		CanBeFocused: canBeFocused,
		IsFocused:    false,
		vSplit:       vSplit,
	})
}

func (t *tile) GetPos() (int, int) {
	return t.posX, t.posY
}

func (t *tile) GetSize() (int, int) {
	return t.sizeX, t.sizeY
}

func (t *tile) getRootPosition() (int, int) {
	if t.Parent != nil {
		parentX, parentY := t.Parent.getRootPosition()
		return parentX + t.posX, parentY + t.posY
	}
	return t.posX, t.posY
}

func (t *tile) getRootSize() (int, int) {
	if t.Parent != nil {
		return t.Parent.getRootSize()
	}
	return t.sizeX, t.sizeY
}

func (t *tile) GetCursorPos() (int, int) {
	return t.cursorPosX, t.cursorPosY
}

func (t *tile) DrawBorder() {
	x, y := t.getRootPosition()
	sx, sy := t.GetSize()
	cursor.MoveTo(x, y)
	c := color.New(color.FgBlack).Add(color.BgRed)
	for i := 0; i < sx; i++ {
		if t.IsFocused {
			c.Print("-")
		} else {
			print("-")
		}
	}
	cursor.MoveTo(x, y+sy-1)
	for i := 0; i < sx; i++ {
		if t.IsFocused {
			c.Print("-")
		} else {
			print("-")
		}
	}
	cursor.MoveTo(x, y+1)
	for i := 0; i < sy; i++ {
		if t.IsFocused {
			c.Print("|")
		} else {
			print("|")
		}
		cursor.MoveTo(x, y+1+i)
	}
	cursor.MoveTo(x+sx, y+1)
	for i := 0; i < sy; i++ {
		if t.IsFocused {
			c.Print("|")
		} else {
			print("|")
		}
		cursor.MoveTo(x+sx, y+1+i)
	}
}

func (t *tile) FindFocused() (*tile, error) {
	if t.IsFocused {
		return t, nil
	}
	for _, c := range t.Children {
		tmp, err := c.FindFocused()
		if err == nil {
			return tmp, nil
		}
	}
	return t, errors.New("no focused tiles")
}

func (t *tile) ChangeSplitDirection(vSplit bool) error {
	if len(t.Children) != 0 {
		return errors.New("cannot split a tile with children")
	}
	t.vSplit = vSplit
	return nil
}

func (t *tile) PrevFocus() error {
	return t.switchFocus(true)
}

func (t *tile) NextFocus() error {
	return t.switchFocus(false)
}

func (t *tile) switchFocus(prev bool) error {
	for t.Parent != nil {
		t = t.Parent
	}
	focused, err := t.FindFocused()
	if err != nil {
		return err
	}
	for i, c := range focused.Parent.Children {
		if c.IsFocused {
			if i == 0 && prev {
				if c.Parent == nil {
					return nil
				}
				_ = c.Parent
				// TODO
				return nil
			}
			if prev {
				tmp := c.Parent.Children[i-1]
				for len(tmp.Children) > 0 {
					tmp = tmp.Children[0]
				}
				tmp.IsFocused = true
			} else {
				tmp := c.Parent.Children[i+1]
				for len(tmp.Children) > 0 {
					tmp = tmp.Children[len(tmp.Children)-1]
				}
				tmp.IsFocused = true
			}
			c.IsFocused = false
			return nil
		}
	}
	return errors.New("no focused objects")
}

func (t *tile) ClearFocus() {
	if t.IsFocused {
		t.IsFocused = false
		return
	}
	for _, c := range t.Children {
		if c.IsFocused {
			c.IsFocused = false
			return
		}
		c.ClearFocus()
	}
}

func (t *tile) DrawCanBeFocusedTiles() {
	redBackground := color.New(color.FgBlack).Add(color.BgRed)
	whiteBackground := color.New(color.FgBlack).Add(color.BgHiWhite)
	for _, c := range t.Children {
		if c.CanBeFocused {
			x, y := c.getRootPosition()
			cursor.MoveTo(x, y)
			for i := 0; i < x+c.sizeX; i++ {
				if c.IsFocused {
					redBackground.Print("-")
				} else {
					whiteBackground.Print("-")
				}
			}
			cursor.MoveTo(x, y+c.sizeY-1)
			for i := 0; i < x+c.sizeX; i++ {
				if c.IsFocused {
					redBackground.Print("-")
				} else {
					whiteBackground.Print("-")
				}
			}
			cursor.MoveTo(x, y+1)
			for i := 1; i < y+c.sizeY-1; i++ {
				if c.IsFocused {
					redBackground.Print("|")
				} else {
					whiteBackground.Print("|")
				}
				cursor.MoveTo(x, y+1+i)

			}
			cursor.MoveTo(x+c.sizeX, y)
			for i := 1; i < y+c.sizeY-1; i++ {
				if c.IsFocused {
					redBackground.Print("|")
				} else {
					whiteBackground.Print("|")
				}
				cursor.MoveTo(x+c.sizeX, y+i)
			}
		} else {
			for _, cc := range c.Children {
				cc.DrawCanBeFocusedTiles()
			}
		}
	}
}

func (t *tile) Information() {
	x, y := t.getRootPosition()
	fmt.Println("Parent:", t.Parent, "\nChildren amount:", len(t.Children), "\nposX, posY:", t.posX, t.posY, "\nSizeX, SizeY:", t.sizeX, t.sizeY, "\nCursorPosX, CursorPosY:", t.cursorPosX, t.cursorPosY, "\nCanBeFocused, IsFocused", t.CanBeFocused, t.IsFocused, "\nVSplit:", t.vSplit, "\nreal pos x, y:", x, y)
}
