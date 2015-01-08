package main

import "testing"

var flipStringTests = []struct {
	args []string
	out  string
}{
	{[]string{"ab"}, "qɐ"},
	{[]string{"table", "abcde"}, "(╯°□°）╯︵ǝpɔqɐ"},
	{[]string{"gopher", "abcde"}, "ʕ╯◔ϖ◔ʔ╯︵ǝpɔqɐ"},
	{[]string{"(ノಠ益ಠ)ノ︵", "custom"}, "(ノಠ益ಠ)ノ︵ɯoʇsnɔ"},
}

func TestFlipStringDown(t *testing.T) {
	for _, tt := range flipStringTests {
		if got := flipString(tt.args); got != tt.out {
			t.Errorf("flipString(%v) = %v, want %v", tt.args, got, tt.out)
		}
	}
}
