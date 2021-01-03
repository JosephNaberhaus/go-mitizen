package prompt

type Group struct {
	Prompts []Prompt
	curPromptIndex int
}

func (g *Group) Show() {
	g.curPromptIndex = 0
}

func (g *Group) HandleInput(input Key) {
	if !g.Finished() {
		g.curPrompt().HandleInput(input)

		if g.curPrompt().Finished() {
			g.curPromptIndex++
		}
	}
}

func (g *Group) curPrompt() Prompt {
	return g.Prompts[g.curPromptIndex]
}

func (g *Group) Showing() bool {
	return g.curPrompt().Showing()
}

func (g *Group) Finish() {
	g.curPrompt().Finish()
}

func (g *Group) Finished() bool {
	return g.curPrompt().Finished()
}

