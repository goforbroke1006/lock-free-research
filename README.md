# lock-free-research

### Benchmarks

```bash
bash ./benchmark.sh 1
bash ./benchmark.sh 10
bash ./benchmark.sh 50
bash ./benchmark.sh 100

```

### Numbers counter benchmark (duration comparison)

```text
model name	: AMD Ryzen 7 3700X 8-Core Processor
cpu MHz		: 2223.645
cache size	: 512 KB
```

| Samples count | Mutex based  | Lock free   |
|---------------|--------------|-------------|
| 10'000        | 103874226    | 53783700    | 1.93
| 100'000       | 1200054903   | 169444553   | 7.08
| 1'000'000     | 12326890996  | 1724432715  | 7.14
| 10'000'000    | 124739376899 | 17250465301 | 7.2

```text
model name	: Intel(R) Core(TM) i7-10850H CPU @ 2.70GHz
cpu MHz		: 897.823
cache size	: 12288 KB
```

| Samples count | CC  | Mutex based   | Lock free   |
|---------------|-----|---------------|-------------|
| 10'000        | 100 | 96760456      | 103742195    | v 1.07
| 100'000       | 100 | 962737425     | 873253398    | 1.10
| 1'000'000     | 100 | 9822636050    | 9493856299   | 1.03
| 10'000'000    | 100 | 96703462565   | 81477100998  | 1.18
| 100'000'000   | 100 | 1026246266034 | 904331483288  | 1.13