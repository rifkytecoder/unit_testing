package services

import "testing"

func Test_grade(t *testing.T) {
	type args struct {
		nilai int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test return C",
			args: args{
				nilai: 30,
			},
			want: "C",
		},
		{
			name: "test return B",
			args: args{
				nilai: 69,
			},
			want: "B",
		},
		{
			name: "test return A",
			args: args{
				nilai: 75,
			},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grade(tt.args.nilai); got != tt.want {
				t.Errorf("grade() = %v, want %v", got, tt.want)
			}
		})
	}
}
