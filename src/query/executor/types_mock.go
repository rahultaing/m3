// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/query/executor (interfaces: Engine)

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package executor is a generated GoMock package.
package executor

import (
	"context"
	"reflect"

	"github.com/m3db/m3/src/query/models"
	"github.com/m3db/m3/src/query/parser"
	"github.com/m3db/m3/src/query/storage"

	"github.com/golang/mock/gomock"
)

// MockEngine is a mock of Engine interface
type MockEngine struct {
	ctrl     *gomock.Controller
	recorder *MockEngineMockRecorder
}

// MockEngineMockRecorder is the mock recorder for MockEngine
type MockEngineMockRecorder struct {
	mock *MockEngine
}

// NewMockEngine creates a new mock instance
func NewMockEngine(ctrl *gomock.Controller) *MockEngine {
	mock := &MockEngine{ctrl: ctrl}
	mock.recorder = &MockEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEngine) EXPECT() *MockEngineMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockEngine) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockEngineMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockEngine)(nil).Close))
}

// ExecuteExpr mocks base method
func (m *MockEngine) ExecuteExpr(arg0 context.Context, arg1 parser.Parser, arg2 *QueryOptions, arg3 *storage.FetchOptions, arg4 models.RequestParams) (Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteExpr", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteExpr indicates an expected call of ExecuteExpr
func (mr *MockEngineMockRecorder) ExecuteExpr(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteExpr", reflect.TypeOf((*MockEngine)(nil).ExecuteExpr), arg0, arg1, arg2, arg3, arg4)
}

// ExecuteProm mocks base method
func (m *MockEngine) ExecuteProm(arg0 context.Context, arg1 *storage.FetchQuery, arg2 *QueryOptions, arg3 *storage.FetchOptions) (storage.PromResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecuteProm", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(storage.PromResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExecuteProm indicates an expected call of ExecuteProm
func (mr *MockEngineMockRecorder) ExecuteProm(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteProm", reflect.TypeOf((*MockEngine)(nil).ExecuteProm), arg0, arg1, arg2, arg3)
}

// Options mocks base method
func (m *MockEngine) Options() EngineOptions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Options")
	ret0, _ := ret[0].(EngineOptions)
	return ret0
}

// Options indicates an expected call of Options
func (mr *MockEngineMockRecorder) Options() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Options", reflect.TypeOf((*MockEngine)(nil).Options))
}