package engine

import (
	"projects/twpsx/guppy/term"
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
