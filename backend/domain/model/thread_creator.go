package model

type ThreadCreator struct {
	User User
}

func NewThreadCreator(user User) (*ThreadCreator, error) {
	// TODO: exist check

	return &ThreadCreator{
		User: user,
	}, nil
}
