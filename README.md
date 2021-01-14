# lock-free-research

### Benchmarks

```bash
bash ./benchmark.sh
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