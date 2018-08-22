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
	"strings"

	"github.com/invin/kkchain/consensus"
	"github.com/invin/kkchain/consensus/algorithm/poa"
	"github.com/invin/kkchain/consensus/pbft"
	"github.com/op/go-logging"
)

var logger *logging.Logger // package-level logger
var consenter consensus.Consenter

func init() {
	logger = logging.MustGetLogger("consensus/controller")
}

// NewConsenter constructs a Consenter object if not already present
func NewConsenter(stack consensus.Stack) consensus.Consenter {

	plugin := strings.ToLower(AppConfig.consensus.GetString("consensus.plugin"))
	if plugin == "pow" {
		logger.Infof("Creating consensus plugin %s", plugin)
		return pow.GetPlugin(stack)
	}
	if plugin == "pos" {
		logger.Infof("Creating consensus plugin %s", plugin)
		return pos.GetPlugin(stack)
	}
	if plugin == "dpos" {
		logger.Infof("Creating consensus plugin %s", plugin)
		return dpos.GetPlugin(stack)
	}
	if plugin == "raft" {
		logger.Infof("Creating consensus plugin %s", plugin)
		return raft.GetPlugin(stack)
	}
	if plugin == "pbft" {
		logger.Infof("Creating consensus plugin %s", plugin)
		return pbft.GetPlugin(stack)
	}
	logger.Info("Creating default consensus plugin (POA)")
	return poa.GetPlugin(stack)

}
