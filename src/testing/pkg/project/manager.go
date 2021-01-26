// Code generated by mockery v2.1.0. DO NOT EDIT.

package project

import (
	context "context"

	models "github.com/goharbor/harbor/src/common/models"
	mock "github.com/stretchr/testify/mock"

	q "github.com/goharbor/harbor/src/lib/q"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx, query
func (_m *Manager) Count(ctx context.Context, query *q.Query) (int64, error) {
	ret := _m.Called(ctx, query)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) int64); ok {
		r0 = rf(ctx, query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
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

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *models.Project) int64); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
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

	var r0 *models.Project
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) *models.Project); ok {
		r0 = rf(ctx, idOrName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Project)
		}
	}

	var r1 error
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

	var r0 []*models.Project
	if rf, ok := ret.Get(0).(func(context.Context, *q.Query) []*models.Project); ok {
		r0 = rf(ctx, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *q.Query) error); ok {
		r1 = rf(ctx, query)
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

	var r0 []int
	if rf, ok := ret.Get(0).(func(context.Context, int64, int, ...int) []int); ok {
		r0 = rf(ctx, projectID, userID, groupIDs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, int, ...int) error); ok {
		r1 = rf(ctx, projectID, userID, groupIDs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
