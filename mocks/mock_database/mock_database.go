// Code generated by MockGen. DO NOT EDIT.
// Source: models/database.go

// Package mock_models is a generated GoMock package.
package mock_models

import (
	gomock "github.com/golang/mock/gomock"
	models "github.com/jamestjw/coup-vin/models"
	reflect "reflect"
)

// MockDatastore is a mock of Datastore interface
type MockDatastore struct {
	ctrl     *gomock.Controller
	recorder *MockDatastoreMockRecorder
}

// MockDatastoreMockRecorder is the mock recorder for MockDatastore
type MockDatastoreMockRecorder struct {
	mock *MockDatastore
}

// NewMockDatastore creates a new mock instance
func NewMockDatastore(ctrl *gomock.Controller) *MockDatastore {
	mock := &MockDatastore{ctrl: ctrl}
	mock.recorder = &MockDatastoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatastore) EXPECT() *MockDatastoreMockRecorder {
	return m.recorder
}

// AllJoinableRooms mocks base method
func (m *MockDatastore) AllJoinableRooms() ([]models.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllJoinableRooms")
	ret0, _ := ret[0].([]models.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllJoinableRooms indicates an expected call of AllJoinableRooms
func (mr *MockDatastoreMockRecorder) AllJoinableRooms() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllJoinableRooms", reflect.TypeOf((*MockDatastore)(nil).AllJoinableRooms))
}

// FindRoomByID mocks base method
func (m *MockDatastore) FindRoomByID(arg0 uint) (*models.Room, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindRoomByID", arg0)
	ret0, _ := ret[0].(*models.Room)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindRoomByID indicates an expected call of FindRoomByID
func (mr *MockDatastoreMockRecorder) FindRoomByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRoomByID", reflect.TypeOf((*MockDatastore)(nil).FindRoomByID), arg0)
}

// FindUserByUsername mocks base method
func (m *MockDatastore) FindUserByUsername(arg0 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByUsername", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByUsername indicates an expected call of FindUserByUsername
func (mr *MockDatastoreMockRecorder) FindUserByUsername(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByUsername", reflect.TypeOf((*MockDatastore)(nil).FindUserByUsername), arg0)
}

// FindUserByID mocks base method
func (m *MockDatastore) FindUserByID(arg0 uint) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByID", arg0)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByID indicates an expected call of FindUserByID
func (mr *MockDatastoreMockRecorder) FindUserByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByID", reflect.TypeOf((*MockDatastore)(nil).FindUserByID), arg0)
}

// UsernameExists mocks base method
func (m *MockDatastore) UsernameExists(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsernameExists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UsernameExists indicates an expected call of UsernameExists
func (mr *MockDatastoreMockRecorder) UsernameExists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsernameExists", reflect.TypeOf((*MockDatastore)(nil).UsernameExists), arg0)
}

// CreateUser mocks base method
func (m *MockDatastore) CreateUser(arg0, arg1 string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockDatastoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockDatastore)(nil).CreateUser), arg0, arg1)
}