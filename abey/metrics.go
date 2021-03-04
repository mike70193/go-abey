// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package abey

import (
	"github.com/abeychain/go-abey/metrics"
	"github.com/abeychain/go-abey/p2p"
)

var (
	propTxnInPacketsMeter     = metrics.NewRegisteredMeter("abey/prop/txns/in/packets", nil)
	propTxnInTxsMeter         = metrics.NewRegisteredMeter("abey/prop/txns/in/txs", nil)
	propTxnInTrafficMeter     = metrics.NewRegisteredMeter("abey/prop/txns/in/traffic", nil)
	propTxnOutPacketsMeter    = metrics.NewRegisteredMeter("abey/prop/txns/out/packets", nil)
	propTxnOutTrafficMeter    = metrics.NewRegisteredMeter("abey/prop/txns/out/traffic", nil)
	propFtnInPacketsMeter     = metrics.NewRegisteredMeter("abey/prop/ftns/in/packets", nil)
	propFtnInTrafficMeter     = metrics.NewRegisteredMeter("abey/prop/ftns/in/traffic", nil)
	propFtnOutPacketsMeter    = metrics.NewRegisteredMeter("abey/prop/ftns/out/packets", nil)
	propFtnOutTrafficMeter    = metrics.NewRegisteredMeter("abey/prop/ftns/out/traffic", nil)
	propFHashInPacketsMeter   = metrics.NewRegisteredMeter("abey/prop/fhashes/in/packets", nil)
	propFHashInTrafficMeter   = metrics.NewRegisteredMeter("abey/prop/fhashes/in/traffic", nil)
	propFHashOutPacketsMeter  = metrics.NewRegisteredMeter("abey/prop/fhashes/out/packets", nil)
	propFHashOutTrafficMeter  = metrics.NewRegisteredMeter("abey/prop/fhashes/out/traffic", nil)
	propSHashInPacketsMeter   = metrics.NewRegisteredMeter("abey/prop/shashes/in/packets", nil)
	propSHashInTrafficMeter   = metrics.NewRegisteredMeter("abey/prop/shashes/in/traffic", nil)
	propSHashOutPacketsMeter  = metrics.NewRegisteredMeter("abey/prop/shashes/out/packets", nil)
	propSHashOutTrafficMeter  = metrics.NewRegisteredMeter("abey/prop/shashes/out/traffic", nil)
	propFBlockInPacketsMeter  = metrics.NewRegisteredMeter("abey/prop/fblocks/in/packets", nil)
	propFBlockInTrafficMeter  = metrics.NewRegisteredMeter("abey/prop/fblocks/in/traffic", nil)
	propFBlockOutPacketsMeter = metrics.NewRegisteredMeter("abey/prop/fblocks/out/packets", nil)
	propFBlockOutTrafficMeter = metrics.NewRegisteredMeter("abey/prop/fblocks/out/traffic", nil)
	propSBlockInPacketsMeter  = metrics.NewRegisteredMeter("abey/prop/sblocks/in/packets", nil)
	propSBlockInTrafficMeter  = metrics.NewRegisteredMeter("abey/prop/sblocks/in/traffic", nil)
	propSBlockOutPacketsMeter = metrics.NewRegisteredMeter("abey/prop/sblocks/out/packets", nil)
	propSBlockOutTrafficMeter = metrics.NewRegisteredMeter("abey/prop/sblocks/out/traffic", nil)

	propNodeInfoInPacketsMeter    = metrics.NewRegisteredMeter("abey/prop/nodeinfo/in/packets", nil)
	propNodeInfoInTrafficMeter  = metrics.NewRegisteredMeter("abey/prop/nodeinfo/in/traffic", nil)
	propNodeInfoOutPacketsMeter = metrics.NewRegisteredMeter("abey/prop/nodeinfo/out/packets", nil)
	propNodeInfoOutTrafficMeter = metrics.NewRegisteredMeter("abey/prop/nodeinfo/out/traffic", nil)

	propNodeInfoHashInPacketsMeter    = metrics.NewRegisteredMeter("abey/prop/nodeinfohash/in/packets", nil)
	propNodeInfoHashInTrafficMeter  = metrics.NewRegisteredMeter("abey/prop/nodeinfohash/in/traffic", nil)
	propNodeInfoHashOutPacketsMeter = metrics.NewRegisteredMeter("abey/prop/nodeinfohash/out/packets", nil)
	propNodeInfoHashOutTrafficMeter = metrics.NewRegisteredMeter("abey/prop/nodeinfohash/out/traffic", nil)



	reqFHeaderInPacketsMeter  = metrics.NewRegisteredMeter("abey/req/headers/in/packets", nil)
	reqFHeaderInTrafficMeter  = metrics.NewRegisteredMeter("abey/req/headers/in/traffic", nil)
	reqFHeaderOutPacketsMeter = metrics.NewRegisteredMeter("abey/req/headers/out/packets", nil)
	reqFHeaderOutTrafficMeter = metrics.NewRegisteredMeter("abey/req/headers/out/traffic", nil)
	reqSHeaderInPacketsMeter  = metrics.NewRegisteredMeter("abey/req/sheaders/in/packets", nil)
	reqSHeaderInTrafficMeter  = metrics.NewRegisteredMeter("abey/req/sheaders/in/traffic", nil)
	reqSHeaderOutPacketsMeter = metrics.NewRegisteredMeter("abey/req/sheaders/out/packets", nil)
	reqSHeaderOutTrafficMeter = metrics.NewRegisteredMeter("abey/req/sheaders/out/traffic", nil)

	reqFBodyInPacketsMeter  = metrics.NewRegisteredMeter("abey/req/fbodies/in/packets", nil)
	reqFBodyInTrafficMeter  = metrics.NewRegisteredMeter("abey/req/fbodies/in/traffic", nil)
	reqFBodyOutPacketsMeter = metrics.NewRegisteredMeter("abey/req/fbodies/out/packets", nil)
	reqFBodyOutTrafficMeter = metrics.NewRegisteredMeter("abey/req/fbodies/out/traffic", nil)
	reqSBodyInPacketsMeter  = metrics.NewRegisteredMeter("abey/req/sbodies/in/packets", nil)
	reqSBodyInTrafficMeter  = metrics.NewRegisteredMeter("abey/req/sbodies/in/traffic", nil)
	reqSBodyOutPacketsMeter = metrics.NewRegisteredMeter("abey/req/sbodies/out/packets", nil)
	reqSBodyOutTrafficMeter = metrics.NewRegisteredMeter("abey/req/sbodies/out/traffic", nil)

	reqStateInPacketsMeter    = metrics.NewRegisteredMeter("abey/req/states/in/packets", nil)
	reqStateInTrafficMeter    = metrics.NewRegisteredMeter("abey/req/states/in/traffic", nil)
	reqStateOutPacketsMeter   = metrics.NewRegisteredMeter("abey/req/states/out/packets", nil)
	reqStateOutTrafficMeter   = metrics.NewRegisteredMeter("abey/req/states/out/traffic", nil)
	reqReceiptInPacketsMeter  = metrics.NewRegisteredMeter("abey/req/receipts/in/packets", nil)
	reqReceiptInTrafficMeter  = metrics.NewRegisteredMeter("abey/req/receipts/in/traffic", nil)
	reqReceiptOutPacketsMeter = metrics.NewRegisteredMeter("abey/req/receipts/out/packets", nil)
	reqReceiptOutTrafficMeter = metrics.NewRegisteredMeter("abey/req/receipts/out/traffic", nil)

	getHeadInPacketsMeter  = metrics.NewRegisteredMeter("abey/get/head/in/packets", nil)
	getHeadInTrafficMeter  = metrics.NewRegisteredMeter("abey/get/head/in/traffic", nil)
	getHeadOutPacketsMeter = metrics.NewRegisteredMeter("abey/get/head/out/packets", nil)
	getHeadOutTrafficMeter = metrics.NewRegisteredMeter("abey/get/head/out/traffic", nil)
	getBodyInPacketsMeter  = metrics.NewRegisteredMeter("abey/get/bodies/in/packets", nil)
	getBodyInTrafficMeter  = metrics.NewRegisteredMeter("abey/get/bodies/in/traffic", nil)
	getBodyOutPacketsMeter = metrics.NewRegisteredMeter("abey/get/bodies/out/packets", nil)
	getBodyOutTrafficMeter = metrics.NewRegisteredMeter("abey/get/bodies/out/traffic", nil)

	getNodeInfoInPacketsMeter    = metrics.NewRegisteredMeter("abey/get/nodeinfo/in/packets", nil)
	getNodeInfoInTrafficMeter  = metrics.NewRegisteredMeter("abey/get/nodeinfo/in/traffic", nil)
	getNodeInfoOutPacketsMeter = metrics.NewRegisteredMeter("abey/get/nodeinfo/out/packets", nil)
	getNodeInfoOutTrafficMeter = metrics.NewRegisteredMeter("abey/get/nodeinfo/out/traffic", nil)

	miscInPacketsMeter  = metrics.NewRegisteredMeter("abey/misc/in/packets", nil)
	miscInTrafficMeter  = metrics.NewRegisteredMeter("abey/misc/in/traffic", nil)
	miscOutPacketsMeter = metrics.NewRegisteredMeter("abey/misc/out/packets", nil)
	miscOutTrafficMeter = metrics.NewRegisteredMeter("abey/misc/out/traffic", nil)
)

// meteredMsgReadWriter is a wrapper around a p2p.MsgReadWriter, capable of
// accumulating the above defined metrics based on the data stream contents.
type meteredMsgReadWriter struct {
	p2p.MsgReadWriter     // Wrapped message stream to meter
	version           int // Protocol version to select correct meters
}

// newMeteredMsgWriter wraps a p2p MsgReadWriter with metering support. If the
// metrics system is disabled, this function returns the original object.
func newMeteredMsgWriter(rw p2p.MsgReadWriter) p2p.MsgReadWriter {
	if !metrics.Enabled {
		return rw
	}
	return &meteredMsgReadWriter{MsgReadWriter: rw}
}

// Init sets the protocol version used by the stream to know which meters to
// increment in case of overlapping message ids between protocol versions.
func (rw *meteredMsgReadWriter) Init(version int) {
	rw.version = version
}

func (rw *meteredMsgReadWriter) ReadMsg() (p2p.Msg, error) {
	// Read the message and short circuit in case of an error
	msg, err := rw.MsgReadWriter.ReadMsg()
	if err != nil {
		return msg, err
	}
	// Account for the data traffic
	packets, traffic := miscInPacketsMeter, miscInTrafficMeter
	switch {
	case msg.Code == FastBlockHeadersMsg:
		packets, traffic = reqFHeaderInPacketsMeter, reqFHeaderInTrafficMeter
	case msg.Code == SnailBlockHeadersMsg:
		packets, traffic = reqSHeaderInPacketsMeter, reqSHeaderInTrafficMeter
	case msg.Code == FastBlockBodiesMsg:
		packets, traffic = reqFBodyInPacketsMeter, reqFBodyInTrafficMeter
	case msg.Code == SnailBlockBodiesMsg:
		packets, traffic = reqSBodyInPacketsMeter, reqSBodyInTrafficMeter

	case msg.Code == NodeDataMsg:
		packets, traffic = reqStateInPacketsMeter, reqStateInTrafficMeter
	case msg.Code == ReceiptsMsg:
		packets, traffic = reqReceiptInPacketsMeter, reqReceiptInTrafficMeter

	case msg.Code == NewFastBlockHashesMsg:
		packets, traffic = propFHashInPacketsMeter, propFHashInTrafficMeter
	case msg.Code == NewSnailBlockHashesMsg:
		packets, traffic = propSHashInPacketsMeter, propSHashInTrafficMeter
	case msg.Code == NewFastBlockMsg:
		packets, traffic = propFBlockInPacketsMeter, propFBlockInTrafficMeter
	case msg.Code == NewSnailBlockMsg:
		packets, traffic = propSBlockInPacketsMeter, propSBlockInTrafficMeter
	case msg.Code == TxMsg:
		packets, traffic = propTxnInPacketsMeter, propTxnInTrafficMeter
	case msg.Code == NewFruitMsg:
		packets, traffic = propFtnInPacketsMeter, propFtnInTrafficMeter
	case msg.Code == TbftNodeInfoMsg:
		packets, traffic = propNodeInfoInPacketsMeter, propNodeInfoInTrafficMeter
	case msg.Code == TbftNodeInfoHashMsg:
		packets, traffic = propNodeInfoHashInPacketsMeter, propNodeInfoHashInTrafficMeter
	case msg.Code == GetTbftNodeInfoMsg:
		packets, traffic = getNodeInfoInPacketsMeter, getNodeInfoInTrafficMeter
	case msg.Code == GetFastBlockHeadersMsg:
		packets, traffic = getHeadInPacketsMeter, getHeadInTrafficMeter
	case msg.Code == GetFastBlockBodiesMsg:
		packets, traffic = getHeadInPacketsMeter, getHeadInTrafficMeter
	case msg.Code == GetSnailBlockHeadersMsg:
		packets, traffic = getBodyInPacketsMeter, getBodyInTrafficMeter
	case msg.Code == GetSnailBlockBodiesMsg:
		packets, traffic = getBodyInPacketsMeter, getBodyInTrafficMeter
	}
	packets.Mark(1)
	traffic.Mark(int64(msg.Size))

	return msg, err
}

func (rw *meteredMsgReadWriter) WriteMsg(msg p2p.Msg) error {
	// Account for the data traffic
	packets, traffic := miscOutPacketsMeter, miscOutTrafficMeter
	switch {
	case msg.Code == FastBlockHeadersMsg:
		packets, traffic = reqFHeaderOutPacketsMeter, reqFHeaderOutTrafficMeter
	case msg.Code == SnailBlockHeadersMsg:
		packets, traffic = reqSHeaderOutPacketsMeter, reqSHeaderOutTrafficMeter
	case msg.Code == FastBlockBodiesMsg:
		packets, traffic = reqFBodyOutPacketsMeter, reqFBodyOutTrafficMeter
	case msg.Code == SnailBlockBodiesMsg:
		packets, traffic = reqSBodyOutPacketsMeter, reqSBodyOutTrafficMeter

	case msg.Code == NodeDataMsg:
		packets, traffic = reqStateOutPacketsMeter, reqStateOutTrafficMeter
	case msg.Code == ReceiptsMsg:
		packets, traffic = reqReceiptOutPacketsMeter, reqReceiptOutTrafficMeter

	case msg.Code == NewFastBlockHashesMsg:
		packets, traffic = propFHashOutPacketsMeter, propFHashOutTrafficMeter
	case msg.Code == NewSnailBlockHashesMsg:
		packets, traffic = propSHashOutPacketsMeter, propSHashOutTrafficMeter
	case msg.Code == NewFastBlockMsg:
		packets, traffic = propFBlockOutPacketsMeter, propFBlockOutTrafficMeter
	case msg.Code == NewSnailBlockMsg:
		packets, traffic = propSBlockOutPacketsMeter, propSBlockOutTrafficMeter
	case msg.Code == TxMsg:
		packets, traffic = propTxnOutPacketsMeter, propTxnOutTrafficMeter
	case msg.Code == NewFruitMsg:
		packets, traffic = propFtnOutPacketsMeter, propFtnOutTrafficMeter
	case msg.Code == TbftNodeInfoMsg:
		packets, traffic = propNodeInfoOutPacketsMeter, propNodeInfoOutTrafficMeter
	case msg.Code == TbftNodeInfoHashMsg:
		packets, traffic = propNodeInfoHashOutPacketsMeter, propNodeInfoHashOutTrafficMeter
	case msg.Code == GetTbftNodeInfoMsg:
		packets, traffic = getNodeInfoOutPacketsMeter, getNodeInfoOutTrafficMeter
	case msg.Code == GetFastBlockHeadersMsg:
		packets, traffic = getHeadOutPacketsMeter, getHeadOutTrafficMeter
	case msg.Code == GetFastBlockBodiesMsg:
		packets, traffic = getHeadInPacketsMeter, getHeadOutTrafficMeter
	case msg.Code == GetSnailBlockHeadersMsg:
		packets, traffic = getBodyOutPacketsMeter, getBodyOutTrafficMeter
	case msg.Code == GetSnailBlockBodiesMsg:
		packets, traffic = getBodyOutPacketsMeter, getBodyOutTrafficMeter
	}
	packets.Mark(1)
	traffic.Mark(int64(msg.Size))

	// Send the packet to the p2p layer
	return rw.MsgReadWriter.WriteMsg(msg)
}
