package TwitterWalk

import (
	"fmt"
	"testing"
)

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		fmt.Println(got)
		got = append(got, input)
	})

	if got[0] != expected {
		t.Errorf("got %q want %q", got[0], expected)
	}
}
