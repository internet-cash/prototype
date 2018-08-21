package netsync

import (
	"github.com/internet-cash/prototype/blockchain"
	"sync"
	"github.com/internet-cash/prototype/peer"
)

type NetSync struct {
	started  int32
	shutdown int32

	Chain      *blockchain.BlockChain
	ChainParam *blockchain.Params

	wg   sync.WaitGroup
	quit chan struct{}

	//
	syncPeer *peer.Peer
}
