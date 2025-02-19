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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChain_GetBlockLatest(t *testing.T) {
	rv, err := testChain.GetBlockLatest(context.Background())
	assert.NoError(t, err)
	assert.True(t, rv.Block.Header.Number > 0)
}

func TestChain_GetBlock(t *testing.T) {
	rv, err := testChain.GetBlockLatest(context.Background())
	assert.NoError(t, err)

	// fetch previous block and verify the number
	latest := rv.Block.Header.Number
	rv, err = testChain.GetBlock(context.Background(), rv.Block.Header.ParentHash)
	assert.NoError(t, err)
	assert.Equal(t, latest-1, rv.Block.Header.Number)
}
