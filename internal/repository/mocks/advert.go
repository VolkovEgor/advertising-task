// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/VolkovEgor/advertising-task/internal/repository (interfaces: Advert)

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"

	model "github.com/VolkovEgor/advertising-task/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockAdvert is a mock of Advert interface.
type MockAdvert struct {
	ctrl     *gomock.Controller
	recorder *MockAdvertMockRecorder
}

// MockAdvertMockRecorder is the mock recorder for MockAdvert.
type MockAdvertMockRecorder struct {
	mock *MockAdvert
}

// NewMockAdvert creates a new mock instance.
func NewMockAdvert(ctrl *gomock.Controller) *MockAdvert {
	mock := &MockAdvert{ctrl: ctrl}
	mock.recorder = &MockAdvertMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdvert) EXPECT() *MockAdvertMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAdvert) Create(arg0 *model.DetailedAdvert) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAdvertMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAdvert)(nil).Create), arg0)
}

// GetAll mocks base method.
func (m *MockAdvert) GetAll(arg0 int, arg1, arg2 string) ([]*model.Advert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*model.Advert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockAdvertMockRecorder) GetAll(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockAdvert)(nil).GetAll), arg0, arg1, arg2)
}

// GetById mocks base method.
func (m *MockAdvert) GetById(arg0 int, arg1 bool) (*model.DetailedAdvert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0, arg1)
	ret0, _ := ret[0].(*model.DetailedAdvert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockAdvertMockRecorder) GetById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockAdvert)(nil).GetById), arg0, arg1)
}
