package a5er

import "testing"

func TestLogical2Physical(t *testing.T) {
	type args struct {
		logicalName string
		dict        *Dictionary
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "マッチする",
			args: args{logicalName: "あ", dict: &Dictionary{
				{Key: "あ", Value: "a"},
			}},
			want: "a",
		},
		{
			name: "マッチする 複数",
			args: args{logicalName: "あい", dict: &Dictionary{
				{Key: "あ", Value: "a"},
				{Key: "い", Value: "i"},
			}},
			want: "a_i",
		},
		{
			name: "マッチする 複数",
			args: args{logicalName: "あい", dict: &Dictionary{
				{Key: "あい", Value: "love"},
				{Key: "あ", Value: "a"},
				{Key: "い", Value: "i"},
			}},
			want: "love",
		},
		{
			name: "マッチしない",
			args: args{logicalName: "あ", dict: &Dictionary{
				{Key: "あい", Value: "love"},
			}},
			want: "あ",
		},
		{
			name: "一部マッチ",
			args: args{logicalName: "あいう", dict: &Dictionary{
				{Key: "い", Value: "i"},
				{Key: "う", Value: "u"},
			}},
			want: "あいう",
		},
		{
			name: "一部マッチ",
			args: args{logicalName: "あいう", dict: &Dictionary{
				{Key: "い", Value: "i"},
			}},
			want: "あいう",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConvertor()
			if got := c.Logical2Physical(tt.args.logicalName, tt.args.dict); got != tt.want {
				t.Errorf("Logical2Physical() = %v, want %v", got, tt.want)
			}
		})
	}
}
