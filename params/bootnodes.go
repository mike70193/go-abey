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

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Abeychain network.
var MainnetBootnodes = []string{
	"enode://cd99daa76de43e2b7a5806c3455d33012cd127bca9b2e271be3af5d78e402c153a77e1d408f708770fb390e597621407f963f1c444090c21f91e03e03caa2110@39.98.216.197:30313", // CN
	"enode://e95937d68263a59c95ac1199eecc450b3590624accaf1542c7e51d8dc3ca3bfa6d3f60785b021c408b4a9a67b2869da33237c75448ae29b70506164a2bfe6931@13.52.156.74:30313",  // US WEST
	"enode://9032cc37954363b4d2dd37a898959aadf213718ff1bdb146848fb8c9a5adfd31d543ca870a08a223b27da2309051d0ce41775fa6de9337ed519b64cfa85b5b0c@47.241.22.155:30313", // SG
	"enode://4c64220af42271b6a6ea5463e97a125fef86d0bbb077db7d669af9d020d8ccf8ef4b617e3b36bbb9c10096404ecc1a7e06bcec3210a2cdf49b2bce5a0e1c7eb5@8.209.88.41:30313",   // DE

	"enode://fb331ff6aded86b393d9de2f9c449d313b356af0c4c0b9500e0f6c51bcb4ed31ca45dc2ab64c6182d1876eb9e3fd073d488277a40a6d357bc6e63350a2e00ffc@101.132.183.35:30313", // CN
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://e8e1af516589b5b49d0110bf52dcbc8c4fe592697cfd94eb01fdc302e2b4f1f9c7d3667ca4be64d094a15635204195681219dc59febe17c7c83f3b71b6258454@18.138.171.105:30313", // SG
	"enode://8b1af471dd393f2f228cefd663e3176963592adcd452e69b868ffb2cbb2f7b67ca44f60269e031952a00ba708799c8c1c2a41c279df80461e9b4d384ea49f9fc@18.140.45.222:30313",  // SG
	"enode://1554bd2e50456dbf1e5b94aabf88fc7eb1f274d39ae52067541d46e2b8043cb102bffffecb0ba42657c8904aec6bde4f1728110107f27f50616c18ede4af983c@52.220.163.103:30313", // SG
	"enode://d3b5fb4283424e6011d6ad1bcad7e3890fc94db4e6d221571a61985b1f48b6ed26733b9871debb18924cb299600611b683f08e1be08e9a320ffba44494388d1f@54.151.132.19:30313",  // SG
}

// DevnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the dev Abeychain network.
var DevnetBootnodes = []string{
	"enode://ec1e13e3d0177196a55570dfc1c810b2ea05109cb310c4dc7397ae6f3109467ec0d13a5f28ebdfb553511d492a4892ffa3a8283ce69bc5f93fce079dbfbfa5f4@39.100.120.25:30310",
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://ebb007b1efeea668d888157df36cf8fe49aa3f6fd63a0a67c45e4745dc081feea031f49de87fa8524ca29343a21a249d5f656e6daeda55cbe5800d973b75e061@39.98.171.41:30315",
	"enode://b5062c25dc78f8d2a8a216cebd23658f170a8f6595df16a63adfabbbc76b81b849569145a2629a65fe50bfd034e38821880f93697648991ba786021cb65fb2ec@39.98.43.179:30312",
}
