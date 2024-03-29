// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ConfigUnmarshaler is an autogenerated mock type for the ConfigUnmarshaler type
type ConfigUnmarshaler struct {
	mock.Mock
}

// Unmarshal provides a mock function with given fields: content, v
func (_m *ConfigUnmarshaler) Unmarshal(content []byte, v interface{}) error {
	ret := _m.Called(content, v)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, interface{}) error); ok {
		r0 = rf(content, v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewConfigUnmarshaler interface {
	mock.TestingT
	Cleanup(func())
}

// NewConfigUnmarshaler creates a new instance of ConfigUnmarshaler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConfigUnmarshaler(t mockConstructorTestingTNewConfigUnmarshaler) *ConfigUnmarshaler {
	mock := &ConfigUnmarshaler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
