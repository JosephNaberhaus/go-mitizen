package prompt

import (
	"errors"
	"fmt"
	"github.com/eiannone/keyboard"
)

func loopUntilFinished(p Prompt) error {
	for !p.Finished() {
		r, key, err := keyboard.GetKey()
		if key == keyboard.KeyCtrlC {
			p.Finish()
			return errors.New("prompt loop aborted")
		}

		if err != nil {
			return fmt.Errorf("error getting key input: %w", err)
		}

		p.handleInput(ToKey(r, key))
	}

	return nil
}
