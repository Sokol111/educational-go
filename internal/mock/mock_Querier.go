// Code generated by mockery v2.38.0. DO NOT EDIT.

package mock

import (
	context "context"

	db "github.com/Sokol111/educational-go/internal/repository/db"
	mock "github.com/stretchr/testify/mock"
)

// MockQuerier is an autogenerated mock type for the Querier type
type MockQuerier struct {
	mock.Mock
}

type MockQuerier_Expecter struct {
	mock *mock.Mock
}

func (_m *MockQuerier) EXPECT() *MockQuerier_Expecter {
	return &MockQuerier_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: ctx, arg
func (_m *MockQuerier) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateUserParams) (db.User, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.CreateUserParams) db.User); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.CreateUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockQuerier_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.CreateUserParams
func (_e *MockQuerier_Expecter) CreateUser(ctx interface{}, arg interface{}) *MockQuerier_CreateUser_Call {
	return &MockQuerier_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, arg)}
}

func (_c *MockQuerier_CreateUser_Call) Run(run func(ctx context.Context, arg db.CreateUserParams)) *MockQuerier_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.CreateUserParams))
	})
	return _c
}

func (_c *MockQuerier_CreateUser_Call) Return(_a0 db.User, _a1 error) *MockQuerier_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_CreateUser_Call) RunAndReturn(run func(context.Context, db.CreateUserParams) (db.User, error)) *MockQuerier_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetAllUsers provides a mock function with given fields: ctx
func (_m *MockQuerier) GetAllUsers(ctx context.Context) ([]db.User, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllUsers")
	}

	var r0 []db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]db.User, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []db.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetAllUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAllUsers'
type MockQuerier_GetAllUsers_Call struct {
	*mock.Call
}

// GetAllUsers is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockQuerier_Expecter) GetAllUsers(ctx interface{}) *MockQuerier_GetAllUsers_Call {
	return &MockQuerier_GetAllUsers_Call{Call: _e.mock.On("GetAllUsers", ctx)}
}

func (_c *MockQuerier_GetAllUsers_Call) Run(run func(ctx context.Context)) *MockQuerier_GetAllUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockQuerier_GetAllUsers_Call) Return(_a0 []db.User, _a1 error) *MockQuerier_GetAllUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetAllUsers_Call) RunAndReturn(run func(context.Context) ([]db.User, error)) *MockQuerier_GetAllUsers_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserById provides a mock function with given fields: ctx, id
func (_m *MockQuerier) GetUserById(ctx context.Context, id int64) (db.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserById")
	}

	var r0 db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (db.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) db.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetUserById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserById'
type MockQuerier_GetUserById_Call struct {
	*mock.Call
}

// GetUserById is a helper method to define mock.On call
//   - ctx context.Context
//   - id int64
func (_e *MockQuerier_Expecter) GetUserById(ctx interface{}, id interface{}) *MockQuerier_GetUserById_Call {
	return &MockQuerier_GetUserById_Call{Call: _e.mock.On("GetUserById", ctx, id)}
}

func (_c *MockQuerier_GetUserById_Call) Run(run func(ctx context.Context, id int64)) *MockQuerier_GetUserById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockQuerier_GetUserById_Call) Return(_a0 db.User, _a1 error) *MockQuerier_GetUserById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetUserById_Call) RunAndReturn(run func(context.Context, int64) (db.User, error)) *MockQuerier_GetUserById_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByName provides a mock function with given fields: ctx, name
func (_m *MockQuerier) GetUserByName(ctx context.Context, name string) (db.User, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByName")
	}

	var r0 db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (db.User, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) db.User); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetUserByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByName'
type MockQuerier_GetUserByName_Call struct {
	*mock.Call
}

// GetUserByName is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockQuerier_Expecter) GetUserByName(ctx interface{}, name interface{}) *MockQuerier_GetUserByName_Call {
	return &MockQuerier_GetUserByName_Call{Call: _e.mock.On("GetUserByName", ctx, name)}
}

func (_c *MockQuerier_GetUserByName_Call) Run(run func(ctx context.Context, name string)) *MockQuerier_GetUserByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockQuerier_GetUserByName_Call) Return(_a0 db.User, _a1 error) *MockQuerier_GetUserByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetUserByName_Call) RunAndReturn(run func(context.Context, string) (db.User, error)) *MockQuerier_GetUserByName_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: ctx, arg
func (_m *MockQuerier) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 db.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateUserParams) (db.User, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, db.UpdateUserParams) db.User); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(db.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, db.UpdateUserParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type MockQuerier_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - arg db.UpdateUserParams
func (_e *MockQuerier_Expecter) UpdateUser(ctx interface{}, arg interface{}) *MockQuerier_UpdateUser_Call {
	return &MockQuerier_UpdateUser_Call{Call: _e.mock.On("UpdateUser", ctx, arg)}
}

func (_c *MockQuerier_UpdateUser_Call) Run(run func(ctx context.Context, arg db.UpdateUserParams)) *MockQuerier_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(db.UpdateUserParams))
	})
	return _c
}

func (_c *MockQuerier_UpdateUser_Call) Return(_a0 db.User, _a1 error) *MockQuerier_UpdateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_UpdateUser_Call) RunAndReturn(run func(context.Context, db.UpdateUserParams) (db.User, error)) *MockQuerier_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockQuerier creates a new instance of MockQuerier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQuerier(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQuerier {
	mock := &MockQuerier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
