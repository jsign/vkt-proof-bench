# vkt-proof-bench
This repository contains benchmarks to measure the performance of proof generation and verification in Verkle Trees.

## Results 
The following are the results run on an AMD Ryzen 7 3800XT processor and `go-ipa@ce4a969`.

```
$ go run main.go
##### VKT proof benchmark #####
Setup: tree with 50000 random key-values...

For 1 random-key values of the tree:
        Generating proof took 52ms
        Serializing proof took 0ms
        Deserializing proof took 0ms

For 1000 random-key values of the tree:
        Generating proof took 368ms
        Serializing proof took 10ms
        Deserializing proof took 21ms

For 2000 random-key values of the tree:
        Generating proof took 633ms
        Serializing proof took 17ms
        Deserializing proof took 39ms

For 4000 random-key values of the tree:
        Generating proof took 1186ms
        Serializing proof took 28ms
        Deserializing proof took 76ms

For 8000 random-key values of the tree:
        Generating proof took 2269ms
        Serializing proof took 49ms
        Deserializing proof took 148ms

For 16000 random-key values of the tree:
        Generating proof took 4394ms
        Serializing proof took 93ms
        Deserializing proof took 291ms


##### Raw polynomials multiproof benchmark #####
For 1 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 52ms
        Duration per milestone:
                        Generate challenge r and powers : 0.02ms
                        Calculate t, g(x) and D         : 3.97ms
                        Calculate h(x) and E            : 3.09ms
                        Calculate (h-g)(x) and E-D      : 0.00ms
                        IPA for (h-g)(x) and E-D on t   : 45.90ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 27ms
                Time per milestone:
                        Generate challenge r and powers                       : 0.01ms
                        Calculating helper_scalars r^i/(t-z_i)                : 0.02ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.00ms
                        Compute E                                             : 0.15ms
                        Compute E-D and verify IPA                            : 27.61ms

For 1000 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 87ms
        Duration per milestone:
                        Generate challenge r and powers : 3.02ms
                        Calculate t, g(x) and D         : 24.99ms
                        Calculate h(x) and E            : 16.05ms
                        Calculate (h-g)(x) and E-D      : 0.00ms
                        IPA for (h-g)(x) and E-D on t   : 43.31ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 106ms
                Time per milestone:
                        Generate challenge r and powers                       : 5.61ms
                        Calculating helper_scalars r^i/(t-z_i)                : 6.66ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.05ms
                        Compute E                                             : 67.53ms
                        Compute E-D and verify IPA                            : 26.25ms

For 2000 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 129ms
        Duration per milestone:
                        Generate challenge r and powers : 5.96ms
                        Calculate t, g(x) and D         : 47.73ms
                        Calculate h(x) and E            : 31.27ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.62ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 183ms
                Time per milestone:
                        Generate challenge r and powers                       : 9.62ms
                        Calculating helper_scalars r^i/(t-z_i)                : 9.94ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.09ms
                        Compute E                                             : 136.98ms
                        Compute E-D and verify IPA                            : 26.58ms

For 4000 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 205ms
        Duration per milestone:
                        Generate challenge r and powers : 11.93ms
                        Calculate t, g(x) and D         : 92.31ms
                        Calculate h(x) and E            : 57.01ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.66ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 329ms
                Time per milestone:
                        Generate challenge r and powers                       : 16.69ms
                        Calculating helper_scalars r^i/(t-z_i)                : 17.87ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.19ms
                        Compute E                                             : 267.91ms
                        Compute E-D and verify IPA                            : 26.50ms

For 8000 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 365ms
        Duration per milestone:
                        Generate challenge r and powers : 24.07ms
                        Calculate t, g(x) and D         : 183.97ms
                        Calculate h(x) and E            : 112.32ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.97ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 627ms
                Time per milestone:
                        Generate challenge r and powers                       : 28.98ms
                        Calculating helper_scalars r^i/(t-z_i)                : 34.84ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.49ms
                        Compute E                                             : 537.09ms
                        Compute E-D and verify IPA                            : 26.18ms

For 16000 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 679ms
        Duration per milestone:
                        Generate challenge r and powers : 47.53ms
                        Calculate t, g(x) and D         : 365.53ms
                        Calculate h(x) and E            : 221.43ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.75ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 1213ms
                Time per milestone:
                        Generate challenge r and powers                       : 51.81ms
                        Calculating helper_scalars r^i/(t-z_i)                : 67.70ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.98ms
                        Compute E                                             : 1067.55ms
                        Compute E-D and verify IPA                            : 25.78ms
```
