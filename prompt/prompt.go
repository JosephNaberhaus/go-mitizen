package prompt

type Prompt interface {
	Show() error
	handleInput(input Key)
	Showing() bool
	Finish()
	Finished() bool
}
