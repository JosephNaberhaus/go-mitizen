package prompt

type Prompt interface {
	Show()
	HandleInput(input Key)
	Showing() bool
	Finish()
	Finished() bool
}
