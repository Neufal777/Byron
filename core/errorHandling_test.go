package core

import "testing"

func TestErrorsHandling(t *testing.T) {
	type args struct {
		html string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				html: `
				<script nonce="FCz9gnI86pYzjdZBlKLI5g">window.Polymer=window.Polymer||{};window.Polymer.legacyOptimizations=true;window.Polymer.setPassiveTouchGestures=true;window.ShadyDOM={force:true,preferPerformance:true,noPatch:true};
				window.ShadyCSS = a database error has occurred {disableRuntime: true};</script><link rel="shortcut icon" href="https://www.youtube.com/s/desktop/07cfff5b/img/favicon.ico" type="image/x-icon">
				`,
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				html: `
				<script nonce="FCz9gnI86pYzjdZBlKLI5g">window.Polymer=window.Polymer||{};window.Polymer.legacyOptimizations=true;window.Polymer.setPassiveTouchGestures=true;window.ShadyDOM={force:true,preferPerformance:true,noPatch:true};
				window.ShadyCSS = 503 Service Temporarily Unavailable {disableRuntime: true};</script><link rel="shortcut icon" href="https://www.youtube.com/s/desktop/07cfff5b/img/favicon.ico" type="image/x-icon">
				`,
			},
			want: true,
		},
		{
			name: "Test 3",
			args: args{
				html: `
				<script nonce="FCz9gnI86pYzjdZBlKLI5g">window.Polymer=window.Polymer||{};window.Polymer.legacyOptimizations=true;window.Polymer.setPassiveTouchGestures=true;window.ShadyDOM={force:true,preferPerformance:true,noPatch:true};
				window.ShadyCSS = {disableRuntime: true};</script><link rel="shortcut icon" href="https://www.youtube.com/s/desktop/07cfff5b/img/favicon.ico" type="image/x-icon">
				`,
			},
			want: false,
		},
		{
			name: "Test 4",
			args: args{
				html: `
				<script nonce="FCz9gnI86pYzjdZBlKLI5g">window.Polymer=window.Polymer||{};window.Polymer.legacyOptimizations=true;window.Polymer.setPassiveTouchGestures=true;window.ShadyDOM={force:true,preferPerformance:true,noPatch:true};
				window.ShadyCSS = Could not connect to the database {disableRuntime: true};</script><link rel="shortcut icon" href="https://www.youtube.com/s/desktop/07cfff5b/img/favicon.ico" type="image/x-icon">
				`,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrorsHandling(tt.args.html); got != tt.want {
				t.Errorf("ErrorsHandling() = %v, want %v", got, tt.want)
			}
		})
	}
}
