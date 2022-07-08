package a5er

import "testing"

func TestLogical2Physical(t *testing.T) {
	type args struct {
		logicalName string
		dict        *Dictionary
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want_ok bool
	}{
		{
			name: "マッチする",
			args: args{logicalName: "あ", dict: &Dictionary{
				{Key: "あ", Value: "a"},
			}},
			want:    "a",
			want_ok: true,
		},
		{
			name: "マッチする 複数",
			args: args{logicalName: "あい", dict: &Dictionary{
				{Key: "あ", Value: "a"},
				{Key: "い", Value: "i"},
			}},
			want:    "a_i",
			want_ok: true,
		},
		{
			name: "マッチする 複数",
			args: args{logicalName: "あい", dict: &Dictionary{
				{Key: "あい", Value: "love"},
				{Key: "あ", Value: "a"},
				{Key: "い", Value: "i"},
			}},
			want:    "love",
			want_ok: true,
		},
		{
			name: "マッチしない",
			args: args{logicalName: "あ", dict: &Dictionary{
				{Key: "あい", Value: "love"},
			}},
			want:    "",
			want_ok: false,
		},
		{
			name: "一部マッチ",
			args: args{logicalName: "あいう", dict: &Dictionary{
				{Key: "い", Value: "i"},
				{Key: "う", Value: "u"},
			}},
			want:    "",
			want_ok: false,
		},
		{
			name: "一部マッチ",
			args: args{logicalName: "あいう", dict: &Dictionary{
				{Key: "い", Value: "i"},
			}},
			want:    "",
			want_ok: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConvertor()
			got, ok := c.Logical2Physical(tt.args.logicalName, tt.args.dict)
			if ok != tt.want_ok {
				t.Errorf("got %v, want %v", ok, tt.want_ok)
			}
			if got != tt.want {
				t.Errorf("Logical2Physical() = %v, want %v", got, tt.want)
			}
		})
	}
}
