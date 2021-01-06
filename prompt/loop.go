package prompt

import (
	"errors"
	"fmt"
	"github.com/eiannone/keyboard"
	"log"
)

func loopUntilFinished(p Prompt) error {
	for !p.Finished() {
		r, key, err := keyboard.GetKey()
		log.Printf("Input rune: %v key: %v", r, key)
		if err != nil {
			if err.Error() == "Unrecognized escape sequence" {
				continue
			}
			return fmt.Errorf("error getting key input: %w", err)
		}

		if key == keyboard.KeyCtrlC {
			p.Finish()
			return errors.New("prompt loop aborted")
		}

		p.handleInput(ToKey(r, key))
	}

	return nil
}
