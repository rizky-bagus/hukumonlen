// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/usecases/member_usecase.go

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	entity "hukum-onlen-go/entity"
)

// MockIMemberUsecase is a mock of IMemberUsecase interface.
type MockIMemberUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockIMemberUsecaseMockRecorder
}

// MockIMemberUsecaseMockRecorder is the mock recorder for MockIMemberUsecase.
type MockIMemberUsecaseMockRecorder struct {
	mock *MockIMemberUsecase
}

// NewMockIMemberUsecase creates a new mock instance.
func NewMockIMemberUsecase(ctrl *gomock.Controller) *MockIMemberUsecase {
	mock := &MockIMemberUsecase{ctrl: ctrl}
	mock.recorder = &MockIMemberUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMemberUsecase) EXPECT() *MockIMemberUsecaseMockRecorder {
	return m.recorder
}

// CreateMember mocks base method.
func (m *MockIMemberUsecase) CreateMember(ctx context.Context, email, firstName, lastName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMember", ctx, email, firstName, lastName)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMember indicates an expected call of CreateMember.
func (mr *MockIMemberUsecaseMockRecorder) CreateMember(ctx, email, firstName, lastName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMember", reflect.TypeOf((*MockIMemberUsecase)(nil).CreateMember), ctx, email, firstName, lastName)
}

// FindAllMembers mocks base method.
func (m *MockIMemberUsecase) FindAllMembers(ctx context.Context, email string) ([]*entity.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllMembers", ctx, email)
	ret0, _ := ret[0].([]*entity.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllMembers indicates an expected call of FindAllMembers.
func (mr *MockIMemberUsecaseMockRecorder) FindAllMembers(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllMembers", reflect.TypeOf((*MockIMemberUsecase)(nil).FindAllMembers), ctx, email)
}

// FindMemberByID mocks base method.
func (m *MockIMemberUsecase) FindMemberByID(ctx context.Context, id string) (*entity.Member, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindMemberByID", ctx, id)
	ret0, _ := ret[0].(*entity.Member)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindMemberByID indicates an expected call of FindMemberByID.
func (mr *MockIMemberUsecaseMockRecorder) FindMemberByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindMemberByID", reflect.TypeOf((*MockIMemberUsecase)(nil).FindMemberByID), ctx, id)
}
