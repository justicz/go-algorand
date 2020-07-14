package apply

import (
	"fmt"

	"github.com/algorand/go-algorand/data/basics"
	"github.com/algorand/go-algorand/data/transactions"
)

// Allocate the map of AppParams if it is nil, and return a copy. We do *not*
// call clone on each AppParams -- callers must do that for any values where
// they intend to modify a contained reference type e.g. the GlobalState.
func cloneAppParams(m map[basics.AppIndex]basics.AppParams) map[basics.AppIndex]basics.AppParams {
	res := make(map[basics.AppIndex]basics.AppParams, len(m))
	for k, v := range m {
		res[k] = v
	}
	return res
}

// getAppParams fetches the creator address and AppParams for the app index,
// if they exist. It does *not* clone the AppParams, so the returned params
// must not be modified directly.
func getAppParams(balances transactions.Balances, aidx basics.AppIndex) (params basics.AppParams, creator basics.Address, exists bool, err error) {
	creator, exists, err = balances.GetCreator(basics.CreatableIndex(aidx), basics.AppCreatable)
	if err != nil {
		return
	}

	// App doesn't exist. Not an error, but return straight away
	if !exists {
		return
	}

	record, err := balances.Get(creator, false)
	if err != nil {
		return
	}

	params, ok := record.AppParams[aidx]
	if !ok {
		// This should never happen. If app exists then we should have
		// found the creator successfully.
		err = fmt.Errorf("app %d not found in account %s", aidx, creator.String())
		return
	}

	return
}

func checkPrograms(ac *transactions.ApplicationCallTxnFields, steva transactions.StateEvaluator, maxCost int) error {
	cost, err := steva.Check(ac.ApprovalProgram)
	if err != nil {
		return fmt.Errorf("check failed on ApprovalProgram: %v", err)
	}

	if cost > maxCost {
		return fmt.Errorf("ApprovalProgram too resource intensive. Cost is %d, max %d", cost, maxCost)
	}

	cost, err = steva.Check(ac.ClearStateProgram)
	if err != nil {
		return fmt.Errorf("check failed on ClearStateProgram: %v", err)
	}

	if cost > maxCost {
		return fmt.Errorf("ClearStateProgram too resource intensive. Cost is %d, max %d", cost, maxCost)
	}

	return nil
}

// createApplication writes a new AppParams entry and returns application ID
func createApplication(ac *transactions.ApplicationCallTxnFields, balances transactions.Balances, creator basics.Address, txnCounter uint64) (appIdx basics.AppIndex, err error) {

	// Fetch the creator's (sender's) balance record
	record, err := balances.Get(creator, false)
	if err != nil {
		return
	}

	// Make sure the creator isn't already at the app creation max
	maxAppsCreated := balances.ConsensusParams().MaxAppsCreated
	if len(record.AppParams) >= maxAppsCreated {
		err = fmt.Errorf("cannot create app for %s: max created apps per acct is %d", creator.String(), maxAppsCreated)
		return
	}

	// Clone app params, so that we have a copy that is safe to modify
	record.AppParams = cloneAppParams(record.AppParams)

	// Allocate the new app params (+ 1 to match Assets Idx namespace)
	appIdx = basics.AppIndex(txnCounter + 1)
	record.AppParams[appIdx] = basics.AppParams{
		ApprovalProgram:   ac.ApprovalProgram,
		ClearStateProgram: ac.ClearStateProgram,
		StateSchemas: basics.StateSchemas{
			LocalStateSchema:  ac.LocalStateSchema,
			GlobalStateSchema: ac.GlobalStateSchema,
		},
	}

	// Update the cached TotalStateSchema for this account, used
	// when computing MinBalance, since the creator has to store
	// the global state
	totalSchema := record.TotalAppSchema
	totalSchema = totalSchema.AddSchema(ac.GlobalStateSchema)
	record.TotalAppSchema = totalSchema

	// Tell the cow what app we created
	created := &basics.CreatableLocator{
		Creator: creator,
		Type:    basics.AppCreatable,
		Index:   basics.CreatableIndex(appIdx),
	}

	// Write back to the creator's balance record and continue
	err = balances.PutWithCreatable(record, created, nil)
	if err != nil {
		return 0, err
	}

	return
}

func deleteApplication(balances Balances, creator basics.Address) error {
	// Deleting the application. Fetch the creator's balance record
	record, err := balances.Get(creator, false)
	if err != nil {
		return err
	}

	// Update the TotalAppSchema used for MinBalance calculation,
	// since the creator no longer has to store the GlobalState
	totalSchema := record.TotalAppSchema
	globalSchema := record.AppParams[appIdx].GlobalStateSchema
	totalSchema = totalSchema.SubSchema(globalSchema)
	record.TotalAppSchema = totalSchema

	// Delete the AppParams
	record.AppParams = cloneAppParams(record.AppParams)
	delete(record.AppParams, appIdx)

	// Tell the cow what app we deleted
	deleted := &basics.CreatableLocator{
		Creator: creator,
		Type:    basics.AppCreatable,
		Index:   basics.CreatableIndex(appIdx),
	}

	// Write back to cow
	return balances.PutWithCreatable(record, nil, deleted)
}

func updateApplication(balances Balances, creator basics.Address) error {
	// Updating the application. Fetch the creator's balance record
	record, err := balances.Get(creator, false)
	if err != nil {
		return err
	}

	// Fill in the new programs
	record.AppParams = cloneAppParams(record.AppParams)
	params := record.AppParams[appIdx]
	params.ApprovalProgram = ac.ApprovalProgram
	params.ClearStateProgram = ac.ClearStateProgram

	record.AppParams[appIdx] = params
	return balances.Put(record)
}

func ApplicationCall(ac *transactions.ApplicationCallTxnFields, header transactions.Header, balances transactions.Balances, ad *transactions.ApplyData, evalParams logic.EvalParams, txnCounter uint64) (err error) {
	defer func() {
		// If we are returning a non-nil error, then don't return a
		// non-empty EvalDelta. Not required for correctness.
		if err != nil && ad != nil {
			ad.EvalDelta = basics.EvalDelta{}
		}
	}()

	// Ensure we are always passed a non-nil ApplyData
	if ad == nil {
		err = fmt.Errorf("ApplyData cannot be nil")
		return
	}

	// Keep track of the application ID we're working on
	appIdx := ac.ApplicationID

	// Specifying an application ID of 0 indicates application creation
	if ac.ApplicationID == 0 {
		appIdx, err = ac.createApplication(balances, header.Sender, txnCounter)
		if err != nil {
			return
		}
	}

	// Fetch the application parameters, if they exist
	params, creator, exists, err := getAppParams(balances, appIdx)
	if err != nil {
		return err
	}

	// Ensure that the only operation we can do is ClearState if the application
	// does not exist
	if !exists && ac.OnCompletion != ClearStateOC {
		return fmt.Errorf("only clearing out is supported for applications that do not exist")
	}

	// If this txn is going to set new programs (either for creation or
	// update), check that the programs are valid and not too expensive
	if ac.ApplicationID == 0 || ac.OnCompletion == UpdateApplicationOC {
		maxCost := balances.ConsensusParams().MaxAppProgramCost
		err = ac.checkPrograms(steva, maxCost)
		if err != nil {
			return err
		}
	}

	// Clear out our LocalState. In this case, we don't execute the
	// ApprovalProgram, since clearing out is always allowed. We only
	// execute the ClearStateProgram, whose failures are ignored.
	if ac.OnCompletion == ClearStateOC {
		evalDelta, err := balances.StatefulEval(evalParams, appIdx, params.ClearStateProgram, false /* err on failure */)
		if err != nil {
			return err
		}

		// Fill in applyData, so that consumers don't have to implement a
		// stateful TEAL interpreter to apply state changes
		ad.EvalDelta = evalDelta
		return balances.CloseOutApp(header.Sender, appIdx)
	}

	// If this is an OptIn transaction, ensure that the sender has
	// LocalState allocated prior to TEAL execution, so that it may be
	// initialized in the same transaction.
	if ac.OnCompletion == OptInOC {
		err = balances.OptInApp(header.Sender, appIdx)
		if err != nil {
			return err
		}
	}

	// Execute the Approval program
	approved, evalDelta, err := balances.StatefulEval(evalParams, appIdx, params.ApprovalProgram, true /* err on failure */)
	if err != nil {
		return err
	}

	if !approved {
		return fmt.Errorf("transaction rejected by ApprovalProgram")
	}

	// Fill in applyData, so that consumers don't have to implement a
	// stateful TEAL interpreter to apply state changes
	ad.EvalDelta = evalDelta

	switch ac.OnCompletion {
	case NoOpOC:
		// Nothing to do

	case OptInOC:
		// Handled above

	case CloseOutOC:
		err = balances.CloseOutApp(header.Sender, appIdx)
		if err != nil {
			return err
		}

	case DeleteApplicationOC:
		err = deleteApplication(&ac, balances, creator)
		if err != nil {
			return err
		}

	case UpdateApplicationOC:
		err = updateApplication(&ac, balances, creator)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("invalid application action")
	}

	return nil
}
