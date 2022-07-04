package a5er

import (
	"reflect"
	"testing"
)

func TestDictionary_sort(t *testing.T) {
	tests := []struct {
		name string
		d    Dictionary
		want Dictionary
	}{
		{
			name: "long name order",
			d: Dictionary{
				{Key: "あ", Value: "a"},
				{Key: "ああ", Value: "aa"},
			},
			want: Dictionary{
				{Key: "ああ", Value: "aa"},
				{Key: "あ", Value: "a"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.sort()
			if !reflect.DeepEqual(tt.d, tt.want) {
				t.Errorf("sort() = %v, want %v", tt.d, tt.want)
			}
		})
	}
}
