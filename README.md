# vkt-proof-bench
This repository contains benchmarks to measure the performance of proof generation and verification in Verkle Trees.

## Run

You can run this benchmark by just doing: `go run main.go`.

The first time it runs it may take longer than usual since it will generate the `precomp` file needed for fast MSM 
required by `go-verkle`. This file is saved and not recomputed in subsequent runs.


## Results 
The following are the results run on an AMD Ryzen 7 3800XT processor and `go-ipa@ce4a969`.

```
$ go run main.go
##### VKT proof benchmark #####
Setup: tree with 1000000 random key-values...

For 1 random-key values of the tree:
        Generating proof took 53ms:
                Collected 8 polynomials evaluations to prove
                Collecting those polys evals (comm, evals, etc) took 0ms
                The rest (Multiproof + nits) took 53ms
        Serializing proof took 0ms
        Deserializing proof took 0ms

For 1000 random-key values of the tree:
        Generating proof took 435ms:
                Collected 7307 polynomials evaluations to prove
                Collecting those polys evals (comm, evals, etc) took 106ms
                The rest (Multiproof + nits) took 328ms
        Serializing proof took 12ms
        Deserializing proof took 24ms

For 2000 random-key values of the tree:
        Generating proof took 786ms:
                Collected 14352 polynomials evaluations to prove
                Collecting those polys evals (comm, evals, etc) took 191ms
                The rest (Multiproof + nits) took 594ms
        Serializing proof took 18ms
        Deserializing proof took 47ms

For 4000 random-key values of the tree:
        Generating proof took 1506ms:
                Collected 28377 polynomials evaluations to prove
                Collecting those polys evals (comm, evals, etc) took 364ms
                The rest (Multiproof + nits) took 1142ms
        Serializing proof took 33ms
        Deserializing proof took 93ms

For 8000 random-key values of the tree:
        Generating proof took 2922ms:
                Collected 56284 polynomials evaluations to prove
                Collecting those polys evals (comm, evals, etc) took 697ms
                The rest (Multiproof + nits) took 2224ms
        Serializing proof took 59ms
        Deserializing proof took 183ms

For 16000 random-key values of the tree:
        Generating proof took 11385ms:
                Collected 111399 polynomials evaluations to prove
                Collecting those polys evals (comm, evals, etc) took 4456ms
                The rest (Multiproof + nits) took 6929ms
        Serializing proof took 152ms
        Deserializing proof took 410ms


##### Raw polynomials multiproof benchmark #####
For 1 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 174ms
                Duration per step:
                        Generate challenge r and powers : 0.02ms
                        Calculate t, g(x) and D         : 71.38ms
                        Calculate h(x) and E            : 55.64ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 47.01ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 28ms
                Duration per step:
                        Generate challenge r and powers                       : 0.01ms
                        Calculating helper_scalars r^i/(t-z_i)                : 0.02ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.00ms
                        Compute E                                             : 0.15ms
                        Compute E-D and verify IPA                            : 27.85ms

For 1000 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 89ms
                Duration per step:
                        Generate challenge r and powers : 3.35ms
                        Calculate t, g(x) and D         : 25.63ms
                        Calculate h(x) and E            : 16.33ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.62ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 106ms
                Duration per step:
                        Generate challenge r and powers                       : 5.70ms
                        Calculating helper_scalars r^i/(t-z_i)                : 6.29ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.03ms
                        Compute E                                             : 66.43ms
                        Compute E-D and verify IPA                            : 27.62ms

For 2000 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 127ms
                Duration per step:
                        Generate challenge r and powers : 6.03ms
                        Calculate t, g(x) and D         : 47.41ms
                        Calculate h(x) and E            : 29.88ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.62ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 180ms
                Duration per step:
                        Generate challenge r and powers                       : 10.35ms
                        Calculating helper_scalars r^i/(t-z_i)                : 9.69ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.07ms
                        Compute E                                             : 134.13ms
                        Compute E-D and verify IPA                            : 25.96ms

For 4000 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 205ms
                Duration per step:
                        Generate challenge r and powers : 11.79ms
                        Calculate t, g(x) and D         : 93.04ms
                        Calculate h(x) and E            : 56.08ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.61ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 328ms
                Duration per step:
                        Generate challenge r and powers                       : 16.56ms
                        Calculating helper_scalars r^i/(t-z_i)                : 17.07ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.12ms
                        Compute E                                             : 267.59ms
                        Compute E-D and verify IPA                            : 26.81ms

For 8000 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 365ms
                Duration per step:
                        Generate challenge r and powers : 24.23ms
                        Calculate t, g(x) and D         : 185.18ms
                        Calculate h(x) and E            : 111.77ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.38ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 644ms
                Duration per step:
                        Generate challenge r and powers                       : 29.68ms
                        Calculating helper_scalars r^i/(t-z_i)                : 35.04ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.71ms
                        Compute E                                             : 552.71ms
                        Compute E-D and verify IPA                            : 26.42ms

For 16000 polynomials:
        Proving:
                Proof serialization duration: 0.05ms
                Proof creation duration: 676ms
                Duration per step:
                        Generate challenge r and powers : 47.89ms
                        Calculate t, g(x) and D         : 366.09ms
                        Calculate h(x) and E            : 217.84ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.84ms
        Verification:
                Proof deserialization duration: 0.05ms
                Total duration: 1215ms
                Duration per step:
                        Generate challenge r and powers                       : 52.49ms
                        Calculating helper_scalars r^i/(t-z_i)                : 68.15ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 1.04ms
                        Compute E                                             : 1067.29ms
                        Compute E-D and verify IPA                            : 26.74ms
```
