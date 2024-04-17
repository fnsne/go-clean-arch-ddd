// Code generated by mockery v2.36.0. DO NOT EDIT.

package mocks

import (
	chagnepassword "go-clean-arch-ddd/account/usecase/interface/in/changePassword"

	mock "github.com/stretchr/testify/mock"
)

// MockUseCase is an autogenerated mock type for the UseCase type
type MockUseCase struct {
	mock.Mock
}

type MockUseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUseCase) EXPECT() *MockUseCase_Expecter {
	return &MockUseCase_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: input
func (_m *MockUseCase) Execute(input chagnepassword.Input) error {
	ret := _m.Called(input)

	var r0 error
	if rf, ok := ret.Get(0).(func(chagnepassword.Input) error); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUseCase_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockUseCase_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - input chagnepassword.Input
func (_e *MockUseCase_Expecter) Execute(input interface{}) *MockUseCase_Execute_Call {
	return &MockUseCase_Execute_Call{Call: _e.mock.On("Execute", input)}
}

func (_c *MockUseCase_Execute_Call) Run(run func(input chagnepassword.Input)) *MockUseCase_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(chagnepassword.Input))
	})
	return _c
}

func (_c *MockUseCase_Execute_Call) Return(_a0 error) *MockUseCase_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUseCase_Execute_Call) RunAndReturn(run func(chagnepassword.Input) error) *MockUseCase_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUseCase creates a new instance of MockUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUseCase {
	mock := &MockUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
