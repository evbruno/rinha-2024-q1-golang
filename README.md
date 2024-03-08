# Golang impl

[ref](https://github.com/zanfranceschi/rinha-de-backend-2024-q1/tree/main)

### Impl

- go 1.22 std lib
- 1 instância
- salvando em memória

### Rodandoooo 🏃🏻‍♂️🏃🏻‍♂️

```bash
 docker compose -f docker-compose-dev.yml up
```

### Resultado local

```
================================================================================
---- Global Information --------------------------------------------------------
> request count                                      61503 (OK=61503  KO=0     )
> min response time                                      0 (OK=0      KO=-     )
> max response time                                    172 (OK=172    KO=-     )
> mean response time                                     5 (OK=5      KO=-     )
> std deviation                                         16 (OK=16     KO=-     )
> response time 50th percentile                          1 (OK=1      KO=-     )
> response time 75th percentile                          2 (OK=2      KO=-     )
> response time 95th percentile                         24 (OK=24     KO=-     )
> response time 99th percentile                         94 (OK=94     KO=-     )
> mean requests/sec                                251.033 (OK=251.033 KO=-     )
---- Response Time Distribution ------------------------------------------------
> t < 800 ms                                         61503 (100%)
> 800 ms <= t < 1200 ms                                  0 (  0%)
> t >= 1200 ms                                           0 (  0%)
> failed                                                 0 (  0%)
================================================================================
```