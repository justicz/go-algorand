// Copyright (C) 2019 Algorand, Inc.
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

package logic

type stringString struct {
	a string
	b string
}

func stringStringListToMap(they []stringString) map[string]string {
	out := make(map[string]string, len(they))
	for _, v := range they {
		out[v.a] = v.b
	}
	return out
}

// short description of every op
var opDocList = []stringString{
	{"err", "Error. Panic immediately. This is primarily a fencepost against accidental zero bytes getting compiled into programs."},
	{"sha256", "SHA256 hash of value, yields [32]byte"},
	{"keccak256", "Keccac256 hash of value, yields [32]byte"},
	{"sha512_256", "SHA512_256 hash of value, yields [32]byte"},
	{"ed25519verify", "for (data, signature, pubkey) verify the signature of the data against the pubkey"},
	{"+", "A plus B. Panic on overflow."},
	{"-", "A minus B. Panic if B > A."},
	{"/", "A divided by B. Panic if B == 0."},
	{"*", "A times B. Panic on overflow."},
	{"<", "A less than B => {0 or 1}"},
	{">", "A greater than B => {0 or 1}"},
	{"<=", "A less than or equal to B => {0 or 1}"},
	{">=", "A greater than or equal to B => {0 or 1}"},
	{"&&", "A is not zero and B is not zero => {0 or 1}"},
	{"||", "A is not zero or B is not zero => {0 or 1}"},
	{"==", "A is equal to B => {0 or 1}"},
	{"!=", "A is not equal to B => {0 or 1}"},
	{"!", "X == 0 yields 1; else 0"},
	{"len", "yields length of byte value"},
	{"btoi", "converts bytes as big endian to uint64"},
	{"%", "A modulo B. Panic if B == 0."},
	{"|", "A bitwise-or B"},
	{"&", "A bitwise-and B"},
	{"^", "A bitwise-xor B"},
	{"~", "bitwise invert value"},
	{"intcblock", "load block of uint64 constants"},
	{"intc", "push value from uint64 constants to stack by index into constants"},
	{"intc_0", "push uint64 constant 0 to stack"},
	{"intc_1", "push uint64 constant 1 to stack"},
	{"intc_2", "push uint64 constant 2 to stack"},
	{"intc_3", "push uint64 constant 3 to stack"},
	{"bytecblock", "load block of byte-array constants"},
	{"bytec", "push bytes constant to stack by index into constants"},
	{"bytec_0", "push bytes constant 0 to stack"},
	{"bytec_1", "push bytes constant 1 to stack"},
	{"bytec_2", "push bytes constant 2 to stack"},
	{"bytec_3", "push bytes constant 3 to stack"},
	{"arg", "push LogicSig.Args[N] value to stack by index"},
	{"arg_0", "push LogicSig.Args[0] to stack"},
	{"arg_1", "push LogicSig.Args[1] to stack"},
	{"arg_2", "push LogicSig.Args[2] to stack"},
	{"arg_3", "push LogicSig.Args[3] to stack"},
	{"txn", "push field from current transaction to stack"},
	{"global", "push value from globals to stack"},
	{"bnz", "branch if value is not zero"},
	{"pop", "discard value from stack"},
	{"dup", "duplicate last value on stack"},
}

var opDocByName map[string]string

func OpDoc(opName string) string {
	if opDocByName == nil {
		opDocByName = stringStringListToMap(opDocList)
	}
	return opDocByName[opName]
}

// notes on immediate bytes following the opcode
var opcodeImmediateNoteList = []stringString{
	{"intcblock", "{varuint length} [{varuint value}, ...]"},
	{"intc", "{uint8 int constant index}"},
	{"bytecblock", "{varuint length} [({varuint value length} bytes), ...]"},
	{"bytec", "{uint8 byte constant index}"},
	{"arg", "{uint8 arg index N}"},
	{"txn", "{uint8 transaction field index}"},
	{"global", "{uint8 global field index}"},
	{"bnz", "{0..0x7fff forward branch offset, big endian}"},
}
var opcodeImmediateNotes map[string]string

func OpImmediateNote(opName string) string {
	if opcodeImmediateNotes == nil {
		opcodeImmediateNotes = stringStringListToMap(opcodeImmediateNoteList)
	}
	return opcodeImmediateNotes[opName]
}

// further documentation on the function of the opcode
var opDocExtraList = []stringString{
	{"bnz", "for a bnz instruction at `pc`, if the last element of the stack is not zero then branch to instruction at `pc + 3 + N`, else proceed to next instruction at `pc + 3`"},
	{"intcblock", "`intcblock` loads following program bytes into an array of integer constants in the evaluator. These integer constants can be referred to by `intc` and `intc_*` which will push the value onto the stack."},
	{"bytecblock", "`bytecblock` loads the following program bytes into an array of byte string constants in the evaluator. These constants can be referred to by `bytec` and `bytec_*` which will push the value onto the stack."},
	{"*", "It is worth noting that there are 10,000,000,000,000,000 micro-Algos in the total supply, or a bit less than 2^54. When doing rational math, e.g. (A * (N/D)) as ((A * N) / D) one should limit the numerator to less than 2^10 to be completely sure there won't be overflow."},
}

var opDocExtras map[string]string

func OpDocExtra(opName string) string {
	if opDocExtras == nil {
		opDocExtras = stringStringListToMap(opDocExtraList)
	}
	return opDocExtras[opName]
}

// OpGroup is a grouping of ops for documentation purposes.
// e.g. "Arithmetic", ["+", "-", ...]
type OpGroup struct {
	GroupName string
	Ops       []string
}

var OpGroupList = []OpGroup{
	{"Arithmetic", []string{"sha256", "keccak256", "sha512_256", "ed25519verify", "+", "-", "/", "*", "<", ">", "<=", ">=", "&&", "||", "==", "!=", "!", "len", "btoi", "%", "|", "&", "^", "~"}},
	{"Loading Values", []string{"intcblock", "intc", "intc_0", "intc_1", "intc_2", "intc_3", "bytecblock", "bytec", "bytec_0", "bytec_1", "bytec_2", "bytec_3", "arg", "arg_0", "arg_1", "arg_2", "arg_3", "txn", "global"}},
	{"Flow Control", []string{"err", "bnz", "pop", "dup"}},
}
