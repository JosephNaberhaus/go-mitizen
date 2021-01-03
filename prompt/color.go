package prompt

import escapes "github.com/snugfox/ansi-escapes"

type Color uint8

const (
	ColorBlack Color = iota
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

func (c Color) ToTextEscapes() string {
	switch c {
	case ColorBlack: return escapes.TextColorBlack
	case ColorRed: return escapes.TextColorRed
	case ColorGreen: return escapes.TextColorGreen
	case ColorYellow: return escapes.TextColorYellow
	case ColorBlue: return escapes.TextColorBlue
	case ColorMagenta: return escapes.TextColorMagenta
	case ColorCyan: return escapes.TextColorCyan
	case ColorWhite: return escapes.TextColorWhite
	}
	
	panic("invalid color")
}

func (c Color) ToBackgroundEscapes() string {
	switch c {
	case ColorBlack: return escapes.BackgroundColorBlack
	case ColorRed: return escapes.BackgroundColorRed
	case ColorGreen: return escapes.BackgroundColorGreen
	case ColorYellow: return escapes.BackgroundColorYellow
	case ColorBlue: return escapes.BackgroundColorBlue
	case ColorMagenta: return escapes.BackgroundColorMagenta
	case ColorCyan: return escapes.BackgroundColorCyan
	case ColorWhite: return escapes.BackgroundColorWhite
	}

	panic("invalid color")
}