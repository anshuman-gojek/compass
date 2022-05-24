// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	tag "github.com/odpf/compass/core/tag"
)

// TagService is an autogenerated mock type for the TagService type
type TagService struct {
	mock.Mock
}

type TagService_Expecter struct {
	mock *mock.Mock
}

func (_m *TagService) EXPECT() *TagService_Expecter {
	return &TagService_Expecter{mock: &_m.Mock}
}

// CreateTag provides a mock function with given fields: ctx, _a1
func (_m *TagService) CreateTag(ctx context.Context, _a1 *tag.Tag) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *tag.Tag) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TagService_CreateTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTag'
type TagService_CreateTag_Call struct {
	*mock.Call
}

// CreateTag is a helper method to define mock.On call
//  - ctx context.Context
//  - _a1 *tag.Tag
func (_e *TagService_Expecter) CreateTag(ctx interface{}, _a1 interface{}) *TagService_CreateTag_Call {
	return &TagService_CreateTag_Call{Call: _e.mock.On("CreateTag", ctx, _a1)}
}

func (_c *TagService_CreateTag_Call) Run(run func(ctx context.Context, _a1 *tag.Tag)) *TagService_CreateTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*tag.Tag))
	})
	return _c
}

func (_c *TagService_CreateTag_Call) Return(_a0 error) *TagService_CreateTag_Call {
	_c.Call.Return(_a0)
	return _c
}

// DeleteTag provides a mock function with given fields: ctx, assetID, templateURN
func (_m *TagService) DeleteTag(ctx context.Context, assetID string, templateURN string) error {
	ret := _m.Called(ctx, assetID, templateURN)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, assetID, templateURN)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TagService_DeleteTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteTag'
type TagService_DeleteTag_Call struct {
	*mock.Call
}

// DeleteTag is a helper method to define mock.On call
//  - ctx context.Context
//  - assetID string
//  - templateURN string
func (_e *TagService_Expecter) DeleteTag(ctx interface{}, assetID interface{}, templateURN interface{}) *TagService_DeleteTag_Call {
	return &TagService_DeleteTag_Call{Call: _e.mock.On("DeleteTag", ctx, assetID, templateURN)}
}

func (_c *TagService_DeleteTag_Call) Run(run func(ctx context.Context, assetID string, templateURN string)) *TagService_DeleteTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *TagService_DeleteTag_Call) Return(_a0 error) *TagService_DeleteTag_Call {
	_c.Call.Return(_a0)
	return _c
}

// FindTagByAssetIDAndTemplateURN provides a mock function with given fields: ctx, assetID, templateURN
func (_m *TagService) FindTagByAssetIDAndTemplateURN(ctx context.Context, assetID string, templateURN string) (tag.Tag, error) {
	ret := _m.Called(ctx, assetID, templateURN)

	var r0 tag.Tag
	if rf, ok := ret.Get(0).(func(context.Context, string, string) tag.Tag); ok {
		r0 = rf(ctx, assetID, templateURN)
	} else {
		r0 = ret.Get(0).(tag.Tag)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, assetID, templateURN)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagService_FindTagByAssetIDAndTemplateURN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindTagByAssetIDAndTemplateURN'
type TagService_FindTagByAssetIDAndTemplateURN_Call struct {
	*mock.Call
}

// FindTagByAssetIDAndTemplateURN is a helper method to define mock.On call
//  - ctx context.Context
//  - assetID string
//  - templateURN string
func (_e *TagService_Expecter) FindTagByAssetIDAndTemplateURN(ctx interface{}, assetID interface{}, templateURN interface{}) *TagService_FindTagByAssetIDAndTemplateURN_Call {
	return &TagService_FindTagByAssetIDAndTemplateURN_Call{Call: _e.mock.On("FindTagByAssetIDAndTemplateURN", ctx, assetID, templateURN)}
}

func (_c *TagService_FindTagByAssetIDAndTemplateURN_Call) Run(run func(ctx context.Context, assetID string, templateURN string)) *TagService_FindTagByAssetIDAndTemplateURN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *TagService_FindTagByAssetIDAndTemplateURN_Call) Return(_a0 tag.Tag, _a1 error) *TagService_FindTagByAssetIDAndTemplateURN_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTagsByAssetID provides a mock function with given fields: ctx, assetID
func (_m *TagService) GetTagsByAssetID(ctx context.Context, assetID string) ([]tag.Tag, error) {
	ret := _m.Called(ctx, assetID)

	var r0 []tag.Tag
	if rf, ok := ret.Get(0).(func(context.Context, string) []tag.Tag); ok {
		r0 = rf(ctx, assetID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]tag.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, assetID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TagService_GetTagsByAssetID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTagsByAssetID'
type TagService_GetTagsByAssetID_Call struct {
	*mock.Call
}

// GetTagsByAssetID is a helper method to define mock.On call
//  - ctx context.Context
//  - assetID string
func (_e *TagService_Expecter) GetTagsByAssetID(ctx interface{}, assetID interface{}) *TagService_GetTagsByAssetID_Call {
	return &TagService_GetTagsByAssetID_Call{Call: _e.mock.On("GetTagsByAssetID", ctx, assetID)}
}

func (_c *TagService_GetTagsByAssetID_Call) Run(run func(ctx context.Context, assetID string)) *TagService_GetTagsByAssetID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *TagService_GetTagsByAssetID_Call) Return(_a0 []tag.Tag, _a1 error) *TagService_GetTagsByAssetID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateTag provides a mock function with given fields: ctx, _a1
func (_m *TagService) UpdateTag(ctx context.Context, _a1 *tag.Tag) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *tag.Tag) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TagService_UpdateTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTag'
type TagService_UpdateTag_Call struct {
	*mock.Call
}

// UpdateTag is a helper method to define mock.On call
//  - ctx context.Context
//  - _a1 *tag.Tag
func (_e *TagService_Expecter) UpdateTag(ctx interface{}, _a1 interface{}) *TagService_UpdateTag_Call {
	return &TagService_UpdateTag_Call{Call: _e.mock.On("UpdateTag", ctx, _a1)}
}

func (_c *TagService_UpdateTag_Call) Run(run func(ctx context.Context, _a1 *tag.Tag)) *TagService_UpdateTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*tag.Tag))
	})
	return _c
}

func (_c *TagService_UpdateTag_Call) Return(_a0 error) *TagService_UpdateTag_Call {
	_c.Call.Return(_a0)
	return _c
}

// Validate provides a mock function with given fields: _a0
func (_m *TagService) Validate(_a0 *tag.Tag) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*tag.Tag) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TagService_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type TagService_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//  - _a0 *tag.Tag
func (_e *TagService_Expecter) Validate(_a0 interface{}) *TagService_Validate_Call {
	return &TagService_Validate_Call{Call: _e.mock.On("Validate", _a0)}
}

func (_c *TagService_Validate_Call) Run(run func(_a0 *tag.Tag)) *TagService_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*tag.Tag))
	})
	return _c
}

func (_c *TagService_Validate_Call) Return(_a0 error) *TagService_Validate_Call {
	_c.Call.Return(_a0)
	return _c
}