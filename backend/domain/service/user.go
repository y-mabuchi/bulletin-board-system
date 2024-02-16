package service

type UserService interface {
	Exist(id string) bool
}
