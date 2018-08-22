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
package poa

import (
	"time"

	"github.com/invin/kkchain/consensus"
	logging "github.com/op/go-logging"
)

var logger *logging.Logger // package-level logger

func init() {
	logger = logging.MustGetLogger("consensus/poa")
}

// Poa is a plugin object implementing the consensus.Consenter interface.
type Poa struct {
	stack    consensus.Stack
	txQ      *txq
	timer    *time.Timer
	duration time.Duration
	channel  chan *pb.Transaction
}

// Setting up a singleton POA consenter
var iPoa consensus.Consenter

// GetPlugin returns a singleton of POA
func GetPlugin(c consensus.Stack) consensus.Consenter {
	if iPlugin == nil {
		iPlugin = newPoa(c)
	}
	return iPoa
}

// newPoa is a constructor returning a consensus.Consenter object.
func newPoa(c consensus.Stack) consensus.Consenter {
}

// HandleNonTransMsg is called for Non transaction messages.
func (i *Poa) HandleNonTransMsg(msg *pb.Message, senderHandle *pb.PeerID) error {
}
