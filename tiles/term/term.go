package term

import (
	"os"

	"golang.org/x/term"
)

func GetSize() (int, int, error) {
	x, y, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return x, y, err
	}
	return x, y, nil
}

func Clear() {
	print("\033[2J")
}
