package ccgrep_test

import (
	"ccgrep/ccgrep"
	"reflect"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Empty",
			args: args{filepath: "../unit-test.txt"},
			want: []byte(""),
		},
		{
			name: "FileWithContent",
			args: args{filepath: "../unit-test-1.txt"},
			want: []byte("unit-test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ccgrep.ReadFromFile(tt.args.filepath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
