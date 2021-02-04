package upload

import (
	"reflect"
	"testing"
)

func Test_removeSkipFile(t *testing.T) {
	type args struct {
		files     []string
		skipFiles []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			"*.md",
			args{
				files:     []string{"README.md", "log/test.log", "test/test.md"},
				skipFiles: []string{`.*\.md`},
			},
			[]string{"log/test.log"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := removeSkipFile(tt.args.files, tt.args.skipFiles)
			if (err != nil) != tt.wantErr {
				t.Errorf("removeSkipFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeSkipFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
