package commands

import (
	"projects/twpsx/guppy/tiles/tiling"
)

func VSplit(root, focused *tiling.TilingTile) error {
	return focused.NewChild(root, true)
}

func HSplit(root, focused *tiling.TilingTile) error {
	return focused.NewChild(root, false)
}
