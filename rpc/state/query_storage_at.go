// Go Substrate RPC Client (GSRPC) provides APIs and types around Polkadot and any Substrate-based chain RPC calls
//
// Copyright 2022 Snowfork
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

package state

import (
	"context"

	"github.com/centrifuge/go-substrate-rpc-client/v4/client"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
)

// QueryStorageAt performs a low-level storage query
func (s *state) QueryStorageAt(ctx context.Context, keys []types.StorageKey, block types.Hash) ([]types.StorageChangeSet, error) {
	return s.queryStorageAt(ctx, keys, &block)
}

// QueryStorageAtLatest performs a low-level storage query
func (s *state) QueryStorageAtLatest(ctx context.Context, keys []types.StorageKey) ([]types.StorageChangeSet, error) {
	return s.queryStorageAt(ctx, keys, nil)
}

func (s *state) queryStorageAt(ctx context.Context, keys []types.StorageKey, block *types.Hash) ([]types.StorageChangeSet, error) {
	hexKeys := make([]string, len(keys))
	for i, key := range keys {
		hexKeys[i] = key.Hex()
	}

	var res []types.StorageChangeSet
	err := client.CallWithBlockHashContext(ctx, s.client, &res, "state_queryStorageAt", block, hexKeys)
	if err != nil {
		return nil, err
	}

	return res, nil
}
