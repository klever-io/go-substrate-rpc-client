// Code generated by mockery v2.51.0. DO NOT EDIT.

package mocks

import (
	context "context"

	state "github.com/centrifuge/go-substrate-rpc-client/v4/rpc/state"
	mock "github.com/stretchr/testify/mock"

	types "github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// State is an autogenerated mock type for the State type
type State struct {
	mock.Mock
}

// GetChildKeys provides a mock function with given fields: ctx, childStorageKey, prefix, blockHash
func (_m *State) GetChildKeys(ctx context.Context, childStorageKey types.StorageKey, prefix types.StorageKey, blockHash types.Hash) ([]types.StorageKey, error) {
	ret := _m.Called(ctx, childStorageKey, prefix, blockHash)

	if len(ret) == 0 {
		panic("no return value specified for GetChildKeys")
	}

	var r0 []types.StorageKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey, types.Hash) ([]types.StorageKey, error)); ok {
		return rf(ctx, childStorageKey, prefix, blockHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey, types.Hash) []types.StorageKey); ok {
		r0 = rf(ctx, childStorageKey, prefix, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.StorageKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey, types.StorageKey, types.Hash) error); ok {
		r1 = rf(ctx, childStorageKey, prefix, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildKeysLatest provides a mock function with given fields: ctx, childStorageKey, prefix
func (_m *State) GetChildKeysLatest(ctx context.Context, childStorageKey types.StorageKey, prefix types.StorageKey) ([]types.StorageKey, error) {
	ret := _m.Called(ctx, childStorageKey, prefix)

	if len(ret) == 0 {
		panic("no return value specified for GetChildKeysLatest")
	}

	var r0 []types.StorageKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey) ([]types.StorageKey, error)); ok {
		return rf(ctx, childStorageKey, prefix)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey) []types.StorageKey); ok {
		r0 = rf(ctx, childStorageKey, prefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.StorageKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey, types.StorageKey) error); ok {
		r1 = rf(ctx, childStorageKey, prefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildStorage provides a mock function with given fields: ctx, childStorageKey, key, target, blockHash
func (_m *State) GetChildStorage(ctx context.Context, childStorageKey types.StorageKey, key types.StorageKey, target interface{}, blockHash types.Hash) (bool, error) {
	ret := _m.Called(ctx, childStorageKey, key, target, blockHash)

	if len(ret) == 0 {
		panic("no return value specified for GetChildStorage")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey, interface{}, types.Hash) (bool, error)); ok {
		return rf(ctx, childStorageKey, key, target, blockHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey, interface{}, types.Hash) bool); ok {
		r0 = rf(ctx, childStorageKey, key, target, blockHash)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey, types.StorageKey, interface{}, types.Hash) error); ok {
		r1 = rf(ctx, childStorageKey, key, target, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildStorageHash provides a mock function with given fields: ctx, childStorageKey, key, blockHash
func (_m *State) GetChildStorageHash(ctx context.Context, childStorageKey types.StorageKey, key types.StorageKey, blockHash types.Hash) (types.Hash, error) {
	ret := _m.Called(ctx, childStorageKey, key, blockHash)

	if len(ret) == 0 {
		panic("no return value specified for GetChildStorageHash")
	}

	var r0 types.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey, types.Hash) (types.Hash, error)); ok {
		return rf(ctx, childStorageKey, key, blockHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey, types.Hash) types.Hash); ok {
		r0 = rf(ctx, childStorageKey, key, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey, types.StorageKey, types.Hash) error); ok {
		r1 = rf(ctx, childStorageKey, key, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildStorageHashLatest provides a mock function with given fields: ctx, childStorageKey, key
func (_m *State) GetChildStorageHashLatest(ctx context.Context, childStorageKey types.StorageKey, key types.StorageKey) (types.Hash, error) {
	ret := _m.Called(ctx, childStorageKey, key)

	if len(ret) == 0 {
		panic("no return value specified for GetChildStorageHashLatest")
	}

	var r0 types.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey) (types.Hash, error)); ok {
		return rf(ctx, childStorageKey, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey) types.Hash); ok {
		r0 = rf(ctx, childStorageKey, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey, types.StorageKey) error); ok {
		r1 = rf(ctx, childStorageKey, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildStorageLatest provides a mock function with given fields: ctx, childStorageKey, key, target
func (_m *State) GetChildStorageLatest(ctx context.Context, childStorageKey types.StorageKey, key types.StorageKey, target interface{}) (bool, error) {
	ret := _m.Called(ctx, childStorageKey, key, target)

	if len(ret) == 0 {
		panic("no return value specified for GetChildStorageLatest")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey, interface{}) (bool, error)); ok {
		return rf(ctx, childStorageKey, key, target)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey, interface{}) bool); ok {
		r0 = rf(ctx, childStorageKey, key, target)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey, types.StorageKey, interface{}) error); ok {
		r1 = rf(ctx, childStorageKey, key, target)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildStorageRaw provides a mock function with given fields: ctx, childStorageKey, key, blockHash
func (_m *State) GetChildStorageRaw(ctx context.Context, childStorageKey types.StorageKey, key types.StorageKey, blockHash types.Hash) (*types.StorageDataRaw, error) {
	ret := _m.Called(ctx, childStorageKey, key, blockHash)

	if len(ret) == 0 {
		panic("no return value specified for GetChildStorageRaw")
	}

	var r0 *types.StorageDataRaw
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey, types.Hash) (*types.StorageDataRaw, error)); ok {
		return rf(ctx, childStorageKey, key, blockHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey, types.Hash) *types.StorageDataRaw); ok {
		r0 = rf(ctx, childStorageKey, key, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.StorageDataRaw)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey, types.StorageKey, types.Hash) error); ok {
		r1 = rf(ctx, childStorageKey, key, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildStorageRawLatest provides a mock function with given fields: ctx, childStorageKey, key
func (_m *State) GetChildStorageRawLatest(ctx context.Context, childStorageKey types.StorageKey, key types.StorageKey) (*types.StorageDataRaw, error) {
	ret := _m.Called(ctx, childStorageKey, key)

	if len(ret) == 0 {
		panic("no return value specified for GetChildStorageRawLatest")
	}

	var r0 *types.StorageDataRaw
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey) (*types.StorageDataRaw, error)); ok {
		return rf(ctx, childStorageKey, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey) *types.StorageDataRaw); ok {
		r0 = rf(ctx, childStorageKey, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.StorageDataRaw)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey, types.StorageKey) error); ok {
		r1 = rf(ctx, childStorageKey, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildStorageSize provides a mock function with given fields: ctx, childStorageKey, key, blockHash
func (_m *State) GetChildStorageSize(ctx context.Context, childStorageKey types.StorageKey, key types.StorageKey, blockHash types.Hash) (types.U64, error) {
	ret := _m.Called(ctx, childStorageKey, key, blockHash)

	if len(ret) == 0 {
		panic("no return value specified for GetChildStorageSize")
	}

	var r0 types.U64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey, types.Hash) (types.U64, error)); ok {
		return rf(ctx, childStorageKey, key, blockHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey, types.Hash) types.U64); ok {
		r0 = rf(ctx, childStorageKey, key, blockHash)
	} else {
		r0 = ret.Get(0).(types.U64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey, types.StorageKey, types.Hash) error); ok {
		r1 = rf(ctx, childStorageKey, key, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChildStorageSizeLatest provides a mock function with given fields: ctx, childStorageKey, key
func (_m *State) GetChildStorageSizeLatest(ctx context.Context, childStorageKey types.StorageKey, key types.StorageKey) (types.U64, error) {
	ret := _m.Called(ctx, childStorageKey, key)

	if len(ret) == 0 {
		panic("no return value specified for GetChildStorageSizeLatest")
	}

	var r0 types.U64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey) (types.U64, error)); ok {
		return rf(ctx, childStorageKey, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.StorageKey) types.U64); ok {
		r0 = rf(ctx, childStorageKey, key)
	} else {
		r0 = ret.Get(0).(types.U64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey, types.StorageKey) error); ok {
		r1 = rf(ctx, childStorageKey, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKeys provides a mock function with given fields: ctx, prefix, blockHash
func (_m *State) GetKeys(ctx context.Context, prefix types.StorageKey, blockHash types.Hash) ([]types.StorageKey, error) {
	ret := _m.Called(ctx, prefix, blockHash)

	if len(ret) == 0 {
		panic("no return value specified for GetKeys")
	}

	var r0 []types.StorageKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.Hash) ([]types.StorageKey, error)); ok {
		return rf(ctx, prefix, blockHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.Hash) []types.StorageKey); ok {
		r0 = rf(ctx, prefix, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.StorageKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey, types.Hash) error); ok {
		r1 = rf(ctx, prefix, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetKeysLatest provides a mock function with given fields: ctx, prefix
func (_m *State) GetKeysLatest(ctx context.Context, prefix types.StorageKey) ([]types.StorageKey, error) {
	ret := _m.Called(ctx, prefix)

	if len(ret) == 0 {
		panic("no return value specified for GetKeysLatest")
	}

	var r0 []types.StorageKey
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey) ([]types.StorageKey, error)); ok {
		return rf(ctx, prefix)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey) []types.StorageKey); ok {
		r0 = rf(ctx, prefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.StorageKey)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey) error); ok {
		r1 = rf(ctx, prefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMetadata provides a mock function with given fields: ctx, blockHash
func (_m *State) GetMetadata(ctx context.Context, blockHash types.Hash) (*types.Metadata, error) {
	ret := _m.Called(ctx, blockHash)

	if len(ret) == 0 {
		panic("no return value specified for GetMetadata")
	}

	var r0 *types.Metadata
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.Hash) (*types.Metadata, error)); ok {
		return rf(ctx, blockHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.Hash) *types.Metadata); ok {
		r0 = rf(ctx, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Metadata)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.Hash) error); ok {
		r1 = rf(ctx, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMetadataLatest provides a mock function with given fields: ctx
func (_m *State) GetMetadataLatest(ctx context.Context) (*types.Metadata, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetMetadataLatest")
	}

	var r0 *types.Metadata
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*types.Metadata, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *types.Metadata); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Metadata)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRuntimeVersion provides a mock function with given fields: ctx, blockHash
func (_m *State) GetRuntimeVersion(ctx context.Context, blockHash types.Hash) (*types.RuntimeVersion, error) {
	ret := _m.Called(ctx, blockHash)

	if len(ret) == 0 {
		panic("no return value specified for GetRuntimeVersion")
	}

	var r0 *types.RuntimeVersion
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.Hash) (*types.RuntimeVersion, error)); ok {
		return rf(ctx, blockHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.Hash) *types.RuntimeVersion); ok {
		r0 = rf(ctx, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.RuntimeVersion)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.Hash) error); ok {
		r1 = rf(ctx, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRuntimeVersionLatest provides a mock function with given fields: ctx
func (_m *State) GetRuntimeVersionLatest(ctx context.Context) (*types.RuntimeVersion, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetRuntimeVersionLatest")
	}

	var r0 *types.RuntimeVersion
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*types.RuntimeVersion, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *types.RuntimeVersion); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.RuntimeVersion)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorage provides a mock function with given fields: ctx, key, target, blockHash
func (_m *State) GetStorage(ctx context.Context, key types.StorageKey, target interface{}, blockHash types.Hash) (bool, error) {
	ret := _m.Called(ctx, key, target, blockHash)

	if len(ret) == 0 {
		panic("no return value specified for GetStorage")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, interface{}, types.Hash) (bool, error)); ok {
		return rf(ctx, key, target, blockHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, interface{}, types.Hash) bool); ok {
		r0 = rf(ctx, key, target, blockHash)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey, interface{}, types.Hash) error); ok {
		r1 = rf(ctx, key, target, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageHash provides a mock function with given fields: ctx, key, blockHash
func (_m *State) GetStorageHash(ctx context.Context, key types.StorageKey, blockHash types.Hash) (types.Hash, error) {
	ret := _m.Called(ctx, key, blockHash)

	if len(ret) == 0 {
		panic("no return value specified for GetStorageHash")
	}

	var r0 types.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.Hash) (types.Hash, error)); ok {
		return rf(ctx, key, blockHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.Hash) types.Hash); ok {
		r0 = rf(ctx, key, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey, types.Hash) error); ok {
		r1 = rf(ctx, key, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageHashLatest provides a mock function with given fields: ctx, key
func (_m *State) GetStorageHashLatest(ctx context.Context, key types.StorageKey) (types.Hash, error) {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for GetStorageHashLatest")
	}

	var r0 types.Hash
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey) (types.Hash, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey) types.Hash); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.Hash)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageLatest provides a mock function with given fields: ctx, key, target
func (_m *State) GetStorageLatest(ctx context.Context, key types.StorageKey, target interface{}) (bool, error) {
	ret := _m.Called(ctx, key, target)

	if len(ret) == 0 {
		panic("no return value specified for GetStorageLatest")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, interface{}) (bool, error)); ok {
		return rf(ctx, key, target)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, interface{}) bool); ok {
		r0 = rf(ctx, key, target)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey, interface{}) error); ok {
		r1 = rf(ctx, key, target)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageRaw provides a mock function with given fields: ctx, key, blockHash
func (_m *State) GetStorageRaw(ctx context.Context, key types.StorageKey, blockHash types.Hash) (*types.StorageDataRaw, error) {
	ret := _m.Called(ctx, key, blockHash)

	if len(ret) == 0 {
		panic("no return value specified for GetStorageRaw")
	}

	var r0 *types.StorageDataRaw
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.Hash) (*types.StorageDataRaw, error)); ok {
		return rf(ctx, key, blockHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.Hash) *types.StorageDataRaw); ok {
		r0 = rf(ctx, key, blockHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.StorageDataRaw)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey, types.Hash) error); ok {
		r1 = rf(ctx, key, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageRawLatest provides a mock function with given fields: ctx, key
func (_m *State) GetStorageRawLatest(ctx context.Context, key types.StorageKey) (*types.StorageDataRaw, error) {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for GetStorageRawLatest")
	}

	var r0 *types.StorageDataRaw
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey) (*types.StorageDataRaw, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey) *types.StorageDataRaw); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.StorageDataRaw)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageSize provides a mock function with given fields: ctx, key, blockHash
func (_m *State) GetStorageSize(ctx context.Context, key types.StorageKey, blockHash types.Hash) (types.U64, error) {
	ret := _m.Called(ctx, key, blockHash)

	if len(ret) == 0 {
		panic("no return value specified for GetStorageSize")
	}

	var r0 types.U64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.Hash) (types.U64, error)); ok {
		return rf(ctx, key, blockHash)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey, types.Hash) types.U64); ok {
		r0 = rf(ctx, key, blockHash)
	} else {
		r0 = ret.Get(0).(types.U64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey, types.Hash) error); ok {
		r1 = rf(ctx, key, blockHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStorageSizeLatest provides a mock function with given fields: ctx, key
func (_m *State) GetStorageSizeLatest(ctx context.Context, key types.StorageKey) (types.U64, error) {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for GetStorageSizeLatest")
	}

	var r0 types.U64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey) (types.U64, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.StorageKey) types.U64); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(types.U64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.StorageKey) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryStorage provides a mock function with given fields: ctx, keys, startBlock, block
func (_m *State) QueryStorage(ctx context.Context, keys []types.StorageKey, startBlock types.Hash, block types.Hash) ([]types.StorageChangeSet, error) {
	ret := _m.Called(ctx, keys, startBlock, block)

	if len(ret) == 0 {
		panic("no return value specified for QueryStorage")
	}

	var r0 []types.StorageChangeSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []types.StorageKey, types.Hash, types.Hash) ([]types.StorageChangeSet, error)); ok {
		return rf(ctx, keys, startBlock, block)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []types.StorageKey, types.Hash, types.Hash) []types.StorageChangeSet); ok {
		r0 = rf(ctx, keys, startBlock, block)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.StorageChangeSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []types.StorageKey, types.Hash, types.Hash) error); ok {
		r1 = rf(ctx, keys, startBlock, block)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryStorageAt provides a mock function with given fields: ctx, keys, block
func (_m *State) QueryStorageAt(ctx context.Context, keys []types.StorageKey, block types.Hash) ([]types.StorageChangeSet, error) {
	ret := _m.Called(ctx, keys, block)

	if len(ret) == 0 {
		panic("no return value specified for QueryStorageAt")
	}

	var r0 []types.StorageChangeSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []types.StorageKey, types.Hash) ([]types.StorageChangeSet, error)); ok {
		return rf(ctx, keys, block)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []types.StorageKey, types.Hash) []types.StorageChangeSet); ok {
		r0 = rf(ctx, keys, block)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.StorageChangeSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []types.StorageKey, types.Hash) error); ok {
		r1 = rf(ctx, keys, block)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryStorageAtLatest provides a mock function with given fields: ctx, keys
func (_m *State) QueryStorageAtLatest(ctx context.Context, keys []types.StorageKey) ([]types.StorageChangeSet, error) {
	ret := _m.Called(ctx, keys)

	if len(ret) == 0 {
		panic("no return value specified for QueryStorageAtLatest")
	}

	var r0 []types.StorageChangeSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []types.StorageKey) ([]types.StorageChangeSet, error)); ok {
		return rf(ctx, keys)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []types.StorageKey) []types.StorageChangeSet); ok {
		r0 = rf(ctx, keys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.StorageChangeSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []types.StorageKey) error); ok {
		r1 = rf(ctx, keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryStorageLatest provides a mock function with given fields: ctx, keys, startBlock
func (_m *State) QueryStorageLatest(ctx context.Context, keys []types.StorageKey, startBlock types.Hash) ([]types.StorageChangeSet, error) {
	ret := _m.Called(ctx, keys, startBlock)

	if len(ret) == 0 {
		panic("no return value specified for QueryStorageLatest")
	}

	var r0 []types.StorageChangeSet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []types.StorageKey, types.Hash) ([]types.StorageChangeSet, error)); ok {
		return rf(ctx, keys, startBlock)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []types.StorageKey, types.Hash) []types.StorageChangeSet); ok {
		r0 = rf(ctx, keys, startBlock)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.StorageChangeSet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []types.StorageKey, types.Hash) error); ok {
		r1 = rf(ctx, keys, startBlock)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeRuntimeVersion provides a mock function with given fields: ctx
func (_m *State) SubscribeRuntimeVersion(ctx context.Context) (*state.RuntimeVersionSubscription, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for SubscribeRuntimeVersion")
	}

	var r0 *state.RuntimeVersionSubscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*state.RuntimeVersionSubscription, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *state.RuntimeVersionSubscription); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.RuntimeVersionSubscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubscribeStorageRaw provides a mock function with given fields: ctx, keys
func (_m *State) SubscribeStorageRaw(ctx context.Context, keys []types.StorageKey) (*state.StorageSubscription, error) {
	ret := _m.Called(ctx, keys)

	if len(ret) == 0 {
		panic("no return value specified for SubscribeStorageRaw")
	}

	var r0 *state.StorageSubscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []types.StorageKey) (*state.StorageSubscription, error)); ok {
		return rf(ctx, keys)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []types.StorageKey) *state.StorageSubscription); ok {
		r0 = rf(ctx, keys)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.StorageSubscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []types.StorageKey) error); ok {
		r1 = rf(ctx, keys)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewState creates a new instance of State. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewState(t interface {
	mock.TestingT
	Cleanup(func())
}) *State {
	mock := &State{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
