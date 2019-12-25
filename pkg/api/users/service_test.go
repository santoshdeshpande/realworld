package users

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessfulRegister(t *testing.T) {
	user := User{ID: 1, Email: "test@gmail.com", Password: "test123", Username: "abcd123"}
	service := newUserService()
	r, err := service.RegisterUser(user)
	assert.Equal(t, r, user, "The email addresses should be same")
	assert.NotEqual(t, r.ID, int64(0), "The id should not be nil")
	assert.Nil(t, err, "There should be no error")
}
