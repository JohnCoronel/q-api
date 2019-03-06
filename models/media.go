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

type Rating struct {
	id string
	score int
	userId string
}

type Review struct {
	id string
	review string
	userId string
}

func (Movie) Rate(score int) (error, int) {
	return nil, score
}
