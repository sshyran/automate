// Code generated by protoc-gen-grpc-mock. DO NOT EDIT.
// source: components/automate-gateway/api/auth/teams/teams.proto

package teams

import (
	"context"

	version "github.com/chef/automate/api/external/common/version"
	request "github.com/chef/automate/components/automate-gateway/api/auth/teams/request"
	response "github.com/chef/automate/components/automate-gateway/api/auth/teams/response"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// verify that the mock satisfies the TeamsServer interface (at compile time)
var _ TeamsServer = &TeamsServerMock{}

// NewTeamsServerMock gives you a fresh instance of TeamsServerMock.
func NewTeamsServerMock() *TeamsServerMock {
	return &TeamsServerMock{validateRequests: true}
}

// NewTeamsServerMockWithoutValidation gives you a fresh instance of
// TeamsServerMock which does not attempt to validate requests before passing
// them to their respective '*Func'.
func NewTeamsServerMockWithoutValidation() *TeamsServerMock {
	return &TeamsServerMock{}
}

// TeamsServerMock is the mock-what-you-want struct that stubs all not-overridden
// methods with "not implemented" returns
type TeamsServerMock struct {
	validateRequests    bool
	GetVersionFunc      func(context.Context, *version.VersionInfoRequest) (*version.VersionInfo, error)
	GetTeamsFunc        func(context.Context, *request.GetTeamsReq) (*response.Teams, error)
	GetTeamFunc         func(context.Context, *request.GetTeamReq) (*response.GetTeamResp, error)
	CreateTeamFunc      func(context.Context, *request.CreateTeamReq) (*response.CreateTeamResp, error)
	UpdateTeamFunc      func(context.Context, *request.UpdateTeamReq) (*response.UpdateTeamResp, error)
	DeleteTeamFunc      func(context.Context, *request.DeleteTeamReq) (*response.DeleteTeamResp, error)
	GetUsersFunc        func(context.Context, *request.GetUsersReq) (*response.GetUsersResp, error)
	AddUsersFunc        func(context.Context, *request.AddUsersReq) (*response.AddUsersResp, error)
	RemoveUsersFunc     func(context.Context, *request.RemoveUsersReq) (*response.RemoveUsersResp, error)
	GetTeamsForUserFunc func(context.Context, *request.GetTeamsForUserReq) (*response.GetTeamsForUserResp, error)
}

func (m *TeamsServerMock) GetVersion(ctx context.Context, req *version.VersionInfoRequest) (*version.VersionInfo, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetVersionFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetVersion' not implemented")
}

func (m *TeamsServerMock) GetTeams(ctx context.Context, req *request.GetTeamsReq) (*response.Teams, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetTeamsFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetTeams' not implemented")
}

func (m *TeamsServerMock) GetTeam(ctx context.Context, req *request.GetTeamReq) (*response.GetTeamResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetTeamFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetTeam' not implemented")
}

func (m *TeamsServerMock) CreateTeam(ctx context.Context, req *request.CreateTeamReq) (*response.CreateTeamResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.CreateTeamFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'CreateTeam' not implemented")
}

func (m *TeamsServerMock) UpdateTeam(ctx context.Context, req *request.UpdateTeamReq) (*response.UpdateTeamResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.UpdateTeamFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'UpdateTeam' not implemented")
}

func (m *TeamsServerMock) DeleteTeam(ctx context.Context, req *request.DeleteTeamReq) (*response.DeleteTeamResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.DeleteTeamFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'DeleteTeam' not implemented")
}

func (m *TeamsServerMock) GetUsers(ctx context.Context, req *request.GetUsersReq) (*response.GetUsersResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetUsersFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetUsers' not implemented")
}

func (m *TeamsServerMock) AddUsers(ctx context.Context, req *request.AddUsersReq) (*response.AddUsersResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.AddUsersFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'AddUsers' not implemented")
}

func (m *TeamsServerMock) RemoveUsers(ctx context.Context, req *request.RemoveUsersReq) (*response.RemoveUsersResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.RemoveUsersFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'RemoveUsers' not implemented")
}

func (m *TeamsServerMock) GetTeamsForUser(ctx context.Context, req *request.GetTeamsForUserReq) (*response.GetTeamsForUserResp, error) {
	if msg, ok := interface{}(req).(interface{ Validate() error }); m.validateRequests && ok {
		if err := msg.Validate(); err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}
	if f := m.GetTeamsForUserFunc; f != nil {
		return f(ctx, req)
	}
	return nil, status.Error(codes.Internal, "mock: 'GetTeamsForUser' not implemented")
}

// Reset resets all overridden functions
func (m *TeamsServerMock) Reset() {
	m.GetVersionFunc = nil
	m.GetTeamsFunc = nil
	m.GetTeamFunc = nil
	m.CreateTeamFunc = nil
	m.UpdateTeamFunc = nil
	m.DeleteTeamFunc = nil
	m.GetUsersFunc = nil
	m.AddUsersFunc = nil
	m.RemoveUsersFunc = nil
	m.GetTeamsForUserFunc = nil
}
