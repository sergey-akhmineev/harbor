// Code generated by mockery v2.53.3. DO NOT EDIT.

package project

import (
	context "context"

	commonmodels "github.com/goharbor/harbor/src/common/models"

	mock "github.com/stretchr/testify/mock"

	models "github.com/goharbor/harbor/src/pkg/project/models"

	q "github.com/goharbor/harbor/src/lib/q"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, query
func (_m *Manager) Count(ctx context.Context, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for Count")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) (int64, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) int64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ctx, _a1
func (_m *Manager) Create(ctx context.Context, _a1 *models.Project) (int64, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Project) (int64, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.Project) int64); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.Project) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Manager) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, idOrName
func (_m *Manager) Get(ctx context.Context, idOrName interface{}) (*models.Project, error) {
	ret := _m.Called(ctx, idOrName)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *models.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) (*models.Project, error)); ok {
		return rf(ctx, idOrName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) *models.Project); ok {
		r0 = rf(ctx, idOrName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		r1 = rf(ctx, idOrName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query
func (_m *Manager) List(ctx context.Context, query *q.Query) ([]*models.Project, error) {
	ret := _m.Called(ctx, query)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*models.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) ([]*models.Project, error)); ok {
		return rf(ctx, query)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*models.Project); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Project)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAdminRolesOfUser provides a mock function with given fields: ctx, user
func (_m *Manager) ListAdminRolesOfUser(ctx context.Context, user commonmodels.User) ([]models.Member, error) {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for ListAdminRolesOfUser")
	}

	var r0 []models.Member
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, commonmodels.User) ([]models.Member, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, commonmodels.User) []models.Member); ok {
		r0 = rf(ctx, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Member)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, commonmodels.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListRoles provides a mock function with given fields: ctx, projectID, userID, groupIDs
func (_m *Manager) ListRoles(ctx context.Context, projectID int64, userID int, groupIDs ...int) ([]int, error) {
	_va := make([]interface{}, len(groupIDs))
	for _i := range groupIDs {
		_va[_i] = groupIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, projectID, userID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ListRoles")
	}

	var r0 []int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int, ...int) ([]int, error)); ok {
		return rf(ctx, projectID, userID, groupIDs...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int, ...int) []int); ok {
		r0 = rf(ctx, projectID, userID, groupIDs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int, ...int) error); ok {
		r1 = rf(ctx, projectID, userID, groupIDs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
