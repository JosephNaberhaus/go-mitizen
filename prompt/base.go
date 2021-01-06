package prompt

type base struct {
	output *output

	showing  bool
	finished bool
}

func (b *base) Show() {
	if b.showing {
		panic("cannot call show multiple times")
	}

	b.output = newOutput()
	b.showing = true
}

func (b *base) Showing() bool {
	return b.showing
}

func (b *base) Finish() {
	b.finished = true
}

func (b *base) Finished() bool {
	return b.finished
}
