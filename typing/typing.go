package typing

type TypingArea struct {
	Content    []string
	PosX       int
	PosY       int
	WindowPosX int
	WindowPosY int
	SizeX      int
	SizeY      int
	CursorPosX int
	CursorPosY int
}

func New(posX int, posY int, sizeX int, sizeY int) *TypingArea {
	return &TypingArea{
		PosX:       posX,
		PosY:       posY,
		SizeX:      sizeX,
		SizeY:      sizeY,
		WindowPosX: 0,
		WindowPosY: 0,
		Content:    make([]string, 1),
		CursorPosX: 0,
		CursorPosY: 0,
	}
}

func (area *TypingArea) Write(str string) {
	area.checkCursorPos()
	area.Content[area.CursorPosY] = area.Content[area.CursorPosY][:area.CursorPosX] + str
	area.CursorPosX += len(str)
	area.AlignToSize(area.CursorPosY)
}

func (area *TypingArea) InsertNewline() {
	area.checkCursorPos()
	area.Content = append(area.Content[:area.CursorPosY+1], append([]string{""}, area.Content[area.CursorPosY+1:]...)...)
	area.CursorPosY++
	area.CursorPosX = 0
}

func (area *TypingArea) checkCursorPos() {
	if area.CursorPosY >= len(area.Content) {
		lastLine := area.Content[len(area.Content)-1]
		area.CursorPosY = len(area.Content) - 1
		area.CursorPosX = len(lastLine) - 1
	}
	if area.CursorPosX > len(area.Content[area.CursorPosY]) {
		area.CursorPosX = len(area.Content[area.CursorPosY]) - 1
	}
}

func (area *TypingArea) RemoveLastCharacter() {
	if area.Content[0] == "" {
		return
	}
	if area.Content[area.CursorPosY] == "" {
		area.CursorPosY--
	}
	area.Content[area.CursorPosY] = area.Content[area.CursorPosY][:len(area.Content[area.CursorPosY])-1]
	if area.CursorPosX == 0 {
		area.CursorPosX = len(area.Content[area.CursorPosY])
	} else {
		area.CursorPosX--
	}
}

func (area *TypingArea) Print() string {
	var bottomEdge int
	if area.WindowPosY+area.SizeY > len(area.Content) {
		bottomEdge = len(area.Content)
	} else {
		bottomEdge = area.WindowPosY + area.SizeY
	}
	result := ""
	for _, line := range area.Content[area.WindowPosY:bottomEdge] {
		result += line + "\n"
	}
	return result
}

func (area *TypingArea) AlignToSize(row int) {
	if len(area.Content[row]) > area.SizeX {
		tmpString := area.Content[row]
		area.Content[row] = tmpString[:area.SizeX]
		area.Content = append(area.Content, tmpString[area.SizeX:])
		area.CursorPosY++
	}
	if row < len(area.Content)-1 {
		area.AlignToSize(row + 1)
	}
}
