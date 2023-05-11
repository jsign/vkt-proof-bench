# vkt-proof-bench
This repository contains benchmarks to measure the performance of proof generation and verification in Verkle Trees.

## Results 
The following are the results run on an AMD Ryzen 7 3800XT processor and `go-ipa@ce4a969`.

```
$ go run main.go
##### VKT proof benchmark #####
Setup: tree with 1000000 random key-values...

For 1 random-key values of the tree:
        Generating proof took 52ms
        Serializing proof took 0ms
        Deserializing proof took 0ms

For 1000 random-key values of the tree:
        Generating proof took 426ms
        Serializing proof took 12ms
        Deserializing proof took 24ms

For 2000 random-key values of the tree:
        Generating proof took 753ms
        Serializing proof took 19ms
        Deserializing proof took 46ms

For 4000 random-key values of the tree:
        Generating proof took 1449ms
        Serializing proof took 32ms
        Deserializing proof took 90ms

For 8000 random-key values of the tree:
        Generating proof took 2941ms
        Serializing proof took 63ms
        Deserializing proof took 189ms

For 16000 random-key values of the tree:
        Generating proof took 10679ms
        Serializing proof took 135ms
        Deserializing proof took 388ms


##### Raw polynomials multiproof benchmark #####
For 1 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 173ms
        Duration per milestone:
                        Generate challenge r and powers : 0.02ms
                        Calculate t, g(x) and D         : 69.26ms
                        Calculate h(x) and E            : 56.75ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 47.21ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 28ms
                Time per milestone:
                        Generate challenge r and powers                       : 0.01ms
                        Calculating helper_scalars r^i/(t-z_i)                : 0.02ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.00ms
                        Compute E                                             : 0.15ms
                        Compute E-D and verify IPA                            : 27.99ms

For 1000 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 90ms
        Duration per milestone:
                        Generate challenge r and powers : 3.21ms
                        Calculate t, g(x) and D         : 26.18ms
                        Calculate h(x) and E            : 16.51ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.71ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 108ms
                Time per milestone:
                        Generate challenge r and powers                       : 5.75ms
                        Calculating helper_scalars r^i/(t-z_i)                : 6.45ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.04ms
                        Compute E                                             : 68.80ms
                        Compute E-D and verify IPA                            : 27.38ms

For 2000 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 130ms
        Duration per milestone:
                        Generate challenge r and powers : 6.87ms
                        Calculate t, g(x) and D         : 48.03ms
                        Calculate h(x) and E            : 30.11ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 45.08ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 180ms
                Time per milestone:
                        Generate challenge r and powers                       : 9.93ms
                        Calculating helper_scalars r^i/(t-z_i)                : 10.05ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.07ms
                        Compute E                                             : 133.47ms
                        Compute E-D and verify IPA                            : 26.71ms

For 4000 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 209ms
        Duration per milestone:
                        Generate challenge r and powers : 12.03ms
                        Calculate t, g(x) and D         : 94.38ms
                        Calculate h(x) and E            : 58.54ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 45.02ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 334ms
                Time per milestone:
                        Generate challenge r and powers                       : 17.05ms
                        Calculating helper_scalars r^i/(t-z_i)                : 18.12ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.27ms
                        Compute E                                             : 272.06ms
                        Compute E-D and verify IPA                            : 26.87ms

For 8000 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 366ms
        Duration per milestone:
                        Generate challenge r and powers : 23.94ms
                        Calculate t, g(x) and D         : 184.90ms
                        Calculate h(x) and E            : 112.81ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.43ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 629ms
                Time per milestone:
                        Generate challenge r and powers                       : 28.66ms
                        Calculating helper_scalars r^i/(t-z_i)                : 34.21ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.55ms
                        Compute E                                             : 539.79ms
                        Compute E-D and verify IPA                            : 26.21ms

For 16000 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 675ms
        Duration per milestone:
                        Generate challenge r and powers : 48.25ms
                        Calculate t, g(x) and D         : 362.54ms
                        Calculate h(x) and E            : 219.71ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.57ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 1214ms
                Time per milestone:
                        Generate challenge r and powers                       : 51.33ms
                        Calculating helper_scalars r^i/(t-z_i)                : 68.23ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 1.28ms
                        Compute E                                             : 1067.57ms
                        Compute E-D and verify IPA                            : 25.77ms
```
