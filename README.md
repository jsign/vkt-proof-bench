# vkt-proof-bench
This repository contains benchmarks to measure the performance of proof generation and verification in Verkle Trees.

## Results 
The following are the results run on a AMD Ryzen 7 3800XT processor and `go-ipa@ce4a969`.

```
$ go run main.go
### Proving ###
For 1 polynomials:
        Avg. total running time: 51ms
        Avg. time per milestone:
                Generate challenge r and powers : 16us
                Calculate g(x) and D            : 3641us
                Calculate h(x) and E            : 3025us
                Calculate (h-g)(x) and E-D      : 4us
                IPA for (h-g)(x) and E-D in r   : 44550us

For 250 polynomials:
        Avg. total running time: 58ms
        Avg. time per milestone:
                Generate challenge r and powers : 766us
                Calculate g(x) and D            : 7965us
                Calculate h(x) and E            : 5826us
                Calculate (h-g)(x) and E-D      : 2us
                IPA for (h-g)(x) and E-D in r   : 44430us

For 500 polynomials:
        Avg. total running time: 67ms
        Avg. time per milestone:
                Generate challenge r and powers : 1528us
                Calculate g(x) and D            : 13740us
                Calculate h(x) and E            : 9351us
                Calculate (h-g)(x) and E-D      : 3us
                IPA for (h-g)(x) and E-D in r   : 43048us

For 1000 polynomials:
        Avg. total running time: 88ms
        Avg. time per milestone:
                Generate challenge r and powers : 3021us
                Calculate g(x) and D            : 25016us
                Calculate h(x) and E            : 16132us
                Calculate (h-g)(x) and E-D      : 4us
                IPA for (h-g)(x) and E-D in r   : 43876us

For 2000 polynomials:
        Avg. total running time: 127ms
        Avg. time per milestone:
                Generate challenge r and powers : 6145us
                Calculate g(x) and D            : 47424us
                Calculate h(x) and E            : 29925us
                Calculate (h-g)(x) and E-D      : 6us
                IPA for (h-g)(x) and E-D in r   : 43873us

For 4000 polynomials:
        Avg. total running time: 205ms
        Avg. time per milestone:
                Generate challenge r and powers : 11925us
                Calculate g(x) and D            : 92073us
                Calculate h(x) and E            : 57069us
                Calculate (h-g)(x) and E-D      : 6us
                IPA for (h-g)(x) and E-D in r   : 44617us

For 10000 polynomials:
        Avg. total running time: 431ms
        Avg. time per milestone:
                Generate challenge r and powers : 29234us
                Calculate g(x) and D            : 222347us
                Calculate h(x) and E            : 135143us
                Calculate (h-g)(x) and E-D      : 6us
                IPA for (h-g)(x) and E-D in r   : 44909us
```
