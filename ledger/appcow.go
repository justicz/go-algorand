// Copyright (C) 2019-2020 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.

package ledger

import (
	"fmt"

	"github.com/algorand/go-algorand/data/basics"
	"github.com/algorand/go-algorand/data/transactions/logic"
)

func (cb *roundCowState) ensureStorageDelta(addr basics.Address, aidx basics.AppIndex, global bool) (*storageDelta, error) {
	// If we already have a storageDelta, return it
	aapp := addrApp{addr, aidx, global}
	lsd, ok := cb.mods.sdeltas[aapp]
	if ok {
		return lsd, nil
	}

	// Otherwise, create a new one, looking up how much storage we are
	// currently using in order to populate `counts` correctly
	counts, err := cb.getStorageCounts(addr, aidx, global)
	if err != nil {
		return nil, err
	}

	lsd = &storageDelta{
		action: noAction,
		kvCow:  make(basics.StateDelta),
		counts: &counts,
	}

	cb.mods.sdeltas[aapp] = lsd
	return lsd, nil
}

func (cb *roundCowState) getStorageCounts(addr basics.Address, aidx basics.AppIndex, global bool) (basics.StateSchema, error) {
	// If we haven't allocated storage, then our used storage count is zero
	allocated, err := cb.Allocated(addr, aidx, global)
	if err != nil {
		return basics.StateSchema{}, err
	}
	if !allocated {
		return basics.StateSchema{}, nil
	}

	// If we already have a storageDelta, return the counts from it
	aapp := addrApp{addr, aidx, global}
	lsd, ok := cb.mods.sdeltas[aapp]
	if ok {
		return *lsd.counts, nil
	}

	// Otherwise, check our parent
	return cb.lookupParent.getStorageCounts(addr, aidx, global)
}

func (cb *roundCowState) Allocated(addr basics.Address, aidx basics.AppIndex, global bool) (bool, error) {
	// Check if we've allocated or deallocate within this very cow
	aapp := addrApp{addr, aidx, global}
	lsd, ok := cb.mods.sdeltas[aapp]
	if ok {
		if lsd.action == allocAction {
			return true, nil
		} else if lsd.action == deallocAction {
			return false, nil
		}
	}

	// Otherwise, check our parent
	return cb.lookupParent.Allocated(addr, aidx, global)
}

func (cb *roundCowState) Allocate(addr basics.Address, aidx basics.AppIndex, global bool) error {
	// Check that account is not already opted in
	allocated, err := cb.Allocated(addr, aidx, global)
	if err != nil {
		return err
	}
	if allocated {
		err = fmt.Errorf("cannot allocate storage, %v already allocated storage for app %d", addr, aidx)
		return err
	}

	lsd, err := cb.ensureStorageDelta(addr, aidx, global)
	if err != nil {
		return err
	}

	lsd.action = allocAction
	return nil
}

func (cb *roundCowState) Deallocate(addr basics.Address, aidx basics.AppIndex, global bool) error {
	// Check that account has allocated storage
	allocated, err := cb.Allocated(addr, aidx, global)
	if err != nil {
		return err
	}
	if !allocated {
		err = fmt.Errorf("cannot deallocate storage, %v has not allocated storage for app %d", addr, aidx)
		return err
	}

	lsd, err := cb.ensureStorageDelta(addr, aidx, global)
	if err != nil {
		return err
	}

	lsd.action = deallocAction
	lsd.counts = &basics.StateSchema{}
	lsd.kvCow = make(basics.StateDelta)
	return nil
}

func (cb *roundCowState) GetStorage(addr basics.Address, aidx basics.AppIndex, global bool, key string) (basics.TealValue, bool, error) {
	// Check that account has allocated storage
	allocated, err := cb.Allocated(addr, aidx, global)
	if err != nil {
		return basics.TealValue{}, false, err
	}
	if !allocated {
		err = fmt.Errorf("cannot fetch key, %v has not allocated storage for app %d", addr, aidx)
		return basics.TealValue{}, false, err
	}

	// Check if key is in a storage delta, if so return it (the "hasDelta"
	// boolean will be true if the kvCow holds _any_ delta for the key,
	// including if that delta is a "delete" delta)
	lsd, ok := cb.mods.sdeltas[addrApp{addr, aidx, global}]
	if ok {
		delta, hasDelta := lsd.kvCow[key]
		if hasDelta {
			val, ok := delta.ToTealValue()
			return val, ok, nil
		}
	}

	// If this storage delta is noAction, then check our parent.
	// Otherwise, the key does not exist.
	if lsd.action == noAction {
		// Check our parent
		return cb.lookupParent.GetStorage(addr, aidx, global, key)
	}

	return basics.TealValue{}, false, nil
}

func (cb *roundCowState) SetStorage(addr basics.Address, aidx basics.AppIndex, global bool, key string, value basics.TealValue) error {
	// Check that account has allocated storage
	allocated, err := cb.Allocated(addr, aidx, global)
	if err != nil {
		return err
	}
	if !allocated {
		err = fmt.Errorf("cannot set key, %v has not allocated storage for app %d", addr, aidx)
		return err
	}

	// Fetch the old value + presence so we know how to update counts
	oldValue, oldOk, err := cb.GetStorage(addr, aidx, global, key)
	if err != nil {
		return err
	}

	// Write the value delta associated with this key/value
	lsd, err := cb.ensureStorageDelta(addr, aidx, global)
	if err != nil {
		return err
	}
	lsd.kvCow[key] = value.ToValueDelta()

	// Fetch the new value + presence so we know how to update counts
	newValue, newOk, err := cb.GetStorage(addr, aidx, global, key)
	if err != nil {
		return err
	}

	// Update counts
	err = updateCounts(lsd, oldValue, oldOk, newValue, newOk)
	if err != nil {
		return err
	}

	return nil
}

func (cb *roundCowState) DelStorage(addr basics.Address, aidx basics.AppIndex, global bool, key string) error {
	// Check that account has allocated storage
	allocated, err := cb.Allocated(addr, aidx, global)
	if err != nil {
		return err
	}
	if !allocated {
		err = fmt.Errorf("cannot del key, %v not opted in to app %d", addr, aidx)
		return err
	}

	// Fetch the old value + presence so we know how to update counts
	oldValue, oldOk, err := cb.GetStorage(addr, aidx, global, key)
	if err != nil {
		return err
	}

	// Write the value delta associated with deleting this key
	lsd, err := cb.ensureStorageDelta(addr, aidx, global)
	if err != nil {
		return nil
	}

	lsd.kvCow[key] = basics.ValueDelta{
		Action: basics.DeleteAction,
	}

	// Fetch the new value + presence so we know how to update counts
	newValue, newOk, err := cb.GetStorage(addr, aidx, global, key)
	if err != nil {
		return err
	}

	// Update counts
	err = updateCounts(lsd, oldValue, oldOk, newValue, newOk)
	if err != nil {
		return err
	}

	return nil
}

func (cb *roundCowState) StatefulEval(params logic.EvalParams, aidx basics.AppIndex, program []byte) (pass bool, err error) {
	// Make a child cow to eval our program in
	calf := cb.child()
	params.Ledger, err = makeLogicLedger(calf, aidx)
	if err != nil {
		return false, err
	}

	// Eval the program
	pass, err = logic.EvalStateful(program, params)
	if err != nil {
		return false, err
	}

	// If program passed, commit to state changes
	if pass {
		calf.commitToParent()
	}

	return pass, nil
}

func updateCounts(lsd *storageDelta, bv basics.TealValue, bok bool, av basics.TealValue, aok bool) error {
	// If the value existed before, decrement the count of the old type.
	if bok {
		switch bv.Type {
		case basics.TealBytesType:
			lsd.counts.NumByteSlice--
		case basics.TealUintType:
			lsd.counts.NumUint--
		default:
			return fmt.Errorf("unknown before type: %v", bv.Type)
		}
	}

	// If the value exists now, increment the count of the new type.
	if aok {
		switch av.Type {
		case basics.TealBytesType:
			lsd.counts.NumByteSlice++
		case basics.TealUintType:
			lsd.counts.NumUint++
		default:
			return fmt.Errorf("unknown after type: %v", av.Type)
		}
	}
	return nil
}
