package color

const (
	pre   = "\033["
	post  = ``
	reset = "\033[0m"
)

// Paint is a color to paint, either as a foreground or background paint
type Paint string

// Valid colors for ANSI terminals
const (
	BlackPaint      Paint = `0;30`
	DarkRedPaint    Paint = `0;31`
	DarkGreenPaint  Paint = `0;32`
	DarkYellowPaint Paint = `0;33`
	DarkBluePaint   Paint = `0;34`
	DarkPurplePaint Paint = `0;35`
	DarkCyanPaint   Paint = `0;36`
	LightGrayPaint  Paint = `0;37`

	DarkGrayPaint Paint = `1;30`
	RedPaint      Paint = `1;31`
	GreenPaint    Paint = `1;32`
	YellowPaint   Paint = `1;33`
	BluePaint     Paint = `1;34`
	PurplePaint   Paint = `1;35`
	CyanPaint     Paint = `1;36`
	WhitePaint    Paint = `1;37`

	nilPaint Paint = `nil`
)

// Brush is a function that let's you colorize strings directly.
type Brush func(string) string

// NewBrush gives you a brush that you can invoke directly to create colorized
// strings
func NewBrush(background, foreground Paint) Brush {
	return NewStyle(background, foreground).Brush()
}

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

// Brush is a function that can be used to color things directly, i.e:
//
//    red := NewStyle(BlackPaint, RedPaint).Brush()
//    fmt.Printf("This is %s\n", red("red"))
func (s Style) Brush() Brush {
	return func(text string) string {
		return s.code + text + reset
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
	if bg == nilPaint {
		return pre + string(fg) + "m" + post
	}

	// The background code is the last color code prefixed by 4
	bgColor := bg[len(bg)-1]
	back := pre + "4" + string(bgColor) + "m" + post

	front := pre + string(fg) + "m" + post
	return back + front
}

// Black gives black text on a white background. Use it like this:
// 		color.Black("I'm color Black")
type Black string

func (b Black) String() string { return NewBrush(WhitePaint, BlackPaint)(string(b)) }

// White gives white text on a dark gray background. Use it like this:
// 		color.White("I'm color White")
type White string

func (w White) String() string { return NewBrush(DarkGrayPaint, WhitePaint)(string(w)) }

// LightGray gives light gray text on a black background. Use it like this:
// 		color.LightGray("I'm color LightGray")
type LightGray string

func (l LightGray) String() string { return NewBrush(nilPaint, LightGrayPaint)(string(l)) }

// Blue gives blue text on a black background. Use it like this:
// 		color.Blue("I'm color Blue")
type Blue string

func (b Blue) String() string { return NewBrush(nilPaint, BluePaint)(string(b)) }

// Cyan gives cyan text on a black background. Use it like this:
// 		color.Cyan("I'm color Cyan")
type Cyan string

func (c Cyan) String() string { return NewBrush(nilPaint, CyanPaint)(string(c)) }

// Green gives green text on a black background. Use it like this:
// 		color.Green("I'm color Green")
type Green string

func (g Green) String() string { return NewBrush(nilPaint, GreenPaint)(string(g)) }

// Purple gives purple text on a black background. Use it like this:
// 		color.Purple("I'm color Purple")
type Purple string

func (p Purple) String() string { return NewBrush(nilPaint, PurplePaint)(string(p)) }

// Red gives red text on a black background. Use it like this:
// 		color.Red("I'm color Red")
type Red string

func (r Red) String() string { return NewBrush(nilPaint, RedPaint)(string(r)) }

// Yellow gives yellow text on a black background. Use it like this:
// 		color.Yellow("I'm color Yellow")
type Yellow string

func (y Yellow) String() string { return NewBrush(nilPaint, YellowPaint)(string(y)) }

// DarkBlue gives dark blue text on a black background. Use it like this:
// 		color.DarkBlue("I'm color DarkBlue")
type DarkBlue string

func (d DarkBlue) String() string { return NewBrush(nilPaint, DarkBluePaint)(string(d)) }

// DarkCyan gives dark cyan text on a black background. Use it like this:
// 		color.DarkCyan("I'm color DarkCyan")
type DarkCyan string

func (d DarkCyan) String() string { return NewBrush(nilPaint, DarkCyanPaint)(string(d)) }

// DarkGray gives dark gray text on a black background. Use it like this:
// 		color.DarkGray("I'm color DarkGray")
type DarkGray string

func (d DarkGray) String() string { return NewBrush(nilPaint, DarkGrayPaint)(string(d)) }

// DarkGreen gives dark green text on a black background. Use it like this:
// 		color.DarkGreen("I'm color DarkGreen")
type DarkGreen string

func (d DarkGreen) String() string { return NewBrush(nilPaint, DarkGreenPaint)(string(d)) }

// DarkPurple gives dark purple text on a black background. Use it like this:
// 		color.DarkPurple("I'm color DarkPurple")
type DarkPurple string

func (d DarkPurple) String() string { return NewBrush(nilPaint, DarkPurplePaint)(string(d)) }

// DarkRed gives dark red text on a black background. Use it like this:
// 		color.DarkRed("I'm color DarkRed")
type DarkRed string

func (d DarkRed) String() string { return NewBrush(nilPaint, DarkRedPaint)(string(d)) }

// DarkYellow gives brown text on a black background. Use it like this:
// 		color.DarkYellow("I'm color DarkYellow")
type DarkYellow string

func (d DarkYellow) String() string { return NewBrush(nilPaint, DarkYellowPaint)(string(d)) }
