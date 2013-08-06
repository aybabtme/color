package color

import (
	"os"
)

const (
	pre   = "\033["
	post  = ``
	reset = "\033[0m"
)

// Paint is a color to paint, either as a foreground or background paint
type Paint string

const (
	nullPaint      Paint = `null`
	BlackPaint     Paint = `0;30`
	BluePaint      Paint = `0;34`
	GreenPaint     Paint = `0;32`
	CyanPaint      Paint = `0;36`
	RedPaint       Paint = `0;31`
	PurplePaint    Paint = `0;35`
	BrownPaint     Paint = `0;33`
	LightGrayPaint Paint = `0;37`

	DarkGrayPaint    Paint = `1;30`
	LightBluePaint   Paint = `1;34`
	LightGreenPaint  Paint = `1;32`
	LightCyanPaint   Paint = `1;36`
	LightRedPaint    Paint = `1;31`
	LightPurplePaint Paint = `1;35`
	YellowPaint      Paint = `1;33`
	WhitePaint       Paint = `1;37`
)

// Style will give you colorized strings.  Styles are immutable.
type Style struct {
	bg   Paint
	fg   Paint
	code string
}

// NewStyle gives you a style ready to produce strings with the given
// background and foreground colors
func NewStyle(background, foreground Paint) Style {
	bg := background
	fg := foreground
	return Style{
		bg,
		fg,
		computeColorCode(bg, fg),
	}
}

// Get the string with this color style applied to it
func (s Style) Get(text string) string {
	return s.code + text + reset
}

// Print will write the current text to os.StdOut.
func (s Style) Print(text string) {
	if _, err := os.Stdout.WriteString(s.Get(text)); err != nil {
		panic(err)
	}
}

// WithBackground copies the current style and return a new Style that
// has the desired background. The original Style is unchanged and you
// must capture the return value.
func (s Style) WithBackground(color Paint) Style {
	newS := s
	newS.bg = color
	newS.code = computeColorCode(newS.bg, newS.fg)
	return newS
}

// WithForeground copies the current style and return a new Style that
// has the desired foreground. The original Style is unchanged and you
// must capture the return value.
func (s Style) WithForeground(color Paint) Style {
	newS := s
	newS.fg = color
	newS.code = computeColorCode(newS.bg, newS.fg)
	return newS
}

func computeColorCode(bg, fg Paint) string {
	if bg == nullPaint {
		return pre + string(fg) + "m" + post
	}

	// The background code is the last color code prefixed by 4
	bgColor := bg[len(bg)-1]
	back := pre + "4" + string(bgColor) + "m" + post

	front := pre + string(fg) + "m" + post
	return back + front
}

// Red gives red text on a black background
func Red() Style {
	return NewStyle(nullPaint, RedPaint)
}

// Blue gives blue text on a black background
func Blue() Style {
	return NewStyle(nullPaint, BluePaint)
}

// Green gives green text on a black background
func Green() Style {
	return NewStyle(nullPaint, GreenPaint)
}

// Cyan gives cyan text on a black background
func Cyan() Style {
	return NewStyle(nullPaint, CyanPaint)
}

// Purple gives purple text on a black background
func Purple() Style {
	return NewStyle(nullPaint, PurplePaint)
}

// Brown gives brown text on a black background
func Brown() Style {
	return NewStyle(nullPaint, BrownPaint)
}

// LightGray gives light gray text on a black background
func LightGray() Style {
	return NewStyle(nullPaint, LightGrayPaint)
}

// DarkGray gives dark gray text on a black background
func DarkGray() Style {
	return NewStyle(nullPaint, DarkGrayPaint)
}

// LightBlue gives light blue text on a black background
func LightBlue() Style {
	return NewStyle(nullPaint, LightBluePaint)
}

// LightGreen gives light green text on a black background
func LightGreen() Style {
	return NewStyle(nullPaint, LightGreenPaint)
}

// LightCyan gives light cyan text on a black background
func LightCyan() Style {
	return NewStyle(nullPaint, LightCyanPaint)
}

// LightRed gives light red text on a black background
func LightRed() Style {
	return NewStyle(nullPaint, LightRedPaint)
}

// LightPurple gives light purple text on a black background
func LightPurple() Style {
	return NewStyle(nullPaint, LightPurplePaint)
}

// Yellow gives light  text on a black background
func Yellow() Style {
	return NewStyle(nullPaint, YellowPaint)
}

// White gives light  text on a dark gray background
func White() Style {
	return NewStyle(DarkGrayPaint, WhitePaint)
}

// Black gives black text on a white background
func Black() Style {
	return NewStyle(WhitePaint, BlackPaint)
}
