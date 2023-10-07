package statusbar

import (
	"time"

	"github.com/gookit/color"
)

type statusbar struct {
	widgets          []string
	bgColor, fgColor color.Color
	side             rune
	SizeX, SizeY     int
}

func New(widgets []string, side rune, sizeX, sizeY int) statusbar {
	return statusbar{
		bgColor: color.White,
		fgColor: color.Black,
		widgets: widgets,
		side:    side,
	}
}

func (sb *statusbar) ToString() string {
	var result string
	for i := 0; i < sb.SizeX; i++ {
		result += " "
	}
	result, err := addClock(result[0:])
	if err != nil {
		panic(err)
	}

	return result
}

func addClock(sb string) (string, error) {
	return time.Now().Format("15:04:05 PM"), nil
}
