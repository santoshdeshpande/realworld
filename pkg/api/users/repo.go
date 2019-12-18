package users

type repo interface {
	CreateUser(User) (User, error)
}
