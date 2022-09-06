package utils

import (
	"os"
	"testing"
)

func TestMustGetEnv(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "env set",
			args: args{
				key: "DATABASE_DNS",
			},
			want: "/home/ochom/go",
		},
		{
			name: "env not set",
			args: args{
				key: "REDIS_URL",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "env set" {
				os.Setenv(tt.args.key, tt.want)
			}
			if got := MustGetEnv(tt.args.key); got != tt.want {
				t.Errorf("MustGetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEnvOrDefault(t *testing.T) {
	type args struct {
		key          string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "env set",
			args: args{
				key:          "DATABASE_DNS",
				defaultValue: "hello world",
			},
			want: "/home/ochom/go",
		},
		{
			name: "env not set",
			args: args{
				key:          "REDIS_URL",
				defaultValue: "redis://localhost:6379",
			},
			want: "redis://localhost:6379",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "env set" {
				os.Setenv(tt.args.key, tt.want)
			}
			if got := GetEnvOrDefault(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetEnvOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}
