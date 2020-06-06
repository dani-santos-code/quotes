package phrase

import "testing"

func TestGlass (t*testing.T) {
	got := Get()
	want := "I can eat glass and it doesn't hurt me."

	if got != want {
		t.Errorf("got %q qant %q", got, want)
	}
}


func TestGetOpt (t*testing.T) {
	got := GetOpt()
	want := "If a program is too slow, it must have a loop."

	if got != want {
		t.Errorf("got %q qant %q", got, want)
	}
}
