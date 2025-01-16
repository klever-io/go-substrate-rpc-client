// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2019 Centrifuge GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:generate mockery --name State --filename state.go

package state

import (
	"context"

	"github.com/centrifuge/go-substrate-rpc-client/v4/client"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

type State interface {
	GetStorage(ctx context.Context, key types.StorageKey, target interface{}, blockHash types.Hash) (ok bool, err error)
	GetStorageLatest(ctx context.Context, key types.StorageKey, target interface{}) (ok bool, err error)
	GetStorageRaw(ctx context.Context, key types.StorageKey, blockHash types.Hash) (*types.StorageDataRaw, error)
	GetStorageRawLatest(ctx context.Context, key types.StorageKey) (*types.StorageDataRaw, error)

	GetChildStorageSize(ctx context.Context, childStorageKey, key types.StorageKey, blockHash types.Hash) (types.U64, error)
	GetChildStorageSizeLatest(ctx context.Context, childStorageKey, key types.StorageKey) (types.U64, error)
	GetChildStorage(ctx context.Context, childStorageKey, key types.StorageKey, target interface{}, blockHash types.Hash) (ok bool, err error)
	GetChildStorageLatest(ctx context.Context, childStorageKey, key types.StorageKey, target interface{}) (ok bool, err error)
	GetChildStorageRaw(ctx context.Context, childStorageKey, key types.StorageKey, blockHash types.Hash) (*types.StorageDataRaw, error)
	GetChildStorageRawLatest(ctx context.Context, childStorageKey, key types.StorageKey) (*types.StorageDataRaw, error)

	GetMetadata(ctx context.Context, blockHash types.Hash) (*types.Metadata, error)
	GetMetadataLatest(ctx context.Context) (*types.Metadata, error)

	GetStorageHash(ctx context.Context, key types.StorageKey, blockHash types.Hash) (types.Hash, error)
	GetStorageHashLatest(ctx context.Context, key types.StorageKey) (types.Hash, error)

	SubscribeStorageRaw(ctx context.Context, keys []types.StorageKey) (*StorageSubscription, error)

	GetRuntimeVersion(ctx context.Context, blockHash types.Hash) (*types.RuntimeVersion, error)
	GetRuntimeVersionLatest(ctx context.Context) (*types.RuntimeVersion, error)

	GetChildKeys(ctx context.Context, childStorageKey, prefix types.StorageKey, blockHash types.Hash) ([]types.StorageKey, error)
	GetChildKeysLatest(ctx context.Context, childStorageKey, prefix types.StorageKey) ([]types.StorageKey, error)

	SubscribeRuntimeVersion(ctx context.Context) (*RuntimeVersionSubscription, error)

	QueryStorage(ctx context.Context, keys []types.StorageKey, startBlock types.Hash, block types.Hash) ([]types.StorageChangeSet, error)
	QueryStorageLatest(ctx context.Context, keys []types.StorageKey, startBlock types.Hash) ([]types.StorageChangeSet, error)

	QueryStorageAt(ctx context.Context, keys []types.StorageKey, block types.Hash) ([]types.StorageChangeSet, error)
	QueryStorageAtLatest(ctx context.Context, keys []types.StorageKey) ([]types.StorageChangeSet, error)

	GetKeys(ctx context.Context, prefix types.StorageKey, blockHash types.Hash) ([]types.StorageKey, error)
	GetKeysLatest(ctx context.Context, prefix types.StorageKey) ([]types.StorageKey, error)

	GetStorageSize(ctx context.Context, key types.StorageKey, blockHash types.Hash) (types.U64, error)
	GetStorageSizeLatest(ctx context.Context, key types.StorageKey) (types.U64, error)

	GetChildStorageHash(ctx context.Context, childStorageKey, key types.StorageKey, blockHash types.Hash) (types.Hash, error)
	GetChildStorageHashLatest(ctx context.Context, childStorageKey, key types.StorageKey) (types.Hash, error)
}

// state exposes methods for querying state
type state struct {
	client client.Client
}

// NewState creates a new state struct
func NewState(c client.Client) State {
	return &state{client: c}
}
