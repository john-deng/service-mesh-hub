// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1 (interfaces: SettingsClient)

// Package mock_zephyr_settings_clients is a generated GoMock package.
package mock_zephyr_settings_clients

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/core.zephyr.solo.io/v1alpha1"
	types "k8s.io/apimachinery/pkg/types"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockSettingsClient is a mock of SettingsClient interface.
type MockSettingsClient struct {
	ctrl     *gomock.Controller
	recorder *MockSettingsClientMockRecorder
}

// MockSettingsClientMockRecorder is the mock recorder for MockSettingsClient.
type MockSettingsClientMockRecorder struct {
	mock *MockSettingsClient
}

// NewMockSettingsClient creates a new mock instance.
func NewMockSettingsClient(ctrl *gomock.Controller) *MockSettingsClient {
	mock := &MockSettingsClient{ctrl: ctrl}
	mock.recorder = &MockSettingsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSettingsClient) EXPECT() *MockSettingsClientMockRecorder {
	return m.recorder
}

// CreateSettings mocks base method.
func (m *MockSettingsClient) CreateSettings(arg0 context.Context, arg1 *v1alpha1.Settings, arg2 ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateSettings", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSettings indicates an expected call of CreateSettings.
func (mr *MockSettingsClientMockRecorder) CreateSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSettings", reflect.TypeOf((*MockSettingsClient)(nil).CreateSettings), varargs...)
}

// DeleteAllOfSettings mocks base method.
func (m *MockSettingsClient) DeleteAllOfSettings(arg0 context.Context, arg1 ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfSettings", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfSettings indicates an expected call of DeleteAllOfSettings.
func (mr *MockSettingsClientMockRecorder) DeleteAllOfSettings(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfSettings", reflect.TypeOf((*MockSettingsClient)(nil).DeleteAllOfSettings), varargs...)
}

// DeleteSettings mocks base method.
func (m *MockSettingsClient) DeleteSettings(arg0 context.Context, arg1 types.NamespacedName, arg2 ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteSettings", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSettings indicates an expected call of DeleteSettings.
func (mr *MockSettingsClientMockRecorder) DeleteSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSettings", reflect.TypeOf((*MockSettingsClient)(nil).DeleteSettings), varargs...)
}

// GetSettings mocks base method.
func (m *MockSettingsClient) GetSettings(arg0 context.Context, arg1 types.NamespacedName) (*v1alpha1.Settings, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSettings", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.Settings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSettings indicates an expected call of GetSettings.
func (mr *MockSettingsClientMockRecorder) GetSettings(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSettings", reflect.TypeOf((*MockSettingsClient)(nil).GetSettings), arg0, arg1)
}

// ListSettings mocks base method.
func (m *MockSettingsClient) ListSettings(arg0 context.Context, arg1 ...client.ListOption) (*v1alpha1.SettingsList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSettings", varargs...)
	ret0, _ := ret[0].(*v1alpha1.SettingsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSettings indicates an expected call of ListSettings.
func (mr *MockSettingsClientMockRecorder) ListSettings(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSettings", reflect.TypeOf((*MockSettingsClient)(nil).ListSettings), varargs...)
}

// PatchSettings mocks base method.
func (m *MockSettingsClient) PatchSettings(arg0 context.Context, arg1 *v1alpha1.Settings, arg2 client.Patch, arg3 ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchSettings", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchSettings indicates an expected call of PatchSettings.
func (mr *MockSettingsClientMockRecorder) PatchSettings(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchSettings", reflect.TypeOf((*MockSettingsClient)(nil).PatchSettings), varargs...)
}

// PatchSettingsStatus mocks base method.
func (m *MockSettingsClient) PatchSettingsStatus(arg0 context.Context, arg1 *v1alpha1.Settings, arg2 client.Patch, arg3 ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchSettingsStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchSettingsStatus indicates an expected call of PatchSettingsStatus.
func (mr *MockSettingsClientMockRecorder) PatchSettingsStatus(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchSettingsStatus", reflect.TypeOf((*MockSettingsClient)(nil).PatchSettingsStatus), varargs...)
}

// UpdateSettings mocks base method.
func (m *MockSettingsClient) UpdateSettings(arg0 context.Context, arg1 *v1alpha1.Settings, arg2 ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSettings", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSettings indicates an expected call of UpdateSettings.
func (mr *MockSettingsClientMockRecorder) UpdateSettings(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSettings", reflect.TypeOf((*MockSettingsClient)(nil).UpdateSettings), varargs...)
}

// UpdateSettingsStatus mocks base method.
func (m *MockSettingsClient) UpdateSettingsStatus(arg0 context.Context, arg1 *v1alpha1.Settings, arg2 ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateSettingsStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateSettingsStatus indicates an expected call of UpdateSettingsStatus.
func (mr *MockSettingsClientMockRecorder) UpdateSettingsStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSettingsStatus", reflect.TypeOf((*MockSettingsClient)(nil).UpdateSettingsStatus), varargs...)
}
