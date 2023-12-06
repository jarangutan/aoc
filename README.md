# AOC 2023

First AOC, and first time using golang

```
perf stat --all-user -B -ddd
```

## Better bench test

Grabbed from https://github.com/erik-adelbert/aoc/blob/main/2023/bench.sh

```sh
hyperfine --warmup 100 "cat ../inputs/input.txt > /dev/null" "./cmd < ../inputs/input.txt" --export-markdown time.md
```
