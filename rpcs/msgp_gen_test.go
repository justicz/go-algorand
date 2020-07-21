// +build !skip_msgp_testing

package rpcs

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"testing"

	"github.com/algorand/msgp/msgp"
)

func TestMarshalUnmarshalEncodedBlockCert(t *testing.T) {
	v := EncodedBlockCert{}
	bts, err := v.MarshalMsg(nil)
	if err != nil {
		t.Fatal(err)
	}
	left, err := v.UnmarshalMsg(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after UnmarshalMsg(): %q", len(left), left)
	}

	left, err = msgp.Skip(bts)
	if err != nil {
		t.Fatal(err)
	}
	if len(left) > 0 {
		t.Errorf("%d bytes left over after Skip(): %q", len(left), left)
	}
}

func TestRandomizedEncodingEncodedBlockCert(t *testing.T) {
	protocol.RunEncodingTest(t, &EncodedBlockCert{})
}

func BenchmarkMarshalMsgEncodedBlockCert(b *testing.B) {
	v := EncodedBlockCert{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgEncodedBlockCert(b *testing.B) {
	v := EncodedBlockCert{}
	bts := make([]byte, 0, v.Msgsize())
	bts, _ = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts, _ = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalEncodedBlockCert(b *testing.B) {
	v := EncodedBlockCert{}
	bts, _ := v.MarshalMsg(nil)
	b.ReportAllocs()
	b.SetBytes(int64(len(bts)))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := v.UnmarshalMsg(bts)
		if err != nil {
			b.Fatal(err)
		}
	}
}
