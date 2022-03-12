package form

import (
	"github.com/JosephNaberhaus/prompt"
	"strconv"
)

type Result int

const (
	ResultSubmit Result = iota
	ResultBack
)

type Question struct {
	prompt     prompt.Prompt
	beforeShow func(q *Form)
}

func (q *Question) Show(questions *Form) (Result, error) {
	if q.beforeShow != nil {
		q.beforeShow(questions)
	}

	if q.prompt.State() == prompt.Finished {
		err := q.prompt.ResetToWaiting()
		if err != nil {
			return 0, err
		}
	}

	err := q.prompt.Show()
	if err != nil {
		return 0, err
	}

	if q.prompt.State() == prompt.Finished {
		return ResultSubmit, nil
	}

	return ResultBack, nil
}

func (q *Question) StringValue() string {
	switch p := q.prompt.(type) {
	case *prompt.Text:
		return p.Response()
	case *prompt.Select:
		return p.Response().Name
	case *prompt.Boolean:
		return strconv.FormatBool(p.Response())
	}

	return ""
}

// A key handler that can be passed into the prompt library to go back when ctrl+b is pressed
func goBackHandler(p prompt.Prompt, key prompt.Key) bool {
	controlKey, isControlKey := key.(prompt.ControlKey)
	if isControlKey && controlKey == prompt.ControlCtrlB {
		_ = p.Pause()
		return true
	}

	return true
}
