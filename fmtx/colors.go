package fmtx

import (
	"fmt"
	"strings"
)

var enabled = true

// EnableColors enables the use of colors globally in the package. By default,
// colors are enabled.
func EnableColors() {
	enabled = true
}

// DisableColors disables the use of colors globally in the package.
func DisableColors() {
	enabled = false
}

type ColorMode uint8

const (
	ColorModeAnsi ColorMode = iota
	ColorMode256
	ColorModeTrue
)

// StyleColor represents a color that can be used in a Style. It can be an
// ANSI color, a 256 color or a true color.
type StyleColor struct {
	mode ColorMode
	bg   bool  // User for 256 and true colors
	r    uint8 // Used for all modes
	g    uint8 // Not used for ansi or 256 modes
	b    uint8 // Not used for ansi or 256 modes
}

// NewColorAnsi creates a new ANSI color with the given code. All ansi colors
// are mapped to constants defined in this package, you may use them instead of
// creating new ones with this function.
//
// Examples: `ColorRed`, `ColorBgBrightBlue`, ...
func NewColorAnsi(code uint8) StyleColor {
	return StyleColor{mode: ColorModeAnsi, r: code}
}

// NewColor creates a new **FOREGROUND** true color with the given RGB values.
// Use `NewColorBg` to create a background color.
func NewColor(r, g, b uint8) StyleColor {
	return StyleColor{mode: ColorModeTrue, r: r, g: g, b: b}
}

// NewColorBg creates a new **BACKGROUND** true color with the given RGB values.
// Use `NewColor` to create a foreground color.
func NewColorBg(r, g, b uint8) StyleColor {
	return StyleColor{mode: ColorModeTrue, bg: true, r: r, g: g, b: b}
}

// NewColor256 creates a new 256 **FOREGROUND** color with the given code.
// Use `NewColor256Bg` to create a background color.
func NewColor256(code uint8) StyleColor {
	return StyleColor{mode: ColorMode256, r: code}
}

// NewColor256Bg creates a new 256 **BACKGROUND** color with the given code.
// Use `NewColor256` to create a foreground color.
func NewColor256Bg(code uint8) StyleColor {
	return StyleColor{mode: ColorMode256, bg: true, r: code}
}

// ToAnsi converts the color to an ANSI color. If the color is already an ANSI
// color, it returns the same color.
func (c *StyleColor) ToAnsi() StyleColor {
	if c.mode == ColorModeAnsi {
		return *c
	}
	if c.mode == ColorMode256 {
		return StyleColor{mode: ColorModeAnsi, r: c.r}
	}
	// Convert true color to 256 color
	r := int(c.r) * 5 / 255
	g := int(c.g) * 5 / 255
	b := int(c.b) * 5 / 255
	code := 16 + r*36 + g*6 + b
	return StyleColor{mode: ColorModeAnsi, r: uint8(code)}
}

// Code returns the ANSI code for the color, depending on its mode.
func (c *StyleColor) Code() string {
	switch c.mode {
	case ColorModeAnsi:
		return fmt.Sprintf("%d", c.r)
	case ColorMode256:
		if c.bg {
			return fmt.Sprintf("48;5;%d", c.r)
		}
		return fmt.Sprintf("38;5;%d", c.r)
	case ColorModeTrue:
		if c.bg {
			return fmt.Sprintf("48;2;%d;%d;%d", c.r, c.g, c.b)
		}
		return fmt.Sprintf("38;2;%d;%d;%d", c.r, c.g, c.b)
	default:
		return "0"
	}
}

const (
	modBold            = "1"
	modDim             = "2"
	modItalic          = "3"
	modUnderline       = "4"
	modSlowBlink       = "5"
	modRapidBlink      = "6"
	modInverse         = "7"
	modHidden          = "8"
	modStrikeThrough   = "9"
	modFraktur         = "20"
	modDoubleUnderline = "21"
	modFramed          = "51"
	modEncircled       = "52"
	modOverline        = "53"
)

// Style represents a combination of text modifiers and colors that can be
// applied to a string.
type Style struct {
	modifiers []string
	colors    []StyleColor
}

// NewStyle creates a new empty Style.
func NewStyle() *Style {
	return &Style{}
}

// WithBold adds the bold modifier to the style.
func (s *Style) WithBold() *Style {
	s.modifiers = append(s.modifiers, modBold)
	return s
}

// WithDim adds the dim modifier to the style.
func (s *Style) WithDim() *Style {
	s.modifiers = append(s.modifiers, modDim)
	return s
}

// WithItalic adds the italic modifier to the style.
func (s *Style) WithItalic() *Style {
	s.modifiers = append(s.modifiers, modItalic)
	return s
}

// WithUnderline adds the underline modifier to the style.
func (s *Style) WithUnderline() *Style {
	s.modifiers = append(s.modifiers, modUnderline)
	return s
}

// WithSlowBlink adds the slow blink modifier to the style. This is rarely
// supported by terminals.
func (s *Style) WithSlowBlink() *Style {
	s.modifiers = append(s.modifiers, modSlowBlink)
	return s
}

// WithRapidBlink adds the rapid blink modifier to the style. This is rarely
// supported by terminals.
func (s *Style) WithRapidBlink() *Style {
	s.modifiers = append(s.modifiers, modRapidBlink)
	return s
}

// WithInverse adds the inverse modifier to the style, which swaps the foreground
// and background colors.
func (s *Style) WithInverse() *Style {
	s.modifiers = append(s.modifiers, modInverse)
	return s
}

// WithHidden adds the hidden modifier to the style, which hides the text.
func (s *Style) WithHidden() *Style {
	s.modifiers = append(s.modifiers, modHidden)
	return s
}

// WithStrikeThrough adds the strike through modifier to the style, which adds a
// line through the text.
func (s *Style) WithStrikeThrough() *Style {
	s.modifiers = append(s.modifiers, modStrikeThrough)
	return s
}

// WithFraktur adds the fraktur modifier to the style, which changes the font to
// a fraktur style. This is rarely supported by terminals.
func (s *Style) WithFraktur() *Style {
	s.modifiers = append(s.modifiers, modFraktur)
	return s
}

// WithDoubleUnderline adds the double underline modifier to the style, which
// adds two lines under the text. This is rarely supported by terminals.
func (s *Style) WithDoubleUnderline() *Style {
	s.modifiers = append(s.modifiers, modDoubleUnderline)
	return s
}

// WithFramed adds the framed modifier to the style, which adds a frame around
// the text. This is rarely supported by terminals.
func (s *Style) WithFramed() *Style {
	s.modifiers = append(s.modifiers, modFramed)
	return s
}

// WithEncircled adds the encircled modifier to the style, which adds a circle
// around the text. This is rarely supported by terminals.
func (s *Style) WithEncircled() *Style {
	s.modifiers = append(s.modifiers, modEncircled)
	return s
}

// WithOverline adds the overline modifier to the style, which adds a line
// above the text.
func (s *Style) WithOverline() *Style {
	s.modifiers = append(s.modifiers, modOverline)
	return s
}

// WithColor sets the foreground color of the style. Be careful to use
func (s *Style) WithColor(color StyleColor) *Style {
	s.colors = append(s.colors, color)
	return s
}

func (s *Style) Apply(text string) string {
	if !enabled {
		return text
	}

	if len(s.modifiers) == 0 && len(s.colors) == 0 {
		return text
	}

	b := strings.Builder{}
	b.WriteString("\x1b[")

	size := 0
	for _, mod := range s.modifiers {
		if size > 0 {
			b.WriteString(";")
		}
		size += 1
		b.WriteString(mod)
	}
	for _, color := range s.colors {
		if size > 0 {
			b.WriteString(";")
		}
		size += 1
		b.WriteString(color.Code())
	}

	b.WriteString("m")
	b.WriteString(text)
	b.WriteString("\x1b[0m")
	return b.String()
}

var (
	ColorBlack           = NewColorAnsi(30)
	ColorRed             = NewColorAnsi(31)
	ColorGreen           = NewColorAnsi(32)
	ColorYellow          = NewColorAnsi(33)
	ColorBlue            = NewColorAnsi(34)
	ColorMagenta         = NewColorAnsi(35)
	ColorCyan            = NewColorAnsi(36)
	ColorWhite           = NewColorAnsi(37)
	ColorDefault         = NewColorAnsi(39)
	ColorBgBlack         = NewColorAnsi(40)
	ColorBgRed           = NewColorAnsi(41)
	ColorBgGreen         = NewColorAnsi(42)
	ColorBgYellow        = NewColorAnsi(43)
	ColorBgBlue          = NewColorAnsi(44)
	ColorBgMagenta       = NewColorAnsi(45)
	ColorBgCyan          = NewColorAnsi(46)
	ColorBgWhite         = NewColorAnsi(47)
	ColorBgDefault       = NewColorAnsi(49)
	ColorBrightBlack     = NewColorAnsi(90)
	ColorBrightRed       = NewColorAnsi(91)
	ColorBrightGreen     = NewColorAnsi(92)
	ColorBrightYellow    = NewColorAnsi(93)
	ColorBrightBlue      = NewColorAnsi(94)
	ColorBrightMagenta   = NewColorAnsi(95)
	ColorBrightCyan      = NewColorAnsi(96)
	ColorBrightWhite     = NewColorAnsi(97)
	ColorBgBrightBlack   = NewColorAnsi(100)
	ColorBgBrightRed     = NewColorAnsi(101)
	ColorBgBrightGreen   = NewColorAnsi(102)
	ColorBgBrightYellow  = NewColorAnsi(103)
	ColorBgBrightBlue    = NewColorAnsi(104)
	ColorBgBrightMagenta = NewColorAnsi(105)
	ColorBgBrightCyan    = NewColorAnsi(106)
	ColorBgBrightWhite   = NewColorAnsi(107)
)

var (
	styleBold            = NewStyle().WithBold()
	styleDim             = NewStyle().WithDim()
	styleItalic          = NewStyle().WithItalic()
	styleUnderline       = NewStyle().WithUnderline()
	styleSlowBlink       = NewStyle().WithSlowBlink()
	styleRapidBlink      = NewStyle().WithRapidBlink()
	styleReverse         = NewStyle().WithInverse()
	styleHidden          = NewStyle().WithHidden()
	styleStrikeThrough   = NewStyle().WithStrikeThrough()
	styleFraktur         = NewStyle().WithFraktur()
	styleDoubleUnderline = NewStyle().WithDoubleUnderline()
	styleFramed          = NewStyle().WithFramed()
	styleEncircled       = NewStyle().WithEncircled()
	styleOverline        = NewStyle().WithOverline()
	styleBlack           = NewStyle().WithColor(ColorBlack)
	styleRed             = NewStyle().WithColor(ColorRed)
	styleGreen           = NewStyle().WithColor(ColorGreen)
	styleYellow          = NewStyle().WithColor(ColorYellow)
	styleBlue            = NewStyle().WithColor(ColorBlue)
	styleMagenta         = NewStyle().WithColor(ColorMagenta)
	styleCyan            = NewStyle().WithColor(ColorCyan)
	styleWhite           = NewStyle().WithColor(ColorWhite)
	styleDefault         = NewStyle().WithColor(ColorDefault)
	styleBgBlack         = NewStyle().WithColor(ColorBgBlack)
	styleBgRed           = NewStyle().WithColor(ColorBgRed)
	styleBgGreen         = NewStyle().WithColor(ColorBgGreen)
	styleBgYellow        = NewStyle().WithColor(ColorBgYellow)
	styleBgBlue          = NewStyle().WithColor(ColorBgBlue)
	styleBgMagenta       = NewStyle().WithColor(ColorBgMagenta)
	styleBgCyan          = NewStyle().WithColor(ColorBgCyan)
	styleBgWhite         = NewStyle().WithColor(ColorBgWhite)
	styleBgDefault       = NewStyle().WithColor(ColorBgDefault)
	styleBrightBlack     = NewStyle().WithColor(ColorBrightBlack)
	styleBrightRed       = NewStyle().WithColor(ColorBrightRed)
	styleBrightGreen     = NewStyle().WithColor(ColorBrightGreen)
	styleBrightYellow    = NewStyle().WithColor(ColorBrightYellow)
	styleBrightBlue      = NewStyle().WithColor(ColorBrightBlue)
	styleBrightMagenta   = NewStyle().WithColor(ColorBrightMagenta)
	styleBrightCyan      = NewStyle().WithColor(ColorBrightCyan)
	styleBrightWhite     = NewStyle().WithColor(ColorBrightWhite)
	styleBgBrightBlack   = NewStyle().WithColor(ColorBgBrightBlack)
	styleBgBrightRed     = NewStyle().WithColor(ColorBgBrightRed)
	styleBgBrightGreen   = NewStyle().WithColor(ColorBgBrightGreen)
	styleBgBrightYellow  = NewStyle().WithColor(ColorBgBrightYellow)
	styleBgBrightBlue    = NewStyle().WithColor(ColorBgBrightBlue)
	styleBgBrightMagenta = NewStyle().WithColor(ColorBgBrightMagenta)
	styleBgBrightCyan    = NewStyle().WithColor(ColorBgBrightCyan)
	styleBgBrightWhite   = NewStyle().WithColor(ColorBgBrightWhite)
)

func Bold(s string) string            { return styleBold.Apply(s) }
func Dim(s string) string             { return styleDim.Apply(s) }
func Italic(s string) string          { return styleItalic.Apply(s) }
func Underline(s string) string       { return styleUnderline.Apply(s) }
func SlowBlink(s string) string       { return styleSlowBlink.Apply(s) }
func RapidBlink(s string) string      { return styleRapidBlink.Apply(s) }
func Reverse(s string) string         { return styleReverse.Apply(s) }
func Hidden(s string) string          { return styleHidden.Apply(s) }
func StrikeThrough(s string) string   { return styleStrikeThrough.Apply(s) }
func Fraktur(s string) string         { return styleFraktur.Apply(s) }
func DoubleUnderline(s string) string { return styleDoubleUnderline.Apply(s) }
func Framed(s string) string          { return styleFramed.Apply(s) }
func Encircled(s string) string       { return styleEncircled.Apply(s) }
func Overline(s string) string        { return styleOverline.Apply(s) }
func Black(s string) string           { return styleBlack.Apply(s) }
func Red(s string) string             { return styleRed.Apply(s) }
func Green(s string) string           { return styleGreen.Apply(s) }
func Yellow(s string) string          { return styleYellow.Apply(s) }
func Blue(s string) string            { return styleBlue.Apply(s) }
func Magenta(s string) string         { return styleMagenta.Apply(s) }
func Cyan(s string) string            { return styleCyan.Apply(s) }
func White(s string) string           { return styleWhite.Apply(s) }
func Default(s string) string         { return styleDefault.Apply(s) }
func BgBlack(s string) string         { return styleBgBlack.Apply(s) }
func BgRed(s string) string           { return styleBgRed.Apply(s) }
func BgGreen(s string) string         { return styleBgGreen.Apply(s) }
func BgYellow(s string) string        { return styleBgYellow.Apply(s) }
func BgBlue(s string) string          { return styleBgBlue.Apply(s) }
func BgMagenta(s string) string       { return styleBgMagenta.Apply(s) }
func BgCyan(s string) string          { return styleBgCyan.Apply(s) }
func BgWhite(s string) string         { return styleBgWhite.Apply(s) }
func BgDefault(s string) string       { return styleBgDefault.Apply(s) }
func BrightBlack(s string) string     { return styleBrightBlack.Apply(s) }
func BrightRed(s string) string       { return styleBrightRed.Apply(s) }
func BrightGreen(s string) string     { return styleBrightGreen.Apply(s) }
func BrightYellow(s string) string    { return styleBrightYellow.Apply(s) }
func BrightBlue(s string) string      { return styleBrightBlue.Apply(s) }
func BrightMagenta(s string) string   { return styleBrightMagenta.Apply(s) }
func BrightCyan(s string) string      { return styleBrightCyan.Apply(s) }
func BrightWhite(s string) string     { return styleBrightWhite.Apply(s) }
func BgBrightBlack(s string) string   { return styleBgBrightBlack.Apply(s) }
func BgBrightRed(s string) string     { return styleBgBrightRed.Apply(s) }
func BgBrightGreen(s string) string   { return styleBgBrightGreen.Apply(s) }
func BgBrightYellow(s string) string  { return styleBgBrightYellow.Apply(s) }
func BgBrightBlue(s string) string    { return styleBgBrightBlue.Apply(s) }
func BgBrightMagenta(s string) string { return styleBgBrightMagenta.Apply(s) }
func BgBrightCyan(s string) string    { return styleBgBrightCyan.Apply(s) }
func BgBrightWhite(s string) string   { return styleBgBrightWhite.Apply(s) }

// Stylize applies multiple style application functions to a string. The styles
// are applied in the order they are given.
func Stylize(s string, v ...func(s string) string) string {
	for _, fn := range v {
		s = fn(s)
	}
	return s
}
