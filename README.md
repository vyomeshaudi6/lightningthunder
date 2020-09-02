## Lightning Network Daemon

[![Build Status](https://img.shields.io/travis/lightningnetwork/lnd.svg)](https://travis-ci.org/lightningnetwork/lnd)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/lightningnetwork/lnd/blob/master/LICENSE)
[![Irc](https://img.shields.io/badge/chat-on%20freenode-brightgreen.svg)](https://webchat.freenode.net/?channels=lnd)
[![Godoc](https://godoc.org/github.com/lightningnetwork/lnd?status.svg)](https://godoc.org/github.com/lightningnetwork/lnd)

<img src="logo.png">

The Lightning Network Daemon (`lnd`) - is a complete implementation of a
[Lightning Network](https://lightning.network) node.  `lnd` has several pluggable back-end
chain services including [`btcd`](https://github.com/btcsuite/btcd) (a
full-node), [`bitcoind`](https://github.com/bitcoin/bitcoin), and
[`neutrino`](https://github.com/lightninglabs/neutrino) (a new experimental light client). The project's codebase uses the
[btcsuite](https://github.com/btcsuite/) set of Bitcoin libraries, and also
exports a large set of isolated re-usable Lightning Network related libraries
within it.  In the current state `lnd` is capable of:
* Creating channels.
* Closing channels.
* Completely managing all channel states (including the exceptional ones!).
* Maintaining a fully authenticated+validated channel graph.
* Performing path finding within the network, passively forwarding incoming payments.
* Sending outgoing [onion-encrypted payments](https://github.com/lightningnetwork/lightning-onion)
through the network.
* Updating advertised fee schedules.
* Automatic channel management ([`autopilot`](https://github.com/lightningnetwork/lnd/tree/master/autopilot)).

## Lightning Network Specification Compliance
`lnd` _fully_ conforms to the [Lightning Network specification
(BOLTs)](https://github.com/lightningnetwork/lightning-rfc). BOLT stands for:
Basis of Lightning Technology. The specifications are currently being drafted
by several groups of implementers based around the world including the
developers of `lnd`. The set of specification documents as well as our
implementation of the specification are still a work-in-progress. With that
said, the current status of `lnd`'s BOLT compliance is:

  - [X] BOLT 1: Base Protocol
  - [X] BOLT 2: Peer Protocol for Channel Management
  - [X] BOLT 3: Bitcoin Transaction and Script Formats
  - [X] BOLT 4: Onion Routing Protocol
  - [X] BOLT 5: Recommendations for On-chain Transaction Handling
  - [X] BOLT 7: P2P Node and Channel Discovery
  - [X] BOLT 8: Encrypted and Authenticated Transport
  - [X] BOLT 9: Assigned Feature Flags
  - [X] BOLT 10: DNS Bootstrap and Assisted Node Location
  - [X] BOLT 11: Invoice Protocol for Lightning Payments

## Developer Resources

The daemon has been designed to be as developer friendly as possible in order
to facilitate application development on top of `lnd`. Two primary RPC
interfaces are exported: an HTTP REST API, and a [gRPC](https://grpc.io/)
service. The exported API's are not yet stable, so be warned: they may change
drastically in the near future.

An automatically generated set of documentation for the RPC APIs can be found
at [api.lightning.community](https://api.lightning.community). A set of developer
resources including talks, articles, and example applications can be found at:
[dev.lightning.community](https://dev.lightning.community).

Finally, we also have an active
[Slack](https://lightning.engineering/slack.html) where protocol developers, application developers, testers and users gather to
discuss various aspects of `lnd` and also Lightning in general.

## Installation
  In order to build from source, please see [the installation
  instructions](docs/INSTALL.md).

## Docker
  To run lnd from Docker, please see the main [Docker instructions](docs/DOCKER.md)

## IRC
  * irc.freenode.net
  * channel #lnd
  * [webchat](https://webchat.freenode.net/?channels=lnd)

## Safety

When operating a mainnet `lnd` node, please refer to our [operational safety
guildelines](docs/safety.md). It is important to note that `lnd` is still
**beta** software and that ignoring these operational guidelines can lead to
loss of funds.

## Security

The developers of `lnd` take security _very_ seriously. The disclosure of
security vulnerabilities helps us secure the health of `lnd`, privacy of our
users, and also the health of the Lightning Network as a whole.  If you find
any issues regarding security or privacy, please disclose the information
responsibly by sending an email to security at lightning dot engineering,
preferably [encrypted using our designated PGP key
(`91FE464CD75101DA6B6BAB60555C6465E5BCB3AF`) which can be found
[here](https://gist.githubusercontent.com/Roasbeef/6fb5b52886183239e4aa558f83d085d3/raw/5ef96c426e3cf20a2443dc9d3c7d6877576da9ca/security@lightning.engineering).

## Further reading
* [Step-by-step send payment guide with docker](https://github.com/lightningnetwork/lnd/tree/master/docker)
* [Contribution guide](https://github.com/lightningnetwork/lnd/blob/master/docs/code_contribution_guidelines.md)

## CUSTOM LND DESCRIPTION

Directory structure
-------------------
    btcd                                                Contains Blockchain, btcd.conf etc.
    lnd                                                 Contains lnd.conf, tls cert etc.
    gocode/dev/test_data_PrvW/graph/simnet/             Directory containing Nodes active on simnet network.
    gocode/dev/test_data_PrvW/graph/testnet/            Directory containing Nodes active on testnet network.
    graph/simnet/User_Id/                               Contains channel.db, wallet.db, sphinxreplay.db, macaroons of respective userid node on simnet.
    graph/testnet/User_Id/                              Contains channel.db, wallet.db, sphinxreplay.db, macaroons of respective userid node on testnet.

 Custom Ports
 ------------
 Custom Ports can be changed accordingly in lnd.conf.
 
    0.0.0.0:10001           RPC listener for wallet actions(Create/Unlock).
    0.0.0.0:10003           REST listener for wallet actions(Create/Unlock).
    0.0.0.0:10002           RPC listener for Lightning service calls(getinfo/sendpayment....)
    0.0.0.0:10004           REST listener for Lightning service calls(getinfo/sendpayment....)
    0.0.0.0:10005           Listener port for the peers to open channel with the respective node got unlocked in the queue first.
    0.0.0.0:10006           Listener port for the peers to open channel with the respective node got unlocked in the queue second.
    0.0.0.0:10007           Listener port for the peers to open channel with the respective node got unlocked in the queue third.
    So on...

LND REST API REFERENCE
----------------------
## Create wallet

### 1.) /v1/genseed
Once the cipherseed is obtained and verified by the user, the InitWallet method should be used to commit the newly generated seed, and create the wallet.

Replace `User_Id` with correct value.

    
    ~/gocode/dev$  curl -X GET --insecure --cacert ~/.lnd/tls.cert  https://127.0.0.1:20001/v1/genseed/User_Id

Output:

    { 
    "cipher_seed_mnemonic": ["abandon","clean","mammal","rebel","again","call","outside","crawl","embrace","buddy","boy","plastic","core","fruit","chicken","warm","village","like","prevent","pudding","laptop","woman","height","little"], 
    "enciphered_seed": <byte>, 
    }

### 2.) /v1/initwallet
This can be used along with the GenSeed RPC to obtain a seed, then present it to the user. Once it has been verified by the user, the seed can be fed into this RPC in order to commit the new wallet.he wallet.

Replace `User_Id , wallet_password, cipher_seed_mnemonic` with correct value.

    
    ~/gocode/dev$  curl -X POST -d '{"wallet_password":"NzE3NzY1NzI3NDc5NzU2OTBhCg==","cipher_seed_mnemonic":["abandon","clean","mammal","rebel","again","call","outside","crawl","embrace","buddy","boy","plastic","core","fruit","chicken","warm","village","like","prevent","pudding","laptop","woman","height","little"]}' --insecure --cacert ~/.lnd/tls.cert  https://127.0.0.1:20001/v1/initwallet/User_Id

Output:

    { 
    }
    
## Unlock wallet
### 1.) /v1/unlockwallet

Replace `User_Id, wallet_password` with correct value.

    
    ~/gocode/dev$  curl -X POST -d '{"wallet_password":"NzE3NzY1NzI3NDc5NzU2OTBhCg=="}' --insecure --cacert ~/.lnd/tls.cert https://127.0.0.1:20001/v1/unlockwallet/User_Id

Output:

    { 
    }
    
## Lightning Service Calls
### 1.) /v1/getinfo

Replace `User_Id` `userPORT` with correct value.

    ~/gocode/dev$  MACAROON_HEADER="Grpc-Metadata-macaroon: $(xxd -ps -u -c 1000 /home/ubuntu/gocode/dev/test_data_PrvW/graph/testnet/User_Id/admin.macaroon)"
    ~/gocode/dev$  curl -X GET  --insecure --cacert ~/.lnd/tls.cert --header "$MACAROON_HEADER" https://127.0.0.1:userPORT/v1/getinfo/User_Id

Output:

    { 
    "version": <string>, 
    "commit_hash": <string>, 
    "identity_pubkey": <string>, 
    "alias": <string>, 
    "color": <string>, 
    "num_pending_channels": <int64>, 
    "num_active_channels": <int64>, 
    "num_inactive_channels": <int64>, 
    "num_peers": <int64>, 
    "block_height": <int64>, 
    "block_hash": <string>, 
    "best_header_timestamp": <string>, 
    "synced_to_chain": <boolean>, 
    "synced_to_graph": <boolean>, 
    "testnet": <boolean>, 
    "chains": <array lnrpcChain>, 
    "uris": <array string>, 
    "features": <object>, 
    }
    
 ### 2.) GET /v1/invoices

Replace `User_Id``userPORT` with correct value.

    ~/gocode/dev$  MACAROON_HEADER="Grpc-Metadata-macaroon: $(xxd -ps -u -c 1000 /home/ubuntu/gocode/dev/test_data_PrvW/graph/testnet/User_Id/admin.macaroon)"
    ~/gocode/dev$  curl -X GET  --insecure --cacert ~/.lnd/tls.cert --header "$MACAROON_HEADER" https://127.0.0.1:userPORT/v1/invoices/User_Id

Output:

    { 
    "invoices": <array lnrpcInvoice>, 
    "last_index_offset": <string>, 
    "first_index_offset": <string>, 
    }
  
### 3.) POST /v1/invoices

Replace `User_Id``userPORT` with correct value.

    ~/gocode/dev$  MACAROON_HEADER="Grpc-Metadata-macaroon: $(xxd -ps -u -c 1000 /home/ubuntu/gocode/dev/test_data_PrvW/graph/testnet/User_Id/admin.macaroon)"
    ~/gocode/dev$  curl -X POST --cacert ~/.lnd/tls.cert --header "$MACAROON_HEADER" https://localhost:userPORT/v1/invoices/User_Id -d '{ "value":"100" }'

Output:

    {
    "invoices":[{"memo":"","r_preimage":"aWGRZd8wQFNdcqY7JcMoltqTlkqIYx45r526eP5o87I=","r_hash":"wNDjbIRmmKAXKlpG+PRWNt9V94hwIIxq+Zx+BhkD7v8=","value":"100","value_msat":"100000      ","settled":false,"creation_date":"1596721759","settle_date":"0","payment_request":"lnsb1u1p0jczjlpp5crgwxmyyv6v2q9e2tfr03azkxm04taugwqsgc6hen3lqvxgramlsdqqcqzpgs  p5tp72r9xjkjqnp4j3y6rmqeht6whh7x5zq42d92x7rn0y0jqzt57s9qy9qsqzrk3yh79nj6tfthjskmjhu6uc8mvuljpz2hh2q2l7xxtwkypf7dhcev3ja59mqf4lm8rxsprn6254yk7wdfxhw3c6ejav7sm3jd3gqgq54suqj","description_hash":null,"expiry":"3600","fallback_addr":"","cltv_expiry":"40","route_hints":[],"private":false,"add_index":"1","settle_index":"0","amt_paid":"0","amt_paid_sat":"0","amt_paid_msat":"0","state":"OPEN","htlcs":[],"features":{"9":{"name":"tlv-onion","is_required":false,"is_known":true},"15":{"name":"payment-addr","is_required":false,"is_known":true},"17":{"name":"multi-path-payments","is_required":false,"is_known":true}},"is_keysend":false,"User_Id":""}],
    "last_index_offset":"1",
    "first_index_offset":"1"
    }

For other API calls  please refer to [LND REST API Reference](https://api.lightning.community/?shell#lnd-rest-api-reference),with the similar User_Id/PORT changes respectively in URI's.
