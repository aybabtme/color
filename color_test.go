package color

import (
	"fmt"
	"testing"
)

var fgTT = []struct {
	name  string
	brush Brush
	fg    Paint
}{
	{"red", Red, RedPaint},
	{"blue", Blue, BluePaint},
	{"green", Green, GreenPaint},
	{"cyan", Cyan, CyanPaint},
	{"purple", Purple, PurplePaint},
	{"light Gray", LightGray, LightGrayPaint},
	{"dark Gray", DarkGray, DarkGrayPaint},
	{"dark Blue", DarkBlue, DarkBluePaint},
	{"dark Yellow", DarkYellow, DarkYellowPaint},
	{"dark Green", DarkGreen, DarkGreenPaint},
	{"dark Cyan", DarkCyan, DarkCyanPaint},
	{"dark Red", DarkRed, DarkRedPaint},
	{"dark Purple", DarkPurple, DarkPurplePaint},
	{"yellow", Yellow, YellowPaint},
	// white and black have different backgrounds
}

func TestAllForegroundStyles(t *testing.T) {
	for _, test := range fgTT {
		want := "\033[" + string(test.fg) + "m" + test.name + "\033[0m"
		got := test.brush(test.name)

		fmt.Println("Want : " + want + ", got : " + got)

		if want != got {
			t.Errorf("Want %#v, got %#v", want, got)
		}
	}
}

func TestStylesImmutable(t *testing.T) {
	yellow := NewStyle(BlackPaint, YellowPaint)
	yel := yellow.Brush()

	msg := "this message has yellow foreground"
	want := yel(msg)

	yellowRedBg := yellow.WithBackground(RedPaint).Brush()
	notWant := yellowRedBg(msg)

	got := yel(msg)

	if got == notWant {
		t.Errorf("Didn't want %#v but got it", notWant)
	}

	if got != want {
		t.Errorf("Want %#v got %#v", want, got)
	}

	fmt.Println("Want : " + want + ", not want : " + notWant + ", got : " + got)

}

var allPaints = []struct {
	color string
	p     Paint
}{
	{"Black", BlackPaint},
	{"Blue", BluePaint},
	{"Green", GreenPaint},
	{"Cyan", CyanPaint},
	{"Red", RedPaint},
	{"Purple", PurplePaint},
	{"DarkYellow", DarkYellowPaint},
	{"LightGray", LightGrayPaint},
	{"DarkGray", DarkGrayPaint},
	{"DarkBlue", DarkBluePaint},
	{"DarkGreen", DarkGreenPaint},
	{"DarkCyan", DarkCyanPaint},
	{"DarkRed", DarkRedPaint},
	{"DarkPurple", DarkPurplePaint},
	{"Yellow", YellowPaint},
	{"White", WhitePaint},
}

type PaintPerm struct {
	name string
	fg   Paint
	bg   Paint
}

func allPaintPermutation() []PaintPerm {
	var perm []PaintPerm
	var name string
	for i, p := range allPaints {
		for j, pp := range allPaints {
			if i == j {
				name = "double-" + p.color
			} else {
				name = p.color + " on " + pp.color
			}
			perm = append(perm, PaintPerm{
				name: name,
				fg:   p.p,
				bg:   pp.p,
			})
		}
	}
	return perm
}

func TestAllPermutationsOfPaint(t *testing.T) {
	for _, perm := range allPaintPermutation() {
		brush := NewBrush(perm.bg, perm.fg)

		want := "" +
			"\033[" + "4" + string(perm.bg[len(perm.bg)-1]) + "m" +
			"\033[" + string(perm.fg) + "m" +
			perm.name + "\033[0m"

		got := brush(perm.name)

		fmt.Printf("Look at %s all the colors %s!!!\n", want, got)
		if got != want {
			t.Errorf("Want %s, got %s.  From %#v to %#v", want, got, want, got)
		}
	}
}

func TestCanDoReadmeExample(t *testing.T) {
	// Default Brush are available for your convenience.  You can invoke
	// them directly
	fmt.Printf("This is %s\n", Red("red"))

	// or rename them and invoke them
	yel := Yellow
	fmt.Printf("This is %s\n", yel("yellow"))

	// or you can create new ones!
	weird := NewBrush(PurplePaint, CyanPaint)
	fmt.Printf("This color is %s\n", weird("weird"))

	// Create a Style, which has convenience methods
	redBg := NewStyle(RedPaint, YellowPaint)

	// Style.WithForeground or WithBackground returns a new Style, with the applied
	// Paint.  Styles are immutable so the original one is left unchanged
	greenFg := redBg.WithForeground(GreenPaint)

	// Style.Brush gives you a Brush that you can invoke directly to colorize strings.
	green := greenFg.Brush()
	fmt.Printf("This is %s\n", green("green"))
}
