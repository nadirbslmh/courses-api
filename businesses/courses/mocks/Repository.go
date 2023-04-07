// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	courses "courses-api/businesses/courses"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, courseDomain
func (_m *Repository) Create(ctx context.Context, courseDomain *courses.Domain) (courses.Domain, error) {
	ret := _m.Called(ctx, courseDomain)

	var r0 courses.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *courses.Domain) courses.Domain); ok {
		r0 = rf(ctx, courseDomain)
	} else {
		r0 = ret.Get(0).(courses.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *courses.Domain) error); ok {
		r1 = rf(ctx, courseDomain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Repository) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ForceDelete provides a mock function with given fields: ctx, id
func (_m *Repository) ForceDelete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields: ctx
func (_m *Repository) GetAll(ctx context.Context) ([]courses.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []courses.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []courses.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]courses.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Repository) GetByID(ctx context.Context, id string) (courses.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 courses.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) courses.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(courses.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Restore provides a mock function with given fields: ctx, id
func (_m *Repository) Restore(ctx context.Context, id string) (courses.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 courses.Domain
	if rf, ok := ret.Get(0).(func(context.Context, string) courses.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(courses.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, courseDomain, id
func (_m *Repository) Update(ctx context.Context, courseDomain *courses.Domain, id string) (courses.Domain, error) {
	ret := _m.Called(ctx, courseDomain, id)

	var r0 courses.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *courses.Domain, string) courses.Domain); ok {
		r0 = rf(ctx, courseDomain, id)
	} else {
		r0 = ret.Get(0).(courses.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *courses.Domain, string) error); ok {
		r1 = rf(ctx, courseDomain, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
