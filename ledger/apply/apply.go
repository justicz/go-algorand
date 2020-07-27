package apply

import (
	"github.com/algorand/go-algorand/config"
	"github.com/algorand/go-algorand/data/basics"
	"github.com/algorand/go-algorand/data/transactions/logic"
)

// Balances allow to move MicroAlgos from one address to another and to update balance records, or to access and modify individual balance records
// After a call to Put (or Move), future calls to Get or Move will reflect the updated balance record(s)
type Balances interface {
	// Get looks up the balance record for an address
	// If the account is known to be empty, then err should be nil and the returned balance record should have the given address and empty AccountData
	// withPendingRewards specifies whether pending rewards should be applied.
	// A non-nil error means the lookup is impossible (e.g., if the database doesn't have necessary state anymore)
	Get(addr basics.Address, withPendingRewards bool) (basics.BalanceRecord, error)

	Put(basics.BalanceRecord) error

	// PutWithCreatable is like Put, but should be used when creating or deleting an asset or application.
	PutWithCreatable(record basics.BalanceRecord, newCreatable *basics.CreatableLocator, deletedCreatable *basics.CreatableLocator) error

	// GetCreator gets the address of the account that created a given creatable
	GetCreator(cidx basics.CreatableIndex, ctype basics.CreatableType) (basics.Address, bool, error)

	Allocated(addr basics.Address, aidx basics.AppIndex, global bool) (bool, error)
	Allocate(addr basics.Address, aidx basics.AppIndex, global bool) error
	Deallocate(addr basics.Address, aidx basics.AppIndex, global bool) error
	GetKey(addr basics.Address, aidx basics.AppIndex, global bool, key string) (basics.TealValue, bool, error)
	SetStorage(addr basics.Address, aidx basics.AppIndex, global bool, key string, value basics.TealValue) error
	DelStorage(addr basics.Address, aidx basics.AppIndex, global bool, key string) error
	StatefulEval(params logic.EvalParams, aidx basics.AppIndex, program []byte) (passed bool, evalDelta basics.EvalDelta, err error)

	// Move MicroAlgos from one account to another, doing all necessary overflow checking (convenience method)
	// TODO: Does this need to be part of the balances interface, or can it just be implemented here as a function that calls Put and Get?
	Move(src, dst basics.Address, amount basics.MicroAlgos, srcRewards *basics.MicroAlgos, dstRewards *basics.MicroAlgos) error

	// Balances correspond to a Round, which mean that they also correspond
	// to a ConsensusParams.  This returns those parameters.
	ConsensusParams() config.ConsensusParams
}
