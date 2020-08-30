package preslovi

import "testing"

var examples = []struct {
	input string
	latin string
	ascii string
}{
	{"ЊЕГОШ", "NJEGOŠ", "NJEGOS"},
	{"Latinica", "Latinica", "Latinica"},
	{"Ђурђевак", "Đurđevak", "Djurdjevak"},
	{"Шабан Шаулић", "Šaban Šaulić", "Saban Saulic"},
	{"електродистрибуција", "elektrodistribucija", "elektrodistribucija"},
	{"ЂОРЂЕ Ђ. Ђорђевић", "ĐORĐE Đ. Đorđević", "DJORDJE Dj. Djordjevic"},
	{"ĐORĐE Đ. Đorđević", "ĐORĐE Đ. Đorđević", "DJORDJE Dj. Djordjevic"},
	{"ЉИЉА Љ. Љиљановић", "LJILJA Lj. Ljiljanović", "LJILJA Lj. Ljiljanovic"},
}

func TestExamples(t *testing.T) {
	for _, tt := range examples {
		latin := Latinicom(tt.input)
		if latin != tt.latin {
			t.Errorf("Latinicom(%s): expected %s, actual %s", tt.input, tt.latin, latin)
		}

		ascii := LatinicomAscii(tt.input)
		if ascii != tt.ascii {
			t.Errorf("LatinicomAscii(%s): expected %s, actual %s", tt.input, tt.ascii, ascii)
		}
	}
}
