# vkt-proof-bench
This repository contains benchmarks to measure the performance of proof generation and verification in Verkle Trees.

## Results 
The following are the results run on a AMD Ryzen 7 3800XT processor and `go-ipa@ce4a969`.

```
$ go run main.go
For 1 polynomials:
        Proving:
                Avg. total duration: 51ms
                Avg. time per milestone:
                        Generate challenge r and powers : 0.02ms
                        Calculate t, g(x) and D         : 3.58ms
                        Calculate h(x) and E            : 2.98ms
                        Calculate (h-g)(x) and E-D      : 0.00ms
                        IPA for (h-g)(x) and E-D on t   : 45.10ms
        Verification:
                Avg. total duration: 27ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 0.01ms
                        Calculating helper_scalars r^i/(t-z_i)                : 0.02ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.00ms
                        Compute E                                             : 0.14ms
                        Compute E-D and verify IPA                            : 27.78ms

For 625 polynomials:
        Proving:
                Avg. total duration: 75ms
                Avg. time per milestone:
                        Generate challenge r and powers : 1.95ms
                        Calculate t, g(x) and D         : 17.07ms
                        Calculate h(x) and E            : 11.16ms
                        Calculate (h-g)(x) and E-D      : 0.00ms
                        IPA for (h-g)(x) and E-D on t   : 45.32ms
        Verification:
                Avg. total duration: 77ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 3.57ms
                        Calculating helper_scalars r^i/(t-z_i)                : 4.36ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.03ms
                        Compute E                                             : 44.19ms
                        Compute E-D and verify IPA                            : 25.57ms

For 1250 polynomials:
        Proving:
                Avg. total duration: 98ms
                Avg. time per milestone:
                        Generate challenge r and powers : 4.22ms
                        Calculate t, g(x) and D         : 30.96ms
                        Calculate h(x) and E            : 19.43ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.09ms
        Verification:
                Avg. total duration: 122ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 6.34ms
                        Calculating helper_scalars r^i/(t-z_i)                : 7.86ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.04ms
                        Compute E                                             : 82.45ms
                        Compute E-D and verify IPA                            : 26.11ms

For 2500 polynomials:
        Proving:
                Avg. total duration: 146ms
                Avg. time per milestone:
                        Generate challenge r and powers : 7.54ms
                        Calculate t, g(x) and D         : 58.33ms
                        Calculate h(x) and E            : 36.58ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.31ms
        Verification:
                Avg. total duration: 214ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 11.89ms
                        Calculating helper_scalars r^i/(t-z_i)                : 10.82ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.09ms
                        Compute E                                             : 165.97ms
                        Compute E-D and verify IPA                            : 26.00ms

For 5000 polynomials:
        Proving:
                Avg. total duration: 243ms
                Avg. time per milestone:
                        Generate challenge r and powers : 15.21ms
                        Calculate t, g(x) and D         : 113.79ms
                        Calculate h(x) and E            : 69.90ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 44.37ms
        Verification:
                Avg. total duration: 401ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 18.69ms
                        Calculating helper_scalars r^i/(t-z_i)                : 21.32ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.24ms
                        Compute E                                             : 335.21ms
                        Compute E-D and verify IPA                            : 26.03ms

For 10000 polynomials:
        Proving:
                Avg. total duration: 433ms
                Avg. time per milestone:
                        Generate challenge r and powers : 29.42ms
                        Calculate t, g(x) and D         : 223.13ms
                        Calculate h(x) and E            : 135.06ms
                        Calculate (h-g)(x) and E-D      : 0.01ms
                        IPA for (h-g)(x) and E-D on t   : 45.61ms
        Verification:
                Avg. total duration: 758ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 33.44ms
                        Calculating helper_scalars r^i/(t-z_i)                : 41.41ms
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0.39ms
                        Compute E                                             : 657.33ms
                        Compute E-D and verify IPA                            : 25.99ms
```
