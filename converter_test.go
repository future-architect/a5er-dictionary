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

func TestConvertor_IsConverted(t *testing.T) {
	type args struct {
		logicalName string
		dict        *Dictionary
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "match(a,i,u)", args: args{logicalName: "a_i_u", dict: &Dictionary{{"あ", "a"}, {"い", "i"}, {"う", "u"}}}, want: true},
		{name: "match(a_i,u)", args: args{logicalName: "a_i_u", dict: &Dictionary{{"あい", "a_i"}, {"う", "u"}}}, want: true},
		{name: "match(a,i_u)", args: args{logicalName: "a_i_u", dict: &Dictionary{{"あ", "a"}, {"いう", "i_u"}}}, want: true},
		{name: "match(a_i_u)", args: args{logicalName: "a_i_u", dict: &Dictionary{{"あいう", "a_i_u"}}}, want: true},
		{name: "match(one word)", args: args{logicalName: "a", dict: &Dictionary{{"あ", "a"}}}, want: true},
		{name: "not match", args: args{logicalName: "a_i_u", dict: &Dictionary{}}, want: false},
		{name: "not match", args: args{logicalName: "a_i_u", dict: &Dictionary{{"あ", "a"}}}, want: false},
		{name: "not match", args: args{logicalName: "a_i_u", dict: &Dictionary{{"い", "i"}}}, want: false},
		{name: "not match", args: args{logicalName: "a_i_u", dict: &Dictionary{{"う", "u"}}}, want: false},
		{name: "not match", args: args{logicalName: "a_i_u", dict: &Dictionary{{"あ", "a"}, {"い", "i"}}}, want: false},
		{name: "not match", args: args{logicalName: "a_i_u", dict: &Dictionary{{"あ", "a"}, {"う", "u"}}}, want: false},
		{name: "not match", args: args{logicalName: "a_i_u", dict: &Dictionary{{"い", "i"}, {"う", "u"}}}, want: false},
		{name: "not match", args: args{logicalName: "a_i_u", dict: &Dictionary{{"あい", "a_i"}}}, want: false},
		{name: "not match", args: args{logicalName: "a_i_u", dict: &Dictionary{{"いう", "i_u"}}}, want: false},
		{name: "not match", args: args{logicalName: "a_i_u", dict: &Dictionary{{"あいうえ", "a_i_u_e"}}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewConvertor()
			if got := c.IsConverted(tt.args.logicalName, tt.args.dict); got != tt.want {
				t.Errorf("IsConverted() = %v, want %v", got, tt.want)
			}
		})
	}
}
