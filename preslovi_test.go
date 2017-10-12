package preslovi

import "testing"

var samples = []struct {
	in  string
	out string
}{
	{"Шабан Шаулић", "Šaban Šaulić"},
	{"ЊЕГОШ", "NJEGOŠ"},
	{"електродистрибуција", "elektrodistribucija"},
	{"Ђурђевак", "Đurđevak"},
	{"Latinica", "Latinica"},
}

func TestLatinicom(t *testing.T) {
	for _, tt := range samples {
		actual := Latinicom(tt.in)
		if actual != tt.out {
			t.Errorf("Latinicom(%s): expected %s, actual %s", tt.in, tt.out, actual)
		}
	}
}
