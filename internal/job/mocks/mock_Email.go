// Code generated by mockery v2.52.4. DO NOT EDIT.

package mocks

import (
	entity "gobdd/internal/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockEmail is an autogenerated mock type for the Email type
type MockEmail struct {
	mock.Mock
}

type MockEmail_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEmail) EXPECT() *MockEmail_Expecter {
	return &MockEmail_Expecter{mock: &_m.Mock}
}

// Send provides a mock function with given fields: email
func (_m *MockEmail) Send(email entity.Email) error {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for Send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.Email) error); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockEmail_Send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Send'
type MockEmail_Send_Call struct {
	*mock.Call
}

// Send is a helper method to define mock.On call
//   - email entity.Email
func (_e *MockEmail_Expecter) Send(email interface{}) *MockEmail_Send_Call {
	return &MockEmail_Send_Call{Call: _e.mock.On("Send", email)}
}

func (_c *MockEmail_Send_Call) Run(run func(email entity.Email)) *MockEmail_Send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(entity.Email))
	})
	return _c
}

func (_c *MockEmail_Send_Call) Return(_a0 error) *MockEmail_Send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockEmail_Send_Call) RunAndReturn(run func(entity.Email) error) *MockEmail_Send_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockEmail creates a new instance of MockEmail. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEmail(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEmail {
	mock := &MockEmail{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
