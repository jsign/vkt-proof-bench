package main

import (
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
	"golang.org/x/sync/errgroup"
)

const (
	vectorSize = 256
	numRounds  = 10
)

var numPolynomials = []int{1, 625, 1_250, 2_500, 5_000, 10_000}

func main() {
	conf := genOrLoadConfig("precomp")

	bench(conf)
}

func bench(conf *ipa.IPAConfig) {
	provingMilestoneNames := []string{
		"Generate challenge r and powers ",
		"Calculate t, g(x) and D         ",
		"Calculate h(x) and E            ",
		"Calculate (h-g)(x) and E-D      ",
		"IPA for (h-g)(x) and E-D on t   ",
	}
	provingNumMilestones := len(provingMilestoneNames)
	verificationMilestoneNames := []string{
		"Generate challenge r and powers                       ",
		"Calculating helper_scalars r^i/(t-z_i)                ",
		"g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars ",
		"Compute E                                             ",
		"Compute E-D and verify IPA                            ",
	}
	verificationNumMilestones := len(verificationMilestoneNames)

	for _, n := range numPolynomials {
		var provingAggrTotalTime time.Duration
		provingAggrMilestoneDuration := make([]time.Duration, provingNumMilestones)

		var verificationAggrTotalTime time.Duration
		verificationAggrMilestoneDuration := make([]time.Duration, verificationNumMilestones)

		for i := 0; i < numRounds; i++ {
			runtime.GC()

			// Setup.
			cs, fs, zs := generateNRandomPolysEvals(conf, n)

			// Proving.
			transcriptProving := common.NewTranscript("bench")
			proof := multiproof.CreateMultiProof(transcriptProving, conf, cs, fs, zs)
			timestamps := transcriptProving.GetTimestamps()
			if len(timestamps) != provingNumMilestones+1 {
				panic("wrong number of timestamps")
			}
			for k := 1; k < len(timestamps); k++ {
				provingAggrMilestoneDuration[k-1] += timestamps[k].Sub(timestamps[k-1])
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
			if len(timestamps) != verificationNumMilestones+1 {
				panic("wrong number of timestamps")
			}
			for k := 1; k < len(timestamps); k++ {
				verificationAggrMilestoneDuration[k-1] += timestamps[k].Sub(timestamps[k-1])
			}
			verificationAggrTotalTime += timestamps[len(timestamps)-1].Sub(timestamps[0])

		}
		fmt.Printf("For %d polynomials:\n", n)
		fmt.Printf("\tProving:\n")
		fmt.Printf("\t\tAvg. total duration: %dms\n", (provingAggrTotalTime / time.Duration(numRounds)).Milliseconds())
		fmt.Printf("\t\tAvg. time per milestone:\n")
		for i := 0; i < provingNumMilestones; i++ {
			fmt.Printf("\t\t\t%s: %.02fms\n", provingMilestoneNames[i], float64((provingAggrMilestoneDuration[i]/time.Duration(numRounds)).Microseconds())/1000)
		}
		fmt.Printf("\tVerification:\n")
		fmt.Printf("\t\tAvg. total duration: %dms\n", (verificationAggrTotalTime / time.Duration(numRounds)).Milliseconds())
		fmt.Printf("\t\tAvg. time per milestone:\n")
		for i := 0; i < verificationNumMilestones; i++ {
			fmt.Printf("\t\t\t%s: %.02fms\n", verificationMilestoneNames[i], float64((verificationAggrMilestoneDuration[i]/time.Duration(numRounds)).Microseconds())/1000)
		}
		fmt.Println()
	}
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
