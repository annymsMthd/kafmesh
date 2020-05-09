// Code generated by MockGen. DO NOT EDIT.
// Source: ./query.go

// Package loaders_test is a generated GoMock package.
package loaders_test

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/syncromatics/kafmesh/internal/graph/model"
	reflect "reflect"
)

// MockQueryRepository is a mock of QueryRepository interface
type MockQueryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockQueryRepositoryMockRecorder
}

// MockQueryRepositoryMockRecorder is the mock recorder for MockQueryRepository
type MockQueryRepositoryMockRecorder struct {
	mock *MockQueryRepository
}

// NewMockQueryRepository creates a new mock instance
func NewMockQueryRepository(ctrl *gomock.Controller) *MockQueryRepository {
	mock := &MockQueryRepository{ctrl: ctrl}
	mock.recorder = &MockQueryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQueryRepository) EXPECT() *MockQueryRepositoryMockRecorder {
	return m.recorder
}

// GetAllServices mocks base method
func (m *MockQueryRepository) GetAllServices(arg0 context.Context) ([]*model.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllServices", arg0)
	ret0, _ := ret[0].([]*model.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllServices indicates an expected call of GetAllServices
func (mr *MockQueryRepositoryMockRecorder) GetAllServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllServices", reflect.TypeOf((*MockQueryRepository)(nil).GetAllServices), arg0)
}

// GetAllPods mocks base method
func (m *MockQueryRepository) GetAllPods(arg0 context.Context) ([]*model.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPods", arg0)
	ret0, _ := ret[0].([]*model.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPods indicates an expected call of GetAllPods
func (mr *MockQueryRepositoryMockRecorder) GetAllPods(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPods", reflect.TypeOf((*MockQueryRepository)(nil).GetAllPods), arg0)
}

// GetAllTopics mocks base method
func (m *MockQueryRepository) GetAllTopics(arg0 context.Context) ([]*model.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTopics", arg0)
	ret0, _ := ret[0].([]*model.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTopics indicates an expected call of GetAllTopics
func (mr *MockQueryRepositoryMockRecorder) GetAllTopics(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTopics", reflect.TypeOf((*MockQueryRepository)(nil).GetAllTopics), arg0)
}

// ServiceByID mocks base method
func (m *MockQueryRepository) ServiceByID(arg0 context.Context, arg1 int) (*model.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ServiceByID", arg0, arg1)
	ret0, _ := ret[0].(*model.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ServiceByID indicates an expected call of ServiceByID
func (mr *MockQueryRepositoryMockRecorder) ServiceByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServiceByID", reflect.TypeOf((*MockQueryRepository)(nil).ServiceByID), arg0, arg1)
}
