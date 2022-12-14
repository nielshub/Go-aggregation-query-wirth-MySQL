// Code generated by MockGen. DO NOT EDIT.
// Source: contentSquare/src/internal/ports (interfaces: DBRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	models "contentSquare/src/internal/models"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDBRepository is a mock of DBRepository interface.
type MockDBRepository struct {
	ctrl     *gomock.Controller
	recorder *MockDBRepositoryMockRecorder
}

// MockDBRepositoryMockRecorder is the mock recorder for MockDBRepository.
type MockDBRepositoryMockRecorder struct {
	mock *MockDBRepository
}

// NewMockDBRepository creates a new mock instance.
func NewMockDBRepository(ctrl *gomock.Controller) *MockDBRepository {
	mock := &MockDBRepository{ctrl: ctrl}
	mock.recorder = &MockDBRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBRepository) EXPECT() *MockDBRepositoryMockRecorder {
	return m.recorder
}

// CountDistinctUsers mocks base method.
func (m *MockDBRepository) CountDistinctUsers(arg0 context.Context, arg1 models.Filters) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountDistinctUsers", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountDistinctUsers indicates an expected call of CountDistinctUsers.
func (mr *MockDBRepositoryMockRecorder) CountDistinctUsers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountDistinctUsers", reflect.TypeOf((*MockDBRepository)(nil).CountDistinctUsers), arg0, arg1)
}

// CountEvents mocks base method.
func (m *MockDBRepository) CountEvents(arg0 context.Context, arg1 models.Filters) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountEvents", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountEvents indicates an expected call of CountEvents.
func (mr *MockDBRepositoryMockRecorder) CountEvents(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountEvents", reflect.TypeOf((*MockDBRepository)(nil).CountEvents), arg0, arg1)
}

// Exists mocks base method.
func (m *MockDBRepository) Exists(arg0 context.Context, arg1 models.Filters) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockDBRepositoryMockRecorder) Exists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockDBRepository)(nil).Exists), arg0, arg1)
}
