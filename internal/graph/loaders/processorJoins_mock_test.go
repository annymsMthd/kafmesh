// Code generated by MockGen. DO NOT EDIT.
// Source: ./processorJoins.go

// Package loaders_test is a generated GoMock package.
package loaders_test

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	model "github.com/syncromatics/kafmesh/internal/graph/model"
	reflect "reflect"
)

// MockProcessorJoinRepository is a mock of ProcessorJoinRepository interface
type MockProcessorJoinRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProcessorJoinRepositoryMockRecorder
}

// MockProcessorJoinRepositoryMockRecorder is the mock recorder for MockProcessorJoinRepository
type MockProcessorJoinRepositoryMockRecorder struct {
	mock *MockProcessorJoinRepository
}

// NewMockProcessorJoinRepository creates a new mock instance
func NewMockProcessorJoinRepository(ctrl *gomock.Controller) *MockProcessorJoinRepository {
	mock := &MockProcessorJoinRepository{ctrl: ctrl}
	mock.recorder = &MockProcessorJoinRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProcessorJoinRepository) EXPECT() *MockProcessorJoinRepositoryMockRecorder {
	return m.recorder
}

// ProcessorByJoins mocks base method
func (m *MockProcessorJoinRepository) ProcessorByJoins(ctx context.Context, joins []int) ([]*model.Processor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessorByJoins", ctx, joins)
	ret0, _ := ret[0].([]*model.Processor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessorByJoins indicates an expected call of ProcessorByJoins
func (mr *MockProcessorJoinRepositoryMockRecorder) ProcessorByJoins(ctx, joins interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessorByJoins", reflect.TypeOf((*MockProcessorJoinRepository)(nil).ProcessorByJoins), ctx, joins)
}

// TopicByJoins mocks base method
func (m *MockProcessorJoinRepository) TopicByJoins(ctx context.Context, joins []int) ([]*model.Topic, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TopicByJoins", ctx, joins)
	ret0, _ := ret[0].([]*model.Topic)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TopicByJoins indicates an expected call of TopicByJoins
func (mr *MockProcessorJoinRepositoryMockRecorder) TopicByJoins(ctx, joins interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopicByJoins", reflect.TypeOf((*MockProcessorJoinRepository)(nil).TopicByJoins), ctx, joins)
}