package a5er

import (
	"reflect"
	"testing"
)

func TestField_String(t *testing.T) {
	type fields struct {
		logicalName           string
		convertedPhysicalName string
		physicalName          string
		other                 string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "not converted",
			fields: fields{
				logicalName:  `"あい"`,
				physicalName: `"a_i"`,
				other:        `"*数値","NOT NULL",0,"","",$00FF0000,""`,
			},
			want: `"あい","a_i","*数値","NOT NULL",0,"","",$00FF0000,""`,
		},
		{
			name: "converted",
			fields: fields{
				logicalName:           `"あい"`,
				convertedPhysicalName: "love",
				physicalName:          `"a_i"`,
				other:                 `"*数値","NOT NULL",0,"","",$00FF0000,""`,
			},
			want: `"あい","love","*数値","NOT NULL",0,"","",$00FF0000,""`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Field{
				logicalName:           tt.fields.logicalName,
				convertedPhysicalName: tt.fields.convertedPhysicalName,
				physicalName:          tt.fields.physicalName,
				other:                 tt.fields.other,
			}
			if got := f.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEntity_extractFields(t *testing.T) {
	tests := []struct {
		name   string
		fields []string
		want   []Field
	}{
		{
			name: "",
			fields: []string{
				`"あい","a_i","*数値","NOT NULL",0,"","",$00FF0000,""`,
			},
			want: []Field{
				{
					logicalName:  `"あい"`,
					physicalName: `"a_i"`,
					other:        `"*数値","NOT NULL",0,"","",$00FF0000,""`,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Entity{fields: tt.fields}
			if got := e.extractFields(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
