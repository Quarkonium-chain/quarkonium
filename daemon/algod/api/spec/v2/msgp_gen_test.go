//go:build !skip_msgp_testing
// +build !skip_msgp_testing

package v2

// Code generated by github.com/algorand/msgp DO NOT EDIT.

import (
	"testing"

	"github.com/algorand/msgp/msgp"

	"github.com/Quarkonium-chain/go-quarkonium/protocol"
	"github.com/Quarkonium-chain/go-quarkonium/test/partitiontest"
)

func TestMarshalUnmarshalAccountApplicationModel(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := AccountApplicationModel{}
	bts := v.MarshalMsg(nil)
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

func TestRandomizedEncodingAccountApplicationModel(t *testing.T) {
	protocol.RunEncodingTest(t, &AccountApplicationModel{})
}

func BenchmarkMarshalMsgAccountApplicationModel(b *testing.B) {
	v := AccountApplicationModel{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgAccountApplicationModel(b *testing.B) {
	v := AccountApplicationModel{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalAccountApplicationModel(b *testing.B) {
	v := AccountApplicationModel{}
	bts := v.MarshalMsg(nil)
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

func TestMarshalUnmarshalAccountAssetModel(t *testing.T) {
	partitiontest.PartitionTest(t)
	v := AccountAssetModel{}
	bts := v.MarshalMsg(nil)
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

func TestRandomizedEncodingAccountAssetModel(t *testing.T) {
	protocol.RunEncodingTest(t, &AccountAssetModel{})
}

func BenchmarkMarshalMsgAccountAssetModel(b *testing.B) {
	v := AccountAssetModel{}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.MarshalMsg(nil)
	}
}

func BenchmarkAppendMsgAccountAssetModel(b *testing.B) {
	v := AccountAssetModel{}
	bts := make([]byte, 0, v.Msgsize())
	bts = v.MarshalMsg(bts[0:0])
	b.SetBytes(int64(len(bts)))
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bts = v.MarshalMsg(bts[0:0])
	}
}

func BenchmarkUnmarshalAccountAssetModel(b *testing.B) {
	v := AccountAssetModel{}
	bts := v.MarshalMsg(nil)
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
