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

//go:generate mockery --name Chain --filename chain.go

package chain

import (
	"context"

	"github.com/centrifuge/go-substrate-rpc-client/v4/client"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types/block"
)

//go:generate mockery --name Chain --filename chain.go

type Chain interface {
	SubscribeFinalizedHeads(ctx context.Context) (*FinalizedHeadsSubscription, error)
	SubscribeNewHeads(ctx context.Context) (*NewHeadsSubscription, error)
	GetBlockHash(ctx context.Context, blockNumber uint64) (types.Hash, error)
	GetBlockHashLatest(ctx context.Context) (types.Hash, error)
	GetFinalizedHead(ctx context.Context) (types.Hash, error)
	GetBlock(ctx context.Context, blockHash types.Hash) (*block.SignedBlock, error)
	GetBlockLatest(ctx context.Context) (*block.SignedBlock, error)
	GetHeader(ctx context.Context, blockHash types.Hash) (*types.Header, error)
	GetHeaderLatest(ctx context.Context) (*types.Header, error)
}

// chain exposes methods for retrieval of chain data
type chain struct {
	client client.Client
}

// NewChain creates a new chain struct
func NewChain(cl client.Client) Chain {
	return &chain{cl}
}
