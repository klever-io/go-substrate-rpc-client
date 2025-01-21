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

package state

import (
	"context"

	"github.com/centrifuge/go-substrate-rpc-client/v4/client"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/codec"
)

// GetChildStorage retreives the child storage for a key and decodes them into the provided interface. Ok is true if the
// value is not empty.
func (s *state) GetChildStorage(ctx context.Context, childStorageKey, key types.StorageKey, target interface{}, blockHash types.Hash) (
	ok bool, err error) {
	raw, err := s.getChildStorageRaw(ctx, childStorageKey, key, &blockHash)
	if err != nil {
		return false, err
	}
	if len(*raw) == 0 {
		return false, nil
	}
	return true, codec.Decode(*raw, target)
}

// GetChildStorageLatest retreives the child storage for a key for the latest block height and decodes them into the
// provided interface. Ok is true if the value is not empty.
func (s *state) GetChildStorageLatest(ctx context.Context, childStorageKey, key types.StorageKey, target interface{}) (ok bool, err error) {
	raw, err := s.getChildStorageRaw(ctx, childStorageKey, key, nil)
	if err != nil {
		return false, err
	}
	if len(*raw) == 0 {
		return false, nil
	}
	return true, codec.Decode(*raw, target)
}

// GetChildStorageRaw retreives the child storage for a key as raw bytes, without decoding them
func (s *state) GetChildStorageRaw(ctx context.Context, childStorageKey, key types.StorageKey, blockHash types.Hash) (
	*types.StorageDataRaw, error) {
	return s.getChildStorageRaw(ctx, childStorageKey, key, &blockHash)
}

// GetChildStorageRawLatest retreives the child storage for a key for the latest block height as raw bytes,
// without decoding them
func (s *state) GetChildStorageRawLatest(ctx context.Context, childStorageKey, key types.StorageKey) (*types.StorageDataRaw, error) {
	return s.getChildStorageRaw(ctx, childStorageKey, key, nil)
}

func (s *state) getChildStorageRaw(ctx context.Context, childStorageKey, key types.StorageKey, blockHash *types.Hash) (
	*types.StorageDataRaw, error) {
	var res string
	err := client.CallWithBlockHashContext(ctx, s.client, &res, "state_getChildStorage", blockHash, childStorageKey.Hex(),
		key.Hex())
	if err != nil {
		return nil, err
	}

	bz, err := codec.HexDecodeString(res)
	if err != nil {
		return nil, err
	}

	data := types.NewStorageDataRaw(bz)
	return &data, nil
}
