package main

import (
	"math/rand"
	"testing"

	multiproof "github.com/crate-crypto/go-ipa"
	"github.com/crate-crypto/go-ipa/bandersnatch/fr"
	"github.com/crate-crypto/go-ipa/common"
	"github.com/gballet/go-verkle"
)

func BenchmarkTreeProofGeneration(b *testing.B) {
	const numKeyValues = 10_000
	const numRandKeyValues = 1_000_000
	keyValues, tree := genRandomTree(rand.New(rand.NewSource(42)), numRandKeyValues)
	tree.Commit()

	keyvals := map[string][]byte{}
	keys := make([][]byte, 0, numKeyValues)
	for i, kv := range keyValues {
		if i == numKeyValues {
			break
		}
		keys = append(keys, kv.key)
		keyvals[string(kv.key)] = kv.value
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _, _, err := verkle.MakeVerkleMultiProof(tree, keys, keyvals)
		if err != nil {
			panic("failed to generate proof")
		}
	}
}

func BenchmarkProofCreation(b *testing.B) {
	const numPolynomials = 128_000
	conf := genOrLoadConfig("precomp")
	cs, fs, zs := generateNRandomPolysEvals(conf, numPolynomials)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transcriptProving := common.NewTranscript("bench")
		multiproof.CreateMultiProof(transcriptProving, conf, cs, fs, zs)
	}
}

func BenchmarkProofVerification(b *testing.B) {
	const numPolynomials = 10_000
	conf := genOrLoadConfig("precomp")
	cs, fs, zs := generateNRandomPolysEvals(conf, numPolynomials)

	transcriptProving := common.NewTranscript("bench")
	proof := multiproof.CreateMultiProof(transcriptProving, conf, cs, fs, zs)

	ys := make([]*fr.Element, len(zs))
	for i, z := range zs {
		ys[i] = &fs[i][z]
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transcriptVerification := common.NewTranscript("bench")
		if ok := multiproof.CheckMultiProof(transcriptVerification, conf, proof, cs, ys, zs); !ok {
			panic("verification failed")
		}
	}
}
