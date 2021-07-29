// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener-extension-provider-aws/pkg/aws/client (interfaces: Interface,Factory)

// Package client is a generated GoMock package.
package client

import (
	context "context"
	reflect "reflect"

	client "github.com/gardener/gardener-extension-provider-aws/pkg/aws/client"
	gomock "github.com/golang/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// CreateBucketIfNotExists mocks base method.
func (m *MockInterface) CreateBucketIfNotExists(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucketIfNotExists", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucketIfNotExists indicates an expected call of CreateBucketIfNotExists.
func (mr *MockInterfaceMockRecorder) CreateBucketIfNotExists(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucketIfNotExists", reflect.TypeOf((*MockInterface)(nil).CreateBucketIfNotExists), arg0, arg1, arg2)
}

// CreateOrUpdateDNSRecordSet mocks base method.
func (m *MockInterface) CreateOrUpdateDNSRecordSet(arg0 context.Context, arg1, arg2, arg3 string, arg4 []string, arg5 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdateDNSRecordSet", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdateDNSRecordSet indicates an expected call of CreateOrUpdateDNSRecordSet.
func (mr *MockInterfaceMockRecorder) CreateOrUpdateDNSRecordSet(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdateDNSRecordSet", reflect.TypeOf((*MockInterface)(nil).CreateOrUpdateDNSRecordSet), arg0, arg1, arg2, arg3, arg4, arg5)
}

// DeleteBucketIfExists mocks base method.
func (m *MockInterface) DeleteBucketIfExists(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBucketIfExists", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBucketIfExists indicates an expected call of DeleteBucketIfExists.
func (mr *MockInterfaceMockRecorder) DeleteBucketIfExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBucketIfExists", reflect.TypeOf((*MockInterface)(nil).DeleteBucketIfExists), arg0, arg1)
}

// DeleteDNSRecordSet mocks base method.
func (m *MockInterface) DeleteDNSRecordSet(arg0 context.Context, arg1, arg2, arg3 string, arg4 []string, arg5 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDNSRecordSet", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDNSRecordSet indicates an expected call of DeleteDNSRecordSet.
func (mr *MockInterfaceMockRecorder) DeleteDNSRecordSet(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDNSRecordSet", reflect.TypeOf((*MockInterface)(nil).DeleteDNSRecordSet), arg0, arg1, arg2, arg3, arg4, arg5)
}

// DeleteELB mocks base method.
func (m *MockInterface) DeleteELB(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteELB", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteELB indicates an expected call of DeleteELB.
func (mr *MockInterfaceMockRecorder) DeleteELB(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteELB", reflect.TypeOf((*MockInterface)(nil).DeleteELB), arg0, arg1)
}

// DeleteELBV2 mocks base method.
func (m *MockInterface) DeleteELBV2(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteELBV2", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteELBV2 indicates an expected call of DeleteELBV2.
func (mr *MockInterfaceMockRecorder) DeleteELBV2(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteELBV2", reflect.TypeOf((*MockInterface)(nil).DeleteELBV2), arg0, arg1)
}

// DeleteObjectsWithPrefix mocks base method.
func (m *MockInterface) DeleteObjectsWithPrefix(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteObjectsWithPrefix", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteObjectsWithPrefix indicates an expected call of DeleteObjectsWithPrefix.
func (mr *MockInterfaceMockRecorder) DeleteObjectsWithPrefix(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjectsWithPrefix", reflect.TypeOf((*MockInterface)(nil).DeleteObjectsWithPrefix), arg0, arg1, arg2)
}

// DeleteSecurityGroup mocks base method.
func (m *MockInterface) DeleteSecurityGroup(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecurityGroup", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecurityGroup indicates an expected call of DeleteSecurityGroup.
func (mr *MockInterfaceMockRecorder) DeleteSecurityGroup(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecurityGroup", reflect.TypeOf((*MockInterface)(nil).DeleteSecurityGroup), arg0, arg1)
}

// GetAccountID mocks base method.
func (m *MockInterface) GetAccountID(arg0 context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountID", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountID indicates an expected call of GetAccountID.
func (mr *MockInterfaceMockRecorder) GetAccountID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountID", reflect.TypeOf((*MockInterface)(nil).GetAccountID), arg0)
}

// GetDNSHostedZones mocks base method.
func (m *MockInterface) GetDNSHostedZones(arg0 context.Context) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDNSHostedZones", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDNSHostedZones indicates an expected call of GetDNSHostedZones.
func (mr *MockInterfaceMockRecorder) GetDNSHostedZones(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDNSHostedZones", reflect.TypeOf((*MockInterface)(nil).GetDNSHostedZones), arg0)
}

// GetInternetGateway mocks base method.
func (m *MockInterface) GetInternetGateway(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInternetGateway", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInternetGateway indicates an expected call of GetInternetGateway.
func (mr *MockInterfaceMockRecorder) GetInternetGateway(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInternetGateway", reflect.TypeOf((*MockInterface)(nil).GetInternetGateway), arg0, arg1)
}

// ListKubernetesELBs mocks base method.
func (m *MockInterface) ListKubernetesELBs(arg0 context.Context, arg1, arg2 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKubernetesELBs", arg0, arg1, arg2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKubernetesELBs indicates an expected call of ListKubernetesELBs.
func (mr *MockInterfaceMockRecorder) ListKubernetesELBs(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKubernetesELBs", reflect.TypeOf((*MockInterface)(nil).ListKubernetesELBs), arg0, arg1, arg2)
}

// ListKubernetesELBsV2 mocks base method.
func (m *MockInterface) ListKubernetesELBsV2(arg0 context.Context, arg1, arg2 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKubernetesELBsV2", arg0, arg1, arg2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKubernetesELBsV2 indicates an expected call of ListKubernetesELBsV2.
func (mr *MockInterfaceMockRecorder) ListKubernetesELBsV2(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKubernetesELBsV2", reflect.TypeOf((*MockInterface)(nil).ListKubernetesELBsV2), arg0, arg1, arg2)
}

// ListKubernetesSecurityGroups mocks base method.
func (m *MockInterface) ListKubernetesSecurityGroups(arg0 context.Context, arg1, arg2 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListKubernetesSecurityGroups", arg0, arg1, arg2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListKubernetesSecurityGroups indicates an expected call of ListKubernetesSecurityGroups.
func (mr *MockInterfaceMockRecorder) ListKubernetesSecurityGroups(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListKubernetesSecurityGroups", reflect.TypeOf((*MockInterface)(nil).ListKubernetesSecurityGroups), arg0, arg1, arg2)
}

// VerifyVPCAttributes mocks base method.
func (m *MockInterface) VerifyVPCAttributes(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyVPCAttributes", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyVPCAttributes indicates an expected call of VerifyVPCAttributes.
func (mr *MockInterfaceMockRecorder) VerifyVPCAttributes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyVPCAttributes", reflect.TypeOf((*MockInterface)(nil).VerifyVPCAttributes), arg0, arg1)
}

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// NewClient mocks base method.
func (m *MockFactory) NewClient(arg0, arg1, arg2 string) (client.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewClient", arg0, arg1, arg2)
	ret0, _ := ret[0].(client.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewClient indicates an expected call of NewClient.
func (mr *MockFactoryMockRecorder) NewClient(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewClient", reflect.TypeOf((*MockFactory)(nil).NewClient), arg0, arg1, arg2)
}