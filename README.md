# vkt-proof-bench
This repository contains benchmarks to measure the performance of proof generation and verification in Verkle Trees.

## Run

You can run this benchmark by just doing: `go run main.go`.

The first time it runs it may take longer than usual since it will generate the `precomp` file needed for fast MSM 
required by `go-verkle`. This file is saved and not recomputed in subsequent runs.


## Results 
The following are the results run on an AMD Ryzen 7 3800XT processor and https://github.com/crate-crypto/go-ipa/commit/a49f82a18f8cc7c6e73a4b7a704ee7c7a4ea1546go-ipa@a49f (with some twists only metric collection).

```
$ go run main.go
##### VKT proof benchmark #####
Setup: tree with 1000000 random key-values...

For 1 random-key values of the tree:
        Generating proof took 51ms:
                Collected 8 polynomials evaluations to prove
                Collecting those polys evals (comm, evals, etc) took 0ms
                The rest (Multiproof + nits) took 51ms
        Serializing proof took 0ms
        Deserializing proof took 0ms

For 1000 random-key values of the tree:
        Generating proof took 217ms:
                Collected 7307 polynomials evaluations to prove
                Collecting those polys evals (comm, evals, etc) took 112ms
                The rest (Multiproof + nits) took 104ms
        Serializing proof took 1ms
        Deserializing proof took 30ms

For 2000 random-key values of the tree:
        Generating proof took 364ms:
                Collected 14352 polynomials evaluations to prove
                Collecting those polys evals (comm, evals, etc) took 203ms
                The rest (Multiproof + nits) took 160ms
        Serializing proof took 1ms
        Deserializing proof took 53ms

For 4000 random-key values of the tree:
        Generating proof took 661ms:
                Collected 28377 polynomials evaluations to prove
                Collecting those polys evals (comm, evals, etc) took 373ms
                The rest (Multiproof + nits) took 287ms
        Serializing proof took 3ms
        Deserializing proof took 102ms

For 8000 random-key values of the tree:
        Generating proof took 1287ms:
                Collected 56284 polynomials evaluations to prove
                Collecting those polys evals (comm, evals, etc) took 685ms
                The rest (Multiproof + nits) took 602ms
        Serializing proof took 5ms
        Deserializing proof took 195ms

For 16000 random-key values of the tree:
        Generating proof took 3802ms:
                Collected 111399 polynomials evaluations to prove
                Collecting those polys evals (comm, evals, etc) took 2336ms
                The rest (Multiproof + nits) took 1465ms
        Serializing proof took 16ms
        Deserializing proof took 385ms


##### Raw polynomials multiproof benchmark #####
For 1 polynomials:
        Proving:
                Proof serialization duration: 0.14ms
                Proof creation duration: 164ms
                Duration per step:
                        Generate challenge r and powers : 0.04ms
                        Calculate t, g(x) and D         : 64.63ms
                        Calculate h(x) and E            : 49.50ms
                        Calculate (h-g)(x) and E-D      : 0.02ms
                        IPA for (h-g)(x) and E-D on t   : 50.13ms
        Verification:
                Proof deserialization duration: 0.14ms
                Total duration: 5ms
                Duration per step:
                        Generate challenge r and powers                       : 0.00ms
                        Calculating helper_scalars r^i/(t-z_i)                : 0.00ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.03ms
                        Compute E                                             : 0.33ms
                        Compute E-D and verify IPA                            : 5.05ms

For 1000 polynomials:
        Proving:
                Proof serialization duration: 0.10ms
                Proof creation duration: 64ms
                Duration per step:
                        Generate challenge r and powers : 1.07ms
                        Calculate t, g(x) and D         : 14.39ms
                        Calculate h(x) and E            : 4.68ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.40ms
        Verification:
                Proof deserialization duration: 0.10ms
                Total duration: 8ms
                Duration per step:
                        Generate challenge r and powers                       : 1.14ms
                        Calculating helper_scalars r^i/(t-z_i)                : 0.05ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.05ms
                        Compute E                                             : 2.66ms
                        Compute E-D and verify IPA                            : 4.17ms

For 2000 polynomials:
        Proving:
                Proof serialization duration: 0.11ms
                Proof creation duration: 69ms
                Duration per step:
                        Generate challenge r and powers : 2.19ms
                        Calculate t, g(x) and D         : 20.68ms
                        Calculate h(x) and E            : 4.07ms
                        Calculate (h-g)(x) and E-D      : 0.00ms
                        IPA for (h-g)(x) and E-D on t   : 42.69ms
        Verification:
                Proof deserialization duration: 0.11ms
                Total duration: 11ms
                Duration per step:
                        Generate challenge r and powers                       : 2.45ms
                        Calculating helper_scalars r^i/(t-z_i)                : 0.10ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.06ms
                        Compute E                                             : 4.55ms
                        Compute E-D and verify IPA                            : 4.11ms

For 4000 polynomials:
        Proving:
                Proof serialization duration: 0.10ms
                Proof creation duration: 82ms
                Duration per step:
                        Generate challenge r and powers : 5.40ms
                        Calculate t, g(x) and D         : 31.12ms
                        Calculate h(x) and E            : 3.95ms
                        Calculate (h-g)(x) and E-D      : 0.00ms
                        IPA for (h-g)(x) and E-D on t   : 42.45ms
        Verification:
                Proof deserialization duration: 0.10ms
                Total duration: 16ms
                Duration per step:
                        Generate challenge r and powers                       : 4.57ms
                        Calculating helper_scalars r^i/(t-z_i)                : 0.25ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.05ms
                        Compute E                                             : 8.41ms
                        Compute E-D and verify IPA                            : 3.32ms

For 8000 polynomials:
        Proving:
                Proof serialization duration: 0.10ms
                Proof creation duration: 116ms
                Duration per step:
                        Generate challenge r and powers : 14.32ms
                        Calculate t, g(x) and D         : 54.71ms
                        Calculate h(x) and E            : 4.07ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 43.18ms
        Verification:
                Proof deserialization duration: 0.10ms
                Total duration: 25ms
                Duration per step:
                        Generate challenge r and powers                       : 8.04ms
                        Calculating helper_scalars r^i/(t-z_i)                : 0.49ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.04ms
                        Compute E                                             : 13.78ms
                        Compute E-D and verify IPA                            : 2.97ms

For 16000 polynomials:
        Proving:
                Proof serialization duration: 0.10ms
                Proof creation duration: 191ms
                Duration per step:
                        Generate challenge r and powers : 43.83ms
                        Calculate t, g(x) and D         : 100.36ms
                        Calculate h(x) and E            : 3.98ms
                        Calculate (h-g)(x) and E-D      : 0.00ms
                        IPA for (h-g)(x) and E-D on t   : 42.86ms
        Verification:
                Proof deserialization duration: 0.10ms
                Total duration: 40ms
                Duration per step:
                        Generate challenge r and powers                       : 15.18ms
                        Calculating helper_scalars r^i/(t-z_i)                : 0.99ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.03ms
                        Compute E                                             : 21.28ms
                        Compute E-D and verify IPA                            : 2.54ms

For 128000 polynomials:
        Proving:
                Proof serialization duration: 0.10ms
                Proof creation duration: 2874ms
                Duration per step:
                        Generate challenge r and powers : 2077.34ms
                        Calculate t, g(x) and D         : 749.15ms
                        Calculate h(x) and E            : 4.04ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.41ms
        Verification:
                Proof deserialization duration: 0.10ms
                Total duration: 236ms
                Duration per step:
                        Generate challenge r and powers                       : 100.18ms
                        Calculating helper_scalars r^i/(t-z_i)                : 18.41ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.05ms
                        Compute E                                             : 114.70ms
                        Compute E-D and verify IPA                            : 3.28ms
```
