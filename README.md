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
                The rest (Multiproof + nits) took 52ms
        Serializing proof took 0ms
        Deserializing proof took 0ms

For 1000 random-key values of the tree:
        Generating proof took 440ms:
                Collected 7307 polynomials evaluations to prove
                Collecting those polys evals (comm, evals, etc) took 107ms
                The rest (Multiproof + nits) took 332ms
        Serializing proof took 12ms
        Deserializing proof took 25ms

For 2000 random-key values of the tree:
        Generating proof took 797ms:
                Collected 14352 polynomials evaluations to prove
                Collecting those polys evals (comm, evals, etc) took 193ms
                The rest (Multiproof + nits) took 604ms
        Serializing proof took 18ms
        Deserializing proof took 48ms

For 4000 random-key values of the tree:
        Generating proof took 1473ms:
                Collected 28377 polynomials evaluations to prove
                Collecting those polys evals (comm, evals, etc) took 341ms
                The rest (Multiproof + nits) took 1132ms
        Serializing proof took 33ms
        Deserializing proof took 93ms

For 8000 random-key values of the tree:
        Generating proof took 4617ms:
                Collected 56284 polynomials evaluations to prove
                Collecting those polys evals (comm, evals, etc) took 1771ms
                The rest (Multiproof + nits) took 2845ms
        Serializing proof took 65ms
        Deserializing proof took 192ms

For 16000 random-key values of the tree:
        Generating proof took 7731ms:
                Collected 111399 polynomials evaluations to prove
                Collecting those polys evals (comm, evals, etc) took 2667ms
                The rest (Multiproof + nits) took 5064ms
        Serializing proof took 134ms
        Deserializing proof took 368ms


##### Raw polynomials multiproof benchmark #####
For 1 polynomials:
        Proving:
                Proof serialization duration: 0.13ms
                Proof creation duration: 219ms
                Duration per step:
                        Generate challenge r and powers : 0.02ms
                        Calculate t, g(x) and D         : 95.73ms
                        Calculate h(x) and E            : 76.28ms
                        Calculate (h-g)(x) and E-D      : 0.02ms
                        IPA for (h-g)(x) and E-D on t   : 47.08ms
        Verification:
                Proof deserialization duration: 0.13ms
                Total duration: 28ms
                Duration per step:
                        Generate challenge r and powers                       : 0.00ms
                        Calculating helper_scalars r^i/(t-z_i)                : 0.01ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.00ms
                        Compute E                                             : 0.14ms
                        Compute E-D and verify IPA                            : 27.91ms

For 1000 polynomials:
        Proving:
                Proof serialization duration: 0.10ms
                Proof creation duration: 89ms
                Duration per step:
                        Generate challenge r and powers : 3.07ms
                        Calculate t, g(x) and D         : 25.80ms
                        Calculate h(x) and E            : 16.57ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.20ms
        Verification:
                Proof deserialization duration: 0.10ms
                Total duration: 101ms
                Duration per step:
                        Generate challenge r and powers                       : 1.01ms
                        Calculating helper_scalars r^i/(t-z_i)                : 6.12ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.04ms
                        Compute E                                             : 68.20ms
                        Compute E-D and verify IPA                            : 26.31ms

For 2000 polynomials:
        Proving:
                Proof serialization duration: 0.10ms
                Proof creation duration: 127ms
                Duration per step:
                        Generate challenge r and powers : 5.94ms
                        Calculate t, g(x) and D         : 47.18ms
                        Calculate h(x) and E            : 29.49ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.40ms
        Verification:
                Proof deserialization duration: 0.10ms
                Total duration: 172ms
                Duration per step:
                        Generate challenge r and powers                       : 1.74ms
                        Calculating helper_scalars r^i/(t-z_i)                : 9.88ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.05ms
                        Compute E                                             : 134.53ms
                        Compute E-D and verify IPA                            : 25.95ms

For 4000 polynomials:
        Proving:
                Proof serialization duration: 0.10ms
                Proof creation duration: 207ms
                Duration per step:
                        Generate challenge r and powers : 11.96ms
                        Calculate t, g(x) and D         : 93.08ms
                        Calculate h(x) and E            : 57.19ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 45.14ms
        Verification:
                Proof deserialization duration: 0.10ms
                Total duration: 317ms
                Duration per step:
                        Generate challenge r and powers                       : 2.50ms
                        Calculating helper_scalars r^i/(t-z_i)                : 17.14ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.15ms
                        Compute E                                             : 271.05ms
                        Compute E-D and verify IPA                            : 26.63ms

For 8000 polynomials:
        Proving:
                Proof serialization duration: 0.11ms
                Proof creation duration: 359ms
                Duration per step:
                        Generate challenge r and powers : 23.63ms
                        Calculate t, g(x) and D         : 181.51ms
                        Calculate h(x) and E            : 110.12ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.68ms
        Verification:
                Proof deserialization duration: 0.11ms
                Total duration: 603ms
                Duration per step:
                        Generate challenge r and powers                       : 5.09ms
                        Calculating helper_scalars r^i/(t-z_i)                : 34.15ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.28ms
                        Compute E                                             : 538.68ms
                        Compute E-D and verify IPA                            : 25.77ms

For 16000 polynomials:
        Proving:
                Proof serialization duration: 0.11ms
                Proof creation duration: 664ms
                Duration per step:
                        Generate challenge r and powers : 46.86ms
                        Calculate t, g(x) and D         : 357.84ms
                        Calculate h(x) and E            : 215.06ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.73ms
        Verification:
                Proof deserialization duration: 0.11ms
                Total duration: 1166ms
                Duration per step:
                        Generate challenge r and powers                       : 10.40ms
                        Calculating helper_scalars r^i/(t-z_i)                : 66.99ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 1.04ms
                        Compute E                                             : 1062.79ms
                        Compute E-D and verify IPA                            : 25.55ms
```
