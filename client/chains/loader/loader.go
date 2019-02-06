package loader

import (
	clientchaintypes "github.com/opennetsys/go-substrate/client/chains/types"
	clienttypes "github.com/opennetsys/go-substrate/client/types"
)

// TODO: https://github.com/polkadot-js/client/blob/master/packages/client-chains/src/loader.ts

// Loader ...
type Loader struct {
	chain       *clientchaintypes.ChainJSON
	ID          string
	GenesisRoot []uint8
}

// NewLoader ...
// TODO: config loader?
func NewLoader(config *clienttypes.ConfigClient) *Loader {
	// TODO
	return &Loader{}
}

// Chain ...
func (l *Loader) Chain() *clientchaintypes.ChainJSON {
	return l.chain
}
