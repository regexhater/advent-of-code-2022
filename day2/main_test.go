package main

import "testing"

func Test_roundResult(t *testing.T) {
	rock := newPlay("A")
	paper := newPlay("B")
	scissors := newPlay("C")
	type args struct {
		oponentPlay Play
		yourPlay    Play
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Rock vs Scissors",
			args{
				*rock,
				*scissors,
			},
			3,
		},
		{
			"Rock vs Paper",
			args{
				*rock,
				*paper,
			},
			8,
		},
		{
			"Rock vs Rock",
			args{
				*rock,
				*rock,
			},
			4,
		},
		{
			"Paper vs Scissors",
			args{
				*paper,
				*scissors,
			},
			9,
		},
		{
			"Paper vs Paper",
			args{
				*paper,
				*paper,
			},
			5,
		},
		{
			"Paper vs Rock",
			args{
				*paper,
				*rock,
			},
			1,
		},
		{
			"Scissors vs Scissors",
			args{
				*scissors,
				*scissors,
			},
			6,
		},
		{
			"Scissors vs Paper",
			args{
				*scissors,
				*paper,
			},
			2,
		},
		{
			"Scissors vs Rock",
			args{
				*scissors,
				*rock,
			},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roundResult(tt.args.oponentPlay, tt.args.yourPlay); got != tt.want {
				t.Errorf("roundResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
