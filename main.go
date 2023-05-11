package main

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"

	multiproof "github.com/crate-crypto/go-ipa"
	"github.com/crate-crypto/go-ipa/bandersnatch/fr"
	"github.com/crate-crypto/go-ipa/banderwagon"
	"github.com/crate-crypto/go-ipa/common"
	"github.com/crate-crypto/go-ipa/ipa"
	"github.com/gballet/go-verkle"
	"golang.org/x/sync/errgroup"
)

const (
	vectorSize = 256
	numRounds  = 10
)

func main() {
	conf := genOrLoadConfig("precomp")

	fmt.Printf("##### VKT proof benchmark #####\n")
	benchVKTProof(conf)

	fmt.Printf("\n##### Raw polynomials multiproof benchmark #####\n")
	benchProvingAndVerification(conf)
}

func benchProvingAndVerification(conf *ipa.IPAConfig) {
	provingStepNames := []string{
		"Generate challenge r and powers ",
		"Calculate t, g(x) and D         ",
		"Calculate h(x) and E            ",
		"Calculate (h-g)(x) and E-D      ",
		"IPA for (h-g)(x) and E-D on t   ",
	}
	provingNumSteps := len(provingStepNames)
	verificationStepNames := []string{
		"Generate challenge r and powers                       ",
		"Calculating helper_scalars r^i/(t-z_i)                ",
		"g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars ",
		"Compute E                                             ",
		"Compute E-D and verify IPA                            ",
	}
	verificationNumSteps := len(verificationStepNames)

	for _, numPolynomials := range []int{1, 1000, 2000, 4000, 8000, 16000} {
		var proofAggrSerialization time.Duration
		var provingAggrTotalTime time.Duration
		provingAggrStepDuration := make([]time.Duration, provingNumSteps)

		var proofAggrDeserialization time.Duration
		var verificationAggrTotalTime time.Duration
		verificationAggrStepDuration := make([]time.Duration, verificationNumSteps)

		for i := 0; i < numRounds; i++ {
			runtime.GC()

			// Setup.
			cs, fs, zs := generateNRandomPolysEvals(conf, numPolynomials)

			// Proving.
			transcriptProving := common.NewTranscript("bench")
			proof := multiproof.CreateMultiProof(transcriptProving, conf, cs, fs, zs)
			timestamps := transcriptProving.GetTimestamps()
			if len(timestamps) != provingNumSteps+1 {
				panic("wrong number of timestamps")
			}
			for k := 1; k < len(timestamps); k++ {
				provingAggrStepDuration[k-1] += timestamps[k].Sub(timestamps[k-1])
			}
			provingAggrTotalTime += timestamps[len(timestamps)-1].Sub(timestamps[0])

			// Verification.
			ys := make([]*fr.Element, len(zs))
			for i, z := range zs {
				ys[i] = &fs[i][z]
			}
			transcriptVerification := common.NewTranscript("bench")
			if ok := multiproof.CheckMultiProof(transcriptVerification, conf, proof, cs, ys, zs); !ok {
				panic("verification failed")
			}
			timestamps = transcriptVerification.GetTimestamps()
			if len(timestamps) != verificationNumSteps+1 {
				panic("wrong number of timestamps")
			}
			for k := 1; k < len(timestamps); k++ {
				verificationAggrStepDuration[k-1] += timestamps[k].Sub(timestamps[k-1])
			}
			verificationAggrTotalTime += timestamps[len(timestamps)-1].Sub(timestamps[0])
			var buf bytes.Buffer
			now := time.Now()
			proof.Write(&buf)
			proofAggrSerialization += time.Since(now)
			now = time.Now()
			proof.Read(&buf)
			proofAggrDeserialization += time.Since(now)
		}
		fmt.Printf("For %d polynomials:\n", numPolynomials)
		fmt.Printf("\tProving:\n")
		fmt.Printf("\t\tProof serialization duration: %.02fms\n", float64((proofAggrSerialization/time.Duration(numRounds)).Microseconds())/1000)
		fmt.Printf("\t\tProof creation duration: %dms\n", (provingAggrTotalTime / time.Duration(numRounds)).Milliseconds())
		fmt.Printf("\t\tDuration per step:\n")
		for i := 0; i < provingNumSteps; i++ {
			fmt.Printf("\t\t\t%s: %.02fms\n", provingStepNames[i], float64((provingAggrStepDuration[i]/time.Duration(numRounds)).Microseconds())/1000)
		}
		fmt.Printf("\tVerification:\n")
		fmt.Printf("\t\tProof deserialization duration: %.02fms\n", float64((proofAggrSerialization/time.Duration(numRounds)).Microseconds())/1000)
		fmt.Printf("\t\tTotal duration: %dms\n", (verificationAggrTotalTime / time.Duration(numRounds)).Milliseconds())
		fmt.Printf("\t\tDuration per step:\n")
		for i := 0; i < verificationNumSteps; i++ {
			fmt.Printf("\t\t\t%s: %.02fms\n", verificationStepNames[i], float64((verificationAggrStepDuration[i]/time.Duration(numRounds)).Microseconds())/1000)
		}
		fmt.Println()
	}
}

func benchVKTProof(conf *ipa.IPAConfig) {
	const numRandKeyValues = 1_000_000
	fmt.Printf("Setup: tree with %d random key-values...\n\n", numRandKeyValues)
	keyValues, tree := genRandomTree(rand.New(rand.NewSource(42)), numRandKeyValues)
	tree.Commit()

	for _, numKeyValues := range []int{1, 1000, 2000, 4000, 8000, 16000} {
		runtime.GC()

		var aggrProofGenTime time.Duration
		var aggrSerializationTime time.Duration
		var aggrDeserializationTime time.Duration
		verkle.BenchCollectPolynomialsDuration = 0
		verkle.BenchNumPolynomials = 0
		for i := 0; i < numRounds; i++ {
			keyvals := map[string][]byte{}
			keys := make([][]byte, 0, numKeyValues)
			for i, kv := range keyValues {
				if i == numKeyValues {
					break
				}
				keys = append(keys, kv.key)
				keyvals[string(kv.key)] = kv.value
			}
			start := time.Now()
			proof, _, _, _, err := verkle.MakeVerkleMultiProof(tree, keys, keyvals)
			if err != nil {
				panic("failed to generate proof")
			}
			aggrProofGenTime += time.Since(start)

			start = time.Now()
			serProof, serStateDiff, err := verkle.SerializeProof(proof)
			if err != nil {
				panic("failed to serialize proof")
			}
			aggrSerializationTime += time.Since(start)

			start = time.Now()
			if _, err := verkle.DeserializeProof(serProof, serStateDiff); err != nil {
				panic("failed to deserialize proof")
			}
			aggrDeserializationTime += time.Since(start)

		}
		fmt.Printf("For %d random-key values of the tree:\n", numKeyValues)
		fmt.Printf("\tGenerating proof took %dms:\n", (aggrProofGenTime / time.Duration(numRounds)).Milliseconds())
		fmt.Printf("\t\tCollected %d polynomials\n", verkle.BenchNumPolynomials/numRounds)
		fmt.Printf("\t\tCollecting those polys (comm, evals, etc) took %dms\n", (verkle.BenchCollectPolynomialsDuration / time.Duration(numRounds)).Milliseconds())
		fmt.Printf("\t\tThe rest (Multiproof + nits) took %dms\n", ((aggrProofGenTime - verkle.BenchCollectPolynomialsDuration) / time.Duration(numRounds)).Milliseconds())
		fmt.Printf("\tSerializing proof took %dms\n", (aggrSerializationTime / time.Duration(numRounds)).Milliseconds())
		fmt.Printf("\tDeserializing proof took %dms\n", (aggrDeserializationTime / time.Duration(numRounds)).Milliseconds())
		fmt.Println()
	}
}

func genRandomTree(rand *rand.Rand, keyValueCount int) ([]keyValue, verkle.VerkleNode) {
	tree := verkle.New()
	keyValues := make([]keyValue, 0, keyValueCount)
	for _, kv := range genRandomKeyValues(rand, keyValueCount) {
		if err := tree.Insert(kv.key, kv.value, nil); err != nil {
			panic(fmt.Sprintf("failed to insert key: %v", err))
		}
		keyValues = append(keyValues, kv)
	}
	return keyValues, tree
}

type keyValue struct {
	key   []byte
	value []byte
}

func genRandomKeyValues(rand *rand.Rand, count int) []keyValue {
	ret := make([]keyValue, count)
	for i := 0; i < count; i++ {
		keyval := make([]byte, 64)
		rand.Read(keyval)
		ret[i].key = keyval[:32]
		ret[i].value = keyval[32:]
	}
	return ret
}

func generateNRandomPolysEvals(conf *ipa.IPAConfig, n int) ([]*banderwagon.Element, [][]fr.Element, []uint8) {
	var lock sync.Mutex
	retCs := make([]*banderwagon.Element, 0, n)
	retEvals := make([][]fr.Element, 0, n)
	retZs := make([]uint8, 0, n)

	batchSize := n/runtime.NumCPU() + 1
	g, _ := errgroup.WithContext(context.Background())
	for i := 0; i < n; i += batchSize {
		num := batchSize
		if i+batchSize > n {
			num = n - i
		}
		g.Go(func() error {
			for i := 0; i < num; i++ {
				frs := make([]fr.Element, vectorSize)
				for j := 0; j < vectorSize; j++ {
					frs[j].SetRandom()
				}
				c := conf.Commit(frs)

				lock.Lock()
				retEvals = append(retEvals, frs)
				retCs = append(retCs, &c)
				retZs = append(retZs, uint8(rand.Uint32()%vectorSize))
				lock.Unlock()
			}
			return nil
		})
	}
	g.Wait()

	return retCs, retEvals, retZs
}

func genOrLoadConfig(fileName string) *ipa.IPAConfig {
	serialized, err := os.ReadFile(fileName)
	if err == nil {
		srs, err := ipa.DeserializeSRSPrecomp(serialized)
		if err != nil {
			panic(err)
		}
		return ipa.NewIPASettingsWithSRSPrecomp(srs)
	}
	conf := ipa.NewIPASettings()
	bts, err := conf.SRSPrecompPoints.SerializeSRSPrecomp()
	if err == nil {
		_ = os.WriteFile(fileName, bts, 0644)
	}
	return conf
}
