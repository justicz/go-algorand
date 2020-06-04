// Package generated provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package generated

import (
	"time"
)

// Account defines model for Account.
type Account struct {

	// the account public key
	Address string `json:"address"`

	// \[algo\] total number of MicroAlgos in the account
	Amount uint64 `json:"amount"`

	// specifies the amount of MicroAlgos in the account, without the pending rewards.
	AmountWithoutPendingRewards uint64 `json:"amount-without-pending-rewards"`

	// \[appl\] applications local data stored in this account.
	//
	// Note the raw object uses `map[int] -> AppLocalState` for this type.
	AppsLocalState *[]ApplicationLocalStates `json:"apps-local-state,omitempty"`

	// Specifies maximums on the number of each type that may be stored.
	AppsTotalSchema *ApplicationStateSchema `json:"apps-total-schema,omitempty"`

	// \[asset\] assets held by this account.
	//
	// Note the raw object uses `map[int] -> AssetHolding` for this type.
	Assets *[]AssetHolding `json:"assets,omitempty"`

	// \[appp\] parameters of applications created by this account including app global data.
	//
	// Note: the raw account uses `map[int] -> AppParams` for this type.
	CreatedApps *[]Application `json:"created-apps,omitempty"`

	// \[apar\] parameters of assets created by this account.
	//
	// Note: the raw account uses `map[int] -> Asset` for this type.
	CreatedAssets *[]Asset `json:"created-assets,omitempty"`

	// AccountParticipation describes the parameters used by this account in consensus protocol.
	Participation *AccountParticipation `json:"participation,omitempty"`

	// amount of MicroAlgos of pending rewards in this account.
	PendingRewards uint64 `json:"pending-rewards"`

	// \[ebase\] used as part of the rewards computation. Only applicable to accounts which are participating.
	RewardBase *uint64 `json:"reward-base,omitempty"`

	// \[ern\] total rewards of MicroAlgos the account has received, including pending rewards.
	Rewards uint64 `json:"rewards"`

	// The round for which this information is relevant.
	Round uint64 `json:"round"`

	// Indicates what type of signature is used by this account, must be one of:
	// * sig
	// * msig
	// * lsig
	SigType *string `json:"sig-type,omitempty"`

	// The address against which signatures/multisigs/logicsigs should be checked
	SpendingKey *string `json:"spending-key,omitempty"`

	// \[onl\] delegation status of the account's MicroAlgos
	// * Offline - indicates that the associated account is delegated.
	// *  Online  - indicates that the associated account used as part of the delegation pool.
	// *   NotParticipating - indicates that the associated account is neither a delegator nor a delegate.
	Status string `json:"status"`
}

// AccountParticipation defines model for AccountParticipation.
type AccountParticipation struct {

	// \[sel\] Selection public key (if any) currently registered for this round.
	SelectionParticipationKey []byte `json:"selection-participation-key"`

	// \[voteFst\] First round for which this participation is valid.
	VoteFirstValid uint64 `json:"vote-first-valid"`

	// \[voteKD\] Number of subkeys in each batch of participation keys.
	VoteKeyDilution uint64 `json:"vote-key-dilution"`

	// \[voteLst\] Last round for which this participation is valid.
	VoteLastValid uint64 `json:"vote-last-valid"`

	// \[vote\] root participation public key (if any) currently registered for this round.
	VoteParticipationKey []byte `json:"vote-participation-key"`
}

// Application defines model for Application.
type Application struct {

	// \[appidx\] application index.
	AppIndex uint64 `json:"app-index"`

	// Stores the global information associated with an application.
	AppParams ApplicationParams `json:"app-params"`
}

// ApplicationLocalState defines model for ApplicationLocalState.
type ApplicationLocalState struct {

	// Represents a key-value store for use in an application.
	KeyValue TealKeyValueStore `json:"key-value"`

	// Specifies maximums on the number of each type that may be stored.
	Schema ApplicationStateSchema `json:"schema"`
}

// ApplicationLocalStates defines model for ApplicationLocalStates.
type ApplicationLocalStates struct {
	AppIndex uint64 `json:"app-index"`

	// Stores local state associated with an application.
	State ApplicationLocalState `json:"state"`
}

// ApplicationParams defines model for ApplicationParams.
type ApplicationParams struct {

	// \[approv\] approval program.
	ApprovalProgram []byte `json:"approval-program"`

	// \[clearp\] approval program.
	ClearStateProgram []byte `json:"clear-state-program"`

	// Represents a key-value store for use in an application.
	GlobalState *TealKeyValueStore `json:"global-state,omitempty"`

	// Specifies maximums on the number of each type that may be stored.
	GlobalStateSchema *ApplicationStateSchema `json:"global-state-schema,omitempty"`

	// Specifies maximums on the number of each type that may be stored.
	LocalStateSchema *ApplicationStateSchema `json:"local-state-schema,omitempty"`
}

// ApplicationStateSchema defines model for ApplicationStateSchema.
type ApplicationStateSchema struct {

	// \[nbs\] num of byte slices.
	NumByteSlice uint64 `json:"num-byte-slice"`

	// \[nui\] num of uints.
	NumUint uint64 `json:"num-uint"`
}

// Asset defines model for Asset.
type Asset struct {

	// unique asset identifier
	Index uint64 `json:"index"`

	// AssetParams specifies the parameters for an asset.
	//
	// \[apar\] when part of an AssetConfig transaction.
	//
	// Definition:
	// data/transactions/asset.go : AssetParams
	Params AssetParams `json:"params"`
}

// AssetHolding defines model for AssetHolding.
type AssetHolding struct {

	// \[a\] number of units held.
	Amount uint64 `json:"amount"`

	// Asset ID of the holding.
	AssetId uint64 `json:"asset-id"`

	// Address that created this asset. This is the address where the parameters for this asset can be found, and also the address where unwanted asset units can be sent in the worst case.
	Creator string `json:"creator"`

	// \[f\] whether or not the holding is frozen.
	IsFrozen bool `json:"is-frozen"`
}

// AssetParams defines model for AssetParams.
type AssetParams struct {

	// \[c\] Address of account used to clawback holdings of this asset.  If empty, clawback is not permitted.
	Clawback *string `json:"clawback,omitempty"`

	// The address that created this asset. This is the address where the parameters for this asset can be found, and also the address where unwanted asset units can be sent in the worst case.
	Creator string `json:"creator"`

	// \[dc\] The number of digits to use after the decimal point when displaying this asset. If 0, the asset is not divisible. If 1, the base unit of the asset is in tenths. If 2, the base unit of the asset is in hundredths, and so on. This value must be between 0 and 19 (inclusive).
	Decimals uint64 `json:"decimals"`

	// \[df\] Whether holdings of this asset are frozen by default.
	DefaultFrozen *bool `json:"default-frozen,omitempty"`

	// \[f\] Address of account used to freeze holdings of this asset.  If empty, freezing is not permitted.
	Freeze *string `json:"freeze,omitempty"`

	// \[m\] Address of account used to manage the keys of this asset and to destroy it.
	Manager *string `json:"manager,omitempty"`

	// \[am\] A commitment to some unspecified asset metadata. The format of this metadata is up to the application.
	MetadataHash *[]byte `json:"metadata-hash,omitempty"`

	// \[an\] Name of this asset, as supplied by the creator.
	Name *string `json:"name,omitempty"`

	// \[r\] Address of account holding reserve (non-minted) units of this asset.
	Reserve *string `json:"reserve,omitempty"`

	// \[t\] The total number of units of this asset.
	Total uint64 `json:"total"`

	// \[un\] Name of a unit of this asset, as supplied by the creator.
	UnitName *string `json:"unit-name,omitempty"`

	// \[au\] URL where more information about the asset can be retrieved.
	Url *string `json:"url,omitempty"`
}

// ErrorResponse defines model for ErrorResponse.
type ErrorResponse struct {
	Data    *string `json:"data,omitempty"`
	Message string  `json:"message"`
}

// TealKeyValue defines model for TealKeyValue.
type TealKeyValue struct {
	Key string `json:"key"`

	// Represents a TEAL value.
	Value TealValue `json:"value"`
}

// TealKeyValueStore defines model for TealKeyValueStore.
type TealKeyValueStore []TealKeyValue

// TealValue defines model for TealValue.
type TealValue struct {

	// \[tb\] bytes value.
	Bytes string `json:"bytes"`

	// \[tt\] value type.
	Type uint64 `json:"type"`

	// \[ui\] uint value.
	Uint uint64 `json:"uint"`
}

// Version defines model for Version.
type Version struct {

	// the current algod build version information.
	Build       VersionBuild `json:"build"`
	GenesisHash []byte       `json:"genesis-hash"`
	GenesisId   string       `json:"genesis-id"`
	Versions    []string     `json:"versions"`
}

// VersionBuild defines model for VersionBuild.
type VersionBuild struct {
	Branch      string `json:"branch"`
	BuildNumber uint64 `json:"build-number"`
	Channel     string `json:"channel"`
	CommitHash  []byte `json:"commit-hash"`
	Major       uint64 `json:"major"`
	Minor       uint64 `json:"minor"`
}

// AccountId defines model for account-id.
type AccountId string

// Address defines model for address.
type Address string

// AddressRole defines model for address-role.
type AddressRole string

// AfterTime defines model for after-time.
type AfterTime time.Time

// AssetId defines model for asset-id.
type AssetId uint64

// BeforeTime defines model for before-time.
type BeforeTime time.Time

// Catchpoint defines model for catchpoint.
type Catchpoint string

// CurrencyGreaterThan defines model for currency-greater-than.
type CurrencyGreaterThan uint64

// CurrencyLessThan defines model for currency-less-than.
type CurrencyLessThan uint64

// ExcludeCloseTo defines model for exclude-close-to.
type ExcludeCloseTo bool

// Format defines model for format.
type Format string

// Limit defines model for limit.
type Limit uint64

// Max defines model for max.
type Max uint64

// MaxRound defines model for max-round.
type MaxRound uint64

// MinRound defines model for min-round.
type MinRound uint64

// Next defines model for next.
type Next string

// NotePrefix defines model for note-prefix.
type NotePrefix string

// Round defines model for round.
type Round uint64

// RoundNumber defines model for round-number.
type RoundNumber uint64

// SigType defines model for sig-type.
type SigType string

// TxId defines model for tx-id.
type TxId string

// TxType defines model for tx-type.
type TxType string

// AccountResponse defines model for AccountResponse.
type AccountResponse Account

// BlockResponse defines model for BlockResponse.
type BlockResponse struct {

	// Block header data.
	Block map[string]interface{} `json:"block"`

	// Optional certificate object. This is only included when the format is set to message pack.
	Cert *map[string]interface{} `json:"cert,omitempty"`
}

// CatchpointAbortResponse defines model for CatchpointAbortResponse.
type CatchpointAbortResponse struct {

	// Catchup abort response string
	CatchupMessage string `json:"catchup-message"`
}

// CatchpointStartResponse defines model for CatchpointStartResponse.
type CatchpointStartResponse struct {

	// Catchup start response string
	CatchupMessage string `json:"catchup-message"`
}

// NodeStatusResponse defines model for NodeStatusResponse.
type NodeStatusResponse struct {

	// The current catchpoint that is being caught up to
	Catchpoint *string `json:"catchpoint,omitempty"`

	// The number of blocks that have already been obtained by the node as part of the catchup
	CatchpointAcquiredBlocks *uint64 `json:"catchpoint-acquired-blocks,omitempty"`

	// The number of account from the current catchpoint that have been processed so far as part of the catchup
	CatchpointProcessedAccounts *uint64 `json:"catchpoint-processed-accounts,omitempty"`

	// The total number of accounts included in the current catchpoint
	CatchpointTotalAccounts *uint64 `json:"catchpoint-total-accounts,omitempty"`

	// The total number of blocks that are required to complete the current catchpoint catchup
	CatchpointTotalBlocks *uint64 `json:"catchpoint-total-blocks,omitempty"`

	// CatchupTime in nanoseconds
	CatchupTime uint64 `json:"catchup-time"`

	// The last catchpoint seen by the node
	LastCatchpoint *string `json:"last-catchpoint,omitempty"`

	// LastRound indicates the last round seen
	LastRound uint64 `json:"last-round"`

	// LastVersion indicates the last consensus version supported
	LastVersion string `json:"last-version"`

	// NextVersion of consensus protocol to use
	NextVersion string `json:"next-version"`

	// NextVersionRound is the round at which the next consensus version will apply
	NextVersionRound uint64 `json:"next-version-round"`

	// NextVersionSupported indicates whether the next consensus version is supported by this node
	NextVersionSupported bool `json:"next-version-supported"`

	// StoppedAtUnsupportedRound indicates that the node does not support the new rounds and has stopped making progress
	StoppedAtUnsupportedRound bool `json:"stopped-at-unsupported-round"`

	// TimeSinceLastRound in nanoseconds
	TimeSinceLastRound uint64 `json:"time-since-last-round"`
}

// PendingTransactionResponse defines model for PendingTransactionResponse.
type PendingTransactionResponse struct {

	// The asset index if the transaction was found and it created an asset.
	AssetIndex *uint64 `json:"asset-index,omitempty"`

	// Rewards in microalgos applied to the close remainder to account.
	CloseRewards *uint64 `json:"close-rewards,omitempty"`

	// Closing amount for the transaction.
	ClosingAmount *uint64 `json:"closing-amount,omitempty"`

	// The round where this transaction was confirmed, if present.
	ConfirmedRound *uint64 `json:"confirmed-round,omitempty"`

	// Indicates that the transaction was kicked out of this node's transaction pool (and specifies why that happened).  An empty string indicates the transaction wasn't kicked out of this node's txpool due to an error.
	PoolError string `json:"pool-error"`

	// Rewards in microalgos applied to the receiver account.
	ReceiverRewards *uint64 `json:"receiver-rewards,omitempty"`

	// Rewards in microalgos applied to the sender account.
	SenderRewards *uint64 `json:"sender-rewards,omitempty"`

	// The raw signed transaction.
	Txn map[string]interface{} `json:"txn"`
}

// PendingTransactionsResponse defines model for PendingTransactionsResponse.
type PendingTransactionsResponse struct {

	// An array of signed transaction objects.
	TopTransactions []map[string]interface{} `json:"top-transactions"`

	// Total number of transactions in the pool.
	TotalTransactions uint64 `json:"total-transactions"`
}

// PostCompileResponse defines model for PostCompileResponse.
type PostCompileResponse struct {

	// base32 SHA512_256 of program bytes (Address style)
	Hash string `json:"hash"`

	// base64 encoded program bytes
	Result string `json:"result"`
}

// PostTransactionsResponse defines model for PostTransactionsResponse.
type PostTransactionsResponse struct {

	// encoding of the transaction hash.
	TxId string `json:"txId"`
}

// SupplyResponse defines model for SupplyResponse.
type SupplyResponse struct {

	// Round
	CurrentRound uint64 `json:"current_round"`

	// OnlineMoney
	OnlineMoney uint64 `json:"online-money"`

	// TotalMoney
	TotalMoney uint64 `json:"total-money"`
}

// TransactionParametersResponse defines model for TransactionParametersResponse.
type TransactionParametersResponse struct {

	// ConsensusVersion indicates the consensus protocol version
	// as of LastRound.
	ConsensusVersion string `json:"consensus-version"`

	// Fee is the suggested transaction fee
	// Fee is in units of micro-Algos per byte.
	// Fee may fall to zero but transactions must still have a fee of
	// at least MinTxnFee for the current network protocol.
	Fee uint64 `json:"fee"`

	// GenesisHash is the hash of the genesis block.
	GenesisHash []byte `json:"genesis-hash"`

	// GenesisID is an ID listed in the genesis block.
	GenesisId string `json:"genesis-id"`

	// LastRound indicates the last round seen
	LastRound uint64 `json:"last-round"`

	// The minimum transaction fee (not per byte) required for the
	// txn to validate for the current network protocol.
	MinFee uint64 `json:"min-fee"`
}

// AccountInformationParams defines parameters for AccountInformation.
type AccountInformationParams struct {

	// Configures whether the response object is JSON or MessagePack encoded.
	Format *string `json:"format,omitempty"`
}

// GetPendingTransactionsByAddressParams defines parameters for GetPendingTransactionsByAddress.
type GetPendingTransactionsByAddressParams struct {

	// Truncated number of transactions to display. If max=0, returns all pending txns.
	Max *uint64 `json:"max,omitempty"`

	// Configures whether the response object is JSON or MessagePack encoded.
	Format *string `json:"format,omitempty"`
}

// GetBlockParams defines parameters for GetBlock.
type GetBlockParams struct {

	// Configures whether the response object is JSON or MessagePack encoded.
	Format *string `json:"format,omitempty"`
}

// TransactionDryRunJSONBody defines parameters for TransactionDryRun.
type TransactionDryRunJSONBody string

// GetPendingTransactionsParams defines parameters for GetPendingTransactions.
type GetPendingTransactionsParams struct {

	// Truncated number of transactions to display. If max=0, returns all pending txns.
	Max *uint64 `json:"max,omitempty"`

	// Configures whether the response object is JSON or MessagePack encoded.
	Format *string `json:"format,omitempty"`
}

// PendingTransactionInformationParams defines parameters for PendingTransactionInformation.
type PendingTransactionInformationParams struct {

	// Configures whether the response object is JSON or MessagePack encoded.
	Format *string `json:"format,omitempty"`
}

// TransactionDryRunRequestBody defines body for TransactionDryRun for application/json ContentType.
type TransactionDryRunJSONRequestBody TransactionDryRunJSONBody
