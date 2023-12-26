# `x/spec`

## Abstract

This document specifies the spec module of Lava Protocol.

The spec module is responsible for managing Lava's chains specifications. a spec is a struct difining all the api's of a specific chain and its bahaviours. The specs determines the chain properties like block time, finalization distance, imports and more. Adding a spec is done with a proposal on chain with the gov module. The first step of providers and subscriptions to work in lava is having a specification for the wanted chain.

This document is focuses on the spec' technical aspects. For more information on chain support in lava see pairing README.

## Contents
* [Concepts](#concepts)
  * [Spec](#spec)
  * [ApiCollections](#apicollections)
  * [CollectionData](#collectiondata)
  * [ParseDirective](#parsedirective)
  * [Verification](#verification)
  * [Import](#import)
* [Parameters](#parameters)
* [Queries](#queries)
* [Transactions](#transactions)
* [Proposals](#proposals)

## Concepts

### Spec

A chain spec consists of general properties of the chain and a list of interface it supports. To make specs simple to create and maintain specs can import api's from another, for example, X testnet spec can Import from X mainnet spec and thus not need to redifine all of the interfaces.

```
type Spec struct {
	Index                         string                // chain index (index key for the chain)                             
	Name                          string                // description string of the spec
	Enabled                       bool                  // spec enabled/disable (base specs are usually disabled, like cosmos base specs)
	AverageBlockTime              int64                 // average block time of the chain in msec
	MinStakeProvider              Coin                  // min stake for a provider to be active in the chain
	ProvidersTypes                Spec_ProvidersTypes   // determines if the spec is for lava or others 
	Imports                       []string              // list of chains to import ApiCollections from
	ApiCollections                []*ApiCollection      // list of ApiCollections that defines all the interfaces and api's of the chain
	Contributor                   []string              // list of contributers (public lava address {lava@...})
	ContributorPercentage         *Dec                  // the percentage of coins the contributers will get from each reward a provider get
	Shares                        uint64                // factor for bonus rewards at the end of the month (see rewards module)
    AllowedBlockLagForQosSync     int64                 // defines the accepted blocks a provider can be behind the chain without QOS degradation
	BlockLastUpdated              uint64                // the last block this spec was updated on chain
    ReliabilityThreshold          uint32                // this determines the amount of data reliability for the chain
	DataReliabilityEnabled        bool                  // enables/disables data reliability for the chain
	BlockDistanceForFinalizedData uint32                              
	BlocksInFinalizationProof     uint32                              
}
```
Note, the `Coin` type is from Cosmos-SDK (`cosmos.base.v1beta1.Coin`).
Note, the `Dec` type is from Cosmos-SDK math LegacyDec.

### ApiCollection

ApiCollection is a struct that defines an interface, for example rest, json etc.., and all of its api's and properties

```
type ApiCollection struct {
	Enabled         bool                // enables/disables the collection
	CollectionData  CollectionData      // defines the properties of the collection, also acts as a unique key
	Apis            []*Api              // list of api's in the collection
	Headers         []*Header           // list of headers supported by the interface and their behaviour
	InheritanceApis []*CollectionData   // list of other ApiCollection to inherite from
	ParseDirectives []*ParseDirective   // list of parsing instructions of specific api's
	Extensions      []*Extension        // list of extensions that providers can support in addition to the basic behaviour (for example, archive node)
	Verifications   []*Verification     // list of verifications that providers must pass to make sure they provide full functionality
}
```

### CollectionData

CollectionData defines the api properties and acts as a unique key for the [api collection](#apicollection)

TAKE THIS FROM PLANS

```
type CollectionData struct {
	ApiInterface string         // defines the connection interface (rest/json/grpc etc...)
	InternalPath string         // defines internal path of the node for this specific ApiCollection
	Type         string         // type of api (POST/GET)
	AddOn        string         // 
}
```

### ParseDirective

ParseDirective is a struct that defines a function needed by the provider in a generic way. it describes how for a specific api collection how to get information from the node. for example, how to get the latest block of an EVM node.

```
type ParseDirective struct {
	FunctionTag      FUNCTION_TAG // defines what the function this serves for
	FunctionTemplate string       // api template to fill and send to the node
	ResultParsing    BlockParser  // description for parsing the result of the api
	ApiName          string       // the api name
}
```

FunctionTag can be one of the following :
```
const (
	FUNCTION_TAG_GET_BLOCKNUM           FUNCTION_TAG = 1 // get latest block number
	FUNCTION_TAG_GET_BLOCK_BY_NUM       FUNCTION_TAG = 2 // get a specific block by block numer
	FUNCTION_TAG_SET_LATEST_IN_METADATA FUNCTION_TAG = 3
	FUNCTION_TAG_SET_LATEST_IN_BODY     FUNCTION_TAG = 4
	FUNCTION_TAG_VERIFICATION           FUNCTION_TAG = 5
)
```

### Verification

Verification is a struct that defines how to verify a specific property of the api collection, for example: verify the chain id of the node.

type Verification struct {
	Name           string                               // verification name
	ParseDirective *ParseDirective                      // [ParseDirective](#parsedirective) to get the the value to verify from the node
	Values         []*ParseValue                        // expected value we want from the result
	Severity       Verification_VerificationSeverity    // instructions for what to do if a verification fails
}

### Import

## Parameters

The plans module does not contain parameters.

## Queries

The plans module supports the following queries:

| Query             | Arguments         | What it does                                  |
| ----------        | ---------------   | ----------------------------------------------|
| `list-spec`       | none              | shows all the specs                           |
| `params`          | none              | show the params of the module                 |
| `show-all-chains` | none              | shows all the specs with minimal info         |
| `show-chain-info` | chainid           | shows a spec with minimal info                |
| `show-spec`       | chainid           | shows a full spec                             |

## Transactions

The plans module does not support any transactions.

## Proposals

The plans module provides a proposal to add/overwrite a spec to the chain

```
lavad tx gov submit-legacy-proposal spec-add <spec_json_1>,<spec_json_2> --from alice <gas-flags>

```

A valid `add_spec_json_1` JSON proposal format:

```
{
    "proposal": {
        "title": "Add Specs: Lava",
        "description": "Adding new specification support for relaying Lava data on Lava",
        "specs": [
            {
                "index": "LAV1",
                "name": "lava testnet",
                "enabled": true,
                "imports": [
                    "COSMOSSDK"
                ],
                "providers_types": 1,
                "reliability_threshold": 268435455,
                "data_reliability_enabled": true,
                "block_distance_for_finalized_data": 0,
                "blocks_in_finalization_proof": 1,
                "average_block_time": 30000,
                "allowed_block_lag_for_qos_sync": 2,
                                "shares" : 1,
                "min_stake_provider": {
                    "denom": "ulava",
                    "amount": "50000000000"
                },
                "api_collections": [
                    {
                        "enabled": true,
                        "collection_data": {
                            "api_interface": "rest",
                            "internal_path": "",
                            "type": "GET",
                            "add_on": ""
                        },
                        "apis": [
                            {
                                "name": "/lavanet/lava/spec/show_all_chains",
                                "block_parsing": {
                                    "parser_arg": [
                                        "latest"
                                    ],
                                    "parser_func": "DEFAULT"
                                },
                                "compute_units": 10,
                                "enabled": true,
                                "category": {
                                    "deterministic": true,
                                    "local": false,
                                    "subscription": false,
                                    "stateful": 0
                                },
                                "extra_compute_units": 0
                            }
                        ],
                        "headers": [],
                        "inheritance_apis": [],
                        "parse_directives": [],
                        "verifications": [
                            {
                                "name": "chain-id",
                                "values": [
                                    {
                                        "expected_value": "lava-testnet"
                                    }
                                ]
                            }
                        ]
                    }
                ]
            }
        ]
    },
    "deposit": "10000000ulava"
}
```