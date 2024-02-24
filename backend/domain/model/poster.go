package model

type Poster struct {
	User User
}

func NewPoster(user User) (*Poster, error) {
	// TODO: exist check

	return &Poster{
		User: user,
	}, nil
}
