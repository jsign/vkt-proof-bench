# vkt-proof-bench
This repository contains benchmarks to measure the performance of proof generation and verification in Verkle Trees.

## Results 
The following are the results run on an AMD Ryzen 7 3800XT processor and `go-ipa@ce4a969`.

```
$ go run main.go
For 1 polynomials:
        Proving:
                Avg. proof serialization duration: 0.06ms
                Avg. proof creation duration: 52ms
                Avg. time per milestone:
                        Generate challenge r and powers : 0.02ms
                        Calculate t, g(x) and D         : 3.90ms
                        Calculate h(x) and E            : 3.04ms
                        Calculate (h-g)(x) and E-D      : 0.00ms
                        IPA for (h-g)(x) and E-D on t   : 45.24ms
        Verification:
                Avg. proof deserialization duration: 0.06ms
                Avg. total duration: 27ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 0.01ms
                        Calculating helper_scalars r^i/(t-z_i)                : 0.02ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.00ms
                        Compute E                                             : 0.15ms
                        Compute E-D and verify IPA                            : 27.74ms

For 625 polynomials:
        Proving:
                Avg. proof serialization duration: 0.05ms
                Avg. proof creation duration: 73ms
                Avg. time per milestone:
                        Generate challenge r and powers : 1.98ms
                        Calculate t, g(x) and D         : 16.56ms
                        Calculate h(x) and E            : 11.16ms
                        Calculate (h-g)(x) and E-D      : 0.00ms
                        IPA for (h-g)(x) and E-D on t   : 43.83ms
        Verification:
                Avg. proof deserialization duration: 0.05ms
                Avg. total duration: 78ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 3.35ms
                        Calculating helper_scalars r^i/(t-z_i)                : 4.73ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.03ms
                        Compute E                                             : 43.56ms
                        Compute E-D and verify IPA                            : 26.87ms

For 1250 polynomials:
        Proving:
                Avg. proof serialization duration: 0.05ms
                Avg. proof creation duration: 100ms
                Avg. time per milestone:
                        Generate challenge r and powers : 3.83ms
                        Calculate t, g(x) and D         : 31.50ms
                        Calculate h(x) and E            : 20.68ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.12ms
        Verification:
                Avg. proof deserialization duration: 0.05ms
                Avg. total duration: 126ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 6.06ms
                        Calculating helper_scalars r^i/(t-z_i)                : 7.48ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.07ms
                        Compute E                                             : 86.53ms
                        Compute E-D and verify IPA                            : 26.58ms

For 2500 polynomials:
        Proving:
                Avg. proof serialization duration: 0.05ms
                Avg. proof creation duration: 146ms
                Avg. time per milestone:
                        Generate challenge r and powers : 7.38ms
                        Calculate t, g(x) and D         : 58.51ms
                        Calculate h(x) and E            : 36.27ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.72ms
        Verification:
                Avg. proof deserialization duration: 0.05ms
                Avg. total duration: 217ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 11.96ms
                        Calculating helper_scalars r^i/(t-z_i)                : 11.89ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.08ms
                        Compute E                                             : 167.31ms
                        Compute E-D and verify IPA                            : 26.00ms

For 5000 polynomials:
        Proving:
                Avg. proof serialization duration: 0.05ms
                Avg. proof creation duration: 246ms
                Avg. time per milestone:
                        Generate challenge r and powers : 15.26ms
                        Calculate t, g(x) and D         : 116.11ms
                        Calculate h(x) and E            : 70.32ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.37ms
        Verification:
                Avg. proof deserialization duration: 0.05ms
                Avg. total duration: 405ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 20.50ms
                        Calculating helper_scalars r^i/(t-z_i)                : 22.09ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.29ms
                        Compute E                                             : 336.92ms
                        Compute E-D and verify IPA                            : 26.15ms

For 10000 polynomials:
        Proving:
                Avg. proof serialization duration: 0.05ms
                Avg. proof creation duration: 434ms
                Avg. time per milestone:
                        Generate challenge r and powers : 29.28ms
                        Calculate t, g(x) and D         : 224.31ms
                        Calculate h(x) and E            : 136.31ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.89ms
        Verification:
                Avg. proof deserialization duration: 0.05ms
                Avg. total duration: 769ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 34.32ms
                        Calculating helper_scalars r^i/(t-z_i)                : 42.29ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.47ms
                        Compute E                                             : 666.25ms
                        Compute E-D and verify IPA                            : 26.56ms
```
