package main

type rps int8

const (
	Rock rps = iota
	Paper
	Scissors
)

type Play struct {
	Played   rps
	Weakness rps
}

func newPlay(code string) *Play {
	play := Play{}
	switch code {
	case "A", "X":
		play.Played = Rock
		play.Weakness = Paper
	case "B", "Y":
		play.Played = Paper
		play.Weakness = Scissors
	case "C", "Z":
		play.Played = Scissors
		play.Weakness = Rock
	}
	return &play
}
