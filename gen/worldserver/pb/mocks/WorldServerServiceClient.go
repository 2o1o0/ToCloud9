// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/walkline/ToCloud9/gen/worldserver/pb"
)

// WorldServerServiceClient is an autogenerated mock type for the WorldServerServiceClient type
type WorldServerServiceClient struct {
	mock.Mock
}

// AddExistingItemToPlayer provides a mock function with given fields: ctx, in, opts
func (_m *WorldServerServiceClient) AddExistingItemToPlayer(ctx context.Context, in *pb.AddExistingItemToPlayerRequest, opts ...grpc.CallOption) (*pb.AddExistingItemToPlayerResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.AddExistingItemToPlayerResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.AddExistingItemToPlayerRequest, ...grpc.CallOption) (*pb.AddExistingItemToPlayerResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.AddExistingItemToPlayerRequest, ...grpc.CallOption) *pb.AddExistingItemToPlayerResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.AddExistingItemToPlayerResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.AddExistingItemToPlayerRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CanPlayerInteractWithGameObject provides a mock function with given fields: ctx, in, opts
func (_m *WorldServerServiceClient) CanPlayerInteractWithGameObject(ctx context.Context, in *pb.CanPlayerInteractWithGameObjectRequest, opts ...grpc.CallOption) (*pb.CanPlayerInteractWithGameObjectResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.CanPlayerInteractWithGameObjectResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CanPlayerInteractWithGameObjectRequest, ...grpc.CallOption) (*pb.CanPlayerInteractWithGameObjectResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CanPlayerInteractWithGameObjectRequest, ...grpc.CallOption) *pb.CanPlayerInteractWithGameObjectResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.CanPlayerInteractWithGameObjectResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.CanPlayerInteractWithGameObjectRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CanPlayerInteractWithNPC provides a mock function with given fields: ctx, in, opts
func (_m *WorldServerServiceClient) CanPlayerInteractWithNPC(ctx context.Context, in *pb.CanPlayerInteractWithNPCRequest, opts ...grpc.CallOption) (*pb.CanPlayerInteractWithNPCResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.CanPlayerInteractWithNPCResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CanPlayerInteractWithNPCRequest, ...grpc.CallOption) (*pb.CanPlayerInteractWithNPCResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CanPlayerInteractWithNPCRequest, ...grpc.CallOption) *pb.CanPlayerInteractWithNPCResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.CanPlayerInteractWithNPCResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.CanPlayerInteractWithNPCRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMoneyForPlayer provides a mock function with given fields: ctx, in, opts
func (_m *WorldServerServiceClient) GetMoneyForPlayer(ctx context.Context, in *pb.GetMoneyForPlayerRequest, opts ...grpc.CallOption) (*pb.GetMoneyForPlayerResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.GetMoneyForPlayerResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetMoneyForPlayerRequest, ...grpc.CallOption) (*pb.GetMoneyForPlayerResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetMoneyForPlayerRequest, ...grpc.CallOption) *pb.GetMoneyForPlayerResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GetMoneyForPlayerResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetMoneyForPlayerRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPlayerItemsByGuids provides a mock function with given fields: ctx, in, opts
func (_m *WorldServerServiceClient) GetPlayerItemsByGuids(ctx context.Context, in *pb.GetPlayerItemsByGuidsRequest, opts ...grpc.CallOption) (*pb.GetPlayerItemsByGuidsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.GetPlayerItemsByGuidsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetPlayerItemsByGuidsRequest, ...grpc.CallOption) (*pb.GetPlayerItemsByGuidsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.GetPlayerItemsByGuidsRequest, ...grpc.CallOption) *pb.GetPlayerItemsByGuidsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.GetPlayerItemsByGuidsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.GetPlayerItemsByGuidsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ModifyMoneyForPlayer provides a mock function with given fields: ctx, in, opts
func (_m *WorldServerServiceClient) ModifyMoneyForPlayer(ctx context.Context, in *pb.ModifyMoneyForPlayerRequest, opts ...grpc.CallOption) (*pb.ModifyMoneyForPlayerResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.ModifyMoneyForPlayerResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ModifyMoneyForPlayerRequest, ...grpc.CallOption) (*pb.ModifyMoneyForPlayerResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ModifyMoneyForPlayerRequest, ...grpc.CallOption) *pb.ModifyMoneyForPlayerResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.ModifyMoneyForPlayerResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.ModifyMoneyForPlayerRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveItemsWithGuidsFromPlayer provides a mock function with given fields: ctx, in, opts
func (_m *WorldServerServiceClient) RemoveItemsWithGuidsFromPlayer(ctx context.Context, in *pb.RemoveItemsWithGuidsFromPlayerRequest, opts ...grpc.CallOption) (*pb.RemoveItemsWithGuidsFromPlayerResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *pb.RemoveItemsWithGuidsFromPlayerResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.RemoveItemsWithGuidsFromPlayerRequest, ...grpc.CallOption) (*pb.RemoveItemsWithGuidsFromPlayerResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.RemoveItemsWithGuidsFromPlayerRequest, ...grpc.CallOption) *pb.RemoveItemsWithGuidsFromPlayerResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.RemoveItemsWithGuidsFromPlayerResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.RemoveItemsWithGuidsFromPlayerRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewWorldServerServiceClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewWorldServerServiceClient creates a new instance of WorldServerServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWorldServerServiceClient(t mockConstructorTestingTNewWorldServerServiceClient) *WorldServerServiceClient {
	mock := &WorldServerServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
