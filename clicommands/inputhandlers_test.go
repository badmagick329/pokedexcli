package clicommands

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestCleanInput(t *testing.T) {
	tests := map[string]struct {
		input string
		want  []string
	}{
		"world split":  {input: "hello world", want: []string{"hello", "world"}},
		"empty string": {input: "", want: []string{}},
		"lowercase":    {input: "HEllO woRLD", want: []string{"hello", "world"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := cleanInput(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}

}

func TestHandleInput(t *testing.T) {
	var t1 int
	cb := func() error {
		t1 = 1
		return nil
	}
	commands := map[string]CliCommand{
		"test1": {name: "test1", description: "test 1 desc", callback: cb},
	}
	type Case struct {
		inp        string
		commands   map[string]CliCommand
		wantOutput bool
	}
	cases := []Case{
		{inp: "test1", commands: commands, wantOutput: false},
		{inp: "test_1", commands: commands, wantOutput: true},
	}
	tests := map[string]Case{
		"runs callback":                     cases[0],
		"invalid command returns something": cases[1],
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := HandleInput(tc.inp, tc.commands)
			if tc.wantOutput && got == "" {
				t.Fatalf("%s did not receive error message", name)
			}
			if name == "runs callback" && t1 != 1 {
				t.Fatalf("%s did not run callback", name)
			}
		})
	}
}
