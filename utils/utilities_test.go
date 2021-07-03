package utils

import (
	"testing"
)

func TestGetMD5Hash(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TEST 1",
			args: args{
				text: "naoufal",
			},
			want: "69509f85764b980a4274f298386d255a",
		},
		{
			name: "TEST 2",
			args: args{
				text: "byron",
			},
			want: "8720070ac8f94bbbff5a347eed656925",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMD5Hash(tt.args.text); got != tt.want {
				t.Errorf("GetMD5Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnyTypeToString(t *testing.T) {
	type args struct {
		input interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TEST 2",
			args: args{
				input: 3,
			},
			want: "3",
		},
		{
			name: "TEST 2",
			args: args{
				input: 3.76,
			},
			want: "3.76",
		},
		{
			name: "TEST 2",
			args: args{
				input: 3587,
			},
			want: "3587",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AnyTypeToString(tt.args.input); got != tt.want {
				t.Errorf("AnyTypeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFixUnitedNames(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TEST 1",
			args: args{
				name: "RichardBrandson",
			},
			want: "Richard Brandson",
		},
		{
			name: "TEST 2",
			args: args{
				name: "StephenHawkingJenifer",
			},
			want: "Stephen Hawking Jenifer",
		},
		{
			name: "TEST 2",
			args: args{
				name: "StephenHawking-Jenifer-Stone (Auth)",
			},
			want: "Stephen Hawking Jenifer Stone (Auth)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FixUnitedNames(tt.args.name); got != tt.want {
				t.Errorf("FixUnitedNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
