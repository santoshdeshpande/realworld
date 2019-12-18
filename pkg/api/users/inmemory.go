package users

type inMemoryStore struct {
	users []User
}

func newInMemoryStore() *inMemoryStore {
	return &inMemoryStore{users: []User{}}
}

//CreateUser implements store method
func (i *inMemoryStore) CreateUser(user User) (User, error) {
	i.users = append(i.users, user)
	return user, nil
}
