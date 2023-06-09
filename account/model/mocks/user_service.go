package mocks

import (
	"context"

	"github.com/Marif226/words-learner/account/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

// MockUserService i
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	ret := m.Called(ctx, uid)

	var r0 *model.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}