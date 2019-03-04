package models

type Media interface {
	Rate(score int) (error, int)
	Review(text string) error
	Share() error
}

type Movie struct {
	id    string
	title string
}

func (Movie) Rate(score int) (error, int) {
	return nil, score
}
