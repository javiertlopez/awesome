// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/javiertlopez/awesome/model"
	mock "github.com/stretchr/testify/mock"
)

// AssetRepo is an autogenerated mock type for the AssetRepo type
type AssetRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, source, public
func (_m *AssetRepo) Create(ctx context.Context, source string, public bool) (string, error) {
	ret := _m.Called(ctx, source, public)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) string); ok {
		r0 = rf(ctx, source, public)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, source, public)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *AssetRepo) GetByID(ctx context.Context, id string) (model.Asset, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Asset
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Asset); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Asset)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
