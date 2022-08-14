package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/future-architect/a5er-dictionary/internal/cmd"
	"github.com/google/go-cmp/cmp"
)

func TestExample(t *testing.T) {
	t.Setenv("TABLE_PLURAL", "false")

	tests := []struct {
		name string
		in   string
		dict string
		want string
	}{
		{
			name: "normal",
			in:   filepath.Join("testdata", "in1.a5er"),
			dict: filepath.Join("testdata", "dict.txt"),
			want: filepath.Join("testdata", "want1.a5er"),
		},
		{
			name: "converted",
			in:   filepath.Join("testdata", "in2.a5er"),
			dict: filepath.Join("testdata", "dict.txt"),
			want: filepath.Join("testdata", "want2.a5er"),
		},
		{
			name: "mismatched",
			in:   filepath.Join("testdata", "in3.a5er"),
			dict: filepath.Join("testdata", "dict2.txt"),
			want: filepath.Join("testdata", "want3.a5er"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()

			t.Setenv("ERD_PATH", tt.in)
			t.Setenv("DICT_PATH", tt.dict)
			t.Setenv("OUTPUT_PATH", filepath.Join(tempDir, "out.a5er"))

			if err := cmd.Do(); err != nil {
				t.Fatalf("failed to cmd: %v", err)
			}

			got, err := os.ReadFile(filepath.Join(tempDir, "out.a5er"))
			if err != nil {
				t.Fatal(err)
			}

			want, err := os.ReadFile(tt.want)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(want, got); diff != "" {
				t.Errorf("A5er file is mismatch, (-want +got):\n%s", diff)
			}
		})
	}
}
