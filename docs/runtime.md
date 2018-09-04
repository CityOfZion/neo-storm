# Runtime
A brief overview of NEO smart contract API's that can be used in the neo-storm framework.

# Overview
1. [Account]()
2. [Asset]()
3. [Attribute]()
4. [Block]()
5. [Blockchain]()
6. [Contract]()
7. [Crypto]()
8. [Engine]()
9. [Enumerator]()
10. [Iterator]()
11. [Header]()
12. [Input]()
13. [Output]()
14. [Runtime]()
15. [Storage]()
16. [Transaction]()
17. [Util]()

## Account 
#### GetScriptHash
```
GetScriptHash(a Account) []byte
```
Returns the script hash of the given account.

#### GetVotes
```
GetVotes(a Account) [][]byte
```
Returns the the votes (a slice of public keys) of the given account.

#### GetBalance
```
GetBalance(a Account, assetID []byte) int
```
Returns the balance of the given asset id for the given account.

## Asset
#### GetAssetID 
```
GetAssetID(a Asset) []byte
```
Returns the id of the given asset.

#### GetAmount
```
GetAmount(a Asset) int 
```
Returns the amount of the given asset id.

#### GetAvailable
```
GetAvailable(a Asset) int 
```
Returns the available amount of the given asset.

#### GetPrecision
```
GetPrecision(a Asset) byte
```
Returns the precision of the given Asset.

#### GetOwner
```
GetOwner(a Asset) []byte
```
Returns the owner of the given asset.

#### GetAdmin
```
GetAdmin(a Asset) []byte
```
Returns the admin of the given asset.

#### GetIssuer
```
GetIssuer(a Asset) []byte
```
Returns the issuer of the given asset.

#### Create
```
Create(type byte, name string, amount int, precision byte, owner, admin, issuer []byte)
```
Creates a new asset on the blockchain.

#### Renew
```
Renew(asset Asset, years int)
```
Renews the given asset as long as the given years.

## Attribute
#### GetUsage
```
GetUsage(attr Attribute) []byte
```
Returns the usage of the given attribute.

#### GetData
```
GetData(attr Attribute) []byte
```
Returns the data of the given attribute.

## Block
#### GetTransactionCount
```
GetTransactionCount(b Block) int
```
Returns the number of transactions that are recorded in the given block.

#### GetTransactions
```
GetTransactions(b Block) []transaction.Transaction
```
Returns a slice of the transactions that are recorded in the given block.

#### GetTransaction
```
GetTransaction(b Block, hash []byte) transaction.Transaction
```
Returns the transaction by the given hash that is recorded in the given block.

## Blockchain
#### GetHeight
```
GetHeight() int
```
Returns the current height of the blockchain.

#### GetHeader
```
GetHeader(heightOrHash []interface{}) header.Header
```
Return the header by the given hash or index.

#### GetBlock
```
GetBlock(heightOrHash interface{}) block.Block
```
Returns the block by the given hash or index.

#### GetTransaction
```
GetTransaction(hash []byte) transaction.Transaction
```
Returns a transaction by the given hash.

## Contract

## Crypto

## Engine

## Enumerator

## Iterator

## Header

## Input

## Output

## Runtime

## Storage

## Transaction

## Util
