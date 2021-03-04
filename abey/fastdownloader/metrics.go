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

// Contains the metrics collected by the downloader.

package fastdownloader

import (
	"github.com/abeychain/go-abey/metrics"
)

var (
	headerInMeter      = metrics.NewRegisteredMeter("abey/fastdownloader/headers/in", nil)
	headerReqTimer     = metrics.NewRegisteredTimer("abey/fastdownloader/headers/req", nil)
	headerDropMeter    = metrics.NewRegisteredMeter("abey/fastdownloader/headers/drop", nil)
	headerTimeoutMeter = metrics.NewRegisteredMeter("abey/fastdownloader/headers/timeout", nil)

	bodyInMeter      = metrics.NewRegisteredMeter("abey/fastdownloader/bodies/in", nil)
	bodyReqTimer     = metrics.NewRegisteredTimer("abey/fastdownloader/bodies/req", nil)
	bodyDropMeter    = metrics.NewRegisteredMeter("abey/fastdownloader/bodies/drop", nil)
	bodyTimeoutMeter = metrics.NewRegisteredMeter("abey/fastdownloader/bodies/timeout", nil)

	receiptInMeter      = metrics.NewRegisteredMeter("abey/fastdownloader/receipts/in", nil)
	receiptReqTimer     = metrics.NewRegisteredTimer("abey/fastdownloader/receipts/req", nil)
	receiptDropMeter    = metrics.NewRegisteredMeter("abey/fastdownloader/receipts/drop", nil)
	receiptTimeoutMeter = metrics.NewRegisteredMeter("abey/fastdownloader/receipts/timeout", nil)


)
