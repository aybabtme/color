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
