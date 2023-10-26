package engine

import (
	"errors"
	"strings"

	"projects/twpsx/guppy/engine/commands"
	"projects/twpsx/guppy/term"
	"projects/twpsx/guppy/tiles/draw"
	"projects/twpsx/guppy/tiles/tiling"
	"projects/twpsx/guppy/typing"
)

type Engine struct {
	rootWindow *tiling.TilingTile
	writers    []*typing.TypingArea
}

func New() (*Engine, error) {
	termSizeX, termSizeY, err := term.GetSize()
	if err != nil {
		return nil, err
	}
	rootWindow, err := tiling.NewRoot(termSizeX, termSizeY)
	if err != nil {
		return nil, err
	}
	return &Engine{
		rootWindow,
		make([]*typing.TypingArea, 0),
	}, nil
}

func (engine *Engine) RefreshTerminalSize() error {
	termSizeX, termSizeY, err := term.GetSize()
	if err != nil {
		return err
	}
	return tiling.RefreshSize(engine.rootWindow, termSizeX, termSizeY)
}

func (engine *Engine) Draw() {
	leaves := tiling.GetLeaves(engine.rootWindow)
	for _, leaf := range leaves {
		draw.DrawBorder(leaf.Content)
	}
}

func (engine *Engine) RunCommand(command string) error {
	var focused *tiling.TilingTile
	if engine.rootWindow.Left == nil {
		focused = engine.rootWindow
	} else {
		var err error
		focused, err = tiling.FindFocused(engine.rootWindow)
		if err != nil {
			return err
		}
	}
	commandParts := strings.Split(command, " ")
	switch commandParts[0] {
	case "vSplit":
		return commands.VSplit(engine.rootWindow, focused)
	case "hSplit":
		return commands.HSplit(engine.rootWindow, focused)
	default:
		return errors.New("invalid command")
	}
}

func (engine *Engine) SwitchTileFocusLeft() {
	tiling.SwitchFocus(engine.rootWindow, true)
}

func (engine *Engine) SwitchTileFocusRight() {
	tiling.SwitchFocus(engine.rootWindow, false)
}
