// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repositories/invitation_status_repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	entity "hukum-onlen-go/entity"
)

// MockIInvitationStatusRepository is a mock of IInvitationStatusRepository interface.
type MockIInvitationStatusRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIInvitationStatusRepositoryMockRecorder
}

// MockIInvitationStatusRepositoryMockRecorder is the mock recorder for MockIInvitationStatusRepository.
type MockIInvitationStatusRepositoryMockRecorder struct {
	mock *MockIInvitationStatusRepository
}

// NewMockIInvitationStatusRepository creates a new mock instance.
func NewMockIInvitationStatusRepository(ctrl *gomock.Controller) *MockIInvitationStatusRepository {
	mock := &MockIInvitationStatusRepository{ctrl: ctrl}
	mock.recorder = &MockIInvitationStatusRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIInvitationStatusRepository) EXPECT() *MockIInvitationStatusRepositoryMockRecorder {
	return m.recorder
}

// CreateInvitationStatus mocks base method.
func (m *MockIInvitationStatusRepository) CreateInvitationStatus(ctx context.Context, invitationStatus *entity.InvitationStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInvitationStatus", ctx, invitationStatus)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateInvitationStatus indicates an expected call of CreateInvitationStatus.
func (mr *MockIInvitationStatusRepositoryMockRecorder) CreateInvitationStatus(ctx, invitationStatus interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInvitationStatus", reflect.TypeOf((*MockIInvitationStatusRepository)(nil).CreateInvitationStatus), ctx, invitationStatus)
}

// FindAllInvitationStatuses mocks base method.
func (m *MockIInvitationStatusRepository) FindAllInvitationStatuses(ctx context.Context) ([]*entity.InvitationStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllInvitationStatuses", ctx)
	ret0, _ := ret[0].([]*entity.InvitationStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllInvitationStatuses indicates an expected call of FindAllInvitationStatuses.
func (mr *MockIInvitationStatusRepositoryMockRecorder) FindAllInvitationStatuses(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllInvitationStatuses", reflect.TypeOf((*MockIInvitationStatusRepository)(nil).FindAllInvitationStatuses), ctx)
}

// FindInvitationStatusByID mocks base method.
func (m *MockIInvitationStatusRepository) FindInvitationStatusByID(ctx context.Context, id string) (*entity.InvitationStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindInvitationStatusByID", ctx, id)
	ret0, _ := ret[0].(*entity.InvitationStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindInvitationStatusByID indicates an expected call of FindInvitationStatusByID.
func (mr *MockIInvitationStatusRepositoryMockRecorder) FindInvitationStatusByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindInvitationStatusByID", reflect.TypeOf((*MockIInvitationStatusRepository)(nil).FindInvitationStatusByID), ctx, id)
}

// FindPendingInvitationStatus mocks base method.
func (m *MockIInvitationStatusRepository) FindPendingInvitationStatus(ctx context.Context) (*entity.InvitationStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPendingInvitationStatus", ctx)
	ret0, _ := ret[0].(*entity.InvitationStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPendingInvitationStatus indicates an expected call of FindPendingInvitationStatus.
func (mr *MockIInvitationStatusRepositoryMockRecorder) FindPendingInvitationStatus(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPendingInvitationStatus", reflect.TypeOf((*MockIInvitationStatusRepository)(nil).FindPendingInvitationStatus), ctx)
}