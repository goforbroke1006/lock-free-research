#!/bin/bash

SAMPLES_COUNT=$1

set -e

if [[ -z $SAMPLES_COUNT ]]; then
  SAMPLES_COUNT=1000000
fi

FILENAME="./numbers-${SAMPLES_COUNT}.txt"

rm -f ${FILENAME} | true

for i in $(seq 1 ${SAMPLES_COUNT}); do
  echo $((10 + $RANDOM % 89)) >>${FILENAME}
done
