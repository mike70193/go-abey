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

// Contains the metrics collected by the fetcher.

package fetcher

import (
	"github.com/abeychain/go-abey/metrics"
)

var (
	propAnnounceInMeter   = metrics.NewRegisteredMeter("abey/fetcher/prop/announces/in", nil)
	propAnnounceOutTimer  = metrics.NewRegisteredTimer("abey/fetcher/prop/announces/out", nil)
	propAnnounceDropMeter = metrics.NewRegisteredMeter("abey/fetcher/prop/announces/drop", nil)
	propAnnounceDOSMeter  = metrics.NewRegisteredMeter("abey/fetcher/prop/announces/dos", nil)

	propBroadcastInMeter      = metrics.NewRegisteredMeter("abey/fetcher/prop/broadcasts/in", nil)
	propBroadcastOutTimer     = metrics.NewRegisteredTimer("abey/fetcher/prop/broadcasts/out", nil)
	propBroadcastDropMeter    = metrics.NewRegisteredMeter("abey/fetcher/prop/broadcasts/drop", nil)
	propBroadcastInvaildMeter = metrics.NewRegisteredMeter("abey/fetcher/prop/broadcasts/invaild", nil)
	propBroadcastDOSMeter     = metrics.NewRegisteredMeter("abey/fetcher/prop/broadcasts/dos", nil)

	propSignInvaildMeter = metrics.NewRegisteredMeter("abey/fetcher/prop/signs/invaild", nil)

	headerFetchMeter = metrics.NewRegisteredMeter("abey/fetcher/fetch/headers", nil)
	bodyFetchMeter   = metrics.NewRegisteredMeter("abey/fetcher/fetch/bodies", nil)

	headerFilterInMeter  = metrics.NewRegisteredMeter("abey/fetcher/filter/headers/in", nil)
	headerFilterOutMeter = metrics.NewRegisteredMeter("abey/fetcher/filter/headers/out", nil)
	bodyFilterInMeter    = metrics.NewRegisteredMeter("abey/fetcher/filter/bodies/in", nil)
	bodyFilterOutMeter   = metrics.NewRegisteredMeter("abey/fetcher/filter/bodies/out", nil)
)
