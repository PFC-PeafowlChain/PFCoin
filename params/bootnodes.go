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
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://f11f02bda595f745333fbb981d938ce705531dcc0f6c94b0a7a4dd17355bdcd56b8e5d81c3a7cfdd41cbe8a8ca2795083fd07940294ea865d30279b7304bf2dd@96.126.124.236:38988", // US
	"enode://423d2b5d3cc970b4b0512ecf2dfc05db308eb4ac8522e4b9f6de513841f89c425db1aa7f3c2a1ac59c6c624092d3586ad6c4ccd2be33465b286e3127271f2f1d@47.91.214.69:38988", // HK
	"enode://b7e4c87d948a88ea435e06be48e3c376531e17adccf18e416c0473ed26430009985c0abbe98a3781ba8909470714287802661ffb8b248553b32b1f5dae18fb21@212.71.235.148:38988", // UK
	"enode://e4b28fe42ae986d49f16b18d31632ac40b02dc484038a0f407d0326674c6b14445ddd0ecdd84d9381c9315a088a18e67bdc851a22168fb559e7e217fa57d6ddf@172.104.160.43:38988",  // SG
	"enode://7a49068756375cfa99b23ceb3c13ba5a61a81c81bb687c026c4f595eb6f46ca6af7fa1754b7b2de034a3dbda6713b2f3dc8187d26279b8d208fd70a005590c93@139.162.105.170:38988",  // JP
	"enode://f5a4b00a24ea19685c90fba2c4e00826bc774852590ac44b2855c3961a2cd1398bf6deb7e917e48ad9cecf63c383d232d58252f741296bf6ae72e5aab54e3be2@139.162.113.189:38988", // JP
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
	"enode://9587b91a36b7a5865d68621cdbc91c3d19be01c527da47722eff30d7a6353eff80a2bc0935f951d448108446b816e5b95c1ac360446b16283a95ac565fed16a3@96.126.124.236:30303",
	"enode://59f633d5bea9deba09edf8d0b9d356ba18ef51e3a179a2c2dd8e1d32fcef105e4b1425ca08a6e50abf082928948d3ad7ddd6ad04f3ba32e4bbf41fe55449ee55@139.162.105.170:30303",
	"enode://1c6321479224e21ca3cb7c39bc7ccac38e964aa325e68400564677fc9ee9f2f7ef4e64aec338e2f036fb749e92440191d5ae4e5d4fdf31c057aa99e35f773dae@139.162.113.189:30303",
	"enode://81b1d387ddd5249bd66e522c9c4add05aef3ee050b438d688b3e9daee9dcf0e31125888361abfb4a9b7e4e73f3f8b012fb75d4fd8ac3aac8dcc43e8245f34996@47.91.214.69:30303",
	"enode://13d96969ea3fe39fd5e305c327b994caf0c8f3927f6b93f72a87a5e7c2a59626550b60c8a11cb4e7811ac6127c1e600e00a4fb382dd16c8c20f46818d73e4cc2@172.104.160.43:30303",
	"enode://25f30d1d5f2c571bda072feb6090ee1f37a5d820e4999157ec578cda6e7ba9a8ffdaf342c8fba36ab336ef5af68e7566d26f531e66d466def777d097e4cdf22d@212.71.235.148:30303",
}
