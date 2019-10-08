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

package client

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/rpc"
)

type Client interface {
	Call(result interface{}, method string, args ...interface{}) error

	Subscribe(ctx context.Context, namespace string, channel interface{}, args ...interface{}) (
		*rpc.ClientSubscription, error)

	URL() string
	//MetaData(cache bool) (*MetadataVersioned, error)
}

type client struct {
	*rpc.Client

	url string

	// metadataVersioned is the metadata cache to prevent unnecessary requests
	//metadataVersioned *MetadataVersioned

	//metadataLock sync.RWMutex
}

// Returns the URL the client connects to
func (c client) URL() string {
	return c.url
}

// TODO move to State struct
//func (c *client) MetaData(cache bool) (m *MetadataVersioned, err error) {
//	if cache && c.metadataVersioned != nil {
//		c.metadataLock.RLock()
//		defer c.metadataLock.RUnlock()
//		m = c.metadataVersioned
//	} else {
//		s := NewStateRPC(c)
//		m, err = s.MetaData(nil)
//		if err != nil {
//			return nil, err
//		}
//		// set cache
//		c.metadataLock.Lock()
//		defer c.metadataLock.Unlock()
//		c.metadataVersioned = m
//	}
//	return
//}

// Connect connects to the provided url
func Connect(url string) (Client, error) {
	log.Printf("Connecting to %v...", url)
	c, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}
	cc := client{c, url}
	return &cc, nil
}