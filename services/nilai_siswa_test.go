package services

import "testing"

func Test_hasil(t *testing.T) {
	type args struct {
		nilai int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "Testing hasil C",
			args: args{
				nilai: 39,
			},
			want: "C"},

		{name: "Testing hasil B",
			args: args{
				nilai: 50,
			}, want: "B"},
		{name: "Testing hasil A",
			args: args{
				nilai: 95,
			}, want: "A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasil(tt.args.nilai); got != tt.want {
				t.Errorf("hasil() = %v, want %v", got, tt.want)
			}
		})
	}
}
