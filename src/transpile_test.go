package src

import "testing"

func Test_transpile(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transpile(tt.args.code)
		})
	}
}
