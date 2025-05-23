// Code generated by mockery v2.53.3. DO NOT EDIT.

package allowlist

import (
	context "context"

	models "github.com/goharbor/harbor/src/pkg/allowlist/models"
	mock "github.com/stretchr/testify/mock"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// CreateEmpty provides a mock function with given fields: ctx, projectID
func (_m *Manager) CreateEmpty(ctx context.Context, projectID int64) error {
	ret := _m.Called(ctx, projectID)

	if len(ret) == 0 {
		panic("no return value specified for CreateEmpty")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, projectID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, projectID
func (_m *Manager) Get(ctx context.Context, projectID int64) (*models.CVEAllowlist, error) {
	ret := _m.Called(ctx, projectID)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *models.CVEAllowlist
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*models.CVEAllowlist, error)); ok {
		return rf(ctx, projectID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.CVEAllowlist); ok {
		r0 = rf(ctx, projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.CVEAllowlist)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSys provides a mock function with given fields: ctx
func (_m *Manager) GetSys(ctx context.Context) (*models.CVEAllowlist, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetSys")
	}

	var r0 *models.CVEAllowlist
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*models.CVEAllowlist, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *models.CVEAllowlist); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.CVEAllowlist)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: ctx, projectID, list
func (_m *Manager) Set(ctx context.Context, projectID int64, list models.CVEAllowlist) error {
	ret := _m.Called(ctx, projectID, list)

	if len(ret) == 0 {
		panic("no return value specified for Set")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, models.CVEAllowlist) error); ok {
		r0 = rf(ctx, projectID, list)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetSys provides a mock function with given fields: ctx, list
func (_m *Manager) SetSys(ctx context.Context, list models.CVEAllowlist) error {
	ret := _m.Called(ctx, list)

	if len(ret) == 0 {
		panic("no return value specified for SetSys")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.CVEAllowlist) error); ok {
		r0 = rf(ctx, list)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
