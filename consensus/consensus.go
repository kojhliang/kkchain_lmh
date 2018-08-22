/*
MIT License

Copyright (c) 2018 invin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package consensus

import (
	pb "github.com/invin/kkchain/protos"
)

// ConsensusExector divde consensus procedure to three parts
type ConsensusExector interface {
	Prepare(tag interface{})
	Execute(tag interface{})                           // Called whenever Execute completes
	Commit(tag interface{}, target *pb.BlockchainInfo) // Called whenever Commit completes
}

// Consenter is used to receive messages from the network
// Every consensus plugin needs to implement this interface
type Consenter interface {
	HandleNonTransMsg() error // handle non transaction messages
	ConsensusExector
}

// Stack is the set of stack-facing methods available to the consensus plugin
type Stack interface {
	NetworkStack
	SecurityUtils
	ChainStack
}

// Inquirer is used to retrieve info about the validating network
type Inquirer interface {
	GetNetworkInfo() (self *pb.PeerEndpoint, network []*pb.PeerEndpoint, err error)
	GetNetworkHandles() (self *pb.PeerID, network []*pb.PeerID, err error)
}

// Communicator is used to send messages to other validators
type Communicator interface {
	Broadcast(msg *pb.Message, peerType pb.PeerEndpoint_Type) error
	Unicast(msg *pb.Message, receiverHandle *pb.PeerID) error
}

// NetworkStack is used to retrieve network info and send messages
type NetworkStack interface {
	Communicator
	Inquirer
}

// SecurityUtils is used to access the sign/verify methods from the crypto package
type SecurityUtils interface {
	Sign(msg []byte) ([]byte, error)
	Verify(peerID *pb.PeerID, signature []byte, message []byte) error
}

// ChainStack is used to read leger and write ledger
type ChainStack interface {
	ChainReader
	ChainWriter
}

// ChainReader is used for get data from the blockchain
type ChainReader interface {
	GetBlock(id uint64) (block *pb.Block, err error)
	GetBlockchainSize() uint64
	GetBlockchainInfo() *pb.BlockchainInfo
	GetBlockchainInfoBlob() []byte
	GetBlockHeadMetadata() ([]byte, error)
}

// ChainWriter is used for write data into the blockchain
type ChainWriter interface {
	WriterBlock(block *pb.Block) (err error)
}
