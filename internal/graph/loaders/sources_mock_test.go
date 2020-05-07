// Code generated by MockGen. DO NOT EDIT.
// Source: ./sources.go

// Package loaders_test is a generated GoMock package.
package loaders_test

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/syncromatics/kafmesh/internal/graph/model"
	reflect "reflect"
)

// MockSourceRepository is a mock of SourceRepository interface
type MockSourceRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSourceRepositoryMockRecorder
}

// MockSourceRepositoryMockRecorder is the mock recorder for MockSourceRepository
type MockSourceRepositoryMockRecorder struct {
	mock *MockSourceRepository
}

// NewMockSourceRepository creates a new mock instance
func NewMockSourceRepository(ctrl *gomock.Controller) *MockSourceRepository {
	mock := &MockSourceRepository{ctrl: ctrl}
	mock.recorder = &MockSourceRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSourceRepository) EXPECT() *MockSourceRepositoryMockRecorder {
	return m.recorder
}

// ComponentBySources mocks base method
func (m *MockSourceRepository) ComponentBySources(ctx context.Context, sources []int) ([]*model.Component, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComponentBySources", ctx, sources)
	ret0, _ := ret[0].([]*model.Component)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComponentBySources indicates an expected call of ComponentBySources
func (mr *MockSourceRepositoryMockRecorder) ComponentBySources(ctx, sources interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComponentBySources", reflect.TypeOf((*MockSourceRepository)(nil).ComponentBySources), ctx, sources)
}

// PodsBySources mocks base method
func (m *MockSourceRepository) PodsBySources(ctx context.Context, sources []int) ([][]*model.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PodsBySources", ctx, sources)
	ret0, _ := ret[0].([][]*model.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PodsBySources indicates an expected call of PodsBySources
func (mr *MockSourceRepositoryMockRecorder) PodsBySources(ctx, sources interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PodsBySources", reflect.TypeOf((*MockSourceRepository)(nil).PodsBySources), ctx, sources)
}

// TopicBySources mocks base method
func (m *MockSourceRepository) TopicBySources(ctx context.Context, sources []int) ([]*model.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TopicBySources", ctx, sources)
	ret0, _ := ret[0].([]*model.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TopicBySources indicates an expected call of TopicBySources
func (mr *MockSourceRepositoryMockRecorder) TopicBySources(ctx, sources interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopicBySources", reflect.TypeOf((*MockSourceRepository)(nil).TopicBySources), ctx, sources)
}