// Copyright (c) 2017 Uber Technologies, Inc.
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

// Package mockstransport generated by mockery v1.0.0
package mockstransport

import mock "github.com/stretchr/testify/mock"
import transport "go.uber.org/yarpc/api/transport"

// ClientConfig is an autogenerated mock type for the ClientConfig type
type ClientConfig struct {
	mock.Mock
}

// Caller provides a mock function with given fields:
func (_m *ClientConfig) Caller() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetOnewayOutbound provides a mock function with given fields:
func (_m *ClientConfig) GetOnewayOutbound() transport.OnewayOutbound {
	ret := _m.Called()

	var r0 transport.OnewayOutbound
	if rf, ok := ret.Get(0).(func() transport.OnewayOutbound); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(transport.OnewayOutbound)
		}
	}

	return r0
}

// GetUnaryOutbound provides a mock function with given fields:
func (_m *ClientConfig) GetUnaryOutbound() transport.UnaryOutbound {
	ret := _m.Called()

	var r0 transport.UnaryOutbound
	if rf, ok := ret.Get(0).(func() transport.UnaryOutbound); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(transport.UnaryOutbound)
		}
	}

	return r0
}

// Service provides a mock function with given fields:
func (_m *ClientConfig) Service() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}