package main

import (
	"testing"
)

func Test_validatePattern(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{"hello"}, false},
		{"test3", args{"$.abc"}, true},
		{"test1", args{"$.abc[123].def"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validatePattern(tt.args.s); got != tt.want {
				t.Errorf("validatePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzValidatePattern(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		validatePattern(orig)
	})
}
