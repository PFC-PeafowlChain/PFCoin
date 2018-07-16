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
// the main Pfc network.
var MainnetBootnodes = []string{
	// Pfc Foundation Go Bootnodes
	"enode://56086a46210ea4521190fba70ff234ffbc4351aa8f2d35ca11a39651248505f15fb45e1d1bab917db5daba37e9ae3fc30d7cf00dfa4885b9894b245f84671b6a@50.116.14.19:38988", // US
	"enode://b427862d90bd385b9656c358828cc43db0a0984c11c6393f6d2baf04e56c219a0bc7a3a16ecbc13cb39e31bceb756687bb935a2a18586f95d7d05c4c4ab50213@50.116.6.164:38988", // US
	"enode://a1ddf8cc8e7c538bd591ad814922e360ddd5ba9820c59bec91a2618e5659ccd25da4f7d7e8a13f58a5972d47f99ae3f526bb98b35fa3cdd98ab21dcdecd8a76a@172.105.193.215:38988", // JP
	"enode://0a1713be8ad1cb2c63456cf5b8a4489f47fdb1bb94ed65b98b1d210dfc926e8bfcec022bf7d32254a9e915abd8c1031ebf48ff80e5f968a414f1769bc6d71dcc@139.162.98.26:38988", // JP
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the testnet
var TestnetBootnodes = []string{
	"enode://f11f02bda595f745333fbb981d938ce705531dcc0f6c94b0a7a4dd17355bdcd56b8e5d81c3a7cfdd41cbe8a8ca2795083fd07940294ea865d30279b7304bf2dd@96.126.124.236:38988", // US
	"enode://423d2b5d3cc970b4b0512ecf2dfc05db308eb4ac8522e4b9f6de513841f89c425db1aa7f3c2a1ac59c6c624092d3586ad6c4ccd2be33465b286e3127271f2f1d@47.91.214.69:38988", // HK
	"enode://b7e4c87d948a88ea435e06be48e3c376531e17adccf18e416c0473ed26430009985c0abbe98a3781ba8909470714287802661ffb8b248553b32b1f5dae18fb21@212.71.235.148:38988", // UK
	"enode://7a49068756375cfa99b23ceb3c13ba5a61a81c81bb687c026c4f595eb6f46ca6af7fa1754b7b2de034a3dbda6713b2f3dc8187d26279b8d208fd70a005590c93@139.162.105.170:38988",  // JP
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
var RinkebyBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
}
