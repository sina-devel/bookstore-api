package config

import (
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		path string
		cfg  *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "parse",
			args: args{
				path: "testdata/config.yaml",
				cfg:  &Config{},
			},
			wantErr: false,
		},
		{
			name: "error unknown file extension",
			args: args{
				path: "testdata/config.conf",
				cfg:  &Config{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Parse(tt.args.path, tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFileExtension(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "file extension",
			args: args{
				path: "testdata/config.yaml",
			},
			want: "yaml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileExtension(tt.args.path); got != tt.want {
				t.Errorf("fileExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}
