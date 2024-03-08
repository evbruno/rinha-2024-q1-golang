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
> request count                                      61503 (OK=61199  KO=304   )
> min response time                                      0 (OK=0      KO=10000 )
> max response time                                  28166 (OK=28166  KO=10244 )
> mean response time                                   805 (OK=759    KO=10023 )
> std deviation                                       1607 (OK=1473   KO=57    )
> response time 50th percentile                        178 (OK=175    KO=10002 )
> response time 75th percentile                        944 (OK=928    KO=10005 )
> response time 95th percentile                       3339 (OK=3056   KO=10173 )
> response time 99th percentile                       7938 (OK=7518   KO=10225 )
> mean requests/sec                                229.489 (OK=228.354 KO=1.134 )
---- Response Time Distribution ------------------------------------------------
> t < 800 ms                                         43991 ( 72%)
> 800 ms <= t < 1200 ms                               4625 (  8%)
> t >= 1200 ms                                       12583 ( 20%)
> failed                                               304 (  0%)
---- Errors --------------------------------------------------------------------
> j.i.IOException: Premature close                                  304 (100.0%)
================================================================================
```