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
)

// GetRuntimeVersion returns the runtime version at the given block
func (s *state) GetRuntimeVersion(ctx context.Context, blockHash types.Hash) (*types.RuntimeVersion, error) {
	return s.getRuntimeVersion(ctx, &blockHash)
}

// GetRuntimeVersionLatest returns the latest runtime version
func (s *state) GetRuntimeVersionLatest(ctx context.Context) (*types.RuntimeVersion, error) {
	return s.getRuntimeVersion(ctx, nil)
}

func (s *state) getRuntimeVersion(ctx context.Context, blockHash *types.Hash) (*types.RuntimeVersion, error) {
	var runtimeVersion types.RuntimeVersion
	err := client.CallWithBlockHashContext(ctx, s.client, &runtimeVersion, "state_getRuntimeVersion", blockHash)
	if err != nil {
		return nil, err
	}
	return &runtimeVersion, err
}
