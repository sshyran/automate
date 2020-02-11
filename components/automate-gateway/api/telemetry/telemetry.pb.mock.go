// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: components/automate-gateway/api/telemetry/telemetry.proto

package telemetry

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the TelemetryServer interface (at compile time)
var _ TelemetryServer = &TelemetryServerMock{}

// NewTelemetryServerMock gives you a fresh instance of TelemetryServerMock.
func NewTelemetryServerMock() *TelemetryServerMock {
	return &TelemetryServerMock{validateRequests: true}
}

// NewTelemetryServerMockWithoutValidation gives you a fresh instance of
// TelemetryServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewTelemetryServerMockWithoutValidation() *TelemetryServerMock {
	return &TelemetryServerMock{}
}

// TelemetryServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type TelemetryServerMock struct {
	validateRequests              bool
	GetTelemetryConfigurationFunc func(context.Context, *TelemetryRequest) (*TelemetryResponse, error)
}

func (m *TelemetryServerMock) GetTelemetryConfiguration(ctx context.Context, req *TelemetryRequest) (*TelemetryResponse, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetTelemetryConfigurationFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetTelemetryConfiguration' not implemented")
}

// Reset resets all overridden functions
func (m *TelemetryServerMock) Reset() {
	m.GetTelemetryConfigurationFunc = nil
}
