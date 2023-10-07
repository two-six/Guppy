package cursor

import (
	"strconv"
)

func MoveTo(x, y int) {
	xStr, yStr := strconv.Itoa(x), strconv.Itoa(y)
	print("\033[" + yStr + ";" + xStr + "H")
}

func MoveUp(n int) {
	nStr := strconv.Itoa(n)
	print("\033[" + nStr + "A")
}

func MoveDown(n int) {
	nStr := strconv.Itoa(n)
	print("\033[" + nStr + "B")
}

func MoveRight(n int) {
	nStr := strconv.Itoa(n)
	print("\033[" + nStr + "C")
}

func MoveLeft(n int) {
	nStr := strconv.Itoa(n)
	print("\033[" + nStr + "D")
}

func MoveDownBeginning(n int) {
	nStr := strconv.Itoa(n)
	print("\033[" + nStr + "E")
}

func MoveUpBeginning(n int) {
	nStr := strconv.Itoa(n)
	print("\033[" + nStr + "E")
}

func MoveToColumn(n int) {
	nStr := strconv.Itoa(n)
	print("\033[" + nStr + "G")
}
