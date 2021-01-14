#!/bin/bash

CC=$1
FILENAME="./benchmark-${CC}.txt"

rm -f "${FILENAME}"

uname -a | tee -a "${FILENAME}"
cat /proc/cpuinfo | grep "model name" | head -1 | tee -a "${FILENAME}"
cat /proc/cpuinfo | grep "cpu MHz" | head -1 | tee -a "${FILENAME}"
cat /proc/cpuinfo | grep "cache size" | head -1 | tee -a "${FILENAME}"

echo '--------------------------------------------------' >>"${FILENAME}"

go build

samplesSizes=(
  10000
  100000
  1000000
  10000000
  100000000
)

# Prepare assets

for size in ${samplesSizes[*]}; do
  printf "==== SIZE: %s\n" $size

  if [[ ! -f "./numbers-${size}.txt" ]]; then
    bash ./gen-rand-numbers.sh ${size}
  fi

  ./lock-free-research util run-standard --concurrent=${CC} --filename=numbers-${size}.txt | tee -a "${FILENAME}"
  ./lock-free-research util run-lock-free --concurrent=${CC} --filename=numbers-${size}.txt | tee -a "${FILENAME}"

done
