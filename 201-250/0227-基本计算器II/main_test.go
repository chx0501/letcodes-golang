package _227_基本计算器II

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args:args{s:"3+2*2"},
			want:7,
		},
		{
			name: "t2",
			args:args{s:"3/2"},
			want:1,
		},
		{
			name: "t3",
			args:args{s:" 3+5 / 2 "},
			want:5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}