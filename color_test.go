package color

import (
	"fmt"
	"testing"
)

var fgTT = []struct {
	name      string
	styleFunc func() Style
	fg        Paint
}{
	{"red", Red, RedPaint},
	{"blue", Blue, BluePaint},
	{"green", Green, GreenPaint},
	{"cyan", Cyan, CyanPaint},
	{"purple", Purple, PurplePaint},
	{"brown", Brown, BrownPaint},
	{"light Gray", LightGray, LightGrayPaint},
	{"dark Gray", DarkGray, DarkGrayPaint},
	{"light Blue", LightBlue, LightBluePaint},
	{"light Green", LightGreen, LightGreenPaint},
	{"light Cyan", LightCyan, LightCyanPaint},
	{"light Red", LightRed, LightRedPaint},
	{"light Purple", LightPurple, LightPurplePaint},
	{"yellow", Yellow, YellowPaint},
	// white and black have different backgrounds
}

func TestAllForegroundStyles(t *testing.T) {
	for _, test := range fgTT {
		want := "\033[" + string(test.fg) + "m" + test.name + "\033[0m"
		got := test.styleFunc().Get(test.name)

		fmt.Println("Want : " + want + ", got : " + got)

		if want != got {
			t.Errorf("Want %#v, got %#v", want, got)
		}
	}
}

func TestStylesImmutable(t *testing.T) {
	yellow := Yellow()

	msg := "this message has yellow foreground"
	want := yellow.Get(msg)

	yellowRedBg := yellow.WithBackground(RedPaint)
	notWant := yellowRedBg.Get(msg)

	got := yellow.Get(msg)

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
	{"Brown", BrownPaint},
	{"LightGray", LightGrayPaint},
	{"DarkGray", DarkGrayPaint},
	{"LightBlue", LightBluePaint},
	{"LightGreen", LightGreenPaint},
	{"LightCyan", LightCyanPaint},
	{"LightRed", LightRedPaint},
	{"LightPurple", LightPurplePaint},
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
		style := NewStyle(perm.bg, perm.fg)

		want := "" +
			"\033[" + "4" + string(perm.bg[len(perm.bg)-1]) + "m" +
			"\033[" + string(perm.fg) + "m" +
			perm.name + "\033[0m"

		got := style.Get(perm.name)

		fmt.Printf("Look at %s all the colors %s!!!\n", want, got)
		if got != want {
			t.Errorf("Want %s, got %s.  From %#v to %#v", want, got, want, got)
		}
	}
}
