package prompt

import (
	"github.com/eiannone/keyboard"
)

type Key interface {
	isKey() bool
}

type RuneKey rune

func (r RuneKey) isKey() bool {
	return true
}

type ControlKey uint8

const (
	Noop ControlKey = iota
	ControlLeft
	ControlRight
	ControlUp
	ControlDown
	ControlEnter
	ControlBackspace
	ControlSpace
	ControlHome
	ControlEnd
	ControlNextWord
	ControlPrevWord
)

func ToKey(rune rune, key keyboard.Key) Key {
	if rune != 0 {
		return RuneKey(rune)
	}

	switch key {
	case keyboard.KeyArrowLeft: return ControlLeft
	case keyboard.KeyArrowRight: return ControlRight
	case keyboard.KeyArrowUp: return ControlUp
	case keyboard.KeyArrowDown: return ControlDown
	case keyboard.KeyEnter: return ControlEnter
	case keyboard.KeyBackspace: fallthrough
	case keyboard.KeyBackspace2: return ControlBackspace
	case keyboard.KeySpace: return ControlSpace
	case keyboard.KeyHome: return ControlHome
	case keyboard.KeyEnd: return ControlEnd
	default: return Noop
	}
}

func (c ControlKey) isKey() bool {
	return true
}
