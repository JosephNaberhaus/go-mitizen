package prompt



type YesNo struct {
	Description string

	output *output
	editor *editor

	showing bool
	finished bool
}

func (y *YesNo) Show() {
	y.output = newOutput()
	y.editor = newEditor(y.output.numCols)

	y.render()
}

func (y *YesNo) HandleInput(input Key) {
	if y.finished {
		return
	}

	if input == ControlEnter {
		y.Finish()
		return
	} else {
		y.editor.handleInput(input)
	}

	y.render()
}

func (y *YesNo) render() {
	y.output.clear()

	y.output.writeColor("? ", ColorGreen)
	y.output.write(y.Description)
	y.output.writeColor(" (y/N) ", ColorGreen)

	offsetX := y.output.cursorX
	offsetY := y.output.cursorY

	if y.finished {
		if y.ResponseYes() {
			y.output.writeColorLn("Yes", ColorCyan)
		} else {
			y.output.writeColorLn("No", ColorCyan)
		}
	} else {
		y.output.writeLn(y.editor.curLine())
	}

	y.output.setCursor(y.editor.getRealCursorPosition(offsetX, offsetY))

	y.output.flush()
}

func (y *YesNo) Showing() bool {
	return y.showing
}

func (y *YesNo) Finish() {
	y.finished = true
	y.render()
	y.output.commit()
}

func (y *YesNo) Finished() bool {
	return y.finished
}

func (y *YesNo) ResponseYes() bool {
	return y.editor.curLineLength() > 0 && y.editor.curLine()[0] == 'y'
}
