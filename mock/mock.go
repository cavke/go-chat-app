package mock

import (
	"github.com/cavke/go-chat-app"
)

// UserService represents a mock implementation of myapp.UserService.
type UserService struct {
	UserFn      func(id int) (*chatapp.User, error)
	UserInvoked bool

	UsersFn      func() ([]*chatapp.User, error)
	UsersInvoked bool

	// TODO implement additional functions
}

// User invokes the mock implementation and marks the function as invoked.
func (s *UserService) User(id int) (*chatapp.User, error) {
	s.UserInvoked = true
	return s.UserFn(id)
}

func (s *UserService) Users() ([]*chatapp.User, error) {
	s.UsersInvoked = true
	return s.UsersFn()
}

// TODO additional functions: Users(), CreateUser(), DeleteUser()
