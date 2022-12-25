// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/headsync (interfaces: DiffSyncer)

// Package mock_headsync is a generated GoMock package.
package mock_headsync

import (
	context "context"
	reflect "reflect"

	deletionstate "github.com/anytypeio/go-anytype-infrastructure-experiments/common/commonspace/settings/deletionstate"
	gomock "github.com/golang/mock/gomock"
)

// MockDiffSyncer is a mock of DiffSyncer interface.
type MockDiffSyncer struct {
	ctrl     *gomock.Controller
	recorder *MockDiffSyncerMockRecorder
}

// MockDiffSyncerMockRecorder is the mock recorder for MockDiffSyncer.
type MockDiffSyncerMockRecorder struct {
	mock *MockDiffSyncer
}

// NewMockDiffSyncer creates a new mock instance.
func NewMockDiffSyncer(ctrl *gomock.Controller) *MockDiffSyncer {
	mock := &MockDiffSyncer{ctrl: ctrl}
	mock.recorder = &MockDiffSyncerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDiffSyncer) EXPECT() *MockDiffSyncerMockRecorder {
	return m.recorder
}

// Init mocks base method.
func (m *MockDiffSyncer) Init(arg0 deletionstate.DeletionState) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Init", arg0)
}

// Init indicates an expected call of Init.
func (mr *MockDiffSyncerMockRecorder) Init(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockDiffSyncer)(nil).Init), arg0)
}

// RemoveObjects mocks base method.
func (m *MockDiffSyncer) RemoveObjects(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveObjects", arg0)
}

// RemoveObjects indicates an expected call of RemoveObjects.
func (mr *MockDiffSyncerMockRecorder) RemoveObjects(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveObjects", reflect.TypeOf((*MockDiffSyncer)(nil).RemoveObjects), arg0)
}

// Sync mocks base method.
func (m *MockDiffSyncer) Sync(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sync", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Sync indicates an expected call of Sync.
func (mr *MockDiffSyncerMockRecorder) Sync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sync", reflect.TypeOf((*MockDiffSyncer)(nil).Sync), arg0)
}

// UpdateHeads mocks base method.
func (m *MockDiffSyncer) UpdateHeads(arg0 string, arg1 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateHeads", arg0, arg1)
}

// UpdateHeads indicates an expected call of UpdateHeads.
func (mr *MockDiffSyncerMockRecorder) UpdateHeads(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHeads", reflect.TypeOf((*MockDiffSyncer)(nil).UpdateHeads), arg0, arg1)
}