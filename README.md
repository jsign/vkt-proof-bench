# vkt-proof-bench
This repository contains benchmarks to measure the performance of proof generation and verification in Verkle Trees.

## Results 
The following are the results run on a AMD Ryzen 7 3800XT processor and `go-ipa@ce4a969`.

```
$ go run main.go
For 1 polynomials:
        Proving:
                Avg. total duration: 52ms
                Avg. time per milestone:
                        Generate challenge r and powers : 16us
                        Calculate t, g(x) and D         : 4026us
                        Calculate h(x) and E            : 3096us
                        Calculate (h-g)(x) and E-D      : 3us
                        IPA for (h-g)(x) and E-D on t   : 45341us
        Verification:
                Avg. total duration: 27ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 9us
                        Calculating helper_scalars r^i/(t-z_i)                : 18us
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 0us
                        Compute E                                             : 151us
                        Compute E-D and verify IPA                            : 27024us

For 625 polynomials:
        Proving:
                Avg. total duration: 71ms
                Avg. time per milestone:
                        Generate challenge r and powers : 1898us
                        Calculate t, g(x) and D         : 16429us
                        Calculate h(x) and E            : 10838us
                        Calculate (h-g)(x) and E-D      : 4us
                        IPA for (h-g)(x) and E-D on t   : 42333us
        Verification:
                Avg. total duration: 77ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 3494us
                        Calculating helper_scalars r^i/(t-z_i)                : 4328us
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 23us
                        Compute E                                             : 43997us
                        Compute E-D and verify IPA                            : 25644us

For 1250 polynomials:
        Proving:
                Avg. total duration: 98ms
                Avg. time per milestone:
                        Generate challenge r and powers : 3853us
                        Calculate t, g(x) and D         : 30862us
                        Calculate h(x) and E            : 19575us
                        Calculate (h-g)(x) and E-D      : 4us
                        IPA for (h-g)(x) and E-D on t   : 44088us
        Verification:
                Avg. total duration: 123ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 6342us
                        Calculating helper_scalars r^i/(t-z_i)                : 7750us
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 38us
                        Compute E                                             : 83024us
                        Compute E-D and verify IPA                            : 26418us

For 2500 polynomials:
        Proving:
                Avg. total duration: 146ms
                Avg. time per milestone:
                        Generate challenge r and powers : 7471us
                        Calculate t, g(x) and D         : 58886us
                        Calculate h(x) and E            : 36399us
                        Calculate (h-g)(x) and E-D      : 7us
                        IPA for (h-g)(x) and E-D on t   : 43985us
        Verification:
                Avg. total duration: 216ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 11132us
                        Calculating helper_scalars r^i/(t-z_i)                : 11422us
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 67us
                        Compute E                                             : 168035us
                        Compute E-D and verify IPA                            : 25462us

For 5000 polynomials:
        Proving:
                Avg. total duration: 241ms
                Avg. time per milestone:
                        Generate challenge r and powers : 14655us
                        Calculate t, g(x) and D         : 113242us
                        Calculate h(x) and E            : 68901us
                        Calculate (h-g)(x) and E-D      : 8us
                        IPA for (h-g)(x) and E-D on t   : 45163us
        Verification:
                Avg. total duration: 401ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 19620us
                        Calculating helper_scalars r^i/(t-z_i)                : 20901us
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 133us
                        Compute E                                             : 334262us
                        Compute E-D and verify IPA                            : 27046us

For 10000 polynomials:
        Proving:
                Avg. total duration: 435ms
                Avg. time per milestone:
                        Generate challenge r and powers : 29657us
                        Calculate t, g(x) and D         : 225042us
                        Calculate h(x) and E            : 136023us
                        Calculate (h-g)(x) and E-D      : 7us
                        IPA for (h-g)(x) and E-D on t   : 44849us
        Verification:
                Avg. total duration: 767ms
                Avg. time per milestone:
                        Generate challenge r and powers                       : 34594us
                        Calculating helper_scalars r^i/(t-z_i)                : 41737us
                        g_2(t) = SUM y_i*(r^i/(t-z_i))=SUM y_i*helper_scalars : 353us
                        Compute E                                             : 663987us
                        Compute E-D and verify IPA                            : 26415us
```
