package core

import (
	"testing"
)

func TestTitleFormat(t *testing.T) {
	type args struct {
		title   string
		authors string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TEST1",
			args: args{
				title:   "Tina Lynn Stark stephen hawking Drafting Contracts: How & Why Lawyers Do What They Do",
				authors: "Tina Lynn,Stark stephen hawking",
			},
			want: "Drafting Contracts: How & Why Lawyers Do What They Do",
		},
		{
			name: "TEST2",
			args: args{
				title:   "Ingrid Hotz, Thomas Schultz (eds.) Visualization and Processing of Higher Order Descriptors for MultiValu...",
				authors: "Ingrid Hotz Thomas Schultz (eds.)",
			},
			want: "Visualization and Processing of Higher Order Descriptors for MultiValu...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TitleFormat(tt.args.title, tt.args.authors); got != tt.want {
				t.Errorf("TitleFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
