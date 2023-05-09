package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/crate-crypto/go-ipa"
	"github.com/crate-crypto/go-ipa/bandersnatch/fr"
	"github.com/crate-crypto/go-ipa/banderwagon"
	"github.com/crate-crypto/go-ipa/common"
	"github.com/crate-crypto/go-ipa/ipa"
)

const (
	vectorSize = 256
)

var provingMilestonesName = []string{
	"Generate challenge r and powers ",
	"Calculate g(x) and D            ",
	"Calculate h(x) and E            ",
	"Calculate (h-g)(x) and E-D      ",
	"IPA for (h-g)(x) and E-D in r   ",
}

func main() {
	benchProving()
}

func benchProving() {
	conf := genOrLoadConfig("precomp")

	numMilestones := len(provingMilestonesName)
	numberOfPolys := []int{1, 250, 500, 1_000, 2_000, 4_000, 10_000}
	numRounds := 5

	for _, n := range numberOfPolys {
		var aggrTotalTime time.Duration
		aggrMilestoneDuration := make([]time.Duration, numMilestones)
		for i := 0; i < numRounds; i++ {
			runtime.GC()
			cs, fs, zs := generateNRandomPolysEvals(conf, n)

			transcript := common.NewTranscript("bench_proving")
			multiproof.CreateMultiProof(transcript, conf, cs, fs, zs)

			timestamps := transcript.GetTimestamps()
			if len(timestamps) != numMilestones+1 {
				panic("wrong number of timestamps")
			}
			for k := 1; k < len(timestamps); k++ {
				aggrMilestoneDuration[k-1] += timestamps[k].Sub(timestamps[k-1])
			}
			aggrTotalTime += timestamps[len(timestamps)-1].Sub(timestamps[0])
		}
		fmt.Printf("For %d polynomials:\n", n)
		fmt.Printf("\tAvg. total running time: %dms\n", (aggrTotalTime / time.Duration(numRounds)).Milliseconds())
		fmt.Printf("\tAvg. time per milestone:\n")
		for i := 0; i < numMilestones; i++ {
			fmt.Printf("\t\t%s: %dus\n", provingMilestonesName[i], (aggrMilestoneDuration[i] / time.Duration(numRounds)).Microseconds())
		}
		fmt.Println()
	}
}

func generateNRandomPolysEvals(conf *ipa.IPAConfig, n int) ([]*banderwagon.Element, [][]fr.Element, []uint8) {
	retCs := make([]*banderwagon.Element, n)
	retFrs := make([][]fr.Element, n)
	retZs := make([]uint8, n)

	for i := 0; i < n; i++ {
		retFrs[i] = make([]fr.Element, vectorSize)
		for j := 0; j < vectorSize; j++ {
			retFrs[i][j].SetRandom()
		}
		c := conf.Commit(retFrs[i])
		retCs[i] = &c
		retZs[i] = uint8(rand.Uint32() % vectorSize)
	}

	return retCs, retFrs, retZs
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
