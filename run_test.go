package a5er

import (
	"context"
	"github.com/gertd/go-pluralize"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	// 暫定的にスキップします
	t.Skip()

	type args struct {
		ctx context.Context
		c   *Config
	}
	tests := []struct {
		name     string
		args     args
		wantFile string
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				ctx: context.TODO(),
				c: &Config{
					InputA5erFilePath:       filepath.Join("testdata", "in.a5er"),
					InputDictionaryFilePath: filepath.Join("testdata", "dict.csv"),
					OutputA5erFilePath:      filepath.Join("testdata", "temp", "temp.a5er"),
					TablePlural:             true,
				},
			},
			wantFile: filepath.Join("testdata", "want.a5er"),
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.c.TablePlural {
				tt.args.c.PluralClient = pluralize.NewClient()
			}
			if err := Run(tt.args.ctx, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
