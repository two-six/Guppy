package typing

type TypingArea struct {
	Content       []string
	PosX          int
	PosY          int
	WindowPosX    int
	WindowPosY    int
	SizeX         int
	SizeY         int
	CurrentRow    int
	CurrentColumn int
}

func New(posX int, posY int, sizeX int, sizeY int) *TypingArea {
	return &TypingArea{
		PosX:          posX,
		PosY:          posY,
		SizeX:         sizeX,
		SizeY:         sizeY,
		WindowPosX:    0,
		WindowPosY:    0,
		Content:       make([]string, 1),
		CurrentRow:    0,
		CurrentColumn: 0,
	}
}

func (area *TypingArea) Write(str string) {
	if len(area.Content[area.CurrentRow]) != 0 {
		if area.CurrentColumn == len(area.Content[area.CurrentRow]) {
			area.Content[area.CurrentRow] = area.Content[area.CurrentRow][:area.CurrentColumn] + str
		} else {
			area.Content[area.CurrentRow] = area.Content[area.CurrentRow][:area.CurrentColumn] + str + area.Content[area.CurrentRow][area.CurrentColumn+1:]
		}
		area.CurrentColumn += len(str)
	} else {
		area.Content[area.CurrentColumn] = str
	}
	if len(area.Content[area.CurrentRow]) > area.SizeX {
		additionalLen := area.Content[area.CurrentRow][area.SizeX:]
		area.RepairFrom(area.CurrentRow)
		area.CurrentRow += 1
		area.CurrentColumn = len(additionalLen)
	}
}

func (area *TypingArea) RemoveLastCharacter() {
	if area.Content[0] == "" {
		return
	}
	if area.Content[area.CurrentRow] == "" {
		area.CurrentRow--
	}
	area.Content[area.CurrentRow] = area.Content[area.CurrentRow][:len(area.Content[area.CurrentRow])-1]
	if area.CurrentColumn == 0 {
		area.CurrentColumn = len(area.Content[area.CurrentRow])
	} else {
		area.CurrentColumn--
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

func (area *TypingArea) RepairFrom(row int) {
	if len(area.Content[row]) > area.SizeX {
		tmpString := area.Content[row]
		area.Content[row] = tmpString[:area.SizeX]
		if len(area.Content) == area.CurrentRow+1 {
			area.Content = append(area.Content, tmpString[area.SizeX:])
		} else {
			area.Content[row+1] = tmpString[area.SizeX:]
		}
	}
}
