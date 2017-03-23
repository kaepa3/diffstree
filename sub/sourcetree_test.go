package sourcetree

import (
	"testing"
)

func TestCheckArg(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"Few", args{[]string{"aho"}}, false},
		{"Few", args{[]string{"aho", "hoge"}}, true},
		{"Few", args{[]string{"aho", "hoge", "karanao"}}, true},
		{"Few", args{[]string{"aho", "hoge", "karanao", "heteno"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckArg(tt.args.args); got != tt.want {
				t.Errorf("CheckArg() = %v, want %v", got, tt.want)
			}
		})
	}
}
