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

package chain

import (
	"context"

	"github.com/centrifuge/go-substrate-rpc-client/v4/client"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/block"
)

// GetBlock returns the header and body of the relay chain block with the given hash
func (c *chain) GetBlock(ctx context.Context, blockHash types.Hash) (*block.SignedBlock, error) {
	return c.getBlock(ctx, &blockHash)
}

// GetBlockLatest returns the header and body of the latest relay chain block
func (c *chain) GetBlockLatest(ctx context.Context) (*block.SignedBlock, error) {
	return c.getBlock(ctx, nil)
}

func (c *chain) getBlock(ctx context.Context, blockHash *types.Hash) (*block.SignedBlock, error) {
	var res block.SignedBlock
	err := client.CallWithBlockHashContext(ctx, c.client, &res, "chain_getBlock", blockHash)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
