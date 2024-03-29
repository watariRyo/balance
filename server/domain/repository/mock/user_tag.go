// Code generated by MockGen. DO NOT EDIT.
// Source: user_tag.go

// Package repository is a generated GoMock package.
package repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/watariRyo/balance/server/domain/model"
	repository "github.com/watariRyo/balance/server/domain/repository"
	proto "github.com/watariRyo/balance/server/proto"
)

// MockUserTagRepository is a mock of UserTagRepository interface.
type MockUserTagRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserTagRepositoryMockRecorder
}

// MockUserTagRepositoryMockRecorder is the mock recorder for MockUserTagRepository.
type MockUserTagRepositoryMockRecorder struct {
	mock *MockUserTagRepository
}

// NewMockUserTagRepository creates a new mock instance.
func NewMockUserTagRepository(ctrl *gomock.Controller) *MockUserTagRepository {
	mock := &MockUserTagRepository{ctrl: ctrl}
	mock.recorder = &MockUserTagRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserTagRepository) EXPECT() *MockUserTagRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockUserTagRepository) Delete(ctx context.Context, conn repository.DBConnection, input *proto.DeleteUserTagRequest) (*proto.DeleteUserTagResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, conn, input)
	ret0, _ := ret[0].(*proto.DeleteUserTagResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockUserTagRepositoryMockRecorder) Delete(ctx, conn, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserTagRepository)(nil).Delete), ctx, conn, input)
}

// Get mocks base method.
func (m *MockUserTagRepository) Get(ctx context.Context, conn repository.DBConnection, input *proto.GetUserTagRequest) (*proto.GetUserTagResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, conn, input)
	ret0, _ := ret[0].(*proto.GetUserTagResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserTagRepositoryMockRecorder) Get(ctx, conn, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserTagRepository)(nil).Get), ctx, conn, input)
}

// InitInsert mocks base method.
func (m *MockUserTagRepository) InitInsert(ctx context.Context, conn repository.DBConnection, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitInsert", ctx, conn, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitInsert indicates an expected call of InitInsert.
func (mr *MockUserTagRepositoryMockRecorder) InitInsert(ctx, conn, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitInsert", reflect.TypeOf((*MockUserTagRepository)(nil).InitInsert), ctx, conn, userID)
}

// Insert mocks base method.
func (m *MockUserTagRepository) Insert(ctx context.Context, conn repository.DBConnection, input *proto.RegisterUserTagRequest) (*proto.RegisterUserTagResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", ctx, conn, input)
	ret0, _ := ret[0].(*proto.RegisterUserTagResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Insert indicates an expected call of Insert.
func (mr *MockUserTagRepositoryMockRecorder) Insert(ctx, conn, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockUserTagRepository)(nil).Insert), ctx, conn, input)
}

// List mocks base method.
func (m *MockUserTagRepository) List(tx context.Context, conn repository.DBConnection, userID string) ([]*model.UserTag, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", tx, conn, userID)
	ret0, _ := ret[0].([]*model.UserTag)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockUserTagRepositoryMockRecorder) List(tx, conn, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockUserTagRepository)(nil).List), tx, conn, userID)
}

// Update mocks base method.
func (m *MockUserTagRepository) Update(ctx context.Context, conn repository.DBConnection, input *proto.UpdateUserTagRequest) (*proto.UpdateUserTagResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, conn, input)
	ret0, _ := ret[0].(*proto.UpdateUserTagResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockUserTagRepositoryMockRecorder) Update(ctx, conn, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserTagRepository)(nil).Update), ctx, conn, input)
}
