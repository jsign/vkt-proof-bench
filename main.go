package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/crate-crypto/go-ipa"
	"github.com/crate-crypto/go-ipa/bandersnatch/fr"
	"github.com/crate-crypto/go-ipa/banderwagon"
	"github.com/crate-crypto/go-ipa/common"
	"github.com/crate-crypto/go-ipa/ipa"
)

const vectorSize = 256
const numMilestones = 6

func main() {
	benchProving()
}

func benchProving() {
	conf := ipa.NewIPASettings()

	numberOfPolys := []int{1, 10, 100, 1_000, 10_000}
	numRounds := 10

	for _, n := range numberOfPolys {
		var aggrTotalTime time.Duration
		aggrMilestoneDuration := make([]time.Duration, numMilestones)
		for i := 0; i < numRounds; i++ {
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
		fmt.Printf("\tAvg. total running time: %v\n", aggrTotalTime/time.Duration(numRounds))
		fmt.Printf("\tAvg. time per milestone:\n")
		for i := 0; i < numMilestones; i++ {
			fmt.Printf("\t\tMilestone %d: %v\n", i+1, aggrMilestoneDuration[i]/time.Duration(numRounds))
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
