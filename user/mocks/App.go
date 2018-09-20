// Copyright 2018 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package mocks

import context "context"
import jwt "github.com/mendersoftware/useradm/jwt"
import mock "github.com/stretchr/testify/mock"
import model "github.com/mendersoftware/useradm/model"

// App is an autogenerated mock type for the App type
type App struct {
	mock.Mock
}

// CreateTenant provides a mock function with given fields: ctx, tenant
func (_m *App) CreateTenant(ctx context.Context, tenant model.NewTenant) error {
	ret := _m.Called(ctx, tenant)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.NewTenant) error); ok {
		r0 = rf(ctx, tenant)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: ctx, u
func (_m *App) CreateUser(ctx context.Context, u *model.User) error {
	ret := _m.Called(ctx, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUserInternal provides a mock function with given fields: ctx, u
func (_m *App) CreateUserInternal(ctx context.Context, u *model.UserInternal) error {
	ret := _m.Called(ctx, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.UserInternal) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTokens provides a mock function with given fields: ctx, tenantId, userId
func (_m *App) DeleteTokens(ctx context.Context, tenantId string, userId string) error {
	ret := _m.Called(ctx, tenantId, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, tenantId, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *App) DeleteUser(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUser provides a mock function with given fields: ctx, id
func (_m *App) GetUser(ctx context.Context, id string) (*model.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: ctx
func (_m *App) GetUsers(ctx context.Context) ([]model.User, error) {
	ret := _m.Called(ctx)

	var r0 []model.User
	if rf, ok := ret.Get(0).(func(context.Context) []model.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
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

// Login provides a mock function with given fields: ctx, email, pass
func (_m *App) Login(ctx context.Context, email string, pass string) (*jwt.Token, error) {
	ret := _m.Called(ctx, email, pass)

	var r0 *jwt.Token
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *jwt.Token); ok {
		r0 = rf(ctx, email, pass)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*jwt.Token)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, pass)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetPassword provides a mock function with given fields: ctx, u
func (_m *App) SetPassword(ctx context.Context, u model.UserUpdate) error {
	ret := _m.Called(ctx, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.UserUpdate) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SignToken provides a mock function with given fields: ctx, t
func (_m *App) SignToken(ctx context.Context, t *jwt.Token) (string, error) {
	ret := _m.Called(ctx, t)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *jwt.Token) string); ok {
		r0 = rf(ctx, t)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *jwt.Token) error); ok {
		r1 = rf(ctx, t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, id, u
func (_m *App) UpdateUser(ctx context.Context, id string, u *model.UserUpdate) error {
	ret := _m.Called(ctx, id, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.UserUpdate) error); ok {
		r0 = rf(ctx, id, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Verify provides a mock function with given fields: ctx, token
func (_m *App) Verify(ctx context.Context, token *jwt.Token) error {
	ret := _m.Called(ctx, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *jwt.Token) error); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
