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
			name:     "fail",
			filename: "../../cmd/stupid/testdata/minimal.csv",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exercise(tt.filename)
		})
	}
}

func Test_parseCsv(t *testing.T) {
	type args struct {
		record []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		want2   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := parseCsv(tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCsv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseCsv() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseCsv() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("parseCsv() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
