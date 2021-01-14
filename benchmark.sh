#!/bin/bash

rm -f benchmark.txt

uname -a | tee -a benchmark.txt
cat /proc/cpuinfo | grep "model name" | head -1 | tee -a benchmark.txt
cat /proc/cpuinfo | grep "cpu MHz" | head -1 | tee -a benchmark.txt
cat /proc/cpuinfo | grep "cache size" | head -1 | tee -a benchmark.txt

echo '--------------------------------------------------' >> benchmark.txt

go build

# Prepare assets

if [[ ! -f ./numbers-10000.txt ]]; then
  bash ./gen-rand-numbers.sh 10000
fi
if [[ ! -f ./numbers-100000.txt ]]; then
  bash ./gen-rand-numbers.sh 100000
fi
if [[ ! -f ./numbers-1000000.txt ]]; then
  bash ./gen-rand-numbers.sh 1000000
fi
if [[ ! -f ./numbers-10000000.txt ]]; then
  bash ./gen-rand-numbers.sh 10000000
fi

# run
./lock-free-research util run-standard --filename=numbers-10000.txt | tee -a benchmark.txt
./lock-free-research util run-lock-free --filename=numbers-10000.txt | tee -a benchmark.txt

./lock-free-research util run-standard --filename=numbers-100000.txt | tee -a benchmark.txt
./lock-free-research util run-lock-free --filename=numbers-100000.txt | tee -a benchmark.txt

./lock-free-research util run-standard --filename=numbers-1000000.txt | tee -a benchmark.txt
./lock-free-research util run-lock-free --filename=numbers-1000000.txt | tee -a benchmark.txt

./lock-free-research util run-standard --filename=numbers-10000000.txt | tee -a benchmark.txt
./lock-free-research util run-lock-free --filename=numbers-10000000.txt | tee -a benchmark.txt
