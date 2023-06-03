package fromFile

import (
	"testing"
)

func Test_exercise(t *testing.T) {
	tests := []struct {
		name     string
		filename string
	}{
		{
			name:     "minimal test from a file",
			filename: "../../cmd/stupid/testdata/minimal.csv",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exercise(tt.filename)
		})
	}
}
