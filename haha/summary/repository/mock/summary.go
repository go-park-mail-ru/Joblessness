// Code generated by MockGen. DO NOT EDIT.
// Source: joblessness/haha/summary/interfaces (interfaces: SummaryRepository)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	models "joblessness/haha/models"
	reflect "reflect"
)

// MockSummaryRepository is a mock of SummaryRepository interface
type MockSummaryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockSummaryRepositoryMockRecorder
}

// MockSummaryRepositoryMockRecorder is the mock recorder for MockSummaryRepository
type MockSummaryRepositoryMockRecorder struct {
	mock *MockSummaryRepository
}

// NewMockSummaryRepository creates a new mock instance
func NewMockSummaryRepository(ctrl *gomock.Controller) *MockSummaryRepository {
	mock := &MockSummaryRepository{ctrl: ctrl}
	mock.recorder = &MockSummaryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSummaryRepository) EXPECT() *MockSummaryRepositoryMockRecorder {
	return m.recorder
}

// ChangeSummary mocks base method
func (m *MockSummaryRepository) ChangeSummary(arg0 *models.Summary) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeSummary", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeSummary indicates an expected call of ChangeSummary
func (mr *MockSummaryRepositoryMockRecorder) ChangeSummary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeSummary", reflect.TypeOf((*MockSummaryRepository)(nil).ChangeSummary), arg0)
}

// CreateSummary mocks base method
func (m *MockSummaryRepository) CreateSummary(arg0 *models.Summary) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSummary", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSummary indicates an expected call of CreateSummary
func (mr *MockSummaryRepositoryMockRecorder) CreateSummary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSummary", reflect.TypeOf((*MockSummaryRepository)(nil).CreateSummary), arg0)
}

// DeleteSummary mocks base method
func (m *MockSummaryRepository) DeleteSummary(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSummary", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSummary indicates an expected call of DeleteSummary
func (mr *MockSummaryRepositoryMockRecorder) DeleteSummary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSummary", reflect.TypeOf((*MockSummaryRepository)(nil).DeleteSummary), arg0)
}

// GetAllSummaries mocks base method
func (m *MockSummaryRepository) GetAllSummaries(arg0 int) ([]models.Summary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSummaries", arg0)
	ret0, _ := ret[0].([]models.Summary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSummaries indicates an expected call of GetAllSummaries
func (mr *MockSummaryRepositoryMockRecorder) GetAllSummaries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSummaries", reflect.TypeOf((*MockSummaryRepository)(nil).GetAllSummaries), arg0)
}

// GetOrgSendSummaries mocks base method
func (m *MockSummaryRepository) GetOrgSendSummaries(arg0 uint64) (models.OrgSummaries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrgSendSummaries", arg0)
	ret0, _ := ret[0].(models.OrgSummaries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrgSendSummaries indicates an expected call of GetOrgSendSummaries
func (mr *MockSummaryRepositoryMockRecorder) GetOrgSendSummaries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrgSendSummaries", reflect.TypeOf((*MockSummaryRepository)(nil).GetOrgSendSummaries), arg0)
}

// GetSummary mocks base method
func (m *MockSummaryRepository) GetSummary(arg0 uint64) (*models.Summary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSummary", arg0)
	ret0, _ := ret[0].(*models.Summary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSummary indicates an expected call of GetSummary
func (mr *MockSummaryRepositoryMockRecorder) GetSummary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSummary", reflect.TypeOf((*MockSummaryRepository)(nil).GetSummary), arg0)
}

// GetUserSendSummaries mocks base method
func (m *MockSummaryRepository) GetUserSendSummaries(arg0 uint64) (models.OrgSummaries, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSendSummaries", arg0)
	ret0, _ := ret[0].(models.OrgSummaries)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSendSummaries indicates an expected call of GetUserSendSummaries
func (mr *MockSummaryRepositoryMockRecorder) GetUserSendSummaries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSendSummaries", reflect.TypeOf((*MockSummaryRepository)(nil).GetUserSendSummaries), arg0)
}

// GetUserSummaries mocks base method
func (m *MockSummaryRepository) GetUserSummaries(arg0 uint64) ([]models.Summary, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSummaries", arg0)
	ret0, _ := ret[0].([]models.Summary)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSummaries indicates an expected call of GetUserSummaries
func (mr *MockSummaryRepositoryMockRecorder) GetUserSummaries(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSummaries", reflect.TypeOf((*MockSummaryRepository)(nil).GetUserSummaries), arg0)
}

// IsOrganizationVacancy mocks base method
func (m *MockSummaryRepository) IsOrganizationVacancy(arg0, arg1 uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsOrganizationVacancy", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsOrganizationVacancy indicates an expected call of IsOrganizationVacancy
func (mr *MockSummaryRepositoryMockRecorder) IsOrganizationVacancy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsOrganizationVacancy", reflect.TypeOf((*MockSummaryRepository)(nil).IsOrganizationVacancy), arg0, arg1)
}

// IsPersonSummary mocks base method
func (m *MockSummaryRepository) IsPersonSummary(arg0, arg1 uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsPersonSummary", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsPersonSummary indicates an expected call of IsPersonSummary
func (mr *MockSummaryRepositoryMockRecorder) IsPersonSummary(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsPersonSummary", reflect.TypeOf((*MockSummaryRepository)(nil).IsPersonSummary), arg0, arg1)
}

// RefreshSummary mocks base method
func (m *MockSummaryRepository) RefreshSummary(arg0, arg1 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshSummary", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshSummary indicates an expected call of RefreshSummary
func (mr *MockSummaryRepositoryMockRecorder) RefreshSummary(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshSummary", reflect.TypeOf((*MockSummaryRepository)(nil).RefreshSummary), arg0, arg1)
}

// ResponseSummary mocks base method
func (m *MockSummaryRepository) ResponseSummary(arg0 *models.SendSummary) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResponseSummary", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResponseSummary indicates an expected call of ResponseSummary
func (mr *MockSummaryRepositoryMockRecorder) ResponseSummary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResponseSummary", reflect.TypeOf((*MockSummaryRepository)(nil).ResponseSummary), arg0)
}

// SendSummary mocks base method
func (m *MockSummaryRepository) SendSummary(arg0 *models.SendSummary) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSummary", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendSummary indicates an expected call of SendSummary
func (mr *MockSummaryRepositoryMockRecorder) SendSummary(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSummary", reflect.TypeOf((*MockSummaryRepository)(nil).SendSummary), arg0)
}
