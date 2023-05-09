package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/crate-crypto/go-ipa"
	"github.com/crate-crypto/go-ipa/bandersnatch/fr"
	"github.com/crate-crypto/go-ipa/banderwagon"
	"github.com/crate-crypto/go-ipa/common"
	"github.com/crate-crypto/go-ipa/ipa"
	"golang.org/x/sync/errgroup"
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
	fmt.Printf("### Proving ###\n")
	benchProving()
}

func benchProving() {
	conf := genOrLoadConfig("precomp")

	numMilestones := len(provingMilestonesName)
	numberOfPolys := []int{1, 250, 500, 1_000, 2_000, 4_000, 10_000}
	numRounds := 10

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
	var lock sync.Mutex
	retCs := make([]*banderwagon.Element, 0, n)
	retFrs := make([][]fr.Element, 0, n)
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
				retFrs = append(retFrs, frs)
				retCs = append(retCs, &c)
				retZs = append(retZs, uint8(rand.Uint32()%vectorSize))
				lock.Unlock()
			}
			return nil
		})
	}
	g.Wait()

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
