// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	admins "main.go/business/admins"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// DeleteAdmin provides a mock function with given fields: ctx, id
func (_m *Repository) DeleteAdmin(ctx context.Context, id int) (admins.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 admins.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) admins.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(admins.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAdminById provides a mock function with given fields: ctx, id
func (_m *Repository) GetAdminById(ctx context.Context, id int) (admins.Domain, error) {
	ret := _m.Called(ctx, id)

	var r0 admins.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) admins.Domain); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(admins.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllAdmin provides a mock function with given fields: ctx
func (_m *Repository) GetAllAdmin(ctx context.Context) ([]admins.Domain, error) {
	ret := _m.Called(ctx)

	var r0 []admins.Domain
	if rf, ok := ret.Get(0).(func(context.Context) []admins.Domain); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]admins.Domain)
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

// HardDeleteAllAdmins provides a mock function with given fields: ctx
func (_m *Repository) HardDeleteAllAdmins(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Login provides a mock function with given fields: ctx, domain
func (_m *Repository) Login(ctx context.Context, domain admins.Domain) (admins.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 admins.Domain
	if rf, ok := ret.Get(0).(func(context.Context, admins.Domain) admins.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(admins.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, admins.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterAdmin provides a mock function with given fields: ctx, domain
func (_m *Repository) RegisterAdmin(ctx context.Context, domain admins.Domain) (admins.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 admins.Domain
	if rf, ok := ret.Get(0).(func(context.Context, admins.Domain) admins.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(admins.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, admins.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAdmin provides a mock function with given fields: ctx, domain
func (_m *Repository) UpdateAdmin(ctx context.Context, domain admins.Domain) (admins.Domain, error) {
	ret := _m.Called(ctx, domain)

	var r0 admins.Domain
	if rf, ok := ret.Get(0).(func(context.Context, admins.Domain) admins.Domain); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Get(0).(admins.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, admins.Domain) error); ok {
		r1 = rf(ctx, domain)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
