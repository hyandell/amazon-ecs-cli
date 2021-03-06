// Copyright 2015-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/aws-sdk-go/service/sts/stsiface (interfaces: STSAPI)

package mock_stsiface

import (
	aws "github.com/aws/aws-sdk-go/aws"
	request "github.com/aws/aws-sdk-go/aws/request"
	sts "github.com/aws/aws-sdk-go/service/sts"
	gomock "github.com/golang/mock/gomock"
)

// Mock of STSAPI interface
type MockSTSAPI struct {
	ctrl     *gomock.Controller
	recorder *_MockSTSAPIRecorder
}

// Recorder for MockSTSAPI (not exported)
type _MockSTSAPIRecorder struct {
	mock *MockSTSAPI
}

func NewMockSTSAPI(ctrl *gomock.Controller) *MockSTSAPI {
	mock := &MockSTSAPI{ctrl: ctrl}
	mock.recorder = &_MockSTSAPIRecorder{mock}
	return mock
}

func (_m *MockSTSAPI) EXPECT() *_MockSTSAPIRecorder {
	return _m.recorder
}

func (_m *MockSTSAPI) AssumeRole(_param0 *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error) {
	ret := _m.ctrl.Call(_m, "AssumeRole", _param0)
	ret0, _ := ret[0].(*sts.AssumeRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) AssumeRole(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AssumeRole", arg0)
}

func (_m *MockSTSAPI) AssumeRoleRequest(_param0 *sts.AssumeRoleInput) (*request.Request, *sts.AssumeRoleOutput) {
	ret := _m.ctrl.Call(_m, "AssumeRoleRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.AssumeRoleOutput)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) AssumeRoleRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AssumeRoleRequest", arg0)
}

func (_m *MockSTSAPI) AssumeRoleWithContext(_param0 aws.Context, _param1 *sts.AssumeRoleInput, _param2 ...request.Option) (*sts.AssumeRoleOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "AssumeRoleWithContext", _s...)
	ret0, _ := ret[0].(*sts.AssumeRoleOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) AssumeRoleWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AssumeRoleWithContext", _s...)
}

func (_m *MockSTSAPI) AssumeRoleWithSAML(_param0 *sts.AssumeRoleWithSAMLInput) (*sts.AssumeRoleWithSAMLOutput, error) {
	ret := _m.ctrl.Call(_m, "AssumeRoleWithSAML", _param0)
	ret0, _ := ret[0].(*sts.AssumeRoleWithSAMLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) AssumeRoleWithSAML(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AssumeRoleWithSAML", arg0)
}

func (_m *MockSTSAPI) AssumeRoleWithSAMLRequest(_param0 *sts.AssumeRoleWithSAMLInput) (*request.Request, *sts.AssumeRoleWithSAMLOutput) {
	ret := _m.ctrl.Call(_m, "AssumeRoleWithSAMLRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.AssumeRoleWithSAMLOutput)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) AssumeRoleWithSAMLRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AssumeRoleWithSAMLRequest", arg0)
}

func (_m *MockSTSAPI) AssumeRoleWithSAMLWithContext(_param0 aws.Context, _param1 *sts.AssumeRoleWithSAMLInput, _param2 ...request.Option) (*sts.AssumeRoleWithSAMLOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "AssumeRoleWithSAMLWithContext", _s...)
	ret0, _ := ret[0].(*sts.AssumeRoleWithSAMLOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) AssumeRoleWithSAMLWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AssumeRoleWithSAMLWithContext", _s...)
}

func (_m *MockSTSAPI) AssumeRoleWithWebIdentity(_param0 *sts.AssumeRoleWithWebIdentityInput) (*sts.AssumeRoleWithWebIdentityOutput, error) {
	ret := _m.ctrl.Call(_m, "AssumeRoleWithWebIdentity", _param0)
	ret0, _ := ret[0].(*sts.AssumeRoleWithWebIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) AssumeRoleWithWebIdentity(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AssumeRoleWithWebIdentity", arg0)
}

func (_m *MockSTSAPI) AssumeRoleWithWebIdentityRequest(_param0 *sts.AssumeRoleWithWebIdentityInput) (*request.Request, *sts.AssumeRoleWithWebIdentityOutput) {
	ret := _m.ctrl.Call(_m, "AssumeRoleWithWebIdentityRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.AssumeRoleWithWebIdentityOutput)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) AssumeRoleWithWebIdentityRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AssumeRoleWithWebIdentityRequest", arg0)
}

func (_m *MockSTSAPI) AssumeRoleWithWebIdentityWithContext(_param0 aws.Context, _param1 *sts.AssumeRoleWithWebIdentityInput, _param2 ...request.Option) (*sts.AssumeRoleWithWebIdentityOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "AssumeRoleWithWebIdentityWithContext", _s...)
	ret0, _ := ret[0].(*sts.AssumeRoleWithWebIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) AssumeRoleWithWebIdentityWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AssumeRoleWithWebIdentityWithContext", _s...)
}

func (_m *MockSTSAPI) DecodeAuthorizationMessage(_param0 *sts.DecodeAuthorizationMessageInput) (*sts.DecodeAuthorizationMessageOutput, error) {
	ret := _m.ctrl.Call(_m, "DecodeAuthorizationMessage", _param0)
	ret0, _ := ret[0].(*sts.DecodeAuthorizationMessageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) DecodeAuthorizationMessage(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DecodeAuthorizationMessage", arg0)
}

func (_m *MockSTSAPI) DecodeAuthorizationMessageRequest(_param0 *sts.DecodeAuthorizationMessageInput) (*request.Request, *sts.DecodeAuthorizationMessageOutput) {
	ret := _m.ctrl.Call(_m, "DecodeAuthorizationMessageRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.DecodeAuthorizationMessageOutput)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) DecodeAuthorizationMessageRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DecodeAuthorizationMessageRequest", arg0)
}

func (_m *MockSTSAPI) DecodeAuthorizationMessageWithContext(_param0 aws.Context, _param1 *sts.DecodeAuthorizationMessageInput, _param2 ...request.Option) (*sts.DecodeAuthorizationMessageOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "DecodeAuthorizationMessageWithContext", _s...)
	ret0, _ := ret[0].(*sts.DecodeAuthorizationMessageOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) DecodeAuthorizationMessageWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DecodeAuthorizationMessageWithContext", _s...)
}

func (_m *MockSTSAPI) GetCallerIdentity(_param0 *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
	ret := _m.ctrl.Call(_m, "GetCallerIdentity", _param0)
	ret0, _ := ret[0].(*sts.GetCallerIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) GetCallerIdentity(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCallerIdentity", arg0)
}

func (_m *MockSTSAPI) GetCallerIdentityRequest(_param0 *sts.GetCallerIdentityInput) (*request.Request, *sts.GetCallerIdentityOutput) {
	ret := _m.ctrl.Call(_m, "GetCallerIdentityRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.GetCallerIdentityOutput)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) GetCallerIdentityRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCallerIdentityRequest", arg0)
}

func (_m *MockSTSAPI) GetCallerIdentityWithContext(_param0 aws.Context, _param1 *sts.GetCallerIdentityInput, _param2 ...request.Option) (*sts.GetCallerIdentityOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetCallerIdentityWithContext", _s...)
	ret0, _ := ret[0].(*sts.GetCallerIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) GetCallerIdentityWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCallerIdentityWithContext", _s...)
}

func (_m *MockSTSAPI) GetFederationToken(_param0 *sts.GetFederationTokenInput) (*sts.GetFederationTokenOutput, error) {
	ret := _m.ctrl.Call(_m, "GetFederationToken", _param0)
	ret0, _ := ret[0].(*sts.GetFederationTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) GetFederationToken(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFederationToken", arg0)
}

func (_m *MockSTSAPI) GetFederationTokenRequest(_param0 *sts.GetFederationTokenInput) (*request.Request, *sts.GetFederationTokenOutput) {
	ret := _m.ctrl.Call(_m, "GetFederationTokenRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.GetFederationTokenOutput)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) GetFederationTokenRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFederationTokenRequest", arg0)
}

func (_m *MockSTSAPI) GetFederationTokenWithContext(_param0 aws.Context, _param1 *sts.GetFederationTokenInput, _param2 ...request.Option) (*sts.GetFederationTokenOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetFederationTokenWithContext", _s...)
	ret0, _ := ret[0].(*sts.GetFederationTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) GetFederationTokenWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFederationTokenWithContext", _s...)
}

func (_m *MockSTSAPI) GetSessionToken(_param0 *sts.GetSessionTokenInput) (*sts.GetSessionTokenOutput, error) {
	ret := _m.ctrl.Call(_m, "GetSessionToken", _param0)
	ret0, _ := ret[0].(*sts.GetSessionTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) GetSessionToken(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSessionToken", arg0)
}

func (_m *MockSTSAPI) GetSessionTokenRequest(_param0 *sts.GetSessionTokenInput) (*request.Request, *sts.GetSessionTokenOutput) {
	ret := _m.ctrl.Call(_m, "GetSessionTokenRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*sts.GetSessionTokenOutput)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) GetSessionTokenRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSessionTokenRequest", arg0)
}

func (_m *MockSTSAPI) GetSessionTokenWithContext(_param0 aws.Context, _param1 *sts.GetSessionTokenInput, _param2 ...request.Option) (*sts.GetSessionTokenOutput, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetSessionTokenWithContext", _s...)
	ret0, _ := ret[0].(*sts.GetSessionTokenOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockSTSAPIRecorder) GetSessionTokenWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSessionTokenWithContext", _s...)
}
