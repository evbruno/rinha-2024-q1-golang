# Golang impl

[rinha-de-backend-2024-q1](https://github.com/zanfranceschi/rinha-de-backend-2024-q1/)

### Impl

- go 1.22 std lib
- 2 instâncias
- salvando em memória
- nginx faz hash no ID p/ achar a instância correta

### Rodandoooo 🏃🏻‍♂️🏃🏻‍♂️

```bash
 docker compose -f docker-compose-dev.yml up --build
```

### Resultado local

```
================================================================================
---- Global Information --------------------------------------------------------
> request count                                      61503 (OK=61503  KO=0     )
> min response time                                      2 (OK=2      KO=-     )
> max response time                                    585 (OK=585    KO=-     )
> mean response time                                    11 (OK=11     KO=-     )
> std deviation                                         25 (OK=25     KO=-     )
> response time 50th percentile                          4 (OK=4      KO=-     )
> response time 75th percentile                          5 (OK=5      KO=-     )
> response time 95th percentile                         53 (OK=54     KO=-     )
> response time 99th percentile                        119 (OK=119    KO=-     )
> mean requests/sec                                251.033 (OK=251.033 KO=-     )
---- Response Time Distribution ------------------------------------------------
> t < 800 ms                                         61503 (100%)
> 800 ms <= t < 1200 ms                                  0 (  0%)
> t >= 1200 ms                                           0 (  0%)
> failed                                                 0 (  0%)
================================================================================
```