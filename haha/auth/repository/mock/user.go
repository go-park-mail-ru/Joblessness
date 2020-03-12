// Code generated by MockGen. DO NOT EDIT.
// Source: joblessness/haha/auth (interfaces: UserRepository)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "joblessness/haha/models"
	reflect "reflect"
)

// MockUserRepository is a mock of UserRepository interface
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreatePerson mocks base method
func (m *MockUserRepository) CreatePerson(arg0 *models.Person) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePerson", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreatePerson indicates an expected call of CreatePerson
func (mr *MockUserRepositoryMockRecorder) CreatePerson(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePerson", reflect.TypeOf((*MockUserRepository)(nil).CreatePerson), arg0)
}

// Login mocks base method
func (m *MockUserRepository) Login(arg0, arg1, arg2 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockUserRepositoryMockRecorder) Login(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockUserRepository)(nil).Login), arg0, arg1, arg2)
}

// Logout mocks base method
func (m *MockUserRepository) Logout(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout
func (mr *MockUserRepositoryMockRecorder) Logout(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockUserRepository)(nil).Logout), arg0)
}

// SessionExists mocks base method
func (m *MockUserRepository) SessionExists(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SessionExists", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SessionExists indicates an expected call of SessionExists
func (mr *MockUserRepositoryMockRecorder) SessionExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SessionExists", reflect.TypeOf((*MockUserRepository)(nil).SessionExists), arg0)
}
