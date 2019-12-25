package users

type inMemoryStore struct {
	users []User
}

func newInMemoryStore() *inMemoryStore {
	return &inMemoryStore{users: []User{}}
}

//CreateUser implements store method
func (i *inMemoryStore) CreateUser(user User) (User, error) {
	maxID := i.findMaxID()
	user.ID = maxID + 1
	i.users = append(i.users, user)
	return user, nil
}

func (i *inMemoryStore) findMaxID() int64 {
	maxID := int64(0)
	for _, u := range i.users {
		if u.ID > maxID {
			maxID = u.ID
		}
	}
	return maxID
}
